package main

import "fmt"

func main() {
	// launch goroutine in background
	go func() {
		fmt.Println("Hello, playground")
	}()
	//
	// Does not print anything
	//
	// when main returns
	// the program exits
	// and the goroutine will not be run
	// and gets garbage-collected
}
