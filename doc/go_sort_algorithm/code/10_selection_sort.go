package main

import "fmt"

func main() {
	nums := []int{1, -1, 23, -2, 23, 123, 12, 1}
	selectionSort(nums)
	fmt.Println(nums)
}

// O (n^2)
func selectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		// Min is the index of the minimum element.
		// Swap it with the current position
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
}
