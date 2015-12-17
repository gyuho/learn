package main

import (
	"fmt"
	"time"
)

// defer function runs in Last In First Out order
// after the surrounding function returns.
// NOT AFTER FOR-LOOP
//
// variables that are defined ON or INSIDE for-loop
// should be passed as arguments to the closure

func main() {
	func() {
		for _, i1 := range []int{0, 1, 2, 3} {
			defer func() {
				fmt.Println("defer i1:", i1)
			}()
		}
		fmt.Println()

		for _, i2 := range []int{0, 1, 2, 3} {
			defer func(i2 int) {
				fmt.Println("defer i2:", i2)
			}(i2)
		}
		fmt.Println()

		i := 0
		for _, i3 := range []int{0, 1, 2, 3} {
			i++
			defer func(i3 int) {
				fmt.Println("defer i, i3:", i, i3)
			}(i3)
		}
		fmt.Println()

		j := 0
		for _, i4 := range []int{0, 1, 2, 3} {
			j++
			defer func(j, i4 int) {
				fmt.Println("defer j, i4:", j, i4)
			}(j, i4)
		}
		fmt.Println()
	}()
	/*
		defer j, i4: 4 3
		defer j, i4: 3 2
		defer j, i4: 2 1
		defer j, i4: 1 0
		defer i, i3: 4 3
		defer i, i3: 4 2
		defer i, i3: 4 1
		defer i, i3: 4 0
		defer i2: 3
		defer i2: 2
		defer i2: 1
		defer i2: 0
		defer i1: 3
		defer i1: 3
		defer i1: 3
		defer i1: 3
	*/

	func() {
		for _, i1 := range []int{0, 1, 2, 3} {
			go func() {
				fmt.Println("go i1:", i1)
			}()
		}
		fmt.Println()
		time.Sleep(time.Second)

		for _, i2 := range []int{0, 1, 2, 3} {
			go func(i2 int) {
				fmt.Println("go i2:", i2)
			}(i2)
		}
		fmt.Println()
		time.Sleep(time.Second)

		i := 0
		for _, i3 := range []int{0, 1, 2, 3} {
			i++
			go func(i3 int) {
				fmt.Println("go i, i3:", i, i3)
			}(i3)
		}
		fmt.Println()
		time.Sleep(time.Second)

		j := 0
		for _, i4 := range []int{0, 1, 2, 3} {
			j++
			go func(j, i4 int) {
				fmt.Println("go j, i4:", j, i4)
			}(j, i4)
		}
		fmt.Println()
		time.Sleep(time.Second)
	}()
	/*
		go i1: 3
		go i1: 3
		go i1: 3
		go i1: 3

		go i2: 0
		go i2: 1
		go i2: 3
		go i2: 2
		go i, i3: 4 3
		go i, i3: 4 0
		go i, i3: 4 1
		go i, i3: 4 2


		go j, i4: 1 0
		go j, i4: 2 1
		go j, i4: 4 3
		go j, i4: 3 2
	*/
}
