package main

import "fmt"

func reverse(str string) string {
	if len(str) == 0 {
		return str
	}
	return reverse(str[1:]) + string(str[0])
}

func reverseInRune(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(reverse("Hello"))       // olleH
	fmt.Println(reverseInRune("Hello")) // olleH
}

/*
reverse("Hello")
(reverse("ello")) + "H"
((reverse("llo")) + "e") + "H"
(((reverse("lo")) + "l") + "e") + "H"
((((reverse("o")) + "l") + "l") + "e") + "H"
(((("o") + "l") + "l") + "e") + "H"
"olleH"
*/
