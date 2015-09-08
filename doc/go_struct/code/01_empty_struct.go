package main

import (
	"fmt"
	"unsafe"
)

func N(n int) []struct{} {
	return make([]struct{}, n)
}

func main() {
	var s0 struct{}
	fmt.Println(unsafe.Sizeof(s0)) // 0

	s1 := struct{}{}
	fmt.Println(unsafe.Sizeof(s1)) // 0

	done := make(chan struct{})
	go func() {
		done <- struct{}{}
	}()
	<-done
	fmt.Println("Done")
	// Done

	for i := range N(10) {
		fmt.Print(i, " ")
	}
	// 0 1 2 3 4 5 6 7 8 9
}

/*
Same in Python:

for i in range(10):
    print i
*/
