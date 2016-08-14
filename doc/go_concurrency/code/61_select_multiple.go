package main

import "fmt"

func main() {
	c1, c2, c3 := make(chan struct{}), make(chan struct{}), make(chan struct{})
	close(c1)
	close(c2)
	close(c3)

	select { // select randomly
	case <-c1:
		fmt.Println("c1")
	case <-c2:
		fmt.Println("c2")
	case <-c3:
		fmt.Println("c3")
	}

	select { // select randomly
	case <-c3:
		fmt.Println("c3")
	case <-c2:
		fmt.Println("c2")
	case <-c1:
		fmt.Println("c1")
	}
}

/*
c1
c2
*/
