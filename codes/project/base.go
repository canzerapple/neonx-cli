package project

import (
	"context"
	"fmt"
	"os"

	"github.com/canzerapple/neonx-cli/codes/exec"

	"github.com/canzerapple/neonx-cli/codes"

	"github.com/canzerapple/neonx-cli/location"
)

type base struct {
	info          *codes.ProjectInfo
	configuration location.Location
}

func (m *base) GetName() string {
	return m.info.Name
}

func (m *base) GetDescribe() string {
	return m.info.Describe
}

func (m *base) GetArchitect() codes.ProjectArchitect {
	return m.info.Architect
}

func NewProject(location location.Location, info *codes.ProjectInfo) (codes.Project, error) {

	var (
		err error
	)

	if err = info.Verify(); err != nil {
		return nil, err
	}

	var (
		p = base{
			info:          info,
			configuration: location.Child(codes.NeonXConfigurationFileName),
		}
	)

	switch info.Architect {
	case codes.ProjectSingleProcess:
		return &SPProject{
			p,
		}, nil
	case codes.ProjectSingleGateway:
		return &SGProject{
			p,
		}, nil
	case codes.ProjectMultiSystem:
		return &MSProject{
			p,
		}, nil
	case codes.ProjectCommandLine:
		return &CmdProject{
			p,
		}, nil
	default:
		return nil, fmt.Errorf("Architect [%s] not support ", info.Architect)
	}

}

func makeGoMod(root location.Location, name string) error {

	cmd, err := exec.LoadGoCommand()

	if err != nil {
		return err
	}

	cmd.Command("mod", "init", name)
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	cmd.Dir = root

	err = cmd.Run(context.Background())

	if err != nil {
		return fmt.Errorf("Exec [%s] fail: %s ", cmd, err)
	}

	if cmd.ExitCode() != 0 {
		return fmt.Errorf("Exec [%s] fail, exitcode %d ", cmd, cmd.ExitCode())
	}

	return nil
}

func makeGoModTidy(root location.Location) error {

	cmd, err := exec.LoadGoCommand()

	if err != nil {
		return err
	}

	cmd.Command("mod", "tidy")
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	cmd.Dir = root

	fmt.Println("1-->", root)
	fmt.Println("2-->", cmd)

	err = cmd.Run(context.Background())

	fmt.Println("3-->", err, cmd.ExitCode())

	if err != nil {
		return fmt.Errorf("Exec [%s] fail: %s ", cmd, err)
	}

	if cmd.ExitCode() != 0 {
		return fmt.Errorf("Exec [%s] fail, exitcode %d ", cmd, cmd.ExitCode())
	}

	return nil
}
