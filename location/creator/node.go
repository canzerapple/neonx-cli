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
)

type Node interface {
	GetName() string
	GetNodeType() NodeType
}

type FileNode interface {
	Node
	SaveTo(writer io.Writer) error
}

type DirectoryNode interface {
	Node
	GetChild() Nodes
}

type Nodes []Node

func CreateNodes(location location.Location, root Node, overwrite bool) error {

	switch node := root.(type) {
	case DirectoryNode:

		if err := location.CreateDirectory(root.GetName()); err != nil {
			return err
		}

		var (
			nodes = node.GetChild()
			now   = location.Child(root.GetName())
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
