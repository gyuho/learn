package main

import (
	"bytes"
	"fmt"
	"math"
)

func main() {
	fmt.Println(
		Get(
			[]byte("I love LA and New York"),
			[]byte("I love New York and LA"),
			Cosine, Hamming, Levenshtein,
		),
	) // 1.0906593406593406

	fmt.Println(
		Get(
			[]byte("I love LA and New York"),
			[]byte("string similarity test..."),
			Cosine, Hamming, Levenshtein,
		),
	) // 0.37105513409025503
}

// Get returns the string similarity from the functions.
// Predefined functions in this package use the scale from 0 to 1
// with higher value for more similar texts.
func Get(txt1, txt2 []byte, functions ...func([]byte, []byte) float64) float64 {
	rs := 0.0
	for _, f := range functions {
		rs += f(txt1, txt2)
	}
	return rs
}

// Cosine converts texts to vectors
// associatting each chracter with its frequncy
// and caculates cosine similarities.
// (https://en.wikipedia.org/wiki/Cosine_similarity)
func Cosine(txt1, txt2 []byte) float64 {
	vect1 := make(map[byte]int)
	for _, t := range txt1 {
		vect1[t]++
	}
	vect2 := make(map[byte]int)
	for _, t := range txt2 {
		vect2[t]++
	}
	//
	// dot-product two vectors
	// map[byte]int return 0 for non-existing key
	// and if two texts are equal, product will be highest
	// and if two texts are totally different, it will be 0
	//
	// to calculate A·B
	dotProduct := 0.0
	for k, v := range vect1 {
		dotProduct += float64(v) * float64(vect2[k])
	}
	// to calculate |A|*|B|
	sum1 := 0.0
	for _, v := range vect1 {
		sum1 += math.Pow(float64(v), 2)
	}
	sum2 := 0.0
	for _, v := range vect2 {
		sum2 += math.Pow(float64(v), 2)
	}
	magnitude := math.Sqrt(sum1) * math.Sqrt(sum2)
	if magnitude == 0 {
		return 0.0
	}
	return float64(dotProduct) / float64(magnitude)
}

// Hamming returns the normalized similarity value.
// hamming distance is the number of differing "bits".
// hamming distance is minimum number of substitutions
// required to change one string into the other
// (https://en.wikipedia.org/wiki/Hamming_distance)
func Hamming(txt1, txt2 []byte) float64 {
	switch bytes.Compare(txt1, txt2) {
	case 0: // txt1 == txt2
	case 1: // txt1 > txt2
		temp := make([]byte, len(txt1))
		copy(temp, txt2)
		txt2 = temp
	case -1: // txt1 < txt2
		temp := make([]byte, len(txt2))
		copy(temp, txt1)
		txt1 = temp
	}
	if len(txt1) != len(txt2) {
		panic("Undefined for sequences of unequal length")
	}
	count := 0
	for idx, b1 := range txt1 {
		b2 := txt2[idx]
		xor := b1 ^ b2 // 1 if bits are different
		//
		// bit count (number of 1)
		// http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetNaive
		//
		// repeat shifting from left to right (divide by 2)
		// until all bits are zero
		for x := xor; x > 0; x >>= 1 {
			// check if lowest bit is 1
			if int(x&1) == 1 {
				count++
			}
		}
	}
	if count == 0 {
		// similarity is 1 for equal texts.
		return 1
	}
	return float64(1) / float64(count)
}

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
