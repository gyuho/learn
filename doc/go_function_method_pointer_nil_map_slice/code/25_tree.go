package main

import "fmt"

type Data struct {
	Root *Node
}

type Node struct {
	Left  *Node
	Key   int
	Right *Node
}

func (tr *Tree) String() string {
	return d.Root.String()
}

func (nd *Node) String() string {
	if nd == nil {
		return "[]"
	}
	s := ""
	if nd.Left != nil {
		s += nd.Left.String() + " "
	}
	s += fmt.Sprintf("%v", nd.Key)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

func main() {
	tree := new(Tree)
	root := new(Node)
	root.Key = 3
	tree.Root = root

	left := new(Node)
	left.Key = 1
	root.Left = left

	right := new(Node)
	right.Key = 5
	root.Right = right

	seven := new(Node)
	seven.Key = 7
	root.Right.Right = seven
	/*
	     3
	    / \
	   1   5
	        \
	         7
	*/

	fmt.Println("tree:", tree) // [[1] 3 [5 [7]]]

	delete1(root)
	fmt.Println("delete1(root) root:", root) // delete1(root) root: [[1] 3 [5 [7]]]
	fmt.Println("delete1(root) tree:", tree) // delete1(root) tree: [[1] 3 [5 [7]]]

	// but if we access directly
	// you can update
	root.Right = new(Node)
	fmt.Println("root.Right = new(Node) tree:", tree)
	// root.Right = new(Node) tree: [[1] 3 [0]]

	delete2(root)
	fmt.Println("delete2(root) root:", root) // delete2(root) root: [0]
	fmt.Println("delete2(root) tree:", tree) // delete2(root) tree: [0]
}

func delete1(nd *Node) {
	// even if nd is a pointer
	// it ONLY copies the value
	// of the tree address
	nd = nil
	// or
	// nd = new(Node)

	// setting nil to the address
	// does not affect the original object
}

func delete2(nd *Node) {
	*nd = Node{}
}
