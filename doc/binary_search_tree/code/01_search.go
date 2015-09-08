package main

import "fmt"

// Tree contains a Root node of a binary search tree.
type Tree struct {
	Root *Node
}

// New returns a new Tree with its root Node.
func New(root *Node) *Tree {
	tr := &Tree{}
	tr.Root = root
	return tr
}

// Interface represents a single object in the tree.
type Interface interface {
	// Less returns true when the receiver item(key)
	// is less than the given(than) argument.
	Less(than Interface) bool
}

// Node is a Node and a Tree itself.
type Node struct {
	// Left is a left child Node.
	Left *Node

	Key Interface

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	return nd
}

func (tr *Tree) String() string {
	return tr.Root.String()
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

// Insert inserts a Node to a Tree without replacement.
func (tr *Tree) Insert(nd *Node) {
	if tr.Root == nd {
		return
	}
	tr.Root = tr.Root.insert(nd)
}

func (nd *Node) insert(node *Node) *Node {
	if nd == nil {
		return node
	}
	if nd.Key.Less(node.Key) {
		nd.Right = nd.Right.insert(node)
	} else {
		nd.Left = nd.Left.insert(node)
	}
	return nd
}

// Min returns the minimum key Node in the tree.
func (tr Tree) Min() *Node {
	nd := tr.Root
	if nd == nil {
		return nil
	}
	for nd.Left != nil {
		nd = nd.Left
	}
	return nd
}

// Max returns the maximum key Node in the tree.
func (tr Tree) Max() *Node {
	nd := tr.Root
	if nd == nil {
		return nil
	}
	for nd.Right != nil {
		nd = nd.Right
	}
	return nd
}

// Search does binary-search on a given key and returns the first Node with the key.
func (tr Tree) Search(key Interface) *Node {
	nd := tr.Root
	// just updating the pointer value (address)
	for nd != nil {
		if nd.Key == nil {
			break
		}
		switch {
		case nd.Key.Less(key):
			nd = nd.Right
		case key.Less(nd.Key):
			nd = nd.Left
		default:
			return nd
		}
	}
	return nil
}

// SearchChan does binary-search on a given key and return the first Node with the key.
func (tr Tree) SearchChan(key Interface, ch chan *Node) {
	searchChan(tr.Root, key, ch)
	close(ch)
}

func searchChan(nd *Node, key Interface, ch chan *Node) {
	// leaf node
	if nd == nil {
		return
	}
	// when equal
	if !nd.Key.Less(key) && !key.Less(nd.Key) {
		ch <- nd
		return
	}
	searchChan(nd.Left, key, ch)  // left
	searchChan(nd.Right, key, ch) // right
}

// SearchParent does binary-search on a given key and returns the parent Node.
func (tr Tree) SearchParent(key Interface) *Node {
	nd := tr.Root
	parent := new(Node)
	parent = nil
	// just updating the pointer value (address)
	for nd != nil {
		if nd.Key == nil {
			break
		}
		switch {
		case nd.Key.Less(key):
			parent = nd // copy the pointer(address)
			nd = nd.Right
		case key.Less(nd.Key):
			parent = nd // copy the pointer(address)
			nd = nd.Left
		default:
			return parent
		}
	}
	return nil
}

type nodeStruct struct {
	ID    string
	Value int
}

func (n nodeStruct) Less(b Interface) bool {
	return n.Value < b.(nodeStruct).Value
}

func main() {
	root := NewNode(nodeStruct{"A", 5})
	tr := New(root)
	tr.Insert(NewNode(nodeStruct{"B", 3}))
	tr.Insert(NewNode(nodeStruct{"C", 17}))
	fmt.Printf("%s\n", tr)
	// [[{B 3}] {A 5} [{C 17}]]

	fmt.Println(tr.Search(nodeStruct{"A", 5}))   // [[{B 3}] {A 5} [{C 17}]]
	fmt.Println(tr.Search(nodeStruct{Value: 3})) // [{B 3}]
	fmt.Println(tr.Search(nodeStruct{"C", 17}))  // [{C 17}]

	ch := make(chan *Node)
	go tr.SearchChan(nodeStruct{"A", 5}, ch)
	fmt.Println(<-ch) // [[{B 3}] {A 5} [{C 17}]]

	fmt.Println(tr.Max()) // [{C 17}]
	fmt.Println(tr.Min()) // [{B 3}]
}
