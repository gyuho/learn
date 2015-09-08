/*
Maximum Contiguous Subarray(substring)
-100, 1, 2 => 1, 2

Kadane Algorithm Dynamic Programming: O ( n )

Divide and Conquer method: O ( n lg n )
maximum of the following
	getMCS(A, begin, mid)
	getMCS(A, mid+1, end)
	getMCS(crossing)
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{-2, -5, 6, -2, 3, -10, 5, -6}
	fmt.Println(getMCS(s, 0, len(s)-1)) // 7
	fmt.Println(kadane(s))              // [6 -2 3] 7
}

func kadane(slice []int) ([]int, int) {
	if len(slice) == 0 {
		return []int{}, 0
	}
	temp, maxSum := 0, 0
	lastIdx := 0
	for i, v := range slice {
		temp += v
		if temp < 0 {
			temp = 0 // reset
			continue
		}
		if maxSum < temp {
			maxSum = temp
			lastIdx = i
		}
	}
	check := 0
	firstIdx := 0
	for j := lastIdx; j > 0; j-- {
		check += slice[j]
		if maxSum == check {
			firstIdx = j
		}
	}
	return slice[firstIdx : lastIdx+1], maxSum
}

func max(more ...int) int {
	max := more[0]
	for _, elem := range more {
		if max < elem {
			max = elem
		}
	}
	return max
}

func getMCS(slice []int, first, last int) int {
	if first == last {
		return slice[first]
	}
	mid := (first + last) / 2
	return max(
		getMCS(slice, first, mid),
		getMCS(slice, mid+1, last),
		getMCSAcross(slice, first, mid, last),
	)
}

func getMCSAcross(slice []int, first, mid, last int) int {
	sum1 := 0
	leftSum := math.MinInt32
	for i := mid; first <= i; i-- {
		sum1 += slice[i]
		if leftSum < sum1 {
			leftSum = sum1
		}
	}
	sum2 := 0
	rightSum := math.MinInt32
	for i := mid + 1; i <= last; i++ {
		sum2 += slice[i]
		if rightSum < sum2 {
			rightSum = sum2
		}
	}
	return leftSum + rightSum
}
