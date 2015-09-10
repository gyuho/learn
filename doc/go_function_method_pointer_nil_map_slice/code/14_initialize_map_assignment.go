package main

import "fmt"

func main() {
	mmap1 := map[string]int{
		"hello": 10,
	}
	mmap1 = make(map[string]int)
	mmap1["A"] = 1
	fmt.Println(mmap1)
	// map[A:1]

	mmap2 := map[string]int{
		"hello": 10,
	}
	mmap2 = nil
	mmap2["A"] = 1
	fmt.Println(mmap2)
	// panic: assignment to entry in nil map
}
