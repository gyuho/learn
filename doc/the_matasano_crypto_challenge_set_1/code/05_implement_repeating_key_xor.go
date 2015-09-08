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
