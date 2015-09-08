package main

import "fmt"

type Node struct {
	Key string
}

func noDestroy1(nd Node) {
	nd = Node{"destroyed"}
}

func noDestroy2(nd *Node) {
	nd = nil
}

func noDestroy3(nd *Node) {
	nd = &Node{"destroyed"}
}

func destroy1(nd *Node) {
	nd.Key = "destroyed"
}

func destroy2(nd *Node) {
	*nd = Node{}
}

func main() {
	nd := new(Node)
	nd.Key = "not destroyed"

	noDestroy1(*nd)
	fmt.Println("after noDestroy1:", nd)
	// after noDestroy1: &{not destroyed}

	noDestroy2(nd)
	fmt.Println("after noDestroy2:", nd)
	// after noDestroy2: &{not destroyed}

	// nd = nil    would destroy
	// but in the function, it only
	// updates the address

	noDestroy3(nd)
	fmt.Println("after noDestroy3:", nd)
	// after noDestroy3: &{not destroyed}

	destroy1(nd)
	fmt.Println("after destroy1:", nd)
	// after destroy1: &{destroyed}

	destroy2(nd)
	fmt.Println("after destroy2:", nd)
	// after destroy2: &{}
}
