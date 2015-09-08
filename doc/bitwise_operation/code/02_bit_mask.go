package main

import (
	"fmt"
	"strconv"
)

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

func main() {
	fmt.Println()
	func() {
		fmt.Println("Bitmasking:  x & 0x0F")
		x := toUint64("10001101")
		y := uint64(0x0F)
		z := x & y
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	func() {
		fmt.Println("Bitmasking:  x & 0xF")
		x := toUint64("10001101")
		y := uint64(0xF)
		z := x & y
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	fmt.Printf("%X\n", toUint64("11111"))
	// 1F

	fmt.Println()
	func() {
		fmt.Println("Bitmasking:  x & 0x1F")
		x := toUint64("100011111")
		y := uint64(0x1F)
		z := x & y
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	func() {
		fmt.Println("Bitmasking:  x & y")
		x := toUint64("10101010011111")
		y := toUint64("10101010000000")
		z := x & y
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()
}

/*
Bitmasking:  x & 0x0F
  10001101 (decimal 141)
      1101 (decimal 13)

Bitmasking:  x & 0xF
  10001101 (decimal 141)
      1101 (decimal 13)

     11111 (decimal 31)
1F

Bitmasking:  x & 0x1F
 100011111 (decimal 287)
     11111 (decimal 31)

Bitmasking:  x & y
10101010011111 (decimal 10911)
10101010000000 (decimal 10880)
10101010000000 (decimal 10880)
*/
