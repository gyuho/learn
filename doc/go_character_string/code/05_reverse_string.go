package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

/*
Recursively Reverse a String

Reverse("Hello")
(Reverse("ello")) + "H"
((Reverse("llo")) + "e") + "H"
(((Reverse("lo")) + "l") + "e") + "H"
((((Reverse("o")) + "l") + "l") + "e") + "H"
(((("o") + "l") + "l") + "e") + "H"
"olleH"
*/
func reverseStringResursion(str string) string {
	// Without this, the program generates slice bound error
	if len(str) <= 1 {
		return str
	}
	return reverseStringResursion(string(str[1:])) + string(str[0])
}

func main() {
	// we can't do this because string is immutable
	// prog.go:15: cannot assign to str[i]
	// for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
	// 	str[i], str[j] = str[j], str[i]
	// }
	reverseString := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	fmt.Println(reverseString("Hello"))          // olleH
	fmt.Println(reverseStringResursion("Hello")) // olleH

	bs := []byte("abc")
	fmt.Println(bs[1])
	for _, elem := range []byte("abc") {
		fmt.Println(elem, string(elem))
	}
	for _, elem := range []rune("abc") {
		fmt.Println(elem, string(elem))
	}
	// 97 a
	// 98 b
	// 99 c
	// 97 a
	// 98 b
	// 99 c
	fmt.Println()

	// byte, rune
	var by byte
	by = 10
	// by= -10
	// -10 overflows byte
	fmt.Println(reflect.TypeOf(by))
	// uint8
	// byte is alias for uint8
	// uint8 is the set of all unsigned 8-bit integers.
	// Range: 0 through 255.

	var br rune
	br = 10
	fmt.Println(reflect.TypeOf(br))
	// int32
	// rune is alias for int32
	// int32 is the set of all signed 32-bit integers.
	// Range: -2147483648 through 2147483647.

	fmt.Println(swapCase("Hello Hi"))
	// hELLO hI

	fmt.Println(swapCaseII("Hello Hi"))
	// hELLO hI

	// reverseIntSlice changes the order of slice
	// , without sorting.
	reverseIntSlice := func(s []int) []int {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	}
	fmt.Println(reverseIntSlice([]int{9, -13, 4, -2, 3, 1, -10, 21, 12}))
	// [12 21 -10 1 3 -2 4 -13 9]

	// reverseIntMore changes the order of multiple int arguments
	// , without sorting.
	reverseIntMore := func(more ...int) []int {
		for i, j := 0, len(more)-1; i < j; i, j = i+1, j-1 {
			more[i], more[j] = more[j], more[i]
		}
		return more
	}
	fmt.Println(reverseIntMore(9, -13, 4, -2, 3, 1, -10, 21, 12))
	// [12 21 -10 1 3 -2 4 -13 9]

}

func swapCase(str string) string {
	b := new(bytes.Buffer)

	// traverse character values, without index
	for _, elem := range str {
		if unicode.IsUpper(elem) {
			b.WriteRune(unicode.ToLower(elem))
		} else {
			b.WriteRune(unicode.ToUpper(elem))
		}
	}

	return b.String()
}

// rune is variable-length and can be made up of one or more bytes.
// rune literals are mapped to their unicode codepoint.
// For example, a rune literal 'a' is a number 97.
// 32 is the offset of the uppercase and lowercase characters.
// So if you add 32 to 'A', you get 'a' and vice versa.
func swapRune(r rune) rune {
	switch {
	case 'a' <= r && r <= 'z':
		return r - 'a' + 'A'
	case 'A' <= r && r <= 'Z':
		return r - 'A' + 'a'
	default:
		return r
	}
}

func swapCaseII(str string) string {
	return strings.Map(swapRune, str)
}
