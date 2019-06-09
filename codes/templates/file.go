package templates

import (
	"fmt"
	"io"

	"github.com/canzerapple/neonx-cli/codes"

	"github.com/canzerapple/neonx-cli/location/creator"
)

type TemplateFile struct {
	Name     string
	Project  codes.Project
	Extend   interface{}
	Template string
}

func (m *TemplateFile) GetName() string {
	return m.Name
}

func (m *TemplateFile) GetNodeType() creator.NodeType {
	return creator.NodeTypeFile
}

func (m *TemplateFile) SaveTo(writer io.Writer) error {

	exec := tpl.Lookup(m.Template)

	if exec == nil {
		return fmt.Errorf("template [%s] not found", m.Name)
	}

	return exec.Execute(writer, m)
}
