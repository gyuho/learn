package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 2, 1}
	slices := [][]int{}
	sort.Ints(slice)
	copied0 := make([]int, len(slice))
	copy(copied0, slice)
	slices = append(slices, copied0)
	fmt.Println(slices) // [[1 2 3]]

	slice[0], slice[1] = slice[1], slice[0]
	copied1 := make([]int, len(slice))
	copy(copied1, slice)
	slices = append(slices, copied1)
	fmt.Println(slices) // [[1 2 3] [2 1 3]]

	slice[1], slice[2] = slice[2], slice[1]
	copied2 := make([]int, len(slice))
	copy(copied2, slice)
	slices = append(slices, copied2)
	fmt.Println(slices) // [[1 2 3] [2 1 3] [2 3 1]]
}
