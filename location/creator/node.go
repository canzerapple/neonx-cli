package creator

import (
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/canzerapple/neonx-cli/location"
)

type NodeType int

const (
	NodeTypeDirectory = 0
	NodeTypeFile      = 1
	NodeTypeFiles     = 2
)

type Node interface {
	GetNodeType() NodeType
}

type FileNode interface {
	Node
	GetName() string
	SaveTo(writer io.Writer) error
}

type DirectoryNode interface {
	Node
	GetName() string
	GetChild() Nodes
}

type FilesNode interface {
	Node
	GetNodes() Nodes
}

type Nodes []Node

func (m Nodes) GetNodeType() NodeType {
	return NodeTypeFiles
}

func (m Nodes) GetNodes() Nodes {
	return m
}

func CreateNodes(location location.Location, root Node, overwrite bool) error {

	switch node := root.(type) {
	case DirectoryNode:

		if err := location.CreateDirectory(node.GetName()); err != nil {
			return err
		}

		var (
			nodes = node.GetChild()
			now   = location.Child(node.GetName())
		)

		if nodes == nil {
			return nil
		}

		for _, child := range nodes {
			if err := CreateNodes(now, child, overwrite); err != nil {
				return err
			}
		}

		return nil
	case FilesNode:
		for _, child := range node.GetNodes() {

			err := CreateNodes(location, child, overwrite)

			if err != nil {
				return err
			}
		}

		return nil
	case FileNode:

		var (
			save   = location.Child(node.GetName())
			file   *os.File
			err    error
			exists bool
		)

		if exists, err = save.IsExists(); err != nil {
			return err
		}

		if exists && !overwrite {
			return nil
		}

		if file, err = save.Open(os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755); err != nil {
			return err
		}

		defer file.Close()

		if err = node.SaveTo(file); err != nil {
			return err
		}

		return nil

	default:
		return fmt.Errorf("root type %d not such file or directory ", reflect.TypeOf(root))
	}
}
