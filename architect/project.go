package architect

import (
	"fmt"
	"neonx-cli/protocol"
)

type ProjectArchitect string

const (
	ProjectSingleProcess ProjectArchitect = "SingleProcess"
	ProjectSingleGateway                  = "SingleGateway"
	ProjectMultiSystem                    = "SingleSystem"
	ProjectCommandLine                    = "CommandLine"
)

type Option interface {
	Verify() error
}

type ProjectInfo struct {
	Type      Type
	Architect ProjectArchitect
	Name      string
	Describe  string
}

func (m *ProjectInfo) Verify() error {

	if err := m.Type.Verify(); err != nil {
		return err
	}

	var (
		architects = map[ProjectArchitect]bool{
			ProjectSingleProcess: true,
			ProjectSingleGateway: true,
			ProjectMultiSystem:   true,
			ProjectCommandLine:   true,
		}
	)

	_, ok := architects[m.Architect]

	if !ok {
		return fmt.Errorf("ProjectInfo.Architect [%s] not found", m.Architect)
	}

	return nil
}

type Project interface {
	GetArchitect() ProjectArchitect
	GetName() string
	GetDescribe() string
}

type project struct {
	info          *ProjectInfo
	configuration Location
	pp protocol.PP
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

func NewProject(root string, info *ProjectInfo) (Project, error) {

	var (
		location Location
		err error
	)

	if err := info.Verify();err != nil {
		return nil,err
	}

	if location,err := ToLocation(root);err != nil {
		return nil,err
	}

	var (
		p = &project{
			info: info,
			configuration:
		}
	)

	switch info.Architect {
	case ProjectSingleProcess:
		p = &SPProject{
			project: to_project(),
		}
	case ProjectSingleGateway:
		p = &SGProject{
			project: info.toProject(),
		}
	case ProjectMultiSystem:
		p = &MSProject{
			project: info.toProject(),
		}
	case ProjectCommandLine:
		p = CmdProject{
			project: info.toProject(),
		}
	default:
		return nil, fmt.Errorf("Architect [%s] not support ", info.Architect)
	}

	return p, nil
}

func LoadProject(root string) (Project, error) {

	location, err := ToLocation(root)

	if err != nil {
		return nil, err
	}

}
