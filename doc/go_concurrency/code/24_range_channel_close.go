package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}

	i := 0
	for v := range ch {
		fmt.Println(v)
		i++
		if i == 5 {
			close(ch)
		}
	}
}

/*
5
5
5
5
5
*/
