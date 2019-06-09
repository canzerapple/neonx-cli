package project

import (
	"fmt"

	"github.com/canzerapple/neonx-cli/codes"

	"github.com/canzerapple/neonx-cli/location"
	. "github.com/canzerapple/neonx-cli/location/creator"
)

/*
Signal Gateway Project
+ root
  - gateways
  - services
  - protocol
  - middleware
  - .neonx.yaml
*/

const (
	SGProjectGateways   = "gateways"
	SGProjectServices   = "services"
	SGProjectProtocol   = "protocol"
	SGProjectMiddleware = "middleware"
)

type SGProject struct {
	base
}

func (m *SGProject) ClassName() string {
	return "SGProject"
}

func CreateSGProject(root location.Location, name, describe string) (*SGProject, error) {

	var (
		nodes = &Directory{
			Name: name,
			Child: Nodes{
				Dir(SGProjectGateways),
				Dir(SGProjectServices),
				Dir(SGProjectProtocol),
				Dir(SGProjectMiddleware),
				&codes.ProjectInfo{
					Type:      codes.TypeProject,
					Architect: codes.ProjectSingleGateway,
					Name:      name,
					Describe:  describe,
				},
			},
		}
	)

	contains, err := root.Contains(name)

	if err != nil {
		return nil, err
	}

	if contains {
		return nil, fmt.Errorf("create SGProject fail, directory [%s] alredy exists ", root.Child(name))
	}

	if err := CreateNodes(root, nodes, false); err != nil {
		return nil, err
	}

	return LoadSGProject(root)
}

func LoadSGProject(root location.Location) (*SGProject, error) {
	return nil, nil
}
