package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Copy Slice
	slice01 := []int{1, 2, 3, 4, 5}
	copy01 := make([]int, len(slice01))
	copy(copy01, slice01)
	fmt.Println(copy01) // [1 2 3 4 5]

	// PushFront
	slice02 := []int{1, 2, 3, 4, 5}
	copy02 := make([]int, len(slice02)+1)
	copy02[0] = 10
	copy(copy02[1:], slice02)
	fmt.Println(copy02) // [10 1 2 3 4 5]

	// PushFront
	pushFront := func(s *[]int, elem int) {
		temp := make([]int, len(*s)+1)
		temp[0] = elem
		copy(temp[1:], *s)
		*s = temp
	}
	pushFront(&copy02, 100)
	fmt.Println(copy02) // [100 10 1 2 3 4 5]

	// PushBack
	slice03 := []int{1, 2, 3, 4, 5}
	slice03 = append(slice03, 10)
	fmt.Println(slice03) // [1 2 3 4 5 10]

	// PopFront
	slice04 := []int{1, 2, 3, 4, 5}
	slice04 = slice04[1:len(slice04):len(slice04)]
	fmt.Println(slice04, len(slice04), cap(slice04)) // [2 3 4 5] 4 4

	// PopBack
	slice05 := []int{1, 2, 3, 4, 5}
	slice05 = slice05[:len(slice05)-1 : len(slice05)-1]
	fmt.Println(slice05, len(slice05), cap(slice05)) // [1 2 3 4] 4 4

	// Delete
	slice06 := []int{1, 2, 3, 4, 5}
	copy(slice06[3:], slice06[4:])
	slice06 = slice06[:len(slice06)-1 : len(slice06)-1]
	// copy(d.OutEdges[edge1.Vtx][idx:], d.OutEdges[edge1.Vtx][idx+1:])
	// d.OutEdges[src][len(d.OutEdges[src])-1] = nil // zero value of type or nil
	fmt.Println(slice06, len(slice06), cap(slice06)) // [1 2 3 5] 4 4

	make2DSlice := func(row, column int) [][]string {
		mat := make([][]string, row)
		// for i := 0; i < row; i++ {
		for i := range mat {
			mat[i] = make([]string, column)
		}
		return mat
	}
	mat := make2DSlice(3, 5)
	for key, value := range mat {
		fmt.Println(key, value)
	}
	/*
	   0 [    ]
	   1 [    ]
	   2 [    ]
	*/
	fmt.Println(mat[1], len(mat[1]), cap(mat[1])) // [    ] 5 5

	// iterate over rows
	for r := range mat {
		// iterate over columns
		for c := range mat[r] {
			mat[r][c] = strconv.Itoa(r) + "x" + strconv.Itoa(c)
		}
	}
	for key, value := range mat {
		fmt.Println(key, value)
	}
	/*
	   0 [0x0 0x1 0x2 0x3 0x4]
	   1 [1x0 1x1 1x2 1x3 1x4]
	   2 [2x0 2x1 2x2 2x3 2x4]
	*/
	fmt.Println(mat[1], len(mat[1]), cap(mat[1])) // [1x0 1x1 1x2 1x3 1x4] 5 5
}
