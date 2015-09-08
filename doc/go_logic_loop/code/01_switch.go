package main

import (
	"fmt"
	"log"
	"reflect"
)

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func main() {
	fmt.Println(shouldEscape([]byte("?")[0]))     // true
	fmt.Println(shouldEscape([]byte("abcd#")[4])) // true
	fmt.Println(shouldEscape([]byte("abcd#")[0])) // false

	num := 2
	switch num {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		panic("what's the number?")
	}
	// 2

	st := "b"
	switch {
	case st == "a":
		fmt.Println("a")
	case st == "b":
		fmt.Println("b")
	case st == "c":
		fmt.Println("c")
	default:
		panic("what's the character?")
	}
	// b

	ts := []interface{}{true, 1, 1.5, "A"}
	for _, t := range ts {
		eval(t)
	}
	/*
	   bool: true is bool
	   int: 1 is int
	   float64: 1.5 is float64
	   string: A is string
	*/

	type temp struct {
		a string
	}
	eval(interface{}(temp{}))
	// 2009/11/10 23:00:00 {} is main.temp
}

func eval(t interface{}) {
	switch typedValue := t.(type) {
	default:
		log.Fatalf("%v is %v", typedValue, reflect.TypeOf(typedValue))
	case bool:
		fmt.Println("bool:", typedValue, "is", reflect.TypeOf(typedValue))
	case int:
		fmt.Println("int:", typedValue, "is", reflect.TypeOf(typedValue))
	case float64:
		fmt.Println("float64:", typedValue, "is", reflect.TypeOf(typedValue))
	case string:
		fmt.Println("string:", typedValue, "is", reflect.TypeOf(typedValue))
	}
}
