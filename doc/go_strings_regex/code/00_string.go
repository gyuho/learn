package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "Hello World!"
	fmt.Println(str[1]) // 101

	fmt.Println()

	// rune
	for _, c := range "Hello World!" {
		fmt.Println(c, string(c), reflect.TypeOf(c))
	}
	/*
	   72 H int32
	   101 e int32
	   108 l int32
	   108 l int32
	   111 o int32
	   32   int32
	   87 W int32
	   111 o int32
	   114 r int32
	   108 l int32
	   100 d int32
	   33 ! int32
	*/

	fmt.Println()

	// byte
	for _, c := range []byte("Hello World!") {
		fmt.Println(c, string(c), reflect.TypeOf(c))
	}
	/*
	   72 H uint8
	   101 e uint8
	   108 l uint8
	   108 l uint8
	   111 o uint8
	   32   uint8
	   87 W uint8
	   111 o uint8
	   114 r uint8
	   108 l uint8
	   100 d uint8
	   33 ! uint8
	*/
}
