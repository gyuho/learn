/*
Longest Common Subsequence

Subsequence needs not be contiguous
X = BDCABA
Y = ABCBDAB => LCS is B C B

Dynamic Programming method : O ( m * n )
*/
package main

import "fmt"

func main() {
	fmt.Println(LCS("BDCABA", "ABCBDAB"))
	// 4 BDAB

	fmt.Println(LCS("AXXBCDXXX", "XFXHXKX"))
	// 4 XXXX

	fmt.Println(LCS("AGGTABTABTABTAB", "GXTXAYBTABTABTAB"))
	// 13 GTABTABTABTAB

	fmt.Println(LCS("AGGTABGHSRCBYJSVDWFVDVSBCBVDWFDWVV", "GXTXAYBRGDVCBDVCCXVXCWQRVCBDJXCVQSQQ"))
	// 14 GTABGCBVWVCBDV
}

func LCS(str1, str2 string) (int, string) {
	size1 := len(str1)
	size2 := len(str2)
	mat := create2D(size1+1, size2+1)
	i, j := 0, 0
	for i = 0; i <= size1; i++ {
		for j = 0; j <= size2; j++ {
			if i == 0 || j == 0 {
				mat[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				mat[i][j] = mat[i-1][j-1] + 1
			} else {
				mat[i][j] = max(mat[i-1][j], mat[i][j-1])
			}
		}
	}
	return mat[size1][size2], backTrack(mat, str1, str2, size1-1, size2-1)
}

func backTrack(mat [][]int, str1, str2 string, i, j int) string {
	if i == -1 || j == -1 {
		return ""
	} else if str1[i] == str2[j] {
		return backTrack(mat, str1, str2, i-1, j-1) + string(str1[i])
	}
	if mat[i+1][j] > mat[i][j+1] {
		return backTrack(mat, str1, str2, i, j-1)
	}
	return backTrack(mat, str1, str2, i-1, j)
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

func create2D(size1, size2 int) [][]int {
	mat := make([][]int, size1)
	for i := range mat {
		mat[i] = make([]int, size2)
	}
	return mat
}
