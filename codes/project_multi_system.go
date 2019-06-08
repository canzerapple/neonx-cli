package codes

import (
	"fmt"

	"github.com/canzerapple/neonx-cli/location"
	. "github.com/canzerapple/neonx-cli/location/creator"
)

/*
+ root
  - protocol    => Git Repo
    - <name>_system
		-<service_name>
  - middleware  => Git Repo
  - <name>_system => Git Repo
	 - bff
		- <bff_name>
		- <bff_name>
	 - service
        - <service_name>
	    - <service_name>
  - ...
  - <name>_system => Git Repo
*/

const (
	MSProjectProtocol   = "protocol"
	MSProjectMiddleware = "middleware"
)

type MSProject struct {
	project
}

func CreateMSProject(root location.Location, name, describe string) (*MSProject, error) {

	var (
		nodes = &Directory{
			Name: name,
			Child: Nodes{
				Dir(MSProjectProtocol),
				Dir(MSProjectMiddleware),
				&ProjectInfo{
					Type:      TypeProject,
					Architect: ProjectMultiSystem,
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
		return nil, fmt.Errorf("create MSProject fail, directory [%s] alredy exists ", root.Child(name))
	}

	if err := CreateNodes(root, nodes, false); err != nil {
		return nil, err
	}

	return LoadMSProject(root)
}

func LoadMSProject(root location.Location) (*MSProject, error) {
	return nil, nil
}
