[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# the matasano crypto challenges: set 1

- [Reference](#reference)
- [Convert `hex` to `base64`](#convert-hex-to-base64)
- [Fixed `XOR`](#fixed-xor)
- [Single-byte `XOR` cipher](#single-byte-xor-cipher)
- [Detect single-character `XOR`](#detect-single-character-xor)
- [Implement repeating-key `XOR`](#implement-repeating-key-xor)

[↑ top](#the-matasano-crypto-challenges-set-1)
<br><br><br><br><hr>


#### Reference

- [Journey into cryptography](https://www.khanacademy.org/computing/computer-science/cryptography)
- [Crypto tutorial](https://github.com/joearms/crypto_tutorial)
- [the matasano crypto challenges](http://cryptopals.com/)

[↑ top](#the-matasano-crypto-challenges-set-1)
<br><br><br><br><hr>


#### Convert `hex` to `base64`

In Go, you can do [this](http://play.golang.org/p/PztjI3GIcT):

```go
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

```

[↑ top](#the-matasano-crypto-challenges-set-1)
<br><br><br><br><hr>


#### Fixed `XOR`

In Go, you can do [this](http://play.golang.org/p/N9U4s8h85-):

```go
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

```

[↑ top](#the-matasano-crypto-challenges-set-1)
<br><br><br><br><hr>


#### Single-byte `XOR` cipher

In Go, you can do:

```go
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

```

[↑ top](#the-matasano-crypto-challenges-set-1)
<br><br><br><br><hr>


#### Detect single-character `XOR`

In Go, you can do:

```go
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

```

[↑ top](#the-matasano-crypto-challenges-set-1)
<br><br><br><br><hr>


#### Implement repeating-key `XOR`

In Go, you can do:

```go
/*
Encrypt with key "ICE" using repeating-key XOR:

Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal

In repeating-key XOR, you'll sequentially apply
each byte of the key; the first byte of plaintext
will be XOR'd against I, the next C, the next E,
then I again for the 4th byte, and so on.

Output:
0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272
a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f
*/
package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

const (
	input  = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key    = "ICE"
	output = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
)

func repeatingKeyXOR(input, key []byte) []byte {
	encrypted := make([]byte, len(input))
	for idx, b := range input {
		encrypted[idx] = b ^ key[idx%len(key)]
	}
	return encrypted
}

func main() {
	encrypted := repeatingKeyXOR([]byte(input), []byte(key))
	encodedHex := hex.EncodeToString(encrypted)
	fmt.Println("encodedHex:", encodedHex)
	fmt.Println(output == encodedHex)
	if !bytes.Equal([]byte(output), []byte(encodedHex)) {
		panic("different")
	}
	/*
	   encodedHex: 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f
	   true
	*/
}

```

[↑ top](#the-matasano-crypto-challenges-set-1)
<br><br><br><br><hr>
