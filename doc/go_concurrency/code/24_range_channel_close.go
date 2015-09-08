package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
	// 0
	// 1
	// 2
	// 3
	// 4

	v, ok := <-ch
	fmt.Println(v, ok) // 0 false
	// any value received from closed channel succeeds without blocking
	// , returning the zero value of channel type and false.
}
