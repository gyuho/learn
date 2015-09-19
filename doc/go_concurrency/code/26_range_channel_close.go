package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	cn := 0
	for v := range ch {
		fmt.Println(v)
		cn++
		if cn == 5 {
			close(ch)
		}
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
