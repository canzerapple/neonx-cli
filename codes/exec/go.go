package exec

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
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
	Cmd
}

func findGoBin(root location.Location) location.Location {

	if runtime.GOOS == "windows" {
		return root.Child("bin/go.exe")
	}

	return root.Child("bin/go")

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

	if cmd.bin, err = location.ToLocation(root[0]); err != nil {
		return nil, fmt.Errorf("load go command fail: %s", err)
	}

	cmd.bin = findGoBin(cmd.bin)

	if exists, err = cmd.bin.IsExists(); err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("go command [%s] not found ", cmd.bin)
	}

	return cmd, nil
}

func (m *Go) GetGoVersion() (*GoVersion, error) {

	cmd, err := LoadGoCommand()

	if err != nil {
		return nil, err
	}

	var (
		writer = bytes.NewBuffer(nil)
		v      = new(GoVersion)
	)

	cmd.Stdout = writer
	cmd.Stderr = writer

	err = cmd.Run(context.Background())

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
