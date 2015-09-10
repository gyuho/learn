package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Root *Node
}

func New(root *Node) *Data {
	d := &Data{}
	d.Root = root
	return d
}

type Interface interface {
	Less(than Interface) bool
}

type Node struct {
	Left  *Node
	Key   Interface
	Right *Node
}

func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	return nd
}

func (d *Data) Insert(nd *Node) {
	if d.Root == nd {
		return
	}
	d.Root = d.Root.insert(nd)
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

func (d Data) Search(key Interface) *Node {
	nd := d.Root
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

type Int int

func (n Int) Less(b Interface) bool {
	return n < b.(Int)
}

func main() {
	// nil for error(interface)
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	root := NewNode(Int(5))
	data := New(root)
	data.Insert(NewNode(Int(7)))

	// nil for interface
	fmt.Println(NewNode(Int(7)).Key == nil) // false

	// nil for struct pointer
	root = nil
	fmt.Println(root == nil) // true

	// nil for bytes (not for string)
	b := []byte("abc")
	fmt.Println(b, string(b)) // [97 98 99] abc
	b = nil
	fmt.Println(b, string(b)) // []
	//
	// str := "abc"
	// str = nil (x) value cannot be nil

	// nil for map
	mmap := make(map[string]bool)
	mmap["A"] = true
	fmt.Println(mmap) // map[A:true]
	mmap = nil
	fmt.Println(mmap) // map[]

	// nil for slice (not for array)
	slice := []int{1}
	fmt.Println(slice) // [1]
	slice = nil
	fmt.Println(slice) // []
	//
	// array := [1]int{1}
	// array = nil
	// (X) cannot use nil as type [1]int in assignment
}
