package main

import "fmt"

// Levenshtein returns the similarity using levenshtein distance.
// (https://en.wikipedia.org/wiki/Levenshtein_distance)
func Levenshtein(txt1, txt2 []byte) float64 {
	// initialize the distance array, with position
	mat := create2Dslice(len(txt1)+1, len(txt2)+1)
	for i := 0; i < len(txt1)+1; i++ {
		mat[i][0] = i
	}
	for i := 0; i < len(txt2)+1; i++ {
		mat[0][i] = i
	}
	for i := 0; i < len(txt1); i++ {
		for j := 0; j < len(txt2); j++ {
			edit := 0
			if txt1[i] != txt2[j] {
				edit = 1
			}
			mat[i+1][j+1] = min(
				mat[i][j+1]+1,  // from txt1
				mat[i+1][j]+1,  // from txt2
				mat[i][j]+edit, // from both
			)
		}
	}
	distance := mat[len(txt1)][len(txt2)]
	if distance == 0 {
		// similarity is 1 for equal texts.
		return 1
	}
	return float64(1) / float64(distance)
}

func min(more ...int) int {
	min := more[0]
	for _, elem := range more {
		if min > elem {
			min = elem
		}
	}
	return min
}

func create2Dslice(row, column int) [][]int {
	mat := make([][]int, row)
	for i := range mat {
		mat[i] = make([]int, column)
	}
	return mat
}

func main() {
	fmt.Println(Levenshtein([]byte("Hello"), []byte("Hello")))
	// 1

	fmt.Println(Levenshtein([]byte(""), []byte("Hello")))
	// 0.2

	fmt.Println(Levenshtein([]byte("hello"), []byte("Hello")))
	// 1

	fmt.Println(Levenshtein([]byte("abc"), []byte("bcd")))
	// 0.5

	fmt.Println(Levenshtein([]byte("Hello"), []byte("Hel lo")))
	// 1

	fmt.Println(Levenshtein([]byte("look at"), []byte("google")))
	// 0.2
}
