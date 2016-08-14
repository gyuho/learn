package main

import "fmt"

func main() {
	var c chan struct{}
	select {
	case <-c:
		panic(1)
	default:
		fmt.Println(c == nil) // true
	}
}
