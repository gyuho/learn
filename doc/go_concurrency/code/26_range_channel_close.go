package main

import "fmt"

func main() {
	func() {
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
	}()
	func() {
		ch := make(chan struct{})
		slice := []string{"A", "B", "C", "D", "E"}
		for _, v := range slice {
			go func(v string) {
				fmt.Println("Printing:", v)
				ch <- struct{}{}
			}(v)
		}
		cn := 0
		for range ch {
			cn++
			if cn == len(slice) {
				close(ch)
			}
		}
		/*
			Printing: E
			Printing: A
			Printing: B
			Printing: C
			Printing: D
		*/
	}()
}
