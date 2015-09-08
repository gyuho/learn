package main

import (
	"fmt"
	"reflect"
)

func plusOne(num int) int {
	return num + 1
}

func plusTwo(num int) int {
	return num + 2
}

func plusThree(num int) int {
	return num + 3
}

var funcMap = map[string]func(num int) int{
	"one":   plusOne,
	"two":   plusTwo,
	"three": plusThree,
}

func add(num int, functions ...func(num int) int) int {
	for _, oneFunc := range functions {
		num = oneFunc(num)
	}
	return num
}

func main() {
	chosen := funcMap["two"]
	fmt.Println(reflect.TypeOf(chosen)) // func(int) int
	fmt.Println(chosen(0))              // 2

	fmt.Println(
		add(
			0,
			funcMap["one"],
			funcMap["two"],
			funcMap["three"],
		),
	)
	// 6
}
