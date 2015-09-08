/*
Write a function that converts the string

49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d

to

SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

using only raw bytes. Use hex and base64 for printing.
base64 represents binary data in an ASCII format.
*/
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const (
	input  string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
)

func main() {
	// first we need to figure out how to
	// read the hexadecimal input
	decodedHex, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("decodedHex:", string(decodedHex))
	encodedHex := hex.EncodeToString(decodedHex)
	fmt.Println("encodedHex:", encodedHex)

	fmt.Println()
	encodedBase64 := base64.StdEncoding.EncodeToString(decodedHex)
	fmt.Println("encodedBase64:", encodedBase64)
	decodedBase64, err := base64.StdEncoding.DecodeString(encodedBase64)
	if err != nil {
		panic(err)
	}
	fmt.Println("decodedBase64:", string(decodedBase64))
}

/*
decodedHex: I'm killing your brain like a poisonous mushroom
encodedHex: 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d

encodedBase64: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
decodedBase64: I'm killing your brain like a poisonous mushroom
*/
