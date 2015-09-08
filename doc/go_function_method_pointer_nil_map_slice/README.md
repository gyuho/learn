[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: function, method, pointer, nil, map, slice

- [What are function and method?](#what-are-function-and-method)
- [What is pointer?](#what-is-pointer)
- [Why pointer?](#why-pointer)
- [Where pointer?](#where-pointer)
- [Why `list` as a pointer?](#why-list-as-a-pointer)
- [pointer: copy `struct`](#pointer-copy-struct)
- [pointer: `map` and `slice`](#pointer-map-and-slice)
- [Recap](#recap)
- [swap: `array` and `slice`](#swap-array-and-slice)
- [`nil`](#nil)
- [initialize and empty `map`](#initialize-and-empty-map)
- [*non-deterministic* `range` `map`](#non-deterministic-range-map)
- [slice tricks](#slice-tricks)
- [permute](#example-permute)
- [destroy `struct`](#destroy-struct)
- [tree](#tree)

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>








#### What are function and method?
Think of a **function** as **an operation(on arguments if any)**. And **method as a
behavior of a type**. A **method** is *a function that takes a receiver*, and always
*used in conjunction with types*. Try this
[code](http://play.golang.org/p/tm52WsTj60):

```go
package main
 
import "fmt"
 
func myFunc(num int) {
	fmt.Println(num + 1)
}
 
type Int int
 
func (num Int) myMethod() {
	fmt.Println(num + 2)
}
 
func main() {
	myFunc(1)         // 2
	Int(1).myMethod() // 3
}
```

<br>

Go supports **_function types_**, **_function values_**, and **_function
closures_**. Let's take a look at this
[code](http://play.golang.org/p/vS15bbZZ8e):

```go
package main
 
import (
	"fmt"
	"reflect"
	"runtime"
)
 
func main() {
	// function values
	fmt.Println(
		// function closure (function literal)
		func() {
			fmt.Println("Hello")
		},
	) // 0x20280
 
	// without _ =
	// we get func literal evaluated but not used
	//
	_ = func() {
		fmt.Println("Hello 1")
	}
	// No output
 
	func(str string) {
		fmt.Println(str)
	}("Hello 2")
	// Hello 2
 
	fn := func() {
		fmt.Println("Hello")
	}
	fmt.Println(fn)                                                      // 0x203a0
	fmt.Println(reflect.TypeOf(fn))                                      // func()
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()) // main.func·004
 
	fn() // Hello
 
	fc := func(num int) int {
		num += 1
		return num
	}
	fmt.Println(fc)                 // 0x20720
	fmt.Println(reflect.TypeOf(fc)) // func(int) int
	fmt.Println(fc(1))              // 2
 
	po := plusOne(1)
	fmt.Println(po)                 // 0x209e0
	fmt.Println(reflect.TypeOf(po)) // main.funcType
	fmt.Println(po(3))              // 4
}
 
// function type
// FunctionType   = "func" Signature .
type funcType func(int) int
 
func plusOne(num int) funcType {
	return func(num int) int {
		return num + 1
	}
}
```

A **function name** in the function signature must be **unique**: [**there is
no function overloading**](https://golang.org/ref/spec#Signature) in Go. And a
function name can be omitted in a function literal or an anonymous
function — this is useful in concurrency and enables a functional programming
style in a strongly typed language.


With **functions having values and types**, you can do something like
[here](http://play.golang.org/p/psW7cO37yn):

```go
package main
 
import (
	"fmt"
	"reflect"
)
 
func plusOne(num int) int {
	return num + 1
}
 
func plusTwo(num int) int {
	return num + 2
}
 
func plusThree(num int) int {
	return num + 3
}
 
var funcMap = map[string]func(num int) int{
	"one":   plusOne,
	"two":   plusTwo,
	"three": plusThree,
}
 
func add(num int, functions ...func(num int) int) int {
	for _, oneFunc := range functions {
		num = oneFunc(num)
	}
	return num
}
 
func main() {
	chosen := funcMap["two"]
	fmt.Println(reflect.TypeOf(chosen)) // func(int) int
	fmt.Println(chosen(0))              // 2
 
	fmt.Println(
		add(
			0,
			funcMap["one"],
			funcMap["two"],
			funcMap["three"],
		),
	)
	// 6
}
```

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>






#### What is pointer?

Take a look at
[`container/list`](http://golang.org/pkg/container/list/)
package with *pointer*:

```go
type Element struct {
	next, prev *Element
	list       *List
	Value      interface{}
}
```

*`next`* and *`prev`* Element are defined as pointer _`*Element`_.
- Can we instead use Element *value*?
- What is *pointer*?
- How does Go handle *pointer*s?

<br>
[**Pointer**](https://en.wikipedia.org/wiki/Pointer_%28computer_programming%29)
is a value that refers to another value in memory (it’s an address value). When
you **reference** a variable, you get the **location** of the value in memory.
When you **dereference** a pointer, you get the **value stored** in that
address. In Go, you would:

```go
oneVar := 1
reference := &oneVar
dereference = *reference
fmt.Printf("%p\n", reference) // 0x1043617c
fmt.Printf("%v", dereference) // 1
```

[Go FAQ](http://golang.org/doc/faq#Pointers) explains that **_every function
parameter in Go is passed by value_**: a function receives the copied data from
arguments, not the original object. Suppose you pass a variable of an integer
value, a function gets the copied value of the integer variable. Updating the
passed variable inside the function wouldn’t have any effect on the original
data in function arguments. **_When you pass a pointer variable, a function
gets the copy of the pointer._** And pointer is not the data itself but
**pointer points to original data it refers to. Therefore, changing
dereferenced data by its pointer inside the function can update the original data.**

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>








#### Why pointer?

That's because **_every function parameter in Go is passed by value._** If you
**want to manipulate objects** with **methods** or **functions**, passing a
copied value won't take any effect on the original object: you **_must pass a
pointer to the original data_**, which can be used for direct access. **The
pointer itself will be copied as well, but the pointer is the address that
still points to the original data.** [Go
FAQ](http://golang.org/doc/faq#Pointers) explains if the **receiver is large**,
a *big struct* for instance, it's **cheaper to use a pointer receiver**, than
copying the whole data. Look at this
[code](http://play.golang.org/p/66_KBKuOYH):

```go
package main
 
import "fmt"
 
type myType struct{ val int }
 
func noChange(t myType) {
	t.val = 100
	// this only changes the copied
}
 
func (t myType) noChange() {
	t.val = 200
	// this only changes the copied
}
 
func change(t *myType) {
	t.val = 300
	// this updates the original
	// that the pointer t points to
}
 
func (t *myType) change() {
	t.val = 400
	// this updates the original
	// that the pointer t points to
}
 
func main() {
	one := myType{val: 1}
	noChange(one)
	fmt.Printf("%+v\n", one) // {val:1}
	one.noChange()
	fmt.Printf("%+v\n", one) // {val:1}
 
	change(&one)
	fmt.Printf("%+v\n", one) // {val:300}
	(&one).change()
	fmt.Printf("%+v\n", one) // {val:400}
}
```

A *function* or *method* updates the original data **only when we pass the
pointer to the original data.** Otherwise, only the copied value gets passed.
Again, updates on a copied value do not affect the original data, but updates
through a pointer change the original. **_Pointer is about shared access._**
**If you want to share the value between functions and methods, then use a
pointer. If you don’t need to share, then use a value then** **_copy_**. This
only **matters** when the **function parameters or method receivers get
copied.**. You don't need to worry about it when you have a [direct
access](http://play.golang.org/p/wIp9clV4TO) to the data:

```go
package main
 
import "fmt"
 
type myType struct{ val int }
 
func main() {
	one := myType{val: 1}
 
	// direct access
	one.val = 100
	fmt.Printf("%+v\n", one) // {val:100}
 
	// copy the value
	value := one
	value.val = 200
	fmt.Printf("%+v\n", one) // {val:100}
 
	// access to the original data through pointer
	pointer := &one
	pointer.val = 300
	fmt.Printf("%+v\n", one) // {val:300}
}
```

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>









#### Where pointer?
Go, as a [garbage collected
language](https://en.wikipedia.org/wiki/Garbage_collection_%28computer_science%29),
handles memory allocation for you. Where is pointer? We [do not need to
worry](http://golang.org/doc/faq#stack_or_heap) about it. To be precise:

> there is no way to know this[stack or heap], since the **compiler is free to
> move things from stack to heap and vice versa**
>
> [**_Rob
> Pike_**](https://groups.google.com/forum/#!msg/golang-nuts/PH_pMZqvLN8/xd12p5x5qqUJ)

It depends on the compiler’s [escape
analysis](http://en.wikipedia.org/wiki/Escape_analysis): **_compilers store
variables in stack when they are local to the function (local variables)._**
**If the analysis cannot prove them to be local, Go stores them in
garbage-collected heap.** Look at this
[code](http://play.golang.org/p/7aiPdY60ez):

```go
package main
 
import "fmt"
 
type myType struct{ val int }
 
func a() *myType {
	one := myType{val: 1}
	return &one
}
 
func main() {
	one := a()
	fmt.Printf("%+v\n", one) // &{val:1}
}
```

In C, you get warnings when returning the address of a local variable. The
local variable gets assigned an address(pointer) but once the function exits,
the memory gets deallocated, thus the pointer to this memory becomes invalid
(**dangling pointer**). **But in Go, this is [totally
OK](http://golang.org/doc/effective_go.html#composite_literals) and the data is
still** **_refer-able_** **after function return.**.

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>











#### Why `list` as a pointer?


And back to the original question:
```go
type Element struct {
	next, prev *Element
	list       *List
	Value      interface{}
}
```
Why is it defined as a pointer? That's because [linked
list](http://en.wikipedia.org/wiki/Linked_list) implementation needs to
manipulate the original moving nodes around. If it were defined with
values(*non-pointers*), we can't *insert* or *remove* elements in the linked
list, as seen in this [example](http://play.golang.org/p/P4mEFTs0ZU):

```go
package main
 
import "fmt"
 
type List struct{ root Element }
type Element struct{ val int }
 
func change(l List) { l.root.val = 100 }
 
func main() {
	l := List{}
	l.root = Element{val: 1}
	l.root.val = 2
	fmt.Printf("%+v\n", l) // {root:{val:2}}
	// this updates because we are not passing the copy
	// it's not in the function or method
 
	change(l)
	fmt.Printf("%+v", l) // {root:{val:2}}
	// passing the non-pointer
	// and only the copied data is passed
	// so it can't update the original value
}
```

As you see, without pointer, we cannot change the original data in function or
with methods. Define with# pointer to globally pass things around and to update
it anywhere.

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>








#### pointer: copy `struct`

Try [this](http://play.golang.org/p/yKL4dEnlaH):

```go
package main

import "fmt"

type Data struct {
	value string
}

func main() {
	d1 := Data{}
	d1.value = "A"
	// just copy the value
	d2 := d1
	d2.value = "B"

	fmt.Println(d1) // {A}
	fmt.Println(d2) // {B}

	d3 := &Data{}
	d3.value = "A"
	// copy the pointer of the original
	d4 := d3
	d4.value = "B"

	fmt.Println(d3) // &{B}
	fmt.Println(d4) // &{B}
}
```

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>







#### pointer: `map` and `slice`

Go [FAQ](https://golang.org/doc/faq#pass_by_value) explains: **maps and slices
are** [**_references_** *in Go*](https://golang.org/doc/faq#references) (slice
as reference, array as value). **Map** and **slice values** behave like
**pointers**. Let's take a look at **map** first with this
[code](http://play.golang.org/p/U3A2uOGkip):

```go
package main
 
import "fmt"
 
var mmap = map[int]string{1: "A"}
 
func changeMap1(m map[int]string) { m[100] = "B" }
 
func changeMap1pt(m *map[int]string) { (*m)[100] = "C" }
 
type mapType map[int]string
 
func (m mapType) changeMap2() { m[100] = "D" }
 
func (m *mapType) changeMap2pt() { (*m)[100] = "E" }
 
func main() {
	// (O) change
	changeMap1(mmap)
	fmt.Println(mmap) // map[1:A 100:B]
 
	// (O) change
	changeMap1pt(&mmap)
	fmt.Println(mmap) // map[1:A 100:C]
 
	// (O) change
	mapType(mmap).changeMap2()
	fmt.Println(mmap) // map[1:A 100:D]
 
	// (O) change
	(*mapType)(&mmap).changeMap2pt()
	fmt.Println(mmap) // map[100:E 1:A]
}
```

**The original map still gets changed** from the function and method, **_even
when_** we **DO NOT** *pass* **pointers** *of maps*. This is because
internally
[makemap](https://github.com/golang/go/blob/master/src/runtime/hashmap.go#L187)
function in [runtime](http://golang.org/pkg/runtime/) **initializes a map** and
returns **_hmap_** **pointer**:

```go
func makemap(
	t *maptype,
	hint int64,
	h *hmap,
	bucket unsafe.Pointer,
) *hmap {
	...
```

<br>

**Array** and **slice** are more subtle than map. Likewise, slice contains
pointer but in a different way. **A slice contains a reference to its backing
array**, as
[here](https://github.com/golang/go/blob/master/src%2Fruntime%2Fslice.go):

```go
package runtime
 
import (
	"unsafe"
)
 
type sliceStruct struct {
	array unsafe.Pointer
	len   int
	cap   int
}
 
// TODO: take uintptrs instead of int64s?
func makeslice(t *slicetype, len64 int64, cap64 int64) sliceStruct {
	// NOTE: The len > MaxMem/elemsize check here is not strictly necessary,
	// but it produces a 'len out of range' error instead of a 'cap out of range' error
	// when someone does make([]T, bignumber). 'cap out of range' is true too,
	// but since the cap is only being supplied implicitly, saying len is clearer.
	// See issue 4085.
	len := int(len64)
	if len64 < 0 || int64(len) != len64 || t.elem.size > 0 && uintptr(len) > _MaxMem/uintptr(t.elem.size) {
		panic(errorString("makeslice: len out of range"))
	}
	cap := int(cap64)
	if cap < len || int64(cap) != cap64 || t.elem.size > 0 && uintptr(cap) > _MaxMem/uintptr(t.elem.size) {
		panic(errorString("makeslice: cap out of range"))
	}
	p := newarray(t.elem, uintptr(cap))
	return sliceStruct{p, len, cap}
}
```


<br>
And take a look at this [code](http://play.golang.org/p/g8y8_BGVKo) and
[code](http://play.golang.org/p/T9nqUmp624):

```go
package main
 
import "fmt"
 
var array = [3]string{"A", "B", "C"}
 
func changeArray1(m [3]string) { m[0] = "X" }
 
func changeArray1pt(m *[3]string) { (*m)[0] = "Y" }
 
type arrayType [3]string
 
func (m arrayType) changeArray2() { m[1] = "XX" }
 
func (m *arrayType) changeArray2p()  { m[1] = "YY" }
func (m *arrayType) changeArray2pt() { (*m)[1] = "ZZ" }
 
func main() {
	// (X) no change
	changeArray1(array)
	fmt.Println("changeArray1:", array) // [A B C]
 
	// (O) change
	changeArray1pt(&array)
	fmt.Println("changeArray1pt:", array) // [Y B C]
 
	// (X) no change
	arrayType(array).changeArray2()
	fmt.Println(".changeArray2():", array) // [Y B C]
 
	// (O) change
	(*arrayType)(&array).changeArray2p()
	fmt.Println(".changeArray2pt():", array) // [Y YY C]
	
	// (O) change
	(*arrayType)(&array).changeArray2pt()
	fmt.Println(".changeArray2pt():", array) // [Y ZZ C]
}
```

```go
package main
 
import "fmt"
 
var slice = []string{"A", "B", "C"}
 
func changeSlice1(m []string) { m[0] = "X" }
 
func changeSlice1pt(m *[]string) { (*m)[0] = "Y" }
 
type sliceType []string
 
// var slice = sliceType{"A", "B", "C"}
 
func (m sliceType) changeSlice2() { m[1] = "XX" }
 
// func (m *sliceType) changeSlice2p() { m[1] = "YY" }
 
func (m *sliceType) changeSlice2pt() { (*m)[1] = "YY" }
 
func changeSlice3(m []string) { m = append(m, "XXX") }
 
func changeSlice3pt(m *[]string) { *m = append(*m, "YYY") }
 
func (m sliceType) changeSlice4() { m = append(m, "XXXX") }
 
func (m *sliceType) changeSlice4pt() { *m = append(*m, "YYYY") }
 
func main() {
	// (O) change
	changeSlice1(slice)
	fmt.Println("changeSlice1:", slice) // [X B C]
 
	// (O) change
	changeSlice1pt(&slice)
	fmt.Println("changeSlice1pt:", slice) // [Y B C]
 
	// (O) change
	sliceType(slice).changeSlice2()
	fmt.Println(".changeSlice2():", slice) // [Y XX C]
 
	// (O) change
	(*sliceType)(&slice).changeSlice2pt()
	fmt.Println(".changeSlice2pt():", slice) // [Y YY C]
 
	// (X) no change
	changeSlice3(slice)
	fmt.Println("changeSlice3:", slice) // [Y YY C]
 
	// (O) change
	changeSlice3pt(&slice)
	fmt.Println("changeSlice3pt:", slice) // [Y YY C YYY]
 
	// (X) no change
	sliceType(slice).changeSlice4()
	fmt.Println(".changeSlice4():", slice) // [Y YY C YYY]
 
	// (O) change
	(*sliceType)(&slice).changeSlice4pt()
	fmt.Println(".changeSlice4pt():", slice) // [Y YY C YYY YYYY]
}
```

According to [Go FAQ](https://golang.org/doc/faq#references), **_slice is a
reference_**. Then **_why do we still have to pass the pointer in order to
update a slice?_** [*Arrays, slices (and string) by Rob
Pike*](https://blog.golang.org/slices) and [*Go Data Structures by Russ
Cox*](http://research.swtch.com/godata)
explain in more detail. In short, **a slice in Go is not an array. A slice just
represents a piece in an array.**

> A [slice](http://golang.org/doc/effective_go.html#slices) is a reference to a
> section of an array.
>
> [**_Russ Cox_**](http://research.swtch.com/godata)

Think of a **slice** as a struct typed data of an **element in an array** (Go
slice implementation is
[here](https://go.googlesource.com/go/+/master/src/runtime/slice.go)):

![slice](img/slice.png)

This is why array elements are values and slice elements are references.

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>








#### Recap

```go
func changeArray1(m [3]string) { m[0] = “X” }  // (X) no change
func changeSlice1(m []string)  { m[0] = “X” }  // (O) change
```
**changeArray1** DOES NOT update its original array(value) but **changeSlice1**
updates its original slice.

<br>

```go
func changeArray1pt(m *[3]string) { (*m)[0] = “Y” }  // (O) change
func changeSlice1pt(m *[]string)  { (*m)[0] = “Y” }  // (O) change
```
Both **changeArray1pt** and **changeSlice1pt** update their original array and
slice.

<br>

```go
func (m arrayType) changeArray2() { m[1] = “XX” }  // (X) no change
func (m sliceType) changeSlice2() { m[1] = "XX" }  // (O) change
```

**changeArray2** DOES NOT update its original array(value) but **changeSlice2**
updates its original slice.

<br>

```go
func (m *arrayType) changeArray2p()  { m[1] = “YY” }
// (O) change

func (m *arrayType) changeArray2pt() { (*m)[1] = “YY” }
// (O) change

func (m *sliceType) changeSlice2pt() { (*m)[1] = "YY" }
// (O) change
```

**changeArray2p**, **changeArray2pt** and **changeSlice2pt** update their
original array and slice.

<br>

```go
func changeSlice3(m []string)     { m = append(m, “XXX”) }
func (m sliceType) changeSlice4() { m = append(m, “XXXX”)}
// (X) no change

func changeSlice3pt(m *[]string)     { *m = append(*m, “YYY”) }
func (m *sliceType) changeSlice4pt() { *m = append(*m, "YYYY")}
// (O) change
```

**changeSlice3** and **changeSlice4** DO NOT update their original slices but
**changeSlice3pt** and **changeSlice4pt** update the original slices.

To conclude:
- In functions and methods, an array or its element is passed as a value.
- Therefore, each element in the array is also a value.
- You must pass the pointer to an array to update the original array.
- Array elements are values and slice elements are references.
- In functions and methods, a **slice** is passed as a **value**.
- You must pass the **_pointer_** to update the original slice (**_append_**).
- In functions and methods, slice elements are references.
- No need to pass a pointer to access(index) the original slice elements.

That is, slice data structure contains a pointer to an original array, but
slice itself is not a pointer: it is a **struct data** that contains a pointer.
Therefore, **slice is passed by value**. Since slice includes a pointer to
elements of an array, **slice values** act like a pointer. They are
[descriptors pointing](https://golang.org/doc/faq#Pointers) to the underlying
slice data.


[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>









#### swap: `array` and `slice`

Try [this](http://play.golang.org/p/3yzDHS15ey):


```go
package main

import "fmt"

func main() {
	// direct access doesn't need any pointer
	array := [5]int{0, 1, 2, 3, 4}
	for i := 0; i < len(array); i++ {
		array[i]++ // DOES CHANGE
	}
	fmt.Println(array) // [1 2 3 4 5]

	// array needs pointer to update its element
	swapArray1(2, 3, array) // NO CHANGE
	fmt.Println(array)      // [1 2 3 4 5]

	swapArray2(2, 3, &array) // DOES CHANGE
	fmt.Println(array)       // [1 2 3 4 5]

	// direct access doesn't need any pointer
	slice := []int{0, 1, 2, 3, 4}
	for i := 0; i < len(slice); i++ {
		slice[i]++ // DOES CHANGE
	}
	fmt.Println(slice) // [1 2 3 4 5]

	// slice elements are pointers
	swapSlice1(2, 3, slice) // DOES CHANGE
	fmt.Println(slice)      // [1 2 4 3 5]

	swapSlice2(2, 3, &slice) // DOES CHANGE
	fmt.Println(slice)       // [1 2 3 4 5]
}

func swapArray1(i, j int, array [5]int) {
	array[i], array[j] = array[j], array[i]
}

func swapArray2(i, j int, array *[5]int) {
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
}

func swapSlice1(i, j int, slice []int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func swapSlice2(i, j int, slice *[]int) {
	(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
}
```

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>









#### `nil`

Since **maps are slices are references** in Go, we can assign
[**nil**](https://golang.org/doc/faq#nil_error) **to them**. To sum, `nil` can
be used for:

- [*`error`*](http://golang.org/pkg/builtin/#error) interface
- interface
- `struct` pointer
- `byte` slice (bytes, but not for `string`)
- `map`
- slice


Try this [code](http://play.golang.org/p/8iX7DFKSKA) below:


```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Root *Node
}

func New(root *Node) *Data {
	d := &Data{}
	d.Root = root
	return d
}

type Interface interface {
	Less(than Interface) bool
}

type Node struct {
	Left  *Node
	Key   Interface
	Right *Node
}

func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	return nd
}

func (d *Data) Insert(nd *Node) {
	if d.Root == nd {
		return
	}
	d.Root = d.Root.insert(nd)
}

func (nd *Node) insert(node *Node) *Node {
	if nd == nil {
		return node
	}
	if nd.Key.Less(node.Key) {
		nd.Right = nd.Right.insert(node)
	} else {
		nd.Left = nd.Left.insert(node)
	}
	return nd
}

func (d Data) Search(key Interface) *Node {
	nd := d.Root
	for nd != nil {
		if nd.Key == nil {
			break
		}
		switch {
		case nd.Key.Less(key):
			nd = nd.Right
		case key.Less(nd.Key):
			nd = nd.Left
		default:
			return nd
		}
	}
	return nil
}

type Int int

func (n Int) Less(b Interface) bool {
	return n < b.(Int)
}

func main() {
	// nil for error(interface)
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	root := NewNode(Int(5))
	data := New(root)
	data.Insert(NewNode(Int(7)))

	// nil for interface
	fmt.Println(NewNode(Int(7)).Key == nil) // false

	// nil for struct pointer
	root = nil
	fmt.Println(root == nil) // true

	// nil for bytes (not for string)
	b := []byte("abc")
	fmt.Println(b, string(b)) // [97 98 99] abc
	b = nil
	fmt.Println(b, string(b)) // []
	//
	// str := "abc"
	// str = nil (x) value cannot be nil

	// nil for map
	mmap := make(map[string]bool)
	mmap["A"] = true
	fmt.Println(mmap) // map[A:true]
	mmap = nil
	fmt.Println(mmap) // map[]

	// nil for slice (not for array)
	slice := []int{1}
	fmt.Println(slice) // [1]
	slice = nil
	fmt.Println(slice) // []
	//
	// array := [1]int{1}
	// array = nil
	// (X) cannot use nil as type [1]int in assignment
}
```

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>









#### initialize and empty `map`

Be careful. Following two lines are **different**:

```go
mmap1 = make(map[string]int)
// reassign the map and initializes it to empty map

// (O)
mmap1["A"] = 1


mmap2 = nil
// reassign the nil to map pointer
// and initializes to nil map

// (X)
mmap2["A"] = 1
```

You assign `nil` only when you **_nullify_** the whole map pointer.
If you need to empty(*initialize*) an existing map, you must use `make`
to reassign an empty map, as [here](http://play.golang.org/p/k1DHm_lpIB):

```go
package main

import "fmt"

func main() {
	mmap1 := map[string]int{
		"hello": 10,
	}
	mmap1 = make(map[string]int)
	mmap1["A"] = 1
	fmt.Println(mmap1)
	// map[A:1]

	mmap2 := map[string]int{
		"hello": 10,
	}
	mmap2 = nil
	mmap2["A"] = 1
	fmt.Println(mmap2)
	// panic: assignment to entry in nil map
}
```


And in order to **update a map**, you can either pass a **pointer or the original
map**. Buf if you want to **initialize (with assignment), you must pass
pointer**, as seen [below](http://play.golang.org/p/LOAS1lmlzy):

```go
package main

import (
	"fmt"
	"math/rand"
)

// You can either pass the pointer of map or just map to update.
// But if you want to initialize with assignment, you have to pass pointer.

func updateMap1(m map[int]bool) {
	for {
		num := rand.Intn(150)
		if _, ok := m[num]; !ok {
			m[num] = true
		}
		if len(m) == 5 {
			return
		}
	}
}

func initializeMap1(m map[int]bool) {
	m = nil
	m = make(map[int]bool)
}

type mapType map[int]bool

func (m mapType) updateMap1() {
	m[0] = false
	m[1] = false
}

func (m mapType) initializeMap1() {
	m = nil
	m = make(map[int]bool)
}

func updateMap2(m *map[int]bool) {
	for {
		num := rand.Intn(150)
		if _, ok := (*m)[num]; !ok {
			(*m)[num] = true
		}
		if len(*m) == 5 {
			return
		}
	}
}

func initializeMap2(m *map[int]bool) {
	// *m = nil
	*m = make(map[int]bool)
}

func (m *mapType) updateMap2() {
	(*m)[0] = false
	(*m)[1] = false
}

func (m *mapType) initializeMap2() {
	// *m = nil
	*m = make(map[int]bool)
}

func main() {
	m0 := make(map[int]bool)
	m0[1] = true
	m0[2] = true
	fmt.Println("Done:", m0) // Done: map[1:true 2:true]

	m0 = make(map[int]bool)
	fmt.Println("After:", m0) // After: map[]

	m1 := make(map[int]bool)
	updateMap1(m1)
	fmt.Println("updateMap1:", m1)
	// (o) change
	// updateMap1: map[131:true 87:true 47:true 59:true 31:true]

	initializeMap1(m1)
	fmt.Println("initializeMap1:", m1)
	// (X) no change
	// initializeMap1: map[59:true 31:true 131:true 87:true 47:true]

	mapType(m1).updateMap1()
	fmt.Println("mapType(m1).updateMap1():", m1)
	// (o) change
	// mapType(m1).updateMap1(): map[87:true 47:true 59:true 31:true 0:false 1:false 131:true]

	mapType(m1).initializeMap1()
	fmt.Println("mapType(m1).initializeMap1():", m1)
	// (X) no change
	// mapType(m1).initializeMap1(): map[59:true 31:true 0:false 1:false 131:true 87:true 47:true]

	m2 := make(map[int]bool)
	updateMap2(&m2)
	fmt.Println("updateMap2:", m2)
	// (o) change
	// updateMap2: map[140:true 106:true 0:true 18:true 25:true]

	initializeMap2(&m2)
	fmt.Println("initializeMap2:", m2)
	// (o) change
	// initializeMap2: map[]

	(*mapType)(&m2).updateMap2()
	fmt.Println("(*mapType)(&m2).updateMap2:", m2)
	// (o) change
	// (*mapType)(&m2).updateMap2: map[0:false 1:false]

	(*mapType)(&m2).initializeMap2()
	fmt.Println("(*mapType)(&m2).initializeMap2:", m2)
	// (o) change
	// (*mapType)(&m2).initializeMap2: map[]
}
```

**_All function calls in Go are pass-by-value: values are copied, so the
assignment to the copy only changes the value of the copy, not the origin._**
Go map and slice are values that refer to other values like pointers.

Then why you need pointer for initialization?

```go
func update(m map[int]bool) {
    m[1] = true
}

func initialize(m *map[int]bool) {
    *m = nil
}
```

*initialize* works the same as:

```go
func updateNum(num *int) {
    *num = 10
}
```

**updateNum** function updates the variable pointed by *num* with the value 10.
*num* still points to the same variable but with a new value 10. We cannot
change the address,# where the pointer variable *num* is pointing to, since the
function passes the copy of the pointer. Changing the copy of the address does
not do anything on the original address. This is true of the map as well.

> Like slices, **maps hold references to an underlying data structure**. If you
> pass a **map** to a **function that changes the contents of the map**, the
> changes will be **visible** in the caller.
>
> [**_Effective Go_**](https://golang.org/doc/effective_go.html)

<br>

Since a map itself refers to the underlying data structure, it works like a
pointer. It just does not explicitly have the pointer notation `*`:

```go
func updateMap(m **OriginalMapData) {
    newMap := OriginalMapData
    *m = &newMap
}
```

<br>

More detailed discussion can be found
[here](https://groups.google.com/forum/#!msg/golang-nuts/xzdPCjKORNA/hLO-Sl7mtJkJ):

> The key difference is that when you **append to a slice, you are
> potentially modifying the slice itself**, not just what it holds. **Slices
> are a contiguous view onto an array**, so appending to them might
> require reallocating and copying, which does not affect any other copy
> of that slice that might be looking at the original array. **A map is
> inherently non-contiguous, and you never need to reallocate the map
> itself, only stuff that’s buried inside it, and that means that a
> change to a map is reflected in every copy of that map.
>
> That is# why you must pass a slice as a pointer if you want the
> function to modify your copy of it, but you don't have to do that for
> a map.**
>
> [*David
> Symonds*](https://groups.google.com/d/msg/golang-nuts/xzdPCjKORNA/7qd0BEklVqAJ)


[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>









#### *non-deterministic* `range` `map`

Try [this](http://play.golang.org/p/CgDbiFTJDX):

```go
package main

import (
	"fmt"
	"strings"
)

func nonDeterministicMapUpdateV1() {
	for i := 0; i < 10; i++ {
		fmt.Println("nonDeterministicMapUpdateV1 TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		length := len(mmap)
		for k, v := range mmap {
			mmap[strings.ToUpper(k)] = v * v
			delete(mmap, k)
		}
		if length == len(mmap) {
			fmt.Println("Luckily, Deterministic with nonDeterministicMapUpdateV1:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with nonDeterministicMapUpdateV1:", length, len(mmap))
	}
}

func nonDeterministicMapUpdateV2() {
	for i := 0; i < 10; i++ {
		fmt.Println("nonDeterministicMapUpdateV2 TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		ks := []string{}
		length := len(mmap)
		for k, v := range mmap {
			mmap[strings.ToUpper(k)] = v * v
			ks = append(ks, k)
		}
		for _, k := range ks {
			delete(mmap, k)
		}
		if length == len(mmap) {
			fmt.Println("Luckily, Deterministic with nonDeterministicMapUpdateV2:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with nonDeterministicMapUpdateV2:", length, len(mmap))
	}
}

func nonDeterministicMapUpdateV3() {
	for i := 0; i < 10; i++ {
		fmt.Println("nonDeterministicMapUpdateV3 TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		length := len(mmap)
		for k := range mmap {
			v := mmap[k]
			mmap[strings.ToUpper(k)] = v * v
			delete(mmap, k)
		}
		if length == len(mmap) {
			fmt.Println("Luckily, Deterministic with nonDeterministicMapUpdateV3:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with nonDeterministicMapUpdateV3:", length, len(mmap))
	}
}

func deterministicMapSet() {
	for i := 0; i < 10000; i++ {
		mmap := make(map[int]bool)
		for i := 0; i < 10000; i++ {
			mmap[i] = true
		}
		length := len(mmap)
		for k := range mmap {
			delete(mmap, k)
		}
		if len(mmap) == 0 {
			fmt.Println("Deterministic with deterministicMapSet:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with deterministicMapSet:", length, len(mmap))
	}
}

func deterministicMapDelete() {
	for i := 0; i < 10000; i++ {
		fmt.Println("deterministicMapDelete TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		length := len(mmap)
		for k := range mmap {
			delete(mmap, k)
		}
		if len(mmap) == 0 {
			fmt.Println("Deterministic with deterministicMapDelete:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with deterministicMapDelete:", length, len(mmap))
	}
}

func deterministicMapUpdate() {
	for i := 0; i < 10000; i++ {
		fmt.Println("deterministicMapUpdate TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		mmapCopy := make(map[string]int)
		length := len(mmap)
		for k, v := range mmap {
			mmapCopy[strings.ToUpper(k)] = v * v
		}
		for k := range mmap {
			delete(mmap, k)
		}
		if length == len(mmapCopy) || len(mmap) != 0 {
			fmt.Println("Deterministic with deterministicMapUpdate:", length, len(mmapCopy))
			return
		} else {
			mmapCopy = make(map[string]int) // to initialize(empty)
			//
			// (X)
			// mmapCopy = nil
		}
		fmt.Println("Non-Deterministic with deterministicMapUpdate:", length, len(mmap))
	}
}

func main() {
	nonDeterministicMapUpdateV1()
	fmt.Println()
	nonDeterministicMapUpdateV2()
	fmt.Println()
	nonDeterministicMapUpdateV3()

	fmt.Println()

	deterministicMapSet()
	fmt.Println()
	deterministicMapDelete()
	fmt.Println()
	deterministicMapUpdate()
}

/*
These are all non-deterministic.
If you are lucky, the map gets updated inside range.

nonDeterministicMapUpdateV1 TRY = 0
Non-Deterministic with nonDeterministicMapUpdateV1: 5 4
nonDeterministicMapUpdateV1 TRY = 1
Non-Deterministic with nonDeterministicMapUpdateV1: 5 4
nonDeterministicMapUpdateV1 TRY = 2
Luckily, Deterministic with nonDeterministicMapUpdateV1: 5 5

nonDeterministicMapUpdateV2 TRY = 0
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 1
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 2
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 3
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 4
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 5
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 6
Non-Deterministic with nonDeterministicMapUpdateV2: 5 3
nonDeterministicMapUpdateV2 TRY = 7
Non-Deterministic with nonDeterministicMapUpdateV2: 5 4
nonDeterministicMapUpdateV2 TRY = 8
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 9
Non-Deterministic with nonDeterministicMapUpdateV2: 5 4

nonDeterministicMapUpdateV3 TRY = 0
Non-Deterministic with nonDeterministicMapUpdateV3: 5 4
nonDeterministicMapUpdateV3 TRY = 1
Non-Deterministic with nonDeterministicMapUpdateV3: 5 4
nonDeterministicMapUpdateV3 TRY = 2
Non-Deterministic with nonDeterministicMapUpdateV3: 5 4
nonDeterministicMapUpdateV3 TRY = 3
Luckily, Deterministic with nonDeterministicMapUpdateV3: 5 5

Deterministic with deterministicMapSet: 10000 0

deterministicMapDelete TRY = 0
Deterministic with deterministicMapDelete: 5 0

deterministicMapUpdate TRY = 0
Deterministic with deterministicMapUpdate: 5 5
*/
```

<br>
This tells that `map` is:
- **_Non-deterministic_** on **update** when it *updates* with `for range` of the map. 
- **_Deterministic_** on **`update`** when it *deletes* with `for range`, **NOT** on the map. 
	- `for i := 0; i < 10000; i++ {mmap[i] = true}`
- **_Deterministic_** on **`delete`** when it *deletes* with `for range` of the map. 
- **_Deterministic_** on **`update`** when it *updates* with `for range` of the **COPIED** map. 


<br>
[Go FAQ](http://golang.org/doc/faq#atomic_maps) explains:

> Why are map operations not defined to be atomic?
>
> After long discussion it was decided that the typical use of maps did not
> require safe access from multiple goroutines, and in those cases where it did,
> the map was probably part of some larger data structure or computation that was
> already synchronized. Therefore requiring that all map operations grab a mutex
> would slow down most programs and add safety to few. This was not an easy
> decision, however, since it means uncontrolled map access can crash the
> program.
>
> The language does not preclude atomic map updates. When required, such as when
> hosting an untrusted program, the implementation could interlock map access.
>
> [*Go FAQ*](http://golang.org/doc/faq#atomic_maps)


And about `for` loop:

> The iteration order over maps is not specified and is not guaranteed to be
> the same from one iteration to the next. If map entries that have not yet
> been reached are removed during iteration, the corresponding iteration values
> will not be produced. If map entries are **created during iteration**, that entry
> may be produced during the iteration or **may be skipped**. The choice may vary
> for each entry created and from one iteration to the next. If the map is nil,
> the number of iterations is 0.
>
> [Go Spec](https://golang.org/ref/spec#For_statements)


[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>








#### slice tricks

We should use slice instead `container/list`! Here's
[why](http://www.reddit.com/r/golang/comments/25oxg0/three_reasons_you_should_not_use_martini/chkvkym)
from [*David Symonds*](https://github.com/dsymonds):

![linked_list_reddit](img/linked_list_reddit.png)

Here's a list of [slice tricks](http://play.golang.org/p/uPpzivsWxj):

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Copy Slice
	slice01 := []int{1, 2, 3, 4, 5}
	copy01 := make([]int, len(slice01))
	copy(copy01, slice01)
	fmt.Println(copy01) // [1 2 3 4 5]

	// PushFront
	slice02 := []int{1, 2, 3, 4, 5}
	copy02 := make([]int, len(slice02)+1)
	copy02[0] = 10
	copy(copy02[1:], slice02)
	fmt.Println(copy02) // [10 1 2 3 4 5]

	// PushFront
	pushFront := func(s *[]int, elem int) {
		temp := make([]int, len(*s)+1)
		temp[0] = elem
		copy(temp[1:], *s)
		*s = temp
	}
	pushFront(&copy02, 100)
	fmt.Println(copy02) // [100 10 1 2 3 4 5]

	// PushBack
	slice03 := []int{1, 2, 3, 4, 5}
	slice03 = append(slice03, 10)
	fmt.Println(slice03) // [1 2 3 4 5 10]

	// PopFront
	slice04 := []int{1, 2, 3, 4, 5}
	slice04 = slice04[1:len(slice04):len(slice04)]
	fmt.Println(slice04, len(slice04), cap(slice04)) // [2 3 4 5] 4 4

	// PopBack
	slice05 := []int{1, 2, 3, 4, 5}
	slice05 = slice05[:len(slice05)-1 : len(slice05)-1]
	fmt.Println(slice05, len(slice05), cap(slice05)) // [1 2 3 4] 4 4

	// Delete
	slice06 := []int{1, 2, 3, 4, 5}
	copy(slice06[3:], slice06[4:])
	slice06 = slice06[:len(slice06)-1 : len(slice06)-1]
	// copy(d.OutEdges[edge1.Vtx][idx:], d.OutEdges[edge1.Vtx][idx+1:])
	// d.OutEdges[src][len(d.OutEdges[src])-1] = nil // zero value of type or nil
	fmt.Println(slice06, len(slice06), cap(slice06)) // [1 2 3 5] 4 4

	make2DSlice := func(row, column int) [][]string {
		mat := make([][]string, row)
		// for i := 0; i < row; i++ {
		for i := range mat {
			mat[i] = make([]string, column)
		}
		return mat
	}
	mat := make2DSlice(3, 5)
	for key, value := range mat {
		fmt.Println(key, value)
	}
	/*
	   0 [    ]
	   1 [    ]
	   2 [    ]
	*/
	fmt.Println(mat[1], len(mat[1]), cap(mat[1])) // [    ] 5 5

	// iterate over rows
	for r := range mat {
		// iterate over columns
		for c := range mat[r] {
			mat[r][c] = strconv.Itoa(r) + "x" + strconv.Itoa(c)
		}
	}
	for key, value := range mat {
		fmt.Println(key, value)
	}
	/*
	   0 [0x0 0x1 0x2 0x3 0x4]
	   1 [1x0 1x1 1x2 1x3 1x4]
	   2 [2x0 2x1 2x2 2x3 2x4]
	*/
	fmt.Println(mat[1], len(mat[1]), cap(mat[1])) // [1x0 1x1 1x2 1x3 1x4] 5 5
}
```

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>













#### permute

Now what if you want to permute a slice? [Actual
algorithm](http://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order)
is a bit more complicated. Simple algorithm would write as
[follows](http://play.golang.org/p/v2qL6xR4k7):

```go
package main
 
import (
	"fmt"
	"sort"
)
 
func main() {
	slice := []int{3, 2, 1}
	slices := [][]int{}
	sort.Ints(slice)
	slices = append(slices, slice)
	fmt.Println(slices) // [[1 2 3]]
 
	slice[0], slice[1] = slice[1], slice[0]
	slices = append(slices, slice)
	fmt.Println(slices) // [[2 1 3] [2 1 3]]
 
	slice[1], slice[2] = slice[2], slice[1]
	slices = append(slices, slice)
	fmt.Println(slices) // [[2 3 1] [2 3 1] [2 3 1]]
}
```

You should have expected something like: `[[1 2 3] [2 1 3] [2 3 1]]` but all
elements came out to be same, NOT the permuted list of original slice. Instead
[this](http://play.golang.org/p/w9iJccsWIV) works as expected:

```go
package main
 
import (
	"fmt"
	"sort"
)
 
func main() {
	slice := []int{3, 2, 1}
	slices := [][]int{}
	sort.Ints(slice)
	copied0 := make([]int, len(slice))
	copy(copied0, slice)
	slices = append(slices, copied0)
	fmt.Println(slices) // [[1 2 3]]
 
	slice[0], slice[1] = slice[1], slice[0]
	copied1 := make([]int, len(slice))
	copy(copied1, slice)
	slices = append(slices, copied1)
	fmt.Println(slices) // [[1 2 3] [2 1 3]]
 
	slice[1], slice[2] = slice[2], slice[1]
	copied2 := make([]int, len(slice))
	copy(copied2, slice)
	slices = append(slices, copied2)
	fmt.Println(slices) // [[1 2 3] [2 1 3] [2 3 1]]
}
```

**A slice is just a pointer to the underlying backing array. Without copying
the values of the original slice, you are just duplicating the slice headers
that point to the same backing array.**

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>







#### destroy `struct`

This is how you would `nil`ify or change an object in Go,
[here](http://play.golang.org/p/tvCXqufyMa):

```go
package main

import "fmt"

type Node struct {
	Key string
}

func noDestroy1(nd Node) {
	nd = Node{"destroyed"}
}

func noDestroy2(nd *Node) {
	nd = nil
}

func noDestroy3(nd *Node) {
	nd = &Node{"destroyed"}
}

func destroy1(nd *Node) {
	nd.Key = "destroyed"
}

func destroy2(nd *Node) {
	*nd = Node{}
}

func main() {
	nd := new(Node)
	nd.Key = "not destroyed"

	noDestroy1(*nd)
	fmt.Println("after noDestroy1:", nd)
	// after noDestroy1: &{not destroyed}

	noDestroy2(nd)
	fmt.Println("after noDestroy2:", nd)
	// after noDestroy2: &{not destroyed}

	// nd = nil    would destroy
	// but in the function, it only
	// updates the address

	noDestroy3(nd)
	fmt.Println("after noDestroy3:", nd)
	// after noDestroy3: &{not destroyed}

	destroy1(nd)
	fmt.Println("after destroy1:", nd)
	// after destroy1: &{destroyed}

	destroy2(nd)
	fmt.Println("after destroy2:", nd)
	// after destroy2: &{}
}

```

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>










#### tree

Note that *even if you pass a pointer of a struct data to a function*,
**_the function takes only the copied value of the pointer(address)_**. And
**you need to dereference to update the original data** or **access directly
before you call any function**, as 
[here](http://play.golang.org/p/w0Erst9vuo):


```go
package main

import "fmt"

type Data struct {
	Root *Node
}

type Node struct {
	Left  *Node
	Key   int
	Right *Node
}

func (tr *Tree) String() string {
	return d.Root.String()
}

func (nd *Node) String() string {
	if nd == nil {
		return "[]"
	}
	s := ""
	if nd.Left != nil {
		s += nd.Left.String() + " "
	}
	s += fmt.Sprintf("%v", nd.Key)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

func main() {
	tree := new(Tree)
	root := new(Node)
	root.Key = 3
	tree.Root = root

	left := new(Node)
	left.Key = 1
	root.Left = left

	right := new(Node)
	right.Key = 5
	root.Right = right

	seven := new(Node)
	seven.Key = 7
	root.Right.Right = seven
	/*
	     3
	    / \
	   1   5
	        \
	         7
	*/

	fmt.Println("tree:", tree) // [[1] 3 [5 [7]]]

	delete1(root)
	fmt.Println("delete1(root) root:", root) // delete1(root) root: [[1] 3 [5 [7]]]
	fmt.Println("delete1(root) tree:", tree) // delete1(root) tree: [[1] 3 [5 [7]]]

	// but if we access directly
	// you can update
	root.Right = new(Node)
	fmt.Println("root.Right = new(Node) tree:", tree)
	// root.Right = new(Node) tree: [[1] 3 [0]]

	delete2(root)
	fmt.Println("delete2(root) root:", root) // delete2(root) root: [0]
	fmt.Println("delete2(root) tree:", tree) // delete2(root) tree: [0]
}

func delete1(nd *Node) {
	// even if nd is a pointer
	// it ONLY copies the value
	// of the tree address
	nd = nil
	// or
	// nd = new(Node)

	// setting nil to the address
	// does not affect the original object
}

func delete2(nd *Node) {
	*nd = Node{}
}
```

[↑ top](#go-function-method-pointer-nil-map-slice)
<br><br><br><br>
<hr>
