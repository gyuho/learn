package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 2, 1}
	slices := [][]int{}
	sort.Ints(slice)
	slices = append(slices, slice)
	fmt.Println(slices) // [[1 2 3]]

	slice[0], slice[1] = slice[1], slice[0]
	slices = append(slices, slice)
	fmt.Println(slices) // [[2 1 3] [2 1 3]]

	slice[1], slice[2] = slice[2], slice[1]
	slices = append(slices, slice)
	fmt.Println(slices) // [[2 3 1] [2 3 1] [2 3 1]]
}
