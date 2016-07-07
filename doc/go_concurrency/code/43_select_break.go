package main

import (
	"fmt"
	"time"
)

func main() {
	d1, d2 := time.Millisecond, time.Second
	for {
		select {
		case <-time.After(d1):
			d1 = time.Hour
			fmt.Println("d1 = time.Hour")
			continue // continue to the for-loop
		case <-time.After(d2):
			break // break and go to the lines below select
		}
		d2 = time.Nanosecond
		fmt.Println("d2 = time.Nanosecond")
		break // otherwise infinite for-loop
	}

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
