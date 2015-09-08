package main

import "fmt"

func main() {
	ch := make(chan int, 0) // make channel with buffer 1
	go run(ch)
	fmt.Println(<-ch) // 1
}

func run(ch chan int) {
	ch <- 1
}
