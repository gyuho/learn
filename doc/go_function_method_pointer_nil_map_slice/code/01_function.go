package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	// function values
	fmt.Println(
		// function closure (function literal)
		func() {
			fmt.Println("Hello")
		},
	) // 0x20280

	// without _ =
	// we get func literal evaluated but not used
	//
	_ = func() {
		fmt.Println("Hello 1")
	}
	// No output

	func(str string) {
		fmt.Println(str)
	}("Hello 2")
	// Hello 2

	fn := func() {
		fmt.Println("Hello")
	}
	fmt.Println(fn)                                                      // 0x203a0
	fmt.Println(reflect.TypeOf(fn))                                      // func()
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()) // main.func·004

	fn() // Hello

	fc := func(num int) int {
		num += 1
		return num
	}
	fmt.Println(fc)                 // 0x20720
	fmt.Println(reflect.TypeOf(fc)) // func(int) int
	fmt.Println(fc(1))              // 2

	po := plusOne(1)
	fmt.Println(po)                 // 0x209e0
	fmt.Println(reflect.TypeOf(po)) // main.funcType
	fmt.Println(po(3))              // 4
}

// function type
// FunctionType   = "func" Signature .
type funcType func(int) int

func plusOne(num int) funcType {
	return func(num int) int {
		return num + 1
	}
}
