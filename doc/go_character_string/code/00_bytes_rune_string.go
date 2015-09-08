package main

import "fmt"

func main() {
	bts := []byte("Hello")
	bts[0] = byte(100)
	for _, c := range bts {
		fmt.Println(string(c), c)
	}
	/*
	   d 100
	   e 101
	   l 108
	   l 108
	   o 111
	*/

	rs := []rune("Hello")
	rs[0] = rune(100)
	for _, c := range rs {
		fmt.Println(string(c), c)
	}
	/*
	   d 100
	   e 101
	   l 108
	   l 108
	   o 111
	*/

	str := "Hello"
	// str[0] = byte(100)
	// cannot assign to str[0]
	for _, c := range str {
		fmt.Println(string(c), c)
	}
	/*
	   H 72
	   e 101
	   l 108
	   l 108
	   o 111
	*/
}
