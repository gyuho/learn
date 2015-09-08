/*
One of the 60-character strings in files/4.txt has
been encrypted by single-character XOR.
Find it.
*/
package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

const (
	textPath      string = "files/english.txt"
	inputTextPath string = "files/4.txt"
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
	byteIndexToFreq, err := getFrequency(textPath)
	if err != nil {
		panic(err)
	}
	f, err := openToRead(inputTextPath)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	min := math.MaxFloat64
	var minDiffMessage []byte
	br := bufio.NewReader(f)
	for {
		l, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		// without this
		// panic: encoding/hex: odd length hex string
		l = strings.TrimSpace(l)

		decodedHexInput, err := hex.DecodeString(l)
		if err != nil {
			panic(err)
		}
		msg, score := getMinDiff(byteIndexToFreq, decodedHexInput)
		if score < min {
			min = score
			minDiffMessage = msg
		}
	}
	fmt.Println("minDiffMessage:", string(minDiffMessage))
	// minDiffMessage: Now that the party is jumping
}
