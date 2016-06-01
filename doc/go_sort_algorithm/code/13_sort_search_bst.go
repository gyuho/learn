package main

import (
	"fmt"
	"sort"
)

func main() {
	{
		// given a slice data sorted in ascending order
		names := []string{"a", "b", "c", "d", "e"}
		fmt.Println("IsSorted:", sort.IsSorted(sort.StringSlice(names)))

		idx := sort.Search(
			len(names),
			func(i int) bool {
				fmt.Println("searching at", i)
				return names[i] >= "d"
			})

		if idx < len(names) && names[idx] == "d" {
			fmt.Println(idx, names[idx])
		} else {
			fmt.Println("d is not found")
		}
	}
	/*
		IsSorted: true
		searching at 2
		searching at 4
		searching at 3
		3 d
	*/

	{
		// Searching data sorted in descending order would use
		// the <= operator instead of the >= operator.
		names := []string{"e", "d", "c", "b", "a"}
		fmt.Println("IsSorted:", sort.IsSorted(sort.Reverse(sort.StringSlice(names))))

		idx := sort.Search(
			len(names),
			func(i int) bool {
				fmt.Println("searching at", i)
				return names[i] <= "d"
			})

		if idx < len(names) && names[idx] == "d" {
			fmt.Println(idx, names[idx])
		} else {
			fmt.Println("d is not found")
		}
	}
	/*
		IsSorted: true
		searching at 2
		searching at 1
		searching at 0
		1 d
	*/

	{
		// Searching data sorted in descending order would use
		// the <= operator instead of the >= operator.
		names := []string{"e", "d", "c", "b", "a"}
		fmt.Println("IsSorted:", sort.IsSorted(sort.Reverse(sort.StringSlice(names))))

		idx := sort.Search(
			len(names),
			func(i int) bool {
				fmt.Println("searching at", i)
				return names[i] <= "x"
			})

		if idx < len(names) && names[idx] == "x" {
			fmt.Println(idx, names[idx])
		} else {
			fmt.Println("x is not found")
		}
	}
	/*
		IsSorted: true
		searching at 2
		searching at 1
		searching at 0
		x is not found
	*/
}
