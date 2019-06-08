package codes

import (
	"fmt"

	"github.com/canzerapple/neonx-cli/location"
)

type ProjectArchitect string

const (
	ProjectSingleProcess ProjectArchitect = "SingleProcess"
	ProjectSingleGateway                  = "SingleGateway"
	ProjectMultiSystem                    = "MultiSystem"
	ProjectCommandLine                    = "CommandLine"
)

type Project interface {
	GetArchitect() ProjectArchitect
	GetName() string
	GetDescribe() string
}

type project struct {
	info          *ProjectInfo
	configuration location.Location
}

func (m *project) GetName() string {
	return m.info.Name
}

func (m *project) GetDescribe() string {
	return m.info.Describe
}

func (m *project) GetArchitect() ProjectArchitect {
	return m.info.Architect
}

func NewProject(location location.Location, info *ProjectInfo) (Project, error) {

	var (
		err error
	)

	if err = info.Verify(); err != nil {
		return nil, err
	}

	var (
		p = project{
			info:          info,
			configuration: location.Child(NeonXConfigurationFileName),
		}
	)

	switch info.Architect {
	case ProjectSingleProcess:
		return &SPProject{
			p,
		}, nil
	case ProjectSingleGateway:
		return &SGProject{
			p,
		}, nil
	case ProjectMultiSystem:
		return &MSProject{
			p,
		}, nil
	case ProjectCommandLine:
		return &CmdProject{
			p,
		}, nil
	default:
		return nil, fmt.Errorf("Architect [%s] not support ", info.Architect)
	}

}
