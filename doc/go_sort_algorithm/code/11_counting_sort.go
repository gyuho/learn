package main

import "fmt"

func main() {
	nums := []int{20, 370, 45, 75, 410, 1802, 24, 2, 66}
	fmt.Println(countingSort(nums))
	// [0 2 20 24 45 66 75 370 410 1802]
}

/*
Counting Sort is O(n).

It does not do any comparison.
Instead, counting sort uses the actual values
of the elements to index into an array.
It only works for positive integers.
The running time depends on the largest element.
Therefore, if the maximum value is very large, the sorting takes long time.

range 0 to k, for some integer k:

1. Create an array(slice) of the size of the maximum value + 1.
2. Count each element.
3. Add up the elements.
4. Put them back to result.
*/

func countingSort(nums []int) []int {

	// 1. Create an array(nums) of the size of the maximum value + 1
	k := max(nums)
	count := make([]int, k+1)

	// 2. Count each element
	for i := 0; i < len(nums); i++ {
		count[nums[i]] = count[nums[i]] + 1
	}

	// 3. Add up the elements
	for i := 1; i < k+1; i++ {
		count[i] = count[i] + count[i-1]
	}

	// 4. Put them back to result
	rs := make([]int, len(nums)+1)
	for j := 0; j < len(nums); j++ {
		rs[count[nums[j]]] = nums[j]
		count[nums[j]] = count[nums[j]] - 1
	}

	return rs
}

func max(nums []int) int {
	max := nums[0]
	for _, elem := range nums {
		if max < elem {
			max = elem
		}
	}
	return max
}
