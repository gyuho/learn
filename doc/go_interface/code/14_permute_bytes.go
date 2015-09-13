package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	bslice := randBytes(3)
	fmt.Println(string(bslice))
	fmt.Println(permuteBytes(bslice))
	/*
	   sam
	   [ams asm mas msa sam sma]
	*/
}

const (
	// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randBytes(n int) []byte {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}

func first(data sort.Interface) {
	sort.Sort(data)
}

// next returns false when it cannot permute any more
// http://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func next(data sort.Interface) bool {
	var k, l int
	for k = data.Len() - 2; ; k-- {
		if k < 0 {
			return false
		}
		if data.Less(k, k+1) {
			break
		}
	}
	for l = data.Len() - 1; !data.Less(k, l); l-- {
	}
	data.Swap(k, l)
	for i, j := k+1, data.Len()-1; i < j; i++ {
		data.Swap(i, j)
		j--
	}
	return true
}

// permuteBytes returns all possible permutations of string slice.
func permuteBytes(bts []byte) []string {
	slice := []string{}
	for _, c := range bts {
		slice = append(slice, string(c))
	}
	first(sort.StringSlice(slice))

	copied1 := make([]string, len(slice)) // we need to make a copy!
	copy(copied1, slice)
	result := [][]string{copied1}

	for {
		isDone := next(sort.StringSlice(slice))
		if !isDone {
			break
		}

		// https://groups.google.com/d/msg/golang-nuts/ApXxTALc4vk/z1-2g1AH9jQJ
		// Lesson from Dave Cheney:
		// A slice is just a pointer to the underlying back array, your storing multiple
		// copies of the slice header, but they all point to the same backing array.

		// NOT
		// result = append(result, slice)

		copied2 := make([]string, len(slice))
		copy(copied2, slice)
		result = append(result, copied2)
	}

	combNum := 1
	for i := 0; i < len(slice); i++ {
		combNum *= i + 1
	}
	if len(result) != combNum {
		fmt.Printf("Expected %d combinations but %+v because of duplicate elements", combNum, result)
	}

	rs := []string{}
	for _, elem := range result {
		rs = append(rs, strings.Join(elem, ""))
	}
	return rs
}
