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
		slice := []string{"A", "B", "C", "D", "E"}
		if len(slice) > 0 {
			// this only works when slice length
			// is greater than 0. Otherwise, it will
			// be deadlocking receiving no done message.
			done := make(chan struct{})
			for _, v := range slice {
				go func(v string) {
					fmt.Println("Printing:", v)
					done <- struct{}{}
				}(v)
			}
			cn := 0
			for range done {
				cn++
				if cn == len(slice) {
					close(done)
				}
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
