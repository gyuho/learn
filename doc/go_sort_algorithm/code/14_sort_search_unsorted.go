package main

import (
	"fmt"
	"sort"
)

func main() {
	{
		ns := []int{1, 2, 0, -1, -2}
		num := -1
		// Binary search the smallest index at which the number is greater than the given one (== 10)
		fmt.Println("IsSorted:", sort.IsSorted(sort.IntSlice(ns)))
		idx := sort.Search(len(ns), func(i int) bool {
			fmt.Println("searching", ns[i])
			return ns[i] >= num
		})
		fmt.Println("Result:", ns[idx])
	}
	/*
	   IsSorted: false
	   searching 0
	   searching 2
	   searching 1
	   Result: 1
	*/

	println()
	{
		ns := []int{2, 100, 1, 20, 10, 50}
		num := 5
		// Binary search the smallest index at which the number is greater than the given one (== 10)
		fmt.Println("IsSorted:", sort.IsSorted(sort.IntSlice(ns)))
		idx := sort.Search(len(ns), func(i int) bool {
			fmt.Println("searching", ns[i])
			return ns[i] > num
		})
		fmt.Println("Result:", ns[idx])
	}
	/*
	   IsSorted: false
	   searching 20
	   searching 100
	   searching 2
	   Result: 100
	*/

	println()
	{
		ns := []int{2, 100, 1, 20, 10, 50}
		sort.Ints(ns)
		num := 5
		// Binary search the smallest index at which the number is greater than the given one (== 10)
		fmt.Println("IsSorted:", sort.IsSorted(sort.IntSlice(ns)))
		idx := sort.Search(len(ns), func(i int) bool {
			fmt.Println("searching", ns[i])
			return ns[i] > num
		})
		fmt.Println("Result:", ns[idx])
	}
	/*
		IsSorted: true
		searching 20
		searching 2
		searching 10
		Result: 10
	*/
}
