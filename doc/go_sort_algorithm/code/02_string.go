package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"X", "x", "a", "A", "G"} // unsorted
	sort.Strings(s)
	fmt.Println(s)
	// [A G X a x]
}
