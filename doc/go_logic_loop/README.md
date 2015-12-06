[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: logic, loop

- [logic, `if`](#logic-if)
- [`switch`](#switch)
- [`select`](#select)
- [`for`](#for)
- [review `goto`](#review-goto)
- [review `switch`, `break`](#review-switch-break)
- [review `fallthrough`](#review-fallthrough)
- [empty `struct`](#empty-struct)
- [fizzbuzz](#fizzbuzz)

[↑ top](#go-logic-loop)
<br><br><br><br><hr>


#### logic, `if`
Like any other languages, Go has
[*if-statement*](https://golang.org/doc/effective_go.html#if), as
[*here*](http://play.golang.org/p/hcXQSywapX):

```go
package main
 
import "fmt"
 
func main() {
	fmt.Println(3 < 10) // true
	fmt.Println(3 > 10) // false
 
	fmt.Println(5 <= 5) // true
	fmt.Println(5 >= 5) // true
 
	fmt.Println(5 == 5) // true
	fmt.Println(3 == 1) // false
 
	fmt.Println("A" < "a") // true
	fmt.Println("A" < "b") // true
	fmt.Println("a" < "b") // true
	fmt.Println("a" > "x") // false
 
	// fmt.Println(2 =< 5)
	// syntax error
 
	// fmt.Println(2 => 5)
	// syntax error
 
	// Note that the a is only valid
	// inside this if-statement
	// a is not existent after this if-statement
	if a := 0; a == 0 {
		fmt.Println("a is 0")
	}
	// a is 0
 
	// a is not existent at this point
	if aa := 0; aa == 1 {
		fmt.Println("aa is 1")
	} else {
		fmt.Println("aa is not 1")
	}
	// aa is not 1
 
	if aaa := 0; aaa == 2 {
		fmt.Println("aaa is 2")
	} else if aaa == 3 {
		fmt.Println("aaa is 3")
	} else {
		fmt.Println("aaa is neither 2 nor 3")
	}
	// aaa is neither 2 nor 3
 
	if 1 == 3 {
 
	} else if 1 == 1 {
		fmt.Println("1 == 1")
	}
 
	var emap = map[int]bool{
		0: true,
		1: false,
	}
 
	if emap[0] || emap[1] {
		fmt.Println("true or false is true")
	}
	// true or false is true
 
	if emap[1] || emap[0] {
		fmt.Println("false or true is true")
	}
	// false or true is true
 
	if emap[0] || emap[0] {
		fmt.Println("true or true is true")
	}
	// true or true is true
 
	if emap[0] && emap[0] {
		fmt.Println("true and true is true")
	}
	// true and true is true
 
	if emap[0] && emap[1] {
 
	} else {
		fmt.Println("true and false is false")
	}
	// true and false is false
 
	if emap[0] || emap[100] {
		fmt.Println("emap[100] is not evaluated")
	}
	// emap[100] is not evaluated
 
	if v, ok := emap[0]; !ok {
 
	} else {
		fmt.Println("value is passed through the whole if-control structure:", v)
	}
	// value is passed through the whole if-control structure: true
 
	if v, ok := emap[100]; ok {
 
	} else {
		fmt.Println("if not exist, it returns a zero value:", ok, v)
	}
	// if not exist, it returns a zero value: false false
 
	// fmt.Println(v)
	// undefined: v
 
	if emap[100] || emap[0] {
		fmt.Println("emap[0] gets evaluated")
		fmt.Println("emap[100]:", emap[100])
	}
	// emap[0] gets evaluated
	// emap[100]: false
 
	var smap = map[int]string{
		0: "A",
	}
	if smap[100] || emap[0] {
 
	}
	// invalid operation: smap[100] || emap[0] (mismatched types string and bool)
}
```

[↑ top](#go-logic-loop)
<br><br><br><br><hr>


#### `switch`

Go has [*`switch`*](https://golang.org/doc/effective_go.html#switch) statement
for *if-else-if-else* patterns, as [here](http://play.golang.org/p/Gh6D0kDPtj):

```go
package main
 
import (
	"fmt"
	"log"
	"reflect"
)
 
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}
 
func main() {
	fmt.Println(shouldEscape([]byte("?")[0]))     // true
	fmt.Println(shouldEscape([]byte("abcd#")[4])) // true
	fmt.Println(shouldEscape([]byte("abcd#")[0])) // false
 
	num := 2
	switch num {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		panic("what's the number?")
	}
	// 2
 
	st := "b"
	switch {
	case st == "a":
		fmt.Println("a")
	case st == "b":
		fmt.Println("b")
	case st == "c":
		fmt.Println("c")
	default:
		panic("what's the character?")
	}
	// b
 
	ts := []interface{}{true, 1, 1.5, "A"}
	for _, t := range ts {
		eval(t)
	}
	/*
	   bool: true is bool
	   int: 1 is int
	   float64: 1.5 is float64
	   string: A is string
	*/
 
	type temp struct {
		a string
	}
	eval(interface{}(temp{}))
	// 2009/11/10 23:00:00 {} is main.temp
}
 
func eval(t interface{}) {
	switch typedValue := t.(type) {
	default:
		log.Fatalf("%v is %v", typedValue, reflect.TypeOf(typedValue))
	case bool:
		fmt.Println("bool:", typedValue, "is", reflect.TypeOf(typedValue))
	case int:
		fmt.Println("int:", typedValue, "is", reflect.TypeOf(typedValue))
	case float64:
		fmt.Println("float64:", typedValue, "is", reflect.TypeOf(typedValue))
	case string:
		fmt.Println("string:", typedValue, "is", reflect.TypeOf(typedValue))
	}
}
```

[↑ top](#go-logic-loop)
<br><br><br><br><hr>


#### `select`

[*`select`*](http://blog.golang.org/pipelines) is similar to *`switch`*: **`select`**
controls channels in Go, as [here](http://play.golang.org/p/9OwTUHX7iy):

```go
package main
 
import (
	"fmt"
	"time"
)
 
func send(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			if i == 5 {
				fmt.Println("Sleeping 2 seconds...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	return ch
}
 
func main() {
	ch := send("Hello")
	for {
		select {
		case v := <-ch:
			fmt.Println("Received:", v)
		case <-time.After(time.Second):
			fmt.Println("Done!")
			return
		}
	}
}
 
/*
Received: Hello 0
Received: Hello 1
Received: Hello 2
Received: Hello 3
Received: Hello 4
Received: Hello 5
Sleeping 2 seconds...
Done!
*/
```

[↑ top](#go-logic-loop)
<br><br><br><br><hr>


#### `for`

**_`for`_** is the only keyword for looping in Go (no `while`), as
[here](http://play.golang.org/p/5cVC_KKQEa):


```go
package main

import "fmt"

func main() {
	ts := []int{0, 1, 2}
	for i := 0; i < len(ts); i++ {
		fmt.Print(i, " ")
	}
	// 0 1 2
	fmt.Println()
	for i := len(ts) - 1; i > -1; i-- {
		fmt.Print(i, " ")
	}
	// 2 1 0
	fmt.Println()

	/*
		continue, break, goto, fallthrough

		continue: for-loop
		break: for-loop, switch
		goto: anything
		fallthrough: switch

		break, continue should only be in for-loop (break can be in switch)
		break, continue can be in if-statement
		 only when the if-statement is enclosed by for-loop

		Label can be used with goto, break, continue
	*/

	/**************************************/
	cn := 0
	for {
		cn++
		if cn == 5 {
			fmt.Println("cn is 5. Break(End) this loop!")
			break
		}

		if cn == 3 {
			fmt.Println("cn is 3. Continue this loop!")
			continue
		}
	}
	// cn is 3. Continue this loop!
	// cn is 5. Break(End) this loop!

	println()
	/**************************************/
	b := 0
	for i := 0; i < 10; i++ {
		b++
		if b == 500 {
			fmt.Println("b is 500. Break(End) this loop!")
			break
		}

		if b == 3 {
			fmt.Println("b is 3. Continue this loop!")
			continue
			fmt.Println("This does not print!")
		}
	}
	// b is 3. Continue this loop!
	// b does not reach 500, because before then
	// , we exit the for-loop condition

	println()
	/**************************************/
	goto here

	for {
		fmt.Println("Infinite Looping; Not Printing")
	}

here:
	fmt.Println("After goto")

	for i := 0; i < 2; i++ {
		for j := 0; j < 1000; j++ {
			if j == 3 {
				fmt.Println("Before Break, i == ", i)

				break
				fmt.Println("Not Printing")
			}
		}
		fmt.Println("Before continue, i == ", i)

		continue
		// go back to for i := 0;
		fmt.Println("Not Printing")
	}
	/*
	   After goto
	   Before Break, i ==  0
	   Before continue, i ==  0
	   Before Break, i ==  1
	   Before continue, i ==  1
	*/

	println()
	/**************************************/
	num := 0
Loop:
	for i := 0; i < 5; i++ {
		switch {
		case num < 3:
			fmt.Println("num is", num)
			num = num + 5
			break
			// this only breaks the enclosing switch statement
			// this does not break the for-loop

		case num > 4:
			fmt.Println("num > 4")
			break Loop
			// break the for-loop, which is labeled as "Loop"
			// this does not run this for-loop anymore

			// if we use goto Loop
			// it goes into infinite loop
			// because it starts over from for-loop
		default:
			fmt.Println("a")
		}
	}
	/*
	   num is 0
	   num > 4
	*/

	println()
	/**************************************/
	limit := 0
Cont:
	for {
		limit++
		if limit == 3 {
			fmt.Println("Break(End) this loop!")
			break
		}
		fmt.Print("..")
		continue Cont
		fmt.Println("This does not print!")
	}
	// ....Break(End) this loop!

	println()
	/**************************************/

	// "fallthrough" statement transfers control
	// to the first statement of the next case
	// clause in a expression "switch" statement.
	// It may be used only as the final non-empty
	// statement in such a clause.

	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		// To fall through to a subsequent case
		fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 1")
	}
	// 1 > 10

	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		// fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 2")
	}
	// [no output]

	bf := 2
	switch bf {
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("None 3")
	}
	// 3

	cf := 5
	switch cf {
	case 1:
		fmt.Println("1")
	case 2:
		// fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("None 4")
	}
	// None 4

	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		break
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 5")
	}
	// [no output]

	/*************************************
	  9 ways to use for-loop

	  Go has only one looping construct, the for loop.

	  The basic for loop looks as it does in C or Java
	  , except that we do not use ( )
	  (they are not even optional)
	   and the { } is required.
	  **************************************/
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// end when i == -1
	for i := 2; i >= 0; i-- {
		print(i, ",")
	}
	// 2,1,0,

	sum2 := 1
	for sum2 < 100 {
		sum2 += sum2
	}
	fmt.Println(sum2) // 128

	slice := []int{1, 2, 3, 4, 5}
	if a := 1; a > 0 {
		fmt.Println(a + slice[3])
	}
	// 5

	for b := 2; b != 5; {
		fmt.Print(slice[b])
		b++
	}
	// 345

	/**************************************/
	// infinite loop
	end := 0
	for {
		// infinite loop
		end++
		if end == 10 {
			fmt.Println("Ending(Breaking) for-loop!")
			break
		}
	}

	/**************************************/
	// string
	strk := "헬로우"
	for key, value := range strk {
		fmt.Println("String:", key, value)
	}
	/*
	   String: 0 54764
	   String: 3 47196
	   String: 6 50864
	*/

	for key, value := range strk {
		fmt.Printf("String: %v %c \n", key, value)
	}
	for key, value := range strk {
		fmt.Println("String:", key, string(value))
	}
	/*
	   String: 0 헬
	   String: 3 로
	   String: 6 우
	*/

	for _, value := range strk {
		fmt.Println("String Value:", string(value))
	}
	/*
	   String: 0 헬
	   String: 3 로
	   String: 6 우
	*/

	for key := range strk {
		fmt.Println("String Key:", key)
	}
	/*
	   String Key: 0
	   String Key: 3
	   String Key: 6
	*/

	/**************************************/
	// array
	for key, value := range [...]int{10, 20, 30} {
		fmt.Println("Array:", key, value)
	}
	/*
	   Array: 0 10
	   Array: 1 20
	   Array: 2 30
	*/

	for _, value := range [3]int{10, 20, 30} {
		fmt.Println("Array Value:", value)
	}
	/*
	   Array Value: 10
	   Array Value: 20
	   Array Value: 30
	*/

	for key := range [3]int{10, 20, 30} {
		fmt.Println("Array Key:", key)
	}
	/*
	   Array Key: 0
	   Array Key: 1
	   Array Key: 2
	*/

	/**************************************/
	// slice
	for key, value := range []string{"A", "B"} {
		fmt.Println("Slice:", key, value)
	}
	/*
	   Slice: 0 A
	   Slice: 1 B
	*/

	for _, value := range []string{"A", "B"} {
		fmt.Println("Slice Value:", value)
	}
	/*
	   Slice Value: A
	   Slice Value: B
	*/

	for key := range []string{"A", "B"} {
		fmt.Println("Slice Key:", key)
	}
	/*
	   Slice Key: 0
	   Slice Key: 1
	*/

	/**************************************/
	// map (Unordered Data Structure)
	// map's key values are like index, but not integer
	// array is also map with Hash Function
	// with integers as indices
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
	}
	for key, value := range elements {
		fmt.Println("Map:", key, value)
	}
	/*
	   Map: H Hydrogen
	   Map: He Helium
	   Map: Li Lithium
	*/

	for _, value := range elements {
		fmt.Println("Map Value:", value)
	}
	/*
	   Map Value: Hydrogen
	   Map Value: Helium
	   Map Value: Lithium
	*/

	for key := range elements {
		fmt.Println("Map Key:", key)
	}
	/*
	   Map Key: H
	   Map Key: He
	   Map Key: Li
	*/

	value1, exist1 := elements["B"]
	fmt.Println(value1, exist1)
	//  false

	value2, exist2 := elements["He"]
	fmt.Println(value2, exist2)
	// Helium true
}
```

[↑ top](#go-logic-loop)
<br><br><br><br><hr>


#### review `goto`
[Here](http://play.golang.org/p/5GJE0Dpkb6):

```go
package main
 
import "fmt"
 
func main() {
	goto Here
 
	for {
		fmt.Println("Infinite Looping; Not Printing")
	}
 
Escape:
	fmt.Println("After goto Escape.")
	return
 
Here:
	fmt.Println("After goto Here.")
	for i := 0; i < 2; i++ {
		fmt.Println("Hello")
		goto Escape
	}
}
 
/*
After goto Here.
Hello
After goto Escape.
*/
```

[↑ top](#go-logic-loop)
<br><br><br><br><hr>


#### review `switch`, `break`

[Here](http://play.golang.org/p/GgheijxtIj):

```go
package main
 
import "fmt"
 
func main() {
	num := 0
Loop:
	for i := 0; i < 5; i++ {
		switch {
		case num < 3:
			fmt.Println("num is", num)
			num = num + 5
			break
			// this only breaks the enclosing switch statement
			// this does not break the for-loop
 
		case num > 4:
			fmt.Println("num > 4")
			break Loop
			// break the for-loop, which is labeled as "Loop"
			// this does not run this for-loop anymore
 
			// if we use goto Loop
			// it goes into infinite loop
			// because it starts over from for-loop
		default:
			fmt.Println("a")
		}
	}
}
 
/*
num is 0
num > 4
*/
```

[↑ top](#go-logic-loop)
<br><br><br><br><hr>


#### review `fallthrough`

[Here](http://play.golang.org/p/p_QIKattOL):

```go
package main
 
import "fmt"
 
func main() {
 
	// "fallthrough" statement transfers control
	// to the first statement of the next case
	// clause in a expression "switch" statement.
	// It may be used only as the final non-empty
	// statement in such a clause.
 
	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		// To fall through to a subsequent case
		fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 1")
	}
	// 1 > 10
 
	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		// fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 2")
	}
	// [no output]
 
	bf := 2
	switch bf {
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("None 3")
	}
	// 3
 
	cf := 5
	switch cf {
	case 1:
		fmt.Println("1")
	case 2:
		// fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("None 4")
	}
	// None 4
 
	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		break
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 5")
	}
	// [no output]
}
```

[↑ top](#go-logic-loop)
<br><br><br><br><hr>


#### empty `struct`

Try [this](http://play.golang.org/p/4B9GnIy-FX)
and read:

- [Empty struct by Dave Cheney](http://dave.cheney.net/2014/03/25/the-empty-struct)


```go
package main

import (
	"fmt"
	"unsafe"
)

func N(n int) []struct{} {
	return make([]struct{}, n)
}

func main() {
	var s0 struct{}
	fmt.Println(unsafe.Sizeof(s0)) // 0

	s1 := struct{}{}
	fmt.Println(unsafe.Sizeof(s1)) // 0

	done := make(chan struct{})
	go func() {
		done <- struct{}{}
	}()
	<-done
	fmt.Println("Done")
	// Done

	for i := range N(10) {
		fmt.Print(i, " ")
	}
	// 0 1 2 3 4 5 6 7 8 9
}

/*
Same in Python:

for i in range(10):
    print i
*/
```

[↑ top](#go-logic-loop)
<br><br><br><br><hr>


#### fizzbuzz

> Write a program that prints the numbers from 1 to 100. But for **multiples of
> three print “Fizz”** instead of the number and for the **multiples of**
> **_five_** **print “Buzz”. For numbers which are multiples of both**
> **_three_** and **_five_** **print “FizzBuzz”**.
>
> [**Fizz Buzz Test**](http://c2.com/cgi/wiki?FizzBuzzTest)

You can use **switch**, as [here](http://play.golang.org/p/a1AyeMommb):

```go
package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		switch {
		case i%15 == 0:
			fmt.Println(i, "FizzBuzz")
		case i%3 == 0:
			fmt.Println(i, "Fizz")
		case i%5 == 0:
			fmt.Println(i, "Buzz")
		default:
			fmt.Println(i)
		}
	}
	for i := 1; i < 101; i++ {
		if i%15 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

/*
1
2
3 Fizz
4
5 Buzz
6 Fizz
7
8
9 Fizz
10 Buzz
11
12 Fizz
13
14
15 FizzBuzz
16
...
*/

```

[↑ top](#go-logic-loop)
<br><br><br><br><hr>
