package architect

import "fmt"

type Type string

const (
	TypeProject Type = "Project"
	TypeSystem       = "System"
)

func (m Type) Verify() error {

	switch m {
	case TypeProject, TypeSystem:
		return nil
	default:
		return fmt.Errorf("type [%s] not support ", m)
	}
}
