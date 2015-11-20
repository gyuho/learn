package main

import "fmt"

func main() {
	func() {
		ch := make(chan int, 0) // make channel with buffer 0
		go func() {
			ch <- 1
		}()
		v, ok := <-ch
		fmt.Println(v, ok) // 1 true
		close(ch)
		v2, ok2 := <-ch
		fmt.Println(v2, ok2) // 0 false
	}()

	func() {
		ch := make(chan int, 1)
		ch <- 1
		close(ch)
		v, ok := <-ch
		fmt.Println(v, ok) // 1 true
		v2, ok2 := <-ch
		fmt.Println(v2, ok2) // 0 false
	}()

	func() {
		ch := make(chan int, 1)
		close(ch)
		v, ok := <-ch
		fmt.Println(v, ok) // 0 false
		v2, ok2 := <-ch
		fmt.Println(v2, ok2) // 0 false
	}()
}
