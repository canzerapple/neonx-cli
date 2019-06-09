package exec

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/canzerapple/neonx-cli/location"
)

var (
	EnvironGoRoot = "GOROOT"
)

type GoVersion struct {
	Version string
	Arch    string
	OS      string
}

type Go struct {
	root location.Location
}

func LoadGoCommand() (*Go, error) {

	var (
		env    = GetEnviron()
		root   = env.Find(EnvironGoRoot)
		cmd    = new(Go)
		err    error
		exists bool
	)

	if len(root) != 1 {
		return nil, fmt.Errorf("invalid environ GOROOT ")
	}

	if cmd.root, err = location.ToLocation(root[0]); err != nil {
		return nil, fmt.Errorf("load go command fail: %s", err)
	}

	cmd.root = findGoBin(cmd.root)

	if exists, err = cmd.root.IsExists(); err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("go command [%s] not found ", cmd.root)
	}

	return cmd, nil
}

func (m *Go) GetVersion() (*GoVersion, error) {

	var (
		cmd    = Command(m.root, "version")
		writer = bytes.NewBuffer(nil)
		v      = new(GoVersion)
	)

	cmd.Stdout = writer
	cmd.Stderr = writer

	err := cmd.Run(context.Background())

	if err != nil {
		return nil, err
	}

	if cmd.ExitCode() != 0 {
		return nil, fmt.Errorf("run command %s exit code [%d]", cmd, cmd.ExitCode())
	}

	var (
		info string
		data []string
	)

	_, err = fmt.Fscanf(writer, "go version go%s %s", &v.Version, &info)

	if err != nil {
		return nil, fmt.Errorf("parse go version command fail: %s", err)
	}

	data = strings.Split(info, "/")

	if len(data) == 2 {
		v.OS = data[0]
		v.Arch = data[1]
	}

	return v, nil
}
