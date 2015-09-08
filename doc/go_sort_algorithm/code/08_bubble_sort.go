package main

import "fmt"

func main() {
	nums := []int{1, -1, 23, -2, 23, 123, 12, 1}
	bubbleSort(nums)
	fmt.Println(nums)
	// [-2 -1 1 1 12 23 23 123]
}

/*
O (n^2)

bubbleSort(A)
for i = 1 to A.length - 1
	for j = A.length downto i + 1
		if A[j] < A[j-1]
			exchange A[j] with A[j-1]
*/
func bubbleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := len(nums) - 1; j != i-1; j-- {
			// the bigger value 'bubbles up' to the last position
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
