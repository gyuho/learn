package main

import "fmt"

func main() {
	str := "Hello, playground"
	for _, elem := range str {
		print(elem, "/")
	}
	// 72/101/108/108/111/44/32/112/108/97/121/103/114/111/117/110/100/

	fmt.Println()
	fmt.Println(str[0])         // 72
	fmt.Println(string(str[0])) // H
	// str[0] = byte(101)
	// cannot assign to str[0]

	bstr := []byte("Hello, playground")
	for _, elem := range bstr {
		print(elem, "/")
	}
	// 72/101/108/108/111/44/32/112/108/97/121/103/114/111/117/110/100/

	fmt.Println()
	fmt.Println(bstr[0])         // 72
	fmt.Println(string(bstr[0])) // H
	bstr[0] = byte(101)
	fmt.Println(string(bstr)) // eello, playground

	// rune represents Unicode code points
	// Go language defines the word rune as an alias for the type int32
	rstr := []rune("Hello, playground")
	for _, elem := range rstr {
		print(elem, "/")
	}
	// 72/101/108/108/111/44/32/112/108/97/121/103/114/111/117/110/100/

	fmt.Println()
	fmt.Println(rstr[0])         // 72
	fmt.Println(string(rstr[0])) // H
	rstr[0] = rune(101)
	fmt.Println(string(rstr)) // eello, playground
}
