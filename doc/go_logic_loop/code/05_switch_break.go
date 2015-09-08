package main

import "fmt"

func main() {
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
}

/*
num is 0
num > 4
*/
