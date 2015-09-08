package main

import (
	"bytes"
	"fmt"
)

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

type Int int

// Less returns true if int(a) < int(b).
func (a Int) Less(b Interface) bool {
	return a < b.(Int)
}

func main() {
	tr := New(NewNode(Int(5)))
	tr.Insert(NewNode(Int(3)))
	tr.Insert(NewNode(Int(17)))
	tr.Insert(NewNode(Int(7)))
	tr.Insert(NewNode(Int(1)))
	/*
	       5
	      / \
	     3   17
	    /    /
	   1    7
	*/

	fmt.Println("Min:", tr.Min().Key)             // Min: 1
	fmt.Println("Max:", tr.Max().Key)             // Max: 17
	fmt.Println("Search:", tr.Search(Int(7)).Key) // Search: 7

	buf1 := new(bytes.Buffer)
	ch1 := make(chan string)
	go tr.PreOrder(ch1) // root, left, right
	for {
		v, ok := <-ch1
		if !ok {
			break
		}
		buf1.WriteString(v)
		buf1.WriteString(" ")
	}
	fmt.Println("PreOrder:", buf1.String()) // PreOrder: 5 3 1 17 7

	buf2 := new(bytes.Buffer)
	ch2 := make(chan string)
	go tr.InOrder(ch2) // left, root, right
	for {
		v, ok := <-ch2
		if !ok {
			break
		}
		buf2.WriteString(v)
		buf2.WriteString(" ")
	}
	fmt.Println("InOrder:", buf2.String()) // 1 3 7 17 5

	buf3 := new(bytes.Buffer)
	ch3 := make(chan string)
	go tr.PostOrder(ch3) // left, right, root
	for {
		v, ok := <-ch3
		if !ok {
			break
		}
		buf3.WriteString(v)
		buf3.WriteString(" ")
	}
	fmt.Println("PostOrder:", buf3.String()) // 1 3 7 17 5

	buf4 := new(bytes.Buffer)
	nodes4 := tr.LevelOrder() // from top-level
	for _, v := range nodes4 {
		buf4.WriteString(fmt.Sprintf("%v ", v.Key))
	}
	fmt.Println("LevelOrder:", buf4.String()) // 5 3 17 1 7

	tr2 := New(NewNode(Int(5)))
	tr2.Insert(NewNode(Int(3)))
	tr2.Insert(NewNode(Int(17)))
	tr2.Insert(NewNode(Int(7)))
	tr2.Insert(NewNode(Int(1)))

	fmt.Println("ComparePreOrder:", ComparePreOrder(tr, tr2))   // true
	fmt.Println("CompareInOrder:", CompareInOrder(tr, tr2))     // true
	fmt.Println("ComparePostOrder:", ComparePostOrder(tr, tr2)) // true
}

// PreOrder traverses from Root, Left-SubTree, and Right-SubTree. (DFS)
func (tr *Tree) PreOrder(ch chan string) {
	preOrder(tr.Root, ch)
	close(ch)
}

func preOrder(nd *Node, ch chan string) {
	// leaf node
	if nd == nil {
		return
	}
	ch <- fmt.Sprintf("%v", nd.Key) // root
	preOrder(nd.Left, ch)           // left
	preOrder(nd.Right, ch)          // right
}

// ComparePreOrder returns true if two Trees are same with PreOrder.
func ComparePreOrder(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go t1.PreOrder(ch1)
	go t2.PreOrder(ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

// InOrder traverses from Left-SubTree, Root, and Right-SubTree. (DFS)
func (tr *Tree) InOrder(ch chan string) {
	inOrder(tr.Root, ch)
	close(ch)
}

func inOrder(nd *Node, ch chan string) {
	// leaf node
	if nd == nil {
		return
	}
	inOrder(nd.Left, ch)            // left
	ch <- fmt.Sprintf("%v", nd.Key) // root
	inOrder(nd.Right, ch)           // right
}

// CompareInOrder returns true if two Trees are same with InOrder.
func CompareInOrder(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go t1.InOrder(ch1)
	go t2.InOrder(ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

// PostOrder traverses from Left-SubTree, Right-SubTree, and Root.
func (tr *Tree) PostOrder(ch chan string) {
	postOrder(tr.Root, ch)
	close(ch)
}

func postOrder(nd *Node, ch chan string) {
	// leaf node
	if nd == nil {
		return
	}
	postOrder(nd.Left, ch)          // left
	postOrder(nd.Right, ch)         // right
	ch <- fmt.Sprintf("%v", nd.Key) // root
}

// ComparePostOrder returns true if two Trees are same with PostOrder.
func ComparePostOrder(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go t1.PostOrder(ch1)
	go t2.PostOrder(ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

// LevelOrder traverses the Tree with Breadth First Search.
// (http://en.wikipedia.org/wiki/Tree_traversal#Breadth-first_2)
//
//	levelorder(root)
//	  q = empty queue
//	  q.enqueue(root)
//	  while not q.empty do
//	    node := q.dequeue()
//	    visit(node)
//	    if node.left ≠ null then
//	      q.enqueue(node.left)
//	    if node.right ≠ null then
//	      q.enqueue(node.right)
//
func (tr *Tree) LevelOrder() []*Node {
	visited := []*Node{}
	queue := []*Node{tr.Root}
	for len(queue) != 0 {
		nd := queue[0]
		queue = queue[1:len(queue):len(queue)]
		visited = append(visited, nd)
		if nd.Left != nil {
			queue = append(queue, nd.Left)
		}
		if nd.Right != nil {
			queue = append(queue, nd.Right)
		}
	}
	return visited
}
