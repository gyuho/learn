/*
Write a function that takes two equal-length buffers
and produce their `XOR` combination.

Pass this string

1c0111001f010100061a024b53535009181c

and decode the string. And `XOR` against this string

686974207468652062756c6c277320657965

and it should return

746865206b696420646f6e277420706c6179
*/
package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

const (
	input1 string = "1c0111001f010100061a024b53535009181c"
	input2 string = "686974207468652062756c6c277320657965"
	output string = "746865206b696420646f6e277420706c6179"
)

func main() {
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
		fmt.Println()
	}()

	decodedHexInput1, err := hex.DecodeString(input1)
	if err != nil {
		panic(err)
	}
	fmt.Println("decodedHexInput1:", string(decodedHexInput1), len(decodedHexInput1))
	// decodedHexInput1: KSSP	 18

	decodedHexInput2, err := hex.DecodeString(input2)
	if err != nil {
		panic(err)
	}
	fmt.Println("decodedHexInput2:", string(decodedHexInput2), len(decodedHexInput2))
	// decodedHexInput2: hit the bull's eye 18

	decodedHexOutput, err := hex.DecodeString(output)
	if err != nil {
		panic(err)
	}
	fmt.Println("decodedHexOutput:", string(decodedHexOutput), len(decodedHexOutput))
	// decodedHexOutput: the kid don't play 18

	resultBytes := make([]byte, len(decodedHexOutput))
	for i := 0; i < len(decodedHexOutput); i++ {
		resultBytes[i] = decodedHexInput1[i] ^ decodedHexInput2[i]
	}
	if !bytes.Equal(resultBytes, decodedHexOutput) {
		log.Fatalf("%s %s", resultBytes, decodedHexOutput)
	}
	fmt.Println("resultBytes:", string(resultBytes))
	// resultBytes: the kid don't play
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
