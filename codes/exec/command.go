package exec

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/canzerapple/neonx-cli/location"
)

type Cmd struct {
	bin      location.Location
	args     []arg
	commands []string
	prefix   string
	Stdin    io.Reader
	Stdout   io.Writer
	Stderr   io.Writer
	cmd      *exec.Cmd
	Dir      location.Location
}

func (m *Cmd) makeArgs(quota bool) ([]string, error) {
	var (
		args = append([]string{}, m.commands...)
	)

	for _, arg := range m.args {

		value, err := arg.format(m.prefix, quota)

		if err != nil {
			return nil, err
		}

		args = append(args, value)
	}

	return args, nil
}

func (m *Cmd) String() string {

	var (
		args, err = m.makeArgs(true)
	)

	if err != nil {
		return fmt.Sprintf("Command [%s] format error: %s", m.bin, err)
	}

	return strings.Join(
		append([]string{string(m.bin)}, args...),
		" ",
	)
}

func (m *Cmd) Arg(name string, v ...interface{}) {
	m.args = append(m.args, arg{Name: name, Values: v})
}

func (m *Cmd) Command(commands ...string) *Cmd {
	m.commands = commands
	return m
}

func (m *Cmd) GetCommand() *exec.Cmd {
	return m.cmd
}

func (m *Cmd) ExitCode() int {
	return m.cmd.ProcessState.ExitCode()
}

func (m *Cmd) Run(ctx context.Context) error {

	args, err := m.makeArgs(false)

	if err != nil {
		return fmt.Errorf(
			"make command [%s] args fail: %s", m.bin, err)
	}

	var (
		cmd = exec.CommandContext(
			ctx,
			string(m.bin),
			args...,
		)
	)

	cmd.Stdin = m.Stdin
	cmd.Stdout = m.Stdout
	cmd.Stderr = m.Stderr

	cmd.Dir = string(m.Dir)

	m.cmd = cmd

	return cmd.Run()
}
