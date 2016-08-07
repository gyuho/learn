package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	fmt.Println(len(ch), cap(ch)) // 0 1
	ch <- 1
	fmt.Println(len(ch), cap(ch)) // 1 1
	if len(ch) == cap(ch) {
		fmt.Println("channel is full")
	} // channel is full
	fmt.Println(<-ch)
	fmt.Println(len(ch), cap(ch)) // 0 1
}
