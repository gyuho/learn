package main

import (
	"fmt"
	"sort"
)

func main() {
	rows := [][]string{
		[]string{"1", "a", "1", "10"},
		[]string{"1", "b", "1", "9"},
		[]string{"1", "c", "1", "8"},
		[]string{"1", "d", "1", "7"},
		[]string{"1", "e", "1", "6"},
		[]string{"1", "f", "1", "5"},
		[]string{"1", "g", "1", "4"},
		[]string{"1", "h", "1", "3"},
		[]string{"1", "i", "1", "2"},
		[]string{"1", "j", "1", "1"},
	}
	rs1 := stringsAscending(rows, 1)
	if fmt.Sprintf("%v", rs1) != "[[1 a 1 10] [1 b 1 9] [1 c 1 8] [1 d 1 7] [1 e 1 6] [1 f 1 5] [1 g 1 4] [1 h 1 3] [1 i 1 2] [1 j 1 1]]" {
		fmt.Errorf("rs1 %v", rs1)
	}
	rs2 := stringsAscending(rows, 3)
	if fmt.Sprintf("%v", rs2) != "[[1 j 1 1] [1 a 1 10] [1 i 1 2] [1 h 1 3] [1 g 1 4] [1 f 1 5] [1 e 1 6] [1 d 1 7] [1 c 1 8] [1 b 1 9]]" {
		fmt.Errorf("rs2 %v", rs2)
	}
	rs3 := stringsDescending(rows, 1)
	if fmt.Sprintf("%v", rs3) != "[[1 j 1 1] [1 i 1 2] [1 h 1 3] [1 g 1 4] [1 f 1 5] [1 e 1 6] [1 d 1 7] [1 c 1 8] [1 b 1 9] [1 a 1 10]]" {
		fmt.Errorf("rs3 %v", rs3)
	}
	rs4 := stringsDescending(rows, 3)
	if fmt.Sprintf("%v", rs4) != "[[1 b 1 9] [1 c 1 8] [1 d 1 7] [1 e 1 6] [1 f 1 5] [1 g 1 4] [1 h 1 3] [1 i 1 2] [1 a 1 10] [1 j 1 1]]" {
		fmt.Errorf("rs4 %v", rs4)
	}
}

var sortColumnIndex int

// sortByIndexAscending sorts two-dimensional strings in an ascending order, at a specified index.
type sortByIndexAscending [][]string

func (s sortByIndexAscending) Len() int {
	return len(s)
}

func (s sortByIndexAscending) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByIndexAscending) Less(i, j int) bool {
	return s[i][sortColumnIndex] < s[j][sortColumnIndex]
}

// stringsAscending sorts two dimensional strings in an ascending order.
func stringsAscending(rows [][]string, idx int) [][]string {
	sortColumnIndex = idx
	sort.Sort(sortByIndexAscending(rows))
	return rows
}

// sortByIndexDescending sorts two-dimensional strings in an Descending order, at a specified index.
type sortByIndexDescending [][]string

func (s sortByIndexDescending) Len() int {
	return len(s)
}

func (s sortByIndexDescending) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByIndexDescending) Less(i, j int) bool {
	return s[i][sortColumnIndex] > s[j][sortColumnIndex]
}

// stringsDescending sorts two dimensional strings in a descending order.
func stringsDescending(rows [][]string, idx int) [][]string {
	sortColumnIndex = idx
	sort.Sort(sortByIndexDescending(rows))
	return rows
}
