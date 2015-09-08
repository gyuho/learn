package main

import (
	"fmt"
	"time"
)

func b() {
	fmt.Println("b is still running")
	fmt.Println("because although a exited but main hasn't exited yet!")
}

func a() {
	fmt.Println("a exits")
	go b()
}

func main() {
	a()
	time.Sleep(time.Second)
	// a exits
	// b is still running
	// because although a exited but main hasn't exited yet!

	go func() {
		fmt.Println("Hello, playground")
	}()
	time.Sleep(time.Second)
	// Hello, playground
}
