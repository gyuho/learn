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

type Float float64

func (a Float) Less(b Interface) bool {
	return a < b.(Float)
}

func main() {
	root := NewNode(Float(1))
	tr := New(root)

	slice := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range slice {
		tr.Insert(NewNode(Float(num)))
	}
	fmt.Printf("%s\n", tr)
	// [1 [[2 [2.5]] 3 [9 [13 [[[15] 16] 17 [20 [25 [39]]]]]]]]

	fmt.Println(tr.Search(Float(20)))
	// [20 [25 [39]]]

	fmt.Println(tr.Max().Key)
	// 39

	fmt.Println(tr.Min().Key)
	// 1

	fmt.Println(tr.SearchParent(Float(16)).Key)
	// 17

	fmt.Println(tr.SearchParent(Float(16)).Key)
	// 17

	fmt.Println()
	deletes := []float64{13, 17, 3, 15, 1, 2.5}
	for _, num := range deletes {
		fmt.Println("Deleting", num)
		tr.Delete(Float(num))
		fmt.Println("Deleted", num, ":", tr)
		fmt.Println()
	}
	/*
	   Deleting 13
	   Deleted 13 : [1 [[2 [2.5]] 3 [9 [[[15] 16] 17 [20 [25 [39]]]]]]]

	   Deleting 17
	   Deleted 17 : [1 [[2 [2.5]] 3 [9 [[15] 16 [20 [25 [39]]]]]]]

	   Deleting 3
	   Deleted 3 : [1 [[2] 2.5 [9 [[15] 16 [20 [25 [39]]]]]]]

	   Deleting 15
	   Deleted 15 : [1 [[2] 2.5 [9 [16 [20 [25 [39]]]]]]]

	   Deleting 1
	   Deleted 1 : [[2] 2.5 [9 [16 [20 [25 [39]]]]]]

	   Deleting 2.5
	   Deleted 2.5 : [2 [9 [16 [20 [25 [39]]]]]]
	*/
}

// Delete deletes a Node from a tree.
// It returns nil if it does not exist in the tree.
func (tr *Tree) Delete(key Interface) Interface {
	if key == nil {
		return nil
	}
	nd := tr.Search(key)
	if nd == nil {
		return nil
	}
	parent := tr.SearchParent(key)

	// you need to dereference the pointer
	// and update with a value
	// in order to change the original struct

	if nd.Left != nil && nd.Right != nil {
		// if two children

		// #1. Find the node to substitute
		// the to-be-deleted node
		//
		// either get the biggest of left sub-tree
		tmp := new(Tree)
		tmp.Root = nd.Left
		tmpNode := tmp.Max()
		//
		// OR
		//
		// get the smallest of right sub-tree
		// tmp := new(Data)
		// tmp.Root = nd.Right
		// tmpNode := nd.Right.Min()
		//
		replacingNode := tr.Search(tmpNode.Key)
		parentOfReplacingNode := tr.SearchParent(replacingNode.Key)

		// order matters!
		if replacingNode.Key.Less(nd.Key) {
			// replacing with the left child
			replacingNode.Right = nd.Right

			// inherit the sub-tree
			if nd.Left.Key.Less(replacingNode.Key) ||
				replacingNode.Key.Less(nd.Left.Key) {
				// if different
				replacingNode.Left = nd.Left

				// destroy the old pointer in sub-tree
				if parentOfReplacingNode.Key.Less(replacingNode.Key) {
					// deleting right child of parentOfReplacingNode
					parentOfReplacingNode.Right = nil
				} else {
					// deleting left child of parentOfReplacingNode
					parentOfReplacingNode.Left = nil
				}
			}

		} else {
			// replacing with the right child
			replacingNode.Left = nd.Left

			// inherit the sub-tree
			if nd.Right.Key.Less(replacingNode.Key) ||
				replacingNode.Key.Less(nd.Right.Key) {

				// destroy the old pointer in sub-tree
				if parentOfReplacingNode.Key.Less(replacingNode.Key) {
					// deleting right child of parentOfReplacingNode
					parentOfReplacingNode.Right = nil
				} else {
					// deleting left child of parentOfReplacingNode
					parentOfReplacingNode.Left = nil
				}
			}
		}

		// #2. Update the parent, child node
		if parent == nil {
			// in case of deleting the root Node
			tr.Root = replacingNode
		} else {
			if parent.Key.Less(nd.Key) {
				// deleting right child of parent
				parent.Right = replacingNode
			} else {
				// deleting left child of parent
				parent.Left = replacingNode
			}
		}

	} else if nd.Left != nil && nd.Right == nil {
		// only left child
		// #1. Update the parent node
		if parent == nil {
			// in case of deleting the root Node
			tr.Root = nd.Left
		} else {
			if parent.Key.Less(nd.Key) {
				// right child of parent
				parent.Right = nd.Left
			} else {
				// left child of parent
				parent.Left = nd.Left
			}
		}

	} else if nd.Left == nil && nd.Right != nil {
		// only right child
		// #1. Update the parent node
		if parent == nil {
			// in case of deleting the root Node
			tr.Root = nd.Right
		} else {
			if parent.Key.Less(nd.Key) {
				// right child of parent
				parent.Right = nd.Right
			} else {
				// left child of parent
				parent.Left = nd.Right
			}
		}
	} else {
		// no child
		if parent == nil {
			// in case of deleting the root Node
			tr.Root = nil
		} else {
			if parent.Key.Less(nd.Key) {
				// right child of parent
				parent.Right = nil
			} else {
				// left child of parent
				parent.Left = nil
			}
		}
	}

	k := nd.Key

	// At the end, delete the node
	// this is not necessary because it will be
	// garbage collected
	*nd = Node{}

	// because this is inside function
	// this won't change the actual node
	//
	// nd = new(Node)
	// nd = nil

	return k
}
