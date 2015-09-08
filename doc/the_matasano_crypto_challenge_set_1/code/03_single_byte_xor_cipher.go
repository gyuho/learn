/*
This hex encoded string

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736

has been `XOR`d again a single character. Find the key, decrypt the message.
How? Devise some method for scoring a piece of English plaintext.
Chracter frequency is a good metric. Evaluate each output and choose
the one with the best score.
*/

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"os"
)

const (
	input    string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	textPath string = "files/english.txt"
)

func openToRead(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return f, err
	}
	return f, nil
}

func getFrequency(textPath string) ([256]float64, error) {
	f, err := openToRead(textPath)
	if err != nil {
		return [256]float64{}, err
	}
	defer f.Close()
	var (
		byteIndexToCount [256]int
		byteIndexToFreq  [256]float64
		// byteIndexToFreq = make([]float64, 256)
	)
	br := bufio.NewReader(f)
	var totalCharCount uint64
	for {
		c, err := br.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			return [256]float64{}, err
		}
		byteIndexToCount[int(c)]++
		totalCharCount++
	}
	for idx, count := range byteIndexToCount {
		byteIndexToFreq[idx] = float64(count) / float64(totalCharCount)
	}
	return byteIndexToFreq, nil
}

func getDiff(byteIndexToFreq [256]float64, data []byte) float64 {
	var byteIndexToCount [256]int
	for _, c := range data {
		byteIndexToCount[int(c)]++
	}
	var diff float64
	totalCharCount := len(data)
	for idx, count := range byteIndexToCount {
		if count == 0 {
			continue
		}
		freq := float64(count) / float64(totalCharCount)
		diff += math.Abs(freq - byteIndexToFreq[idx])
	}
	return diff
}

func xorAgainstChar(data []byte, c byte) []byte {
	rs := make([]byte, len(data))
	for i, b := range data {
		rs[i] = b ^ c
	}
	return rs
}

func getMinDiff(byteIndexToFreq [256]float64, decodedHexInput []byte) ([]byte, float64) {
	var (
		xoredBytes     [256][]byte
		xoredDiffScore [256]float64
	)
	for idx := range xoredBytes {
		data := xorAgainstChar(decodedHexInput, byte(idx))
		xoredBytes[idx] = data
		xoredDiffScore[idx] = getDiff(byteIndexToFreq, data)
	}
	bestIndex := 0
	min := math.MaxFloat64
	for i, v := range xoredDiffScore {
		if v < min {
			min = v
			bestIndex = i
		}
	}
	return xoredBytes[bestIndex], min
}

func main() {
	decodedHexInput, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("decodedHexInput:", string(decodedHexInput))

	fmt.Println()
	byteIndexToFreq, err := getFrequency(textPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("byteIndexToFreq:", byteIndexToFreq)

	fmt.Println()
	rs, score := getMinDiff(byteIndexToFreq, decodedHexInput)
	fmt.Println("getMinDiff:", string(rs), score)
	// getMinDiff: Cooking MC's like a pound of bacon
}
