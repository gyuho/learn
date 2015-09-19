package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
