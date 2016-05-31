package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"a", "b", "c", "hello", "d", "e"}
	idx := sort.Search(len(names), func(i int) bool { return names[i] == "hello" })
	fmt.Println(idx, names[idx]) // 3 hello

	idx = sort.Search(len(names), func(i int) bool { return names[i] == "x" })
	if idx == len(names) {
		fmt.Println("x is not found")
	}
}

/*
3 hello
x is not found
*/
