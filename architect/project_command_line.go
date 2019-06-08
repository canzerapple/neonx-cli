package architect

import "gopkg.in/yaml.v3"

type CmdProjectInfo struct {
	ProjectInfo
}

func (m *CmdProjectInfo) SaveTo(root string) error {

}

/*
Project Architect

*/
type CmdProject struct {
	*project
}
