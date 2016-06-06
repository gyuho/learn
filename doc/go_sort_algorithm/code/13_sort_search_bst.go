package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(`numbers[i] >= -1 (ascending)`)
	{
		// given a slice data sorted in ascending order
		numbers := []int{-2, -1, 0, 1, 2}
		fmt.Println("IsSorted:", sort.IsSorted(sort.IntSlice(numbers)))

		idx := sort.Search(
			len(numbers),
			func(i int) bool {
				fmt.Println("searching at", i)
				return numbers[i] >= -1
			})

		if idx < len(numbers) {
			fmt.Println(idx, numbers[idx])
		} else {
			fmt.Println("-1 is not found:", idx)
		}
	}
	/*
	   numbers[i] >= -1 (ascending)
	   IsSorted: true
	   searching at 2
	   searching at 1
	   searching at 0
	   1 -1
	*/

	println()
	fmt.Println(`names[i] == "d" (ascending)`)
	{
		// given a slice data sorted in ascending order
		names := []string{"a", "b", "c", "d", "e"}
		fmt.Println("IsSorted:", sort.IsSorted(sort.StringSlice(names)))

		idx := sort.Search(
			len(names),
			func(i int) bool {
				fmt.Println("searching at", i)
				return names[i] == "d"
			})

		if idx < len(names) && names[idx] == "d" {
			fmt.Println(idx, names[idx])
		} else {
			fmt.Println("d is not found")
		}
	}
	/*
		names[i] == "d" (ascending)
		IsSorted: true
		searching at 2
		searching at 4
		d is not found
	*/

	println()
	fmt.Println(`names[i] >= "d" (ascending)`)
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
		names[i] >= "d" (ascending)
		IsSorted: true
		searching at 2
		searching at 4
		searching at 3
		3 d
	*/

	println()
	fmt.Println(`names[i] == "d" (descending)`)
	{
		// Searching data sorted in descending order would use
		// the <= operator instead of the >= operator.
		names := []string{"e", "d", "c", "b", "a"}
		fmt.Println("IsSorted:", sort.IsSorted(sort.Reverse(sort.StringSlice(names))))

		idx := sort.Search(
			len(names),
			func(i int) bool {
				fmt.Println("searching at", i)
				return names[i] == "d"
			})

		if idx < len(names) && names[idx] == "d" {
			fmt.Println(idx, names[idx])
		} else {
			fmt.Println("d is not found")
		}
	}
	/*
	   names[i] == "d" (descending)
	   IsSorted: true
	   searching at 2
	   searching at 4
	   d is not found
	*/

	println()
	fmt.Println(`names[i] <= "d" (descending)`)
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
		names[i] <= "d" (descending)
		IsSorted: true
		searching at 2
		searching at 1
		searching at 0
		1 d
	*/

	println()
	fmt.Println(`names[i] <= "x" (descending)`)
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
		names[i] <= "x" (descending)
		IsSorted: true
		searching at 2
		searching at 1
		searching at 0
		x is not found
	*/
}
