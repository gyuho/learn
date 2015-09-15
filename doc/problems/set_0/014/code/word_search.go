package main

func make2DSlice(row, column int) [][]string {
	mat := make([][]string, row)
	// for i := 0; i < row; i++ {
	for i := range mat {
		mat[i] = make([]string, column)
	}
	return mat
}

var board [][]string = [][]string{
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

func search(target []string, letterIdx string, row, col int, found map[string]bool) {
	// base case
	if row > len(board)-1 || col > len(board[0])-1 {
		// not valid move
		// because exceeding the slice range
		return
	}
}

func main() {

}
