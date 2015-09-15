package main

import (
	"fmt"
	"strings"
)

func make2DSlice(row, column int) [][]string {
	mat := make([][]string, row)
	// for i := 0; i < row; i++ {
	for i := range mat {
		mat[i] = make([]string, column)
	}
	return mat
}

var (
	board = [][]string{
		{"A", "X", "F", "H", "K", "C", "O", "F", "Q", "R"},
		{"C", "U", "Y", "T", "X", "B", "V", "H", "F", "D"},
		{"U", "J", "X", "B", "O", "D", "E", "N", "D", "S"},
		{"B", "E", "N", "C", "X", "M", "L", "O", "I", "L"},
		{"Q", "B", "D", "O", "Z", "P", "K", "O", "C", "K"},
		{"C", "T", "H", "D", "Y", "X", "E", "R", "T", "M"},
		{"A", "O", "B", "E", "U", "C", "O", "D", "E", "E"},
		{"H", "A", "D", "F", "F", "P", "H", "P", "O", "W"},
		{"P", "L", "G", "E", "V", "F", "G", "I", "C", "V"},
		{"A", "T", "E", "A", "S", "X", "G", "J", "D", "B"},
	}

	target = []string{"C", "O", "D", "E"}
)

// each recursion needs its own storage, so we need to
// make a copy of it.
func copyMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[k] = v
	}
	return n
}

func search(
	target []string,
	letterIdx int,
	row int,
	col int,
	subPath map[string]string,
	found map[string]bool,
) {
	// base case
	if row < 0 || col < 0 || row > len(board)-1 || col > len(board[0])-1 {
		// not valid move
		// because it exceeds array(slice) range
		return
	}

	targetLetter := target[letterIdx]
	currentLetter := board[row][col]
	if targetLetter != currentLetter {
		return
	}
	letterIdx++
	subPath[currentLetter] = fmt.Sprintf("%d,%d", row, col)

	// found the path
	lastTargetLetter := target[len(target)-1]
	if targetLetter == lastTargetLetter && len(target) == len(subPath) {
		ts := []string{}
		for _, v := range target {
			ts = append(ts, subPath[v])
		}
		found[strings.Join(ts, "->")] = true
		return
	}

	// find the next letter
	search(target, letterIdx, row, col-1, copyMap(subPath), found)   // left
	search(target, letterIdx, row, col+1, copyMap(subPath), found)   // right
	search(target, letterIdx, row-1, col, copyMap(subPath), found)   // up
	search(target, letterIdx, row+1, col, copyMap(subPath), found)   // down
	search(target, letterIdx, row-1, col-1, copyMap(subPath), found) // diagonal
	search(target, letterIdx, row+1, col+1, copyMap(subPath), found) // diagonal
	search(target, letterIdx, row-1, col+1, copyMap(subPath), found) // diagonal
	search(target, letterIdx, row+1, col-1, copyMap(subPath), found) // diagonal
	return
}

func main() {
	found := make(map[string]bool)
	for row, val := range board {
		for col := range val {
			subPath := make(map[string]string)
			search(target, 0, row, col, subPath, found)
		}
	}
	for path := range found {
		fmt.Println(path)
	}
}

/*
3,3->2,4->2,5->2,6
5,0->6,1->7,2->8,3
6,5->6,6->6,7->6,8
6,5->6,6->6,7->5,6
8,8->7,8->6,7->6,8
8,8->7,8->6,7->5,6
3,3->4,3->4,2->3,1
5,0->6,1->7,2->6,3
3,3->4,3->5,3->6,3
*/
