package main

import "fmt"

func main() {
	ts := []int{0, 1, 2}
	for i := 0; i < len(ts); i++ {
		fmt.Print(i, " ")
	}
	// 0 1 2
	fmt.Println()
	for i := len(ts) - 1; i > -1; i-- {
		fmt.Print(i, " ")
	}
	// 2 1 0
	fmt.Println()

	/*
		continue, break, goto, fallthrough

		continue: for-loop
		break: for-loop, switch
		goto: anything
		fallthrough: switch

		break, continue should only be in for-loop (break can be in switch)
		break, continue can be in if-statement
		 only when the if-statement is enclosed by for-loop

		Label can be used with goto, break, continue
	*/

	/**************************************/
	cn := 0
	for {
		cn++
		if cn == 5 {
			fmt.Println("cn is 5. Break(End) this loop!")
			break
		}

		if cn == 3 {
			fmt.Println("cn is 3. Continue this loop!")
			continue
		}
	}
	// cn is 3. Continue this loop!
	// cn is 5. Break(End) this loop!

	println()
	/**************************************/
	b := 0
	for i := 0; i < 10; i++ {
		b++
		if b == 500 {
			fmt.Println("b is 500. Break(End) this loop!")
			break
		}

		if b == 3 {
			fmt.Println("b is 3. Continue this loop!")
			continue
			fmt.Println("This does not print!")
		}
	}
	// b is 3. Continue this loop!
	// b does not reach 500, because before then
	// , we exit the for-loop condition

	println()
	/**************************************/
	goto here

	for {
		fmt.Println("Infinite Looping; Not Printing")
	}

here:
	fmt.Println("After goto")

	for i := 0; i < 2; i++ {
		for j := 0; j < 1000; j++ {
			if j == 3 {
				fmt.Println("Before Break, i == ", i)

				break
				fmt.Println("Not Printing")
			}
		}
		fmt.Println("Before continue, i == ", i)

		continue
		// go back to for i := 0;
		fmt.Println("Not Printing")
	}
	/*
	   After goto
	   Before Break, i ==  0
	   Before continue, i ==  0
	   Before Break, i ==  1
	   Before continue, i ==  1
	*/

	println()
	/**************************************/
	num := 0
Loop:
	for i := 0; i < 5; i++ {
		switch {
		case num < 3:
			fmt.Println("num is", num)
			num = num + 5
			break
			// this only breaks the enclosing switch statement
			// this does not break the for-loop

		case num > 4:
			fmt.Println("num > 4")
			break Loop
			// break the for-loop, which is labeled as "Loop"
			// this does not run this for-loop anymore

			// if we use goto Loop
			// it goes into infinite loop
			// because it starts over from for-loop
		default:
			fmt.Println("a")
		}
	}
	/*
	   num is 0
	   num > 4
	*/

	println()
	/**************************************/
	limit := 0
Cont:
	for {
		limit++
		if limit == 3 {
			fmt.Println("Break(End) this loop!")
			break
		}
		fmt.Print("..")
		continue Cont
		fmt.Println("This does not print!")
	}
	// ....Break(End) this loop!

	println()
	/**************************************/

	// "fallthrough" statement transfers control
	// to the first statement of the next case
	// clause in a expression "switch" statement.
	// It may be used only as the final non-empty
	// statement in such a clause.

	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		// To fall through to a subsequent case
		fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 1")
	}
	// 1 > 10

	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		// fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 2")
	}
	// [no output]

	bf := 2
	switch bf {
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("None 3")
	}
	// 3

	cf := 5
	switch cf {
	case 1:
		fmt.Println("1")
	case 2:
		// fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("None 4")
	}
	// None 4

	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		break
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 5")
	}
	// [no output]

	/*************************************
	  9 ways to use for-loop

	  Go has only one looping construct, the for loop.

	  The basic for loop looks as it does in C or Java
	  , except that we do not use ( )
	  (they are not even optional)
	   and the { } is required.
	  **************************************/
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// end when i == -1
	for i := 2; i >= 0; i-- {
		print(i, ",")
	}
	// 2,1,0,

	sum2 := 1
	for sum2 < 100 {
		sum2 += sum2
	}
	fmt.Println(sum2) // 128

	slice := []int{1, 2, 3, 4, 5}
	if a := 1; a > 0 {
		fmt.Println(a + slice[3])
	}
	// 5

	for b := 2; b != 5; {
		fmt.Print(slice[b])
		b++
	}
	// 345

	/**************************************/
	// infinite loop
	end := 0
	for {
		// infinite loop
		end++
		if end == 10 {
			fmt.Println("Ending(Breaking) for-loop!")
			break
		}
	}

	/**************************************/
	// string
	strk := "헬로우"
	for key, value := range strk {
		fmt.Println("String:", key, value)
	}
	/*
	   String: 0 54764
	   String: 3 47196
	   String: 6 50864
	*/

	for key, value := range strk {
		fmt.Printf("String: %v %c \n", key, value)
	}
	for key, value := range strk {
		fmt.Println("String:", key, string(value))
	}
	/*
	   String: 0 헬
	   String: 3 로
	   String: 6 우
	*/

	for _, value := range strk {
		fmt.Println("String Value:", string(value))
	}
	/*
	   String: 0 헬
	   String: 3 로
	   String: 6 우
	*/

	for key := range strk {
		fmt.Println("String Key:", key)
	}
	/*
	   String Key: 0
	   String Key: 3
	   String Key: 6
	*/

	/**************************************/
	// array
	for key, value := range [...]int{10, 20, 30} {
		fmt.Println("Array:", key, value)
	}
	/*
	   Array: 0 10
	   Array: 1 20
	   Array: 2 30
	*/

	for _, value := range [3]int{10, 20, 30} {
		fmt.Println("Array Value:", value)
	}
	/*
	   Array Value: 10
	   Array Value: 20
	   Array Value: 30
	*/

	for key := range [3]int{10, 20, 30} {
		fmt.Println("Array Key:", key)
	}
	/*
	   Array Key: 0
	   Array Key: 1
	   Array Key: 2
	*/

	/**************************************/
	// slice
	for key, value := range []string{"A", "B"} {
		fmt.Println("Slice:", key, value)
	}
	/*
	   Slice: 0 A
	   Slice: 1 B
	*/

	for _, value := range []string{"A", "B"} {
		fmt.Println("Slice Value:", value)
	}
	/*
	   Slice Value: A
	   Slice Value: B
	*/

	for key := range []string{"A", "B"} {
		fmt.Println("Slice Key:", key)
	}
	/*
	   Slice Key: 0
	   Slice Key: 1
	*/

	/**************************************/
	// map (Unordered Data Structure)
	// map's key values are like index, but not integer
	// array is also map with Hash Function
	// with integers as indices
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
	}
	for key, value := range elements {
		fmt.Println("Map:", key, value)
	}
	/*
	   Map: H Hydrogen
	   Map: He Helium
	   Map: Li Lithium
	*/

	for _, value := range elements {
		fmt.Println("Map Value:", value)
	}
	/*
	   Map Value: Hydrogen
	   Map Value: Helium
	   Map Value: Lithium
	*/

	for key := range elements {
		fmt.Println("Map Key:", key)
	}
	/*
	   Map Key: H
	   Map Key: He
	   Map Key: Li
	*/

	value1, exist1 := elements["B"]
	fmt.Println(value1, exist1)
	//  false

	value2, exist2 := elements["He"]
	fmt.Println(value2, exist2)
	// Helium true
}
