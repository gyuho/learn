package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

func hashSha1(data []byte) string {
	sum := sha1.Sum(data)
	// convert [20]byte to []byte
	return base64.StdEncoding.EncodeToString(sum[:])
}

func hashSha256(data []byte) string {
	sum := sha256.Sum256(data)
	return base64.StdEncoding.EncodeToString(sum[:])
}

func hashSha512(data []byte) string {
	sum := sha512.Sum512(data)
	return base64.StdEncoding.EncodeToString(sum[:])
}

func hashMd5(data []byte) string {
	sum := md5.Sum(data)
	// convert [20]byte to []byte
	return base64.StdEncoding.EncodeToString(sum[:])
}

func main() {
	data := []byte("Hello World!")
	fmt.Println("hashSha1:", hashSha1(data))
	fmt.Println("hashSha256:", hashSha256(data))
	fmt.Println("hashSha512:", hashSha512(data))
	fmt.Println("hashMd5:", hashMd5(data))
}

/*
hashSha1: Lve95gjOVATpfV8EL5X4nxwjKHE=
hashSha256: f4OxZX/x/FO5LcGBSKHWXfwtSx+j1ncoSt3SABJtkGk=
hashSha512: hhhE1nBOhXP+w02WfiC8/vPUJM9IvgTm3AjyvVjHKXQzcQFerYkcw88cnTS0kmS1EHUbH/nlN5N7xGtdb/TsyA==
hashMd5: 7Qdih1MuhjZehB6Sv8UNjA==
*/
