package main

import "time"

func main() {
	// goroutine #01 : Queue
	go println(1)

	// goroutine #02
	// Anonymous Function Closure
	// Not function literal
	// So we need parenthesis at the end
	go func() {
		println(2)
	}()

	// goroutine #03
	// Anonymous Function Closure with input
	go func(n int) {
		println(n)
	}(3)

	// 1
	// 2
	// 3

	time.Sleep(time.Nanosecond)
	// main goroutine does not wait(block) for goroutine's return
	// Without this, we just reach the end of main and goroutine does not run
}
