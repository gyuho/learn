package main

import "fmt"

func main() {

	// %U prints Unicode format of a character
	// It prints out the Unicode code points

	str := "Hello"
	for _, c := range str {
		fmt.Printf("%U, %q %v\n", c, c, c)
	}

	fmt.Println()

	bts := []byte("Hello")
	for _, c := range bts {
		fmt.Printf("%U, %q %v\n", c, c, c)
	}
}

/*
U+0048, 'H' 72
U+0065, 'e' 101
U+006C, 'l' 108
U+006C, 'l' 108
U+006F, 'o' 111

U+0048, 'H' 72
U+0065, 'e' 101
U+006C, 'l' 108
U+006C, 'l' 108
U+006F, 'o' 111
*/
