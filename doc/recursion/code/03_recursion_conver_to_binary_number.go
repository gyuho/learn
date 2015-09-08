package main

import "fmt"

func toBinaryNumber(num uint64) uint64 {
	fmt.Println("calling on", num)
	if num == 0 {
		return 0
	}
	return (num % 2) + 10*toBinaryNumber(num/2)
}

func main() {
	fmt.Println(toBinaryNumber(15))
	/*
	   calling on 15
	   calling on 7
	   calling on 3
	   calling on 1
	   calling on 0
	   1111
	*/
}
