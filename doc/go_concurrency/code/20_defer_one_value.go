package main

import "fmt"

func main() {
	// Deferred function runs
	// in Last In First Out order
	// after the surrounding function returns.
	// NOT AFTER FOR-LOOP
	for i := range []int{0, 1, 2, 3, 4, 5} {
		defer func() {
			fmt.Println("i:", i)
		}()
	}
	fmt.Println()
	/*

	*/

	// variables that are defined ON for-loop
	// should be passed as arguments to the closure
	i1 := 0
	for i := 0; i < 3; i++ {
		i1++
		defer func() {
			fmt.Println("i1:", i, i1)
		}()
	}

	i2 := 0
	for i := 0; i < 3; i++ {
		i2++
		defer func(i, i2 int) {
			fmt.Println("i2:", i, i2)
		}(i, i2)
	}
}

/*
i2: 2 3
i2: 1 2
i2: 0 1
i1: 3 3
i1: 3 3
i1: 3 3
i: 5
i: 5
i: 5
i: 5
i: 5
i: 5
*/
