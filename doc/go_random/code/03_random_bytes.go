package main

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	b := make([]byte, 10)
	if _, err := crand.Read(b); err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// ��.���ms#

	fmt.Println(string(randBytes(10)))
	// IdPDZOxast

	fmt.Println(multiRandBytes(3, 5))
	// [[119 121 67] [114 70 70] [112 90 100] [74 85 77] [84 101 101]]
}

// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func randBytes(bytesN int) []byte {
	const (
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, bytesN)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := bytesN-1, src.Int63(), letterIdxMax; i >= 0; {
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

func multiRandBytes(bytesN, sliceN int) [][]byte {
	m := make(map[string]struct{})
	rs := [][]byte{}
	for len(rs) != sliceN {
		b := randBytes(bytesN)
		if _, ok := m[string(b)]; !ok {
			rs = append(rs, b)
			m[string(b)] = struct{}{}
		} else {
			continue
		}
	}
	return rs
}
