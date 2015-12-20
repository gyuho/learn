package main

import "fmt"

func main() {
	{

		ch := make(chan int)
		limit := 5

		for i := 0; i < limit; i++ {
			go func(i int) {
				ch <- i
			}(i)
		}

		cn := 0
		for v := range ch {
			fmt.Println(v)
			cn++
			if cn == limit {
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

	{

		done, errChan := make(chan struct{}), make(chan error)

		limit := 5
		for i := 0; i < limit; i++ {
			go func(i int) {
				fmt.Println("Done at", i)
				done <- struct{}{}
			}(i)
		}

		cn := 0
		for cn != limit {
			select {
			case err := <-errChan:
				panic(err)
			case <-done:
				cn++
			}
		}

		close(done)
		close(errChan)

		/*
			Done at 4
			Done at 0
			Done at 1
			Done at 2
			Done at 3
		*/

	}
}
