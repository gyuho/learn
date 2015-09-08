package main

import (
	"fmt"
	"strconv"
)

// minimum number of substitutions required to change one string into the other
// https://en.wikipedia.org/wiki/Hamming_distance

// hammingDistance returns the number of differing "bits".
func hammingDistance(bt1, bt2 []byte) int {
	if len(bt1) != len(bt2) {
		panic("Undefined for sequences of unequal length")
	}
	count := 0
	for idx, b1 := range bt1 {
		b2 := bt2[idx]
		xor := b1 ^ b2 // 1 if bits are different

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
	return count
}

func main() {
	fmt.Println()
	func() {
		fmt.Println("XOR:  x ^ y")
		x := toUint64("0101")
		y := toUint64("0011")
		z := x ^ y
		fmt.Printf("%10b (decimal %d)\n", z, z)
		/*
		   XOR:  x ^ y
		         0101 (decimal 5)
		         0011 (decimal 3)
		          110 (decimal 6)
		*/
	}()

	fmt.Println(hammingDistance([]byte("A"), []byte("A")))             // 0
	fmt.Println(hammingDistance([]byte("A"), []byte("a")))             // 1
	fmt.Println(hammingDistance([]byte("a"), []byte("A")))             // 1
	fmt.Println(hammingDistance([]byte("aaa"), []byte("aba")))         // 2
	fmt.Println(hammingDistance([]byte("aaa"), []byte("aBa")))         // 3
	fmt.Println(hammingDistance([]byte("aaa"), []byte("a a")))         // 2
	fmt.Println(hammingDistance([]byte("karolin"), []byte("kathrin"))) // 9
}

func toBin(num uint64) uint64 {
	if num == 0 {
		return 0
	}
	return (num % 2) + 10*toBin(num/2)
}

func toUint64(bstr string) uint64 {
	var num uint64
	if i, err := strconv.ParseUint(bstr, 2, 64); err != nil {
		panic(err)
	} else {
		num = i
	}
	fmt.Printf("%10s (decimal %d)\n", bstr, num)
	return num
}
