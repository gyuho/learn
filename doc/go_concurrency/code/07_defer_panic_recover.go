package main

import (
	"fmt"
	"time"
)

func panicAndrecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("Panic!")
}

func main() {
	panicAndrecover()
	fmt.Println("Hello, World!")
	/*
	   Panic!
	   Hello, World!
	*/

	recursiveRecover()
	/*
	   Restarting after error: [ 0 ] Panic
	   Restarting after error: [ 1 ] Panic
	   Restarting after error: [ 2 ] Panic
	   Restarting after error: [ 3 ] Panic
	   Restarting after error: [ 4 ] Panic
	   Too much panic: 5
	*/
}

var count int

func recursiveRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Restarting after error:", err)
			time.Sleep(time.Second)
			count++
			if count == 5 {
				fmt.Printf("Too much panic: %d", count)
				return
			}
			recursiveRecover()
		}
	}()
	panic(fmt.Sprintf("[ %d ] Panic", count))
}
