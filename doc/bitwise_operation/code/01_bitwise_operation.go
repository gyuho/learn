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
	fmt.Println(toBin(5))
	// 101

	fmt.Println()
	func() {
		fmt.Println("AND:  x & y")
		x := toUint64("10001101")
		y := toUint64("01010111")
		z := x & y
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	func() {
		fmt.Println("OR:  x | y")
		x := toUint64("10001101")
		y := toUint64("01010111")
		z := x | y
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	func() {
		fmt.Println("XOR:  x ^ y")
		x := toUint64("0101")
		y := toUint64("0011")
		z := x ^ y
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	func() {
		fmt.Println("XOR:  x ^ y")
		x := toUint64("1111")
		y := toUint64("1111")
		z := x ^ y
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	func() {
		fmt.Println("XOR:  x ^ y")
		x := toUint64("0000")
		y := toUint64("0000")
		z := x ^ y
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	func() {
		fmt.Println("NOT(bit complement):  ^x  or  ~x")
		x := toUint64("0101")
		z := ^x
		// or x ^ 0x0F
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	func() {
		fmt.Println("NOT(bit complement):  ^ 0xF")
		x := toUint64("0101")
		z := x ^ 0xF
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	func() {
		fmt.Println("NOT(bit complement):  ^ 0x1F")
		x := toUint64("010101")
		z := x ^ 0x3F
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	func() {
		fmt.Println("AND NOT:  x &^ y  or  x &~ y")
		x := toUint64("0101")
		y := toUint64("0011")
		z := x &^ y
		fmt.Printf("%10b (decimal %d)\n", z, z)
	}()

	fmt.Println()
	func() {
		fmt.Println("Left Shift:  x << 1")
		x := toUint64("1010")
		y := x << 1
		fmt.Printf("%10b (decimal %d)\n", y, y)
	}()

	fmt.Println()
	func() {
		fmt.Println("Right Shift:  x >> 1")
		x := toUint64("1010")
		y := x >> 1
		fmt.Printf("%10b (decimal %d)\n", y, y)
	}()
}

/*
101

AND:  x & y
  10001101 (decimal 141)
  01010111 (decimal 87)
       101 (decimal 5)

OR:  x | y
  10001101 (decimal 141)
  01010111 (decimal 87)
  11011111 (decimal 223)

XOR:  x ^ y
      0101 (decimal 5)
      0011 (decimal 3)
       110 (decimal 6)

XOR:  x ^ y
      1111 (decimal 15)
      1111 (decimal 15)
         0 (decimal 0)

XOR:  x ^ y
      0000 (decimal 0)
      0000 (decimal 0)
         0 (decimal 0)

NOT(bit complement):  ^x  or  ~x
      0101 (decimal 5)
1111111111111111111111111111111111111111111111111111111111111010 (decimal 18446744073709551610)

NOT(bit complement):  ^ 0xF
      0101 (decimal 5)
      1010 (decimal 10)

NOT(bit complement):  ^ 0x1F
    010101 (decimal 21)
    101010 (decimal 42)

AND NOT:  x &^ y  or  x &~ y
      0101 (decimal 5)
      0011 (decimal 3)
       100 (decimal 4)

Left Shift:  x << 1
      1010 (decimal 10)
     10100 (decimal 20)

Right Shift:  x >> 1
      1010 (decimal 10)
       101 (decimal 5)
*/
