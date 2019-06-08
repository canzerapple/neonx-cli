package architect

import (
	"path/filepath"
)

type Node interface {
	GetParent() Node
	GetName() string
}

type fileNode struct {
	parent Node
	root string
}

func (m *fileNode) GetParent() Node {
	return m.parent
}


func (m *fileNode) GetName() string {
	return filepath.Base(m.root)
}


func newFileNode(parent Node, root string) *fileNode {
	return &fileNode{
		parent:parent,
		root:root,
	}
}