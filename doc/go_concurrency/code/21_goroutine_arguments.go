package main

import "fmt"

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	// variables that are defined ON for-loop
	// should be passed as arguments to the closure

	i1 := 0
	for i := 0; i < 3; i++ {
		i1++
		go func() {
			ch1 <- fmt.Sprintf("i1: %d %d", i, i1)
		}()
	}

	i2 := 0
	for i := 0; i < 3; i++ {
		i2++
		go func(i, i2 int) {
			ch2 <- fmt.Sprintf("i2: %d %d", i, i2)
		}(i, i2)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("ch1:", <-ch1)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("ch2:", <-ch2)
	}

	/*
		ch1: i1: 3 3
		ch1: i1: 3 3
		ch1: i1: 3 3
		ch2: i2: 0 1
		ch2: i2: 1 2
		ch2: i2: 2 3
	*/
}
