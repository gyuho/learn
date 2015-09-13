package main

import (
	"fmt"

	"github.com/gyuho/learn/doc/go_interface/code/implicit"
)

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

func main() {
	nd := StartNode("Gyu-Ho")
	fmt.Println(nd.GetName())
	// Gyu-Ho

	ndi := implicit.StartNode("A")
	fmt.Println(ndi.GetName())
	// A
}
