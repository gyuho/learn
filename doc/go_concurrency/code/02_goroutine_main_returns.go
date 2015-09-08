package main

import (
	"fmt"
	"time"
)

func a() {
	fmt.Println("a() called")
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("go func() called")
		// this is not called
		//
		// you can get this printed with channel
	}()
	go b()
}

func b() {
	time.Sleep(1 * time.Second)
	fmt.Println("b() called")
}

func main() {
	a()
	time.Sleep(5 * time.Second)
	// when main returns all others return as well
}

/*
a() called
b() called
*/
