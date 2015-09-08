// http://codercareer.blogspot.com/2013/02/no-44-maximal-stolen-values.html
// Problem: Maximize the money to rob without robbing adjacent houses.
//
// Example 1. Rob 22(=15+7), not 11(=3+1+7)
// 					 Can't rob 3 + 15
//
// ___[ 3 ]___[ 15 ]___[ 1 ]___[ 4 ]___[ 7 ]___
//
//
// Example 2. Rob 11(=3+1+7), not 9(=5+4)
//
// ___[ 3 ]___[ 5 ]___[ 1 ]___[ 4 ]___[ 7 ]___
//
package main

import "fmt"

// rob returns the maximum sum of money that we can rob from the input slice.
//
// Dynamic Programming:
// if slice contains more than 3 elements
// return the maximum in two cases:
//	1. Choose the first element
//	2. Choose the second
//
func rob(houses []int) int {
	switch len(houses) {
	case 0:
		return 0
	case 1:
		return houses[0]
	case 2:
		return max(houses...)
	case 3:
		case1 := houses[0] + houses[2]
		case2 := houses[1]
		if case1 >= case2 {
			return case1
		}
		return case2
	}
	return max(
		houses[0]+rob(houses[2:]),
		houses[1]+rob(houses[3:]),
	)
}

func main() {
	testCases := []struct {
		Houses   []int
		MaxMoney int
	}{
		{[]int{3}, 3},                                       // 3 = 3
		{[]int{3, 15}, 15},                                  // 15 = 15
		{[]int{3, 15, 13}, 16},                              // 16 = 3 + 13
		{[]int{3, 15, 13, 5}, 20},                           // 20 = 15 + 5
		{[]int{3, 15, 13, 4, 7}, 23},                        // 23 = 3 + 13 + 7
		{[]int{3, 15, 1, 4, 7}, 22},                         // 22 = 15 + 7
		{[]int{3, 5, 1, 4, 7}, 12},                          // 12 = 5 + 7
		{[]int{7, 2, 4, 3, 1, 4}, 15},                       // 15 = 7 + 4 + 4
		{[]int{7, 2, 4, 3, 1, 2}, 13},                       // 13 = 7 + 4 + 2
		{[]int{1, 7, 12, 3, 1, 2}, 15},                      // 15 = 1 + 12 + 2 = 15
		{[]int{1, 7, 12, 3, 1, 2, 8, 3}, 22},                // 22 = 1 + 12 + 1 + 8
		{[]int{1, 7, 12, 3, 1, 11, 8, 1}, 25},               // 25 = 1 + 12 + 11 + 1
		{[]int{6, 7, 2, 1, 3, 9, 6, 12, 1, 2, 6, 7, 1}, 38}, // 38 = ?
	}
	for idx, testCase := range testCases {
		maxmoney1 := rob(testCase.Houses)
		maxmoney2 := testCase.MaxMoney
		fmt.Printf("Max: %3d / Original Slice: %v\n", testCase.MaxMoney, testCase.Houses)
		if maxmoney1 != maxmoney2 {
			fmt.Printf("WRONG %2d: %v != %v / %+v\n", idx, maxmoney1, maxmoney2, testCase)
		}
	}
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
