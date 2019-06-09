package project

import (
	"fmt"

	"github.com/canzerapple/neonx-cli/codes"

	. "github.com/canzerapple/neonx-cli/codes/templates"
	"github.com/canzerapple/neonx-cli/location"
	. "github.com/canzerapple/neonx-cli/location/creator"
)

const (
	CmdProjectCommands = "cmd"
)

/*
+ root
 - cmd
   - cmd.go
 - main.go
*/
type CmdProject struct {
	base
}

func (m *CmdProject) ClassName() string {
	return "CmdProject"
}

func CreateCmdProject(root location.Location, name, describe string) (*CmdProject, error) {

	var (
		nodes = Nodes{
			Dir(CmdProjectCommands),
			CmdMain(),
			&codes.ProjectInfo{
				Type:      codes.TypeProject,
				Architect: codes.ProjectCommandLine,
				Name:      name,
				Describe:  describe,
			},
		}
	)

	exists, err := root.IsExists()

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("create CmdProject fail, directory [%s] alredy exists ", root.Child(name))
	}

	if err := CreateNodes(root, nodes, false); err != nil {
		return nil, err
	}

	err = makeGoMod(root, name)

	if err != nil {
		return nil, err
	}

	err = makeGoModTidy(root)

	if err != nil {
		return nil, err
	}

	return LoadCmdProject(root)

}

func LoadCmdProject(root location.Location) (*CmdProject, error) {
	return nil, nil
}
