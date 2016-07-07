package main

import (
	"fmt"
	"time"
)

func runGoroutine() {
	go func() {
		time.Sleep(time.Hour)
	}()
}

func runDefer() {
	defer func() {
		time.Sleep(time.Hour)
	}()
}

func main() {
	fmt.Println("before runGoroutine #0")
	runGoroutine()
	fmt.Println("after runGoroutine #0") // return regardless of goroutine

	fmt.Println("before runDefer #1")
	runDefer()
	fmt.Println("after runDefer #1") // does not return until defer is done
}

/*
before runGoroutine #0
after runGoroutine #0
before runDefer #1
...
*/
