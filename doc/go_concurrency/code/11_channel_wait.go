package main

import "fmt"

func main() {
	ch := make(chan struct{})
	go func() {
		fmt.Println("Hello, playground")
		ch <- struct{}{}
	}()

	// wait until we receive from channel ch
	<-ch

	// Hello, playground
}
