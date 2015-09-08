package main

import "fmt"

// fib returns nth fibonacci number.
func fib(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d, ", fib(uint(i)))
	}
	fmt.Println()
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377,
}
