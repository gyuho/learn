package main

import (
	"bytes"
	"fmt"
)

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

func main() {
	fmt.Println(Hamming([]byte("A"), []byte("A")))             // 1
	fmt.Println(Hamming([]byte("A"), []byte("a")))             // 1
	fmt.Println(Hamming([]byte("a"), []byte("A")))             // 1
	fmt.Println(Hamming([]byte("aaa"), []byte("aba")))         // 0.5
	fmt.Println(Hamming([]byte("aaa"), []byte("aBa")))         // 0.333
	fmt.Println(Hamming([]byte("aaa"), []byte("a a")))         // 0.5
	fmt.Println(Hamming([]byte("karolin"), []byte("kathrin"))) // 0.1111111111111111

	fmt.Println(Hamming([]byte("Hello"), []byte("Hello")))
	// 1

	fmt.Println(Hamming([]byte("Hello"), []byte("Hel lo")))
	// 0.2

	fmt.Println(Hamming([]byte(""), []byte("Hello")))
	// 0.05

	fmt.Println(Hamming([]byte("hello"), []byte("Hello")))
	// 1

	fmt.Println(Hamming([]byte("abc"), []byte("bcd")))
	// 0.16666666666666666
}
