package architect

import "path/filepath"

func Load(root string) (Node,error) {

	root,err := filepath.Abs(root)

	if err != nil {
		return nil,err
	}

	node := newFileNode(nil,root)

}
