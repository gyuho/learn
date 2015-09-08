package main

import "fmt"

func factorialWithIteration(num int) int {
	result := 1
	if num == 0 {
		return result
	}
	for i := 2; i <= num; i++ {
		result *= i
	}
	return result
}

func factorial(num int) int {
	if num <= 1 {
		fmt.Println("returning: 1")
		return 1
	}
	fmt.Println("returning:", num, num-1)
	return num * factorial(num-1)
}

func main() {
	fmt.Println(factorialWithIteration(5)) // 120
	fmt.Println(factorial(5))              // 120
}

/*
returning: 5 4
returning: 4 3
returning: 3 2
returning: 2 1
returning: 1
*/
