package main

import "fmt"

func main() {
	fmt.Println(3 < 10) // true
	fmt.Println(3 > 10) // false

	fmt.Println(5 <= 5) // true
	fmt.Println(5 >= 5) // true

	fmt.Println(5 == 5) // true
	fmt.Println(3 == 1) // false

	fmt.Println("A" < "a") // true
	fmt.Println("A" < "b") // true
	fmt.Println("a" < "b") // true
	fmt.Println("a" > "x") // false

	// fmt.Println(2 =< 5)
	// syntax error

	// fmt.Println(2 => 5)
	// syntax error

	// Note that the a is only valid
	// inside this if-statement
	// a is not existent after this if-statement
	if a := 0; a == 0 {
		fmt.Println("a is 0")
	}
	// a is 0

	// a is not existent at this point
	if aa := 0; aa == 1 {
		fmt.Println("aa is 1")
	} else {
		fmt.Println("aa is not 1")
	}
	// aa is not 1

	if aaa := 0; aaa == 2 {
		fmt.Println("aaa is 2")
	} else if aaa == 3 {
		fmt.Println("aaa is 3")
	} else {
		fmt.Println("aaa is neither 2 nor 3")
	}
	// aaa is neither 2 nor 3

	if 1 == 3 {

	} else if 1 == 1 {
		fmt.Println("1 == 1")
	}

	var emap = map[int]bool{
		0: true,
		1: false,
	}

	if emap[0] || emap[1] {
		fmt.Println("true or false is true")
	}
	// true or false is true

	if emap[1] || emap[0] {
		fmt.Println("false or true is true")
	}
	// false or true is true

	if emap[0] || emap[0] {
		fmt.Println("true or true is true")
	}
	// true or true is true

	if emap[0] && emap[0] {
		fmt.Println("true and true is true")
	}
	// true and true is true

	if emap[0] && emap[1] {

	} else {
		fmt.Println("true and false is false")
	}
	// true and false is false

	if emap[0] || emap[100] {
		fmt.Println("emap[100] is not evaluated")
	}
	// emap[100] is not evaluated

	if v, ok := emap[0]; !ok {

	} else {
		fmt.Println("value is passed through the whole if-control structure:", v)
	}
	// value is passed through the whole if-control structure: true

	if v, ok := emap[100]; ok {

	} else {
		fmt.Println("if not exist, it returns a zero value:", ok, v)
	}
	// if not exist, it returns a zero value: false false

	// fmt.Println(v)
	// undefined: v

	if emap[100] || emap[0] {
		fmt.Println("emap[0] gets evaluated")
		fmt.Println("emap[100]:", emap[100])
	}
	// emap[0] gets evaluated
	// emap[100]: false

	var smap = map[int]string{
		0: "A",
	}
	if smap[100] || emap[0] {

	}
	// invalid operation: smap[100] || emap[0] (mismatched types string and bool)
}
