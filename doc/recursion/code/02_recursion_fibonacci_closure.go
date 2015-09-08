package main

import "fmt"

func fib() func() int {
	v1, v2 := 0, 1
	return func() int {
		v1, v2 = v2, v1+v2
		return v1
	}
}

func main() {
	f := fib()
	for i := 0; i < 15; i++ {
		fmt.Printf("%d, ", f())
	}
	fmt.Println()
	// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610,
}
