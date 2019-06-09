package exec

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"reflect"
	"strings"

	"github.com/canzerapple/neonx-cli/location"
)

type arg struct {
	Name   string
	Values []interface{}
}

func (m arg) format(prefix string, quota bool) (string, error) {

	if m.Values == nil || len(m.Values) == 0 {
		return fmt.Sprintf("%s%s", prefix, m.Name), nil
	}

	var (
		values = make([]string, 0)
	)

	for _, value := range m.Values {

		var (
			format string
		)

		switch value.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			format = fmt.Sprintf("%d", value)
		case float32, float64:
			format = fmt.Sprintf("%f", value)
		case bool:
			format = fmt.Sprintf("%b", value)
		case string:
			if quota {
				format = fmt.Sprintf(`"%s"`, value)
			} else {
				format = fmt.Sprintf("%s", value)
			}
		default:
			return "", fmt.Errorf(
				"command arg [%s] type [%s] not support ",
				m.Name,
				reflect.TypeOf(value))
		}

		values = append(values, format)
	}

	return fmt.Sprintf(
		"%s=%s",
		m.Name,
		strings.Join(values, ",")), nil
}

type Cmd struct {
	bin      location.Location
	args     []arg
	commands []string
	prefix   string
	Stdin    io.Reader
	Stdout   io.Writer
	Stderr   io.Writer
	cmd      *exec.Cmd
}

func Command(root location.Location, commands ...string) *Cmd {
	return &Cmd{
		bin:      root,
		args:     make([]arg, 0),
		commands: commands,
	}
}

func (m *Cmd) Prefix(prefix string) *Cmd {
	m.prefix = prefix
	return m
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
			"rund command [%s] fail: %s", m.bin, err)
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

	m.cmd = cmd

	err = cmd.Run()

	return err
}
