[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Bitwise operation

- [Reference](#reference)
- [bit, byte](#bit-byte)
- [`bitwise operation`](#bitwise-operation-1)
- [`bit mask`](#bit-mask)
- [`bitset`](#bitset)
- [hamming distance](#hamming-distance)

[↑ top](#bitwise-operation)
<br><br><br><br><hr>


#### Reference

- [Bit masks](https://www.arduino.cc/en/Tutorial/BitMask)
- [Mask (computing)](https://en.wikipedia.org/wiki/Mask_(computing))
- [Bitwise operation](https://en.wikipedia.org/wiki/Bitwise_operation)
- [Bit manipulation](https://en.wikipedia.org/wiki/Bit_manipulation)
- [Bitwise Operators and Bit Masks](http://www.vipan.com/htdocs/bitwisehelp.html)
- [Bit Twiddling Hacks](http://graphics.stanford.edu/~seander/bithacks.html)
- [XOR'd play: Normalized Hamming Distance](http://trustedsignal.blogspot.com/2015/06/xord-play-normalized-hamming-distance.html?view=classic)

[↑ top](#bitwise-operation)
<br><br><br><br><hr>


#### bit, byte

Bit is **B**inary Dig**it**. Bit is the basic unit of information in computing.
Most common representation of bits is 0 and 1.
Code consists of a sequence of **bits**—*a value of 0 or 1*. 
*One* **_8-bit_** *chunk* builds *one* **_byte_** that represents **_one character_**.
**_byte_** is the smallest addressable unit of memory in computing.
Go `string` is a sequence of `bytes`. Try this [code](http://play.golang.org/p/EqXQKOGTJ9):

```go
package main

import "fmt"

func main() {
	bts := []byte("Hello")
	bts[0] = byte(100)
	for _, c := range bts {
		fmt.Println(string(c), c)
	}
	/*
	   d 100
	   e 101
	   l 108
	   l 108
	   o 111
	*/

	rs := []rune("Hello")
	rs[0] = rune(100)
	for _, c := range rs {
		fmt.Println(string(c), c)
	}
	/*
	   d 100
	   e 101
	   l 108
	   l 108
	   o 111
	*/

	str := "Hello"
	// str[0] = byte(100)
	// cannot assign to str[0]
	for _, c := range str {
		fmt.Println(string(c), c)
	}
	/*
	   H 72
	   e 101
	   l 108
	   l 108
	   o 111
	*/
}
```


<br>

C++ has no built-in `byte` type; it instead has `char` array
to represent bytes of data:

```cpp
#include <iostream>
using namespace std;

int main()
{
	// string literals are regular arrays
	//
	// \0 is a null character
	char bt1[] = {'H', 'e', 'l', 'l', 'o', '\0'};
	char bt2[] = "Hello";
	// final null character ('\0') is appended automatically
	cout << (bt1 == bt2) << endl; // 0

	cout << bt1 << endl << bt2 << endl;
	// Hello
	// Hello
	
	int i = 0;
	while (bt1[i] != '\0'){
		cout << bt1[i];
		i++;
	}
	cout << endl; // Hello
	i = 0;
	while (bt2[i] != '\0'){
		cout << bt2[i];
		i++;
	}
	cout << endl; // Hello

	// Is character array mutable? Yes.
	bt1[0] = 'A';
	cout << bt1 << endl;
	// Aello

	typedef unsigned char BYTE;
	BYTE text[] = "text";
	cout << text << endl; // text
}

```

[↑ top](#bitwise-operation)
<br><br><br><br><hr>


#### `bitwise operation`

**AND** (`&`) operator results in 1 at each bit position
only when both inputs are 1:

```
    x:  10001101
    y:  01010111
x & y:  00000101
```

<br>

**OR** (`|`) operator results in 1 at each bit position
when either input is 1:

```
    x:  10001101
    y:  01010111
x | y:  11011111
```

<br>

**XOR** (`^`) results in 1 at each bit position
only when two bits are different:

```
    x:  0101 (decimal 5)
    y:  0011 (decimal 3)
x ^ y:  0110 (decimal 6)
```

<br>

**NOT** (`^` or `~`) clears bits set to 0
where *y*'s bits are 1:

```
     x:  0101 (decimal 5)
    ^x:  0011 (decimal 3)
```

<br>

**AND NOT** (`&^` or `&~`) clears bits set to 0
where *y*'s bits are 1:

```
     x:  0101 (decimal 5)
     y:  0011 (decimal 3)
x &^ y:  0100 (decimal 4)
```

<br>

**Left Shift** (`<<`):

```
        x = 1010
        y = x << 1
yields: y = 10100 
```

<br>

**Right Shift** (`>>`):

```
        x = 1010
        y = x >> 1
yields: y = 101 
```

<br>

In Go, you would:

```go
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

```


<br>

In C++, you would:

```cpp
#include <iostream>
#include <stdio.h>
#include <string>
#include <bitset>
using namespace std;

long unsigned int toBin(long unsigned int num) {
	if (num == 0)
		return 0;
	return (num % 2) + 10*toBin(num/2);
}

long unsigned int toLu(string bstr) {
	// auto size = bstr.length();
	// to_ulong converts to unsigned long integer

	// Templates are evaluated when compiled
	const int size = 100;
	long unsigned int num = bitset<size>(bstr).to_ulong();
	printf("%10s (decimal %lu)\n", bstr.c_str(), num);
	return num;
}

int main()
{
	cout << toBin(5) << endl;
	// 101

	cout << bitset<100>("101").to_ulong() << endl;
	// 5

	cout << toLu("101") << endl;
	//       101 (decimal 5)

	cout << endl;
	cout << "AND:  x & y" << endl;
	unsigned int x = toLu("10001101");
	unsigned int y = toLu("01010111");
	unsigned int z = x & y;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "OR:  x | y" << endl;
	x = toLu("10001101");
	y = toLu("01010111");
	z = x | y;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "XOR:  x ^ y" << endl;
	x = toLu("0101");
	y = toLu("0011");
	z = x ^ y;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "NOT(bit complement):  ^x  or  ~x" << endl;
	x = toLu("0101");
	// ~ flips every bit.
	/*
	NOT 0000 0000 0000 0000 0000 0000 0000 0101
	-------------------------------------------
	    1111 1111 1111 1111 1111 1111 1111 1010
	*/	
	z = ~x;
	printf("%10lu (decimal %d)\n", toBin(z), z);
	z = x ^ 0xF;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "AND NOT:  x &^ y  or  x &~ y" << endl;
	x = toLu("0101");
	y = toLu("0011");
	z = x &~ y;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "Left Shift:  x << 1"<< endl;
	x = toLu("1010");
	y = x << 1;
	printf("%10lu (decimal %d)\n", toBin(y), y);

	cout << endl;
	cout << "Right Shift:  x >> 1" << endl;
	x = toLu("1010");
	y = x >> 1;
	printf("%10lu (decimal %d)\n", toBin(y), y);
}

/*
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

NOT(bit complement):  ^x  or  ~x
      0101 (decimal 5)
13368089053625086306 (decimal -6)
      1010 (decimal 10)

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

```

[↑ top](#bitwise-operation)
<br><br><br><br><hr>


#### `bit mask`

A `mask` is used for [`bitwise operation`](https://en.wikipedia.org/wiki/Bitwise_operation).
Binary numbers are only of `1` and `0`, read from right to left (like decimals).
To switch these numbers on and off or manipulate, you need `bit mask` values
to mask out certain bits by appling `bitwise operation`s. So a `mask` defines 
through `bitwise operation` what bits to keep (`1`) and what bits to remove (`0`).
Here's an example `bit masking` with `bitwise operation`. This only keeps the
bottom 4 bits:

```
	   10101010 = 0xAA
   AND 00001111 = 0x0F  <-- this is the 4-bit mask
	   --------
EQUALS 00001010 = 0x0A
```

<br>

In Go:

```go
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

```

<br>

In C++:

```cpp
#include <iostream>
#include <stdio.h>
#include <string>
#include <bitset>
using namespace std;

long unsigned int toBin(long unsigned int num) {
	if (num == 0)
		return 0;
	return (num % 2) + 10*toBin(num/2);
}

long unsigned int toLu(string bstr) {
	// auto size = bstr.length();
	// to_ulong converts to unsigned long integer

	// Templates are evaluated when compiled
	const int size = 100;
	long unsigned int num = bitset<size>(bstr).to_ulong();
	printf("%10s (decimal %lu)\n", bstr.c_str(), num);
	return num;
}

int main()
{
	cout << toBin(5) << endl;
	// 101

	cout << bitset<100>("101").to_ulong() << endl;
	// 5

	cout << toLu("101") << endl;
	//       101 (decimal 5)

	cout << endl;
	cout << "AND:  x & y" << endl;
	unsigned int x = toLu("10001101");
	unsigned int y = toLu("01010111");
	unsigned int z = x & y;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "Bitmasking:  x & 0x0F" << endl;
	x = toLu("10001101");
	z = x & 0x0F;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "Bitmasking:  x & 0xF" << endl;
	x = toLu("10001101");
	z = x & 0xF;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "Bitmasking:  x & 0x1F" << endl;
	x = toLu("100011111");
	z = x & 0x1F;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "Bitmasking:  x & y" << endl;
	x = toLu("10101010011111");
	y = toLu("10101010000000");
	z = x & y;
	printf("%10lu (decimal %d)\n", toBin(z), z);
}

/*
AND:  x & y
  10001101 (decimal 141)
  01010111 (decimal 87)
       101 (decimal 5)

Bitmasking:  x & 0x0F
  10001101 (decimal 141)
      1101 (decimal 13)

Bitmasking:  x & 0xF
  10001101 (decimal 141)
      1101 (decimal 13)

Bitmasking:  x & 0x1F
 100011111 (decimal 287)
     11111 (decimal 31)

Bitmasking:  x & y
10101010011111 (decimal 10911)
10101010000000 (decimal 10880)
10101010000000 (decimal 10880)
*/

```

[↑ top](#bitwise-operation)
<br><br><br><br><hr>


#### `bitset`

```cpp
#include <iostream>
#include <stdio.h>
#include <string>
#include <bitset>
using namespace std;

long unsigned int toLu(string bstr) {
	// auto size = bstr.length();
	// to_ulong converts to unsigned long integer

	// Templates are evaluated when compiled
	const int size = 100;
	long unsigned int num = bitset<size>(bstr).to_ulong();
	printf("%10s (decimal %lu)\n", bstr.c_str(), num);
	return num;
}

int main()
{
	cout << bitset<100>("101").to_ulong() << endl;
	// 5

	cout << toLu("101") << endl;
	// 5
	//       101 (decimal 5)

	bitset<4> mybits; // mybits: 0000
	cout << "mybits.set(): " << mybits.set() << endl; // mybits: 1111
	cout << "mybits.set(2,0): " << mybits.set(2,0) << endl;
	cout << "mybits.set(2): " << mybits.set(2) << endl;
	cout << "mybits.flip(): " << mybits.flip() << endl;
	cout << "mybits.flip(1): " << mybits.flip(1) << endl;
	cout << "mybits.reset(1): " << mybits.reset(1) << endl;
}

/*
mybits.set(): 1111
mybits.set(2,0): 1011
mybits.set(2): 1111
mybits.flip(): 0000
mybits.flip(1): 0010
mybits.reset(1): 0000
*/

```

[↑ top](#bitwise-operation)
<br><br><br><br><hr>


#### hamming distance

> In information theory, the **Hamming distance** between two strings of *equal
> length* is the number of positions at which the corresponding symbols are
> different. In another way, it measures the **minimum number of
> substitutions** required to *change one string into the other*, or the 
> minimum number of errors that could have transformed one string into 
> the other.
> 
> [*Hamming distance*](https://en.wikipedia.org/wiki/Hamming_distance) *by
> Wikipedia*

<br>

In Go, you would:

```go
package main

import (
	"fmt"
	"strconv"
)

// minimum number of substitutions required to change one string into the other
// https://en.wikipedia.org/wiki/Hamming_distance

// hammingDistance returns the number of differing "bits".
func hammingDistance(txt1, txt2 []byte) int {
	if len(txt1) != len(txt2) {
		panic("Undefined for sequences of unequal length")
	}
	count := 0
	for idx, b1 := range txt1 {
		b2 := txt2[idx]
		xor := b1 ^ b2 // 1 if bits are different

		// bit count (number of 1)
		// http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetNaive
		//
		// repeat shifting from left to right (divide by 2)
		// until all bits are zero
		for x := xor; x > 0; x >>= 1 {
			// check if lowest bit is 1
			if int(x&1) == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println()
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
	}()

	fmt.Println(hammingDistance([]byte("A"), []byte("A")))             // 0
	fmt.Println(hammingDistance([]byte("A"), []byte("a")))             // 1
	fmt.Println(hammingDistance([]byte("a"), []byte("A")))             // 1
	fmt.Println(hammingDistance([]byte("aaa"), []byte("aba")))         // 2
	fmt.Println(hammingDistance([]byte("aaa"), []byte("aBa")))         // 3
	fmt.Println(hammingDistance([]byte("aaa"), []byte("a a")))         // 2
	fmt.Println(hammingDistance([]byte("karolin"), []byte("kathrin"))) // 9
}

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

```

<br>

In C++, you would:

```cpp
#include <iostream>
#include <cstring>
// #include <string.h>
#include <stdexcept>
using namespace std;

int hammingDistance(char txt1[], char txt2[]) {
	if (strlen(txt1) != strlen(txt2))
		throw invalid_argument("Undefined for sequences of unequal length");
	int count = 0;
	size_t size = strlen(txt1);
	for (int idx=0; idx < size; ++idx)
	{
		char b1 = txt1[idx];
		char b2 = txt2[idx];
		unsigned int xorBit = b1 ^ b2;

		for (int x=xorBit; x > 0; x >>= 1)
		{
			if (int(x & 1) == 1)
				count++;
		}
	}
	return count;
}

int main()
{
	char txt1[] = {'H', 'e', 'l', 'l', 'o', '\0'};
	char txt2[] = "Hello";
	// final null character ('\0') is appended automatically
	cout << (txt1 == txt2) << endl; // 0
	
	cout << "hammingDistance: " << hammingDistance(txt1, txt2) << endl;
	// hammingDistance: 0 

	char txt3[] = "A";
	char txt4[] = "a";
	cout << hammingDistance(txt3, txt4) << endl; // 1

	char txt5[] = "aaa";
	char txt6[] = "aba";
	cout << hammingDistance(txt5, txt6) << endl; // 2

	strncpy(txt6, "aBa", sizeof(txt6));
	cout << hammingDistance(txt5, txt6) << endl; // 3

	char txt7[] = "karolin";
	char txt8[] = "kathrin";
	cout << hammingDistance(txt7, txt8) << endl; // 9
}

```

[↑ top](#bitwise-operation)
<br><br><br><br><hr>

