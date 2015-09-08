package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	nums := []int{732, 23, 1, 55, 7130, 321, 223, 5}
	radixSort(nums)
	fmt.Println(nums)
	// [1 5 23 55 223 321 732 7130]
}

/*
Radix Sort
1. Set up an array of initially empty "buckets"
2. Take the smallest of each element
3. Group elements from the smallest
4. Repeat the process
*/
func radixSort(nums []int) {

	// 1. Set up an array of initially empty "buckets"
	// create 10 buckets of which is a list
	var bucketList [10]*list.List
	for i := 0; i < 10; i++ {
		// initialize each bucket
		bucketList[i] = list.New()
	}

	max := max(nums)
	maxdigit := 0
	for max > 0 {
		// 2/10 == 0, 2%10 == 2
		max /= 10

		// if max is 812, maxdigit is 3
		maxdigit++
	}

	/*
		2. Take the smallest of each element
		3. Group elements from the smallest
		4. Repeat the process
	*/
	// if i is 2, then it means 3rd digit
	// if i is 2, in 321, i is 1
	for i := 0; i < maxdigit; i++ {

		// Pow10 returns 10**e, the base-10 exponential of e
		// math.Pow10(2) is 100
		p := int(math.Pow10(i + 1))
		q := int(math.Pow10(i))

		for j := 0; j < len(nums); j++ {
			/*
				x is the i-th digit

				if nums[0] is 123, and i is 0
				then 123 % 10 / 1 ---> x is 3

				if nums[0] is 123, and i is 1
				then 123 % 100 / 10 ---> x is 2
			*/
			x := nums[j] % p / q

			// add nums[j] to x th bucket
			// group by the digit
			bucketList[x].PushBack(nums[j])
		}

		count := 0
		for k := 0; k < 10; k++ {
			for elem := bucketList[k].Front(); elem != nil; elem = elem.Next() {
				nums[count] = elem.Value.(int)
				count++
			}
			bucketList[k].Init()
		}
	}
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
