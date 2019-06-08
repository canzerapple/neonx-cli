package codes

import (
	"fmt"
	"io"

	"github.com/canzerapple/neonx-cli/location"
	"github.com/canzerapple/neonx-cli/location/creator"
	yaml "gopkg.in/yaml.v3"
)

type ProjectInfo struct {
	Type      Type
	Architect ProjectArchitect
	Name      string
	Describe  string
}

func (m *ProjectInfo) GetName() string {
	return NeonXConfigurationFileName
}

func (m *ProjectInfo) GetNodeType() creator.NodeType {
	return creator.NodeTypeFile
}

func (m *ProjectInfo) SaveTo(writer io.Writer) error {

	encoder := yaml.NewEncoder(writer)

	return encoder.Encode(m)
}

func LoadProjectInfo(root location.Location) (*ProjectInfo, error) {

	var (
		info      = new(ProjectInfo)
		err       error
		configure = root.Child(NeonXConfigurationFileName)
	)

	if err = configure.ReadYAML(info); err != nil {
		return nil, err
	}

	return info, nil

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
