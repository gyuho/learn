package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}

	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
}
