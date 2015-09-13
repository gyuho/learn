package implicit

type Node interface {
	GetName() string
}

type node struct {
	name string
}

func (n node) GetName() string {
	return n.name
}

func StartNode(name string) Node {
	nd := node{}
	nd.name = name
	return &nd
}
