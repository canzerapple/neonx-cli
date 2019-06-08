package creator

type Directory struct {
	Name  string
	Child Nodes
}

func (m *Directory) GetChild() Nodes {
	return m.Child
}

func (m *Directory) GetName() string {
	return m.Name
}

func (m *Directory) GetNodeType() NodeType {
	return NodeTypeDirectory
}

type Dir string

func (m Dir) GetChild() Nodes {
	return nil
}

func (m Dir) GetName() string {
	return string(m)
}

func (m Dir) GetNodeType() NodeType {
	return NodeTypeDirectory
}
