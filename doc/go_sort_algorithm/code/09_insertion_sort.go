package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, -1, 23, -2, 23, 123, 12, 1}
	insertionSort(nums1)
	fmt.Println(nums1)
	// [-2 -1 1 1 12 23 23 123]

	nums2 := []int{1, -1, 23, -2, 23, 123, 12, 1}
	insertionSortInterface(sort.IntSlice(nums2), 0, len(nums2))
	fmt.Println(nums2)
	// [-2 -1 1 1 12 23 23 123]
}

// O (n^2)
func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; (j > 0) && (nums[j] < nums[j-1]); j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

func insertionSortInterface(data sort.Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
