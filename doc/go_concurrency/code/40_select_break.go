package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 5000)

	ch <- "a"
	ch <- "a"
	ch <- "a"
	ch <- "a"
	ch <- "a"

	done := make(chan struct{})
	go func() {
	here:
		for {
			select {
			case s, ok := <-ch:
				if !ok {
					fmt.Println("break 1")
					break // closed
				}
				fmt.Println(s, ok)
			case <-time.After(time.Second):
				fmt.Println("break 2")
				break here
			}
		}
		fmt.Println("break 3")
		done <- struct{}{}
	}()

	<-done
	fmt.Println("done")
}

/*
a true
a true
a true
a true
a true
break 2
break 3
done

*/
