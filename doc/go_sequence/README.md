[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: sequence

- [mutable `bytes` and `rune`, immutable `string`](#mutable-bytes-and-rune-immutable-string)
- [interface](#interface)
- [pointer](#pointer)
- [`container/list`, linked list](#containerlist-linked-list)
	- [doubly linked list implementation](#doubly-linked-list-implementation)
	- [singly linked list implementation](#singly-linked-list-implementation)
- [array](#array)
- [`slice`, **slice tricks**](#slice-slice-tricks)
- [thread-safe, generic **`slice`**](#thread-safe-generic-slice)
- [`slice` vs. `container/list`](#slice-vs-containerlist)
- [thread-safe, generic **set**](#thread-safe-generic-set)

[↑ top](#go-sequence)
<br><br><br><br>
<hr>







#### mutable `bytes` and `rune`, immutable `string`

Code consists of a sequence of **bits**—*a value of 0 or 1*. 
*One* **_8-bit_** *chunk* builds *one* **_byte_** that represents **_one character_**.
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

[↑ top](#go-sequence)
<br><br><br><br>
<hr>












#### interface

> **Note too that the elimination of the type hierarchy also eliminates a form
> of dependency hierarchy. Interface satisfaction allows the program to grow
> organically without predetermined contracts.** And it is a linear form of
> growth; a change to an interface affects only the immediate clients of that
> interface; there is no subtree to update. **The lack of implements
> declarations disturbs some people but it enables programs to grow naturally,
> gracefully, and safely.**
>
> [**_Rob Pike_**](https://talks.golang.org/2012/splash.article)

Interface is a set of methods (not functions).
[Polymorphism](http://en.wikipedia.org/wiki/Polymorphism_%28computer_science%29)
can be done via interfaces. Any type that implements those set of methods
automatically satisfies the interface. **Interface is satisfied implicitly**:
you don’t have to specify that *a data type A implements interface B.*
Therefore, we can define our own interface methods for the code that we don’t
own. This has a huge impact on program design and encourages to write
compatible code.

<br>

> One important category of type is interface types, which represent fixed sets
> of methods. **An interface variable can store any concrete (non-interface)
> value as long as that value implements the `interface`’s methods.**
>
> [**_Laws of Reflection by Rob Pike_**](http://blog.golang.org/laws-of-reflection)

[↑ top](#go-sequence)
<br><br><br><br>
<hr>









#### pointer

Take a look at 
[`container/list`](https://go.googlesource.com/go/+/master/src/container/list/list.go):

```go
type Element struct {
   next, prev *Element
   list       *List
   Value      interface{}
}
```

*`next`* and *`prev`* Element is defined as a pointer type __`*Element`__ Can we
instead use value `Element`? No, it's because [linked
list](http://en.wikipedia.org/wiki/Linked_list) needs to manipulate the
original *list* while moving nodes around. It it were defined as a
value(*non-pointer*), we can't insert or remove elements in the linked list, as
[here](http://play.golang.org/p/P4mEFTs0ZU):

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
with methods. Define with pointer to globally pass things around and to update
it anywhere.


[↑ top](#go-sequence)
<br><br><br><br>
<hr>









#### `container/list`, linked list

**Linked list** is similar to an **array** in that they both store the
collection of data in **sequence**. **Array** (or *slice*) is a list of items
that are **contiguously allocated in memory space**. An array allocates memory
for all items in one block of memory. **Linked list** is a *group of nodes*, each
of which is **connected** by **pointers**. A node should contain **_data_** and
**_reference (pointer)_** to its *next* or *previous* node. **A linked list
allocates memory for each node, separately in its own memory space.** A **_linked
list_** is useful when you do lots of *insertions* and *removals*:
**searching** can be *inefficient* because you need to iterate from beginning
whatever item you try to find.

<br>
**_Slice_** is like:

![slice](img/slice.png)

<br>
**_Linked lists_** are like:

![linked_list](img/linked_list.png)

<br>
Go [**`container/list`**](http://golang.org/pkg/container/list/)
implements the **doubly linked list**. To simplify its implementation, Go
`list` is implemented as if it were a *ring*. The **_root node (`Element`)_**
is both **_previous element of the first node_** and **_next element of the
last node_**, as [here](http://play.golang.org/p/NKzqyoYt47):

```go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	func() {
		myList := list.New()
		myList.PushBack(1)
		myList.PushBack(2)
		myList.PushBack(3)
		for e := myList.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value, " ")
		}
		// 1 2 3

		fmt.Println()
		myList.InsertAfter(50, myList.PushFront(10))
		for e := myList.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value, " ")
		}
		// 10 50 1 2 3

		fmt.Println()
		myList.Remove(myList.Front())
		for e := myList.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value, " ")
		}
		// 50 1 2 3
	}()

	func() {
		myList := list.New()
		for i := 0; i < 10; i++ {
			myList.PushBack(i)
		}

		// when deleting, the list gets changed too
		// we need to declare next element outside
		var next *list.Element
		for elem := myList.Front(); elem != nil; elem = next {
			next = elem.Next()
			myList.Remove(myList.Front())
		}
		fmt.Println()
		fmt.Println(myList.Len()) // 0
	}()
}

```


[↑ top](#go-sequence)
<br><br><br><br>
<hr>










#### doubly linked list implementation

Go [**`container/list`**](http://golang.org/pkg/container/list/)
implements the **doubly linked list**, as [here](http://play.golang.org/p/MYkHhhVmaF):

```go
// http://golang.org/pkg/container/list
// package "container/list" is for a doubly-linked list
// https://code.google.com/p/go/source/browse/#hg%2Fsrc%2Fpkg%2Fcontainer%2Flist

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package list implements a doubly linked list.
//
// To iterate over a list (where l is a *List):
//	for e := l.Front(); e != nil; e = e.Next() {
//		// do something with e.Value
//	}
//
package main

import "fmt"

type List struct {
	// when a function exits, all of its variables are popped off of the stack
	// if we want to update outside of this, we need pointer
	// use pointers to access memory on the heap
	// Heap variables are essentially global in scope
	root *Element
	len  int
}

// Element is an element of a linked list.
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element
	list       *List
	Value      int
}

func (l *List) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e

	e.list = l
	l.len++
	return e
}

func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev

	e.next = nil
	e.prev = nil

	e.list = nil
	l.len--

	return e
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != e.list.root {
		return p
	}
	return nil
}

// Init initializes or clears list l.
func (l *List) Init() *List {

	// references to data structures that must be initialized before use.
	//
	// never dereference a nil pointer
	// root will never be nil with this
	l.root = new(Element) // do not need this if 'root Element'

	l.root.next = l.root
	l.root.prev = l.root
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) insertValue(v int, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

func (l *List) InsertAfter(v int, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

func (l *List) Remove(e *Element) int {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

func (l *List) PushFront(v int) *Element {
	l.lazyInit()
	// return l.insertValue(v, &l.root)
	return l.insertValue(v, l.root)
}

func (l *List) PushBack(v int) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

func main() {
	list1 := New()
	list1.PushBack(1)
	list1.PushBack(2)
	list1.PushBack(3)
	for e := list1.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	// 1 2 3

	fmt.Println()
	list1.InsertAfter(50, list1.PushFront(10))
	for e := list1.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	// 10 50 1 2 3

	fmt.Println()
	list1.Remove(list1.Front())
	for e := list1.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	// 50 1 2 3

	list2 := New()
	for i := 0; i < 10; i++ {
		list2.PushBack(i)
	}

	// when deleting, the list gets changed too
	// we need to declare next element outside
	var next *Element
	for elem := list2.Front(); elem != nil; elem = next {
		next = elem.Next()
		list2.Remove(list2.Front())
	}

	fmt.Println()
	fmt.Println(list2.Len()) // 0
}

func (l *List) InsertBefore(v int, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.insert(l.remove(e), l.root)
}

func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.insert(l.remove(e), l.root.prev)
}

func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || e == mark {
		return
	}
	l.insert(l.remove(e), mark)
}

func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, l.root)
	}
}
```

[↑ top](#go-sequence)
<br><br><br><br>
<hr>










#### singly linked list implementation

Try [this](http://play.golang.org/p/g7Vjk-yAa-):

```go
package main

import "fmt"

type List struct {
	root *Element
	tail *Element
	len  int
}

type Element struct {
	next  *Element
	list  *List
	Value int
}

func (l *List) insert(e, at *Element) *Element {
	// add first time
	if l.len == 0 {

		l.root.next = e
		e.next = l.tail

	} else if at == l.tail {

		// push back
		e.next = l.tail

		// update the previous element of tail
		atPrev := l.root
		for p := l.Front(); p != at; p = p.Next() {
			atPrev = p
		}
		atPrev.next = e

	} else {

		// push front or between
		n := at.next
		at.next = e
		e.next = n

	}

	e.list = l
	l.len++
	return e
}

func (l *List) remove(e *Element) *Element {
	if e == l.root || e == l.tail {
		return nil
	}
	ePrev := l.root
	for p := l.Front(); p != e; p = p.Next() {
		ePrev = p
	}
	n := e.next
	ePrev.next = n
	e.next = nil
	e.list = nil
	l.len--
	return e
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != e.list.root {
		return p
	}
	return nil
}

func (l *List) Init() *List {
	l.root = new(Element)
	l.tail = new(Element)
	l.root.next = l.tail
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) insertValue(v int, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

func (l *List) InsertAfter(v int, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

func (l *List) PushFront(v int) *Element {
	return l.insertValue(v, l.root)
}

func (l *List) PushBack(v int) *Element {
	return l.insertValue(v, l.tail)
}

func (l *List) Remove(e *Element) int {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

func main() {
	list1 := New()
	list1.PushBack(1)
	list1.PushBack(2)
	list1.PushBack(3)
	list1.PushFront(5)
	for e := list1.Front(); e != list1.tail; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	// 5 1 2 3

	fmt.Println()
	list1.InsertAfter(50, list1.PushFront(10))
	for e := list1.Front(); e != list1.tail; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	// 10 50 5 1 2 3

	fmt.Println()
	list1.Remove(list1.Front())
	for e := list1.Front(); e != list1.tail; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	// 50 5 1 2 3

	fmt.Println()
	list2 := reverseList(list1)
	for e := list2.Front(); e != list2.tail; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	// 3 2 1 5 50

	a := New()
	for i := 0; i < 10; i++ {
		a.PushBack(i)
	}
	var next *Element
	for elem := a.Front(); elem != a.tail; elem = next {
		next = elem.Next()
		a.Remove(a.Front())
	}

	fmt.Println()
	fmt.Println(a.Len()) // 0
}

func reverseList(l *List) *List {
	tempList := New()
	for e := l.Front(); e != l.tail; e = e.Next() {
		tempList.PushFront(e.Value)
	}
	return tempList
}
```

[↑ top](#go-sequence)
<br><br><br><br>
<hr>









#### array

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

[↑ top](#go-sequence)
<br><br><br><br>
<hr>









#### `slice`, **slice tricks**

We should use slice instead of `container/list`! Here's
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

[↑ top](#go-sequence)
<br><br><br><br>
<hr>










#### thread-safe, generic **`slice`**

[Code](http://play.golang.org/p/cj0E_CJRgT):

```go
package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	func() {
		dt1 := NewData()
		dt1.Init()
		dt2 := NewData()
		if !reflect.DeepEqual(dt1, dt2) {
			fmt.Errorf("%#v %#v", dt1, dt2)
		}
	}()

	func() {
		dt := NewData()
		dt.Insert([]interface{}{1, "A", 3, -.9, "B"}...)
		if dt.GetSize() != 5 {
			fmt.Errorf("Should be '5': %#v", dt)
		}
	}()

	func() {
		d := NewData()
		d.Insert([]interface{}{1, "A", 3, -.9, "B"}...)
		d.Init()
		if d.GetSize() != 0 {
			fmt.Errorf("Should be '0': %#v", d)
		}
		dt := NewData()
		dt.PushBack(1)
		dt.PushBack(2)
		dt.PushBack(3)
		dt.Insert(3, 1, 7, 11)
		isempty1 := dt.IsEmpty()
		dt.Init()
		isempty2 := dt.IsEmpty()
		if isempty1 != false && isempty2 != true {
			fmt.Errorf("Should return 'false' and 'true': %v, %v", isempty1, isempty2)
		}
	}()

	func() {
		dt := NewData()
		if !dt.IsEmpty() {
			fmt.Errorf("Should return 'true': %#v", dt)
		}
	}()

	func() {
		dt := NewData()
		dt.Insert(1, 3, 4, 5, "A", "7", 100)
		if dt.GetSize() != 7 {
			fmt.Errorf("Should return 7: %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, 3, 4, 5, "A", "7", 100)
		if dt.GetSize() != 7 {
			fmt.Errorf("Should return 7: %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		c := dt.Clone()
		if dt.GetSize() != c.GetSize() {
			fmt.Errorf("Should return true but %#v / %#v", dt, c)
		}
		if !dt.IsDeepEqual(*c) || !dt.IsSemiEqual(*c) {
			fmt.Errorf("Should return true but %+v %+v", dt, c)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		idx, ok := dt.Contains("3")
		if ok {
			fmt.Errorf("Should return false but %+v %+v", idx, ok)
		}
		idx, ok = dt.Contains(3)
		if !ok {
			fmt.Errorf("Should return true but %+v %+v", idx, ok)
		}
		a, b := dt.Contains("A")
		if a != 1 && b != true {
			fmt.Errorf("Should return '1, true': %#v", dt)
		}
		c, d := dt.Contains(-.8)
		if c != 0 && d != false {
			fmt.Errorf("Should return '0, false': %#v", dt)
		}
	}()

	func() {
		s1 := CreateData(1, "A", 3, -.9, "B")
		s1c := CreateData()
		s1c.Insert(1, "A", 3, -.9, "B")
		s2 := CreateData(3, -.9, "B", 1, "A")
		s3 := CreateData()
		s3.Insert(-.9, "B", 1, 3, "A")
		if !s1.IsDeepEqual(*s1c) {
			fmt.Errorf("Should return true but %+v %+v", s1, s1c)
		}
		if s1.IsDeepEqual(*s2) {
			fmt.Errorf("Should return false but %+v %+v", s1, s2)
		}
		if !s1.IsSemiEqual(*s3) {
			fmt.Errorf("Should return true but %+v %+v", s1, s3)
		}
		if !s2.IsSemiEqual(*s3) {
			fmt.Errorf("Should return true but %+v %+v", s2, s3)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		dt.PushFront("Front")
		if dt.GetFront() != "Front" {
			fmt.Errorf("dt.GetFront() should be 'Front': %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		dt.PushBack("Back")
		if dt.GetBack() != "Back" {
			fmt.Errorf("dt.GetBack() should be 'Back': %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		dt.DeepSliceDelete(2)
		if dt.m[2] != -0.9 {
			fmt.Errorf("dt.m[2] should be '-0.9': %#v", dt)
		}
		for _ = range dt.m {
			dt.DeepSliceDelete(0)
			// Don't do dt.DeepSliceDelete(k)
			// the slice length decreases at the same time
		}
		if dt.GetSize() != 0 {
			fmt.Errorf("Should be empty but: %#v", dt.GetSize())
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		ok := dt.FindAndDelete(3)
		if !ok || dt.m[2] != -0.9 {
			fmt.Errorf("Should return true, but %#v, and dt.m[2] should be '-0.9': %#v", ok, dt)
		}

		// list := dt
		// this does deep copy
		// (they are the same)
		// so we need to use Copy
		list := dt.Clone()
		for _, v := range list.m {
			dt.FindAndDelete(v)
		}
		l := dt.GetSize()
		if l != 0 {
			fmt.Errorf("Should be empty but: %#v / %#v / %#v", l, dt, list)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		dt.DeepSliceCut(2, 3)
		if dt.m[2] != "B" {
			fmt.Errorf("dt[2] should be 'B': %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		tm := dt.GetFront()
		if tm != 1 {
			fmt.Errorf("dt.GetFront() should be 1: %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		tm := dt.GetBack()
		if tm != "B" {
			fmt.Errorf("dt.GetBack() should be 'B': %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		tm := dt.PopFront()
		if tm != 1 {
			fmt.Errorf("dt.PopFront() should return 1: %#v", dt)
		}
		if dt.m[0] != "A" {
			fmt.Errorf("dt[0] should be 'A': %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		tm := dt.PopBack()
		if tm != "B" {
			fmt.Errorf("dt.PopBack() should return 'B': %#v", dt)
		}
		if dt.m[3] != -0.9 {
			fmt.Errorf("dt[3] should be -0.9: %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		slice := dt.GetElements()
		if len(slice) != 5 {
			fmt.Errorf("len(slice) should return 5: %#v", dt)
		}
		if slice[3] != -0.9 {
			fmt.Errorf("slice[3] should be -0.9: %#v", dt)
		}
	}()

	func() {
		s1 := CreateData(1, "A", 3, -.9, "B", "e", "f", "G")
		s2 := CreateData(1, "A", 3, -.9, "B")
		s3 := CreateData(1, "A", 3, -.9, "H", 2, 3, 4)
		s4 := CreateData(1, "A", 3, -.9, "H", 2, 3, 4)
		s5 := CreateData(1, "A", 3, -.9, "B", "e", "f")
		result := CommonPrefix(s1, s2, s3, s4, s5)
		if len(result) != 4 {
			fmt.Errorf("len(result) should return 4: %#v", result)
		}
		if result[3] != -0.9 {
			fmt.Errorf("result[3] should be -0.9: %#v", result)
		}
	}()
}

// Data can contain any type of values,
// because its data is a slice of interface{} type.
// It is an empty interface, which means that it can
// be satisfied by any type of value.
type Data struct {
	m []interface{}

	// RWMutex is more expensive
	// https://blogs.oracle.com/roch/entry/beware_of_the_performance_of
	// sync.RWMutex
	//
	// to synchronize access to shared state across multiple goroutines.
	//
	sync.Mutex
}

// NewData returns a new object of Data.
func NewData() *Data {
	nslice := []interface{}{}
	return &Data{
		m: nslice,
	}
}

// GetSize returns the GetSizegth of sequence.
// If the method needs to mutate the receiver,
// the receiver must be a pointer.
// (http://golang.org/doc/faq#methods_on_values_or_pointers)
func (d Data) GetSize() int {
	d.Lock()
	size := len(d.m)
	d.Unlock()
	return size
}

// Init initializes the Data.
func (d *Data) Init() {
	// (X) d = NewData()
	// This only updates its pointer
	// , not the Data itself
	//
	*d = *NewData()
}

// IsEmpty returns true if the sequence is empty.
func (d Data) IsEmpty() bool {
	return d.GetSize() == 0
}

// Insert appends values to Data.
func (d *Data) Insert(vals ...interface{}) {
	d.Lock()
	for _, elem := range vals {
		d.m = append(d.m, elem)
	}
	d.Unlock()
}

// CreateData instantiates a set object with initial elements.
func CreateData(vals ...interface{}) *Data {
	data := NewData()
	data.Insert(vals...)
	return data
}

// Clone returns a copy of the sequence.
// This is useful because ":=" operator
// does deep copy and when we manipulate
// either one, then the other one also changed.
func (d Data) Clone() *Data {
	td := NewData()
	for _, v := range d.m {
		td.PushBack(v)
	}
	return td
}

// Contains returns true if the elem exists
// in the Data.
func (d Data) Contains(elem interface{}) (int, bool) {
	d.Lock()
	defer d.Unlock()
	for idx, val := range d.m {
		if reflect.DeepEqual(val, elem) {
			return idx, true
		}
	}
	return 0, false
}

// IsSemiEqual returns true if s1 is equal to s2
// regardless of its ordering of elementd.
func (d Data) IsSemiEqual(a Data) bool {
	if d.GetSize() != a.GetSize() {
		return false
	}
	for _, val := range a.m {
		_, ok := d.Contains(val)
		if !ok {
			return false
		}
	}
	return true
}

// IsDeepEqual returns true if s1 is equal to s2.
func (d Data) IsDeepEqual(a Data) bool {
	if d.GetSize() != a.GetSize() {
		return false
	}
	return reflect.DeepEqual(d.m, a.m)
}

// PushFront adds an element to the front of sequence.
func (d *Data) PushFront(val interface{}) {
	size := d.GetSize()
	ts := make([]interface{}, size+1)
	ts[0] = val
	d.Lock()
	copy(ts[1:], d.m)
	d.m = ts
	d.Unlock()
}

// PushBack adds an element to the back of sequence.
func (d *Data) PushBack(val interface{}) {
	d.Lock()
	d.m = append(d.m, val)
	d.Unlock()
}

// DeepSliceDelete deletes the element in the index.
func (d *Data) DeepSliceDelete(idx int) {
	size := d.GetSize()
	d.Lock()
	copy(d.m[idx:], d.m[idx+1:])
	d.m[size-1] = nil // zero value of type or nil
	d.m = d.m[:size-1 : size-1]
	d.Unlock()
}

// FindAndDelete finds the element and delete it.
func (d *Data) FindAndDelete(val interface{}) bool {
	idx, ok := d.Contains(val)
	if !ok {
		return false
	}
	d.DeepSliceDelete(idx)
	return true
}

// DeepSliceCut deletes the elements from indices a to b.
func (d *Data) DeepSliceCut(a, b int) {
	if b > d.GetSize()-1 || a < 0 || a > b {
		panic("Index out of range! You can cut only inside slice.")
	}
	diff := b - a + 1
	idx := a
	i := 0
	for i < diff {
		d.DeepSliceDelete(idx)
		i++
	}
}

// GetFront returns the first(front) element of sequence.
func (d Data) GetFront() interface{} {
	if d.GetSize() == 0 {
		return nil
	}
	return d.m[0]
}

// GetBack returns the last(back) element of sequence.
func (d Data) GetBack() interface{} {
	if d.GetSize() == 0 {
		return nil
	}
	return d.m[d.GetSize()-1]
}

// PopFront removes the front(first) element of sequence
// and return it at the same time.
func (d *Data) PopFront() interface{} {
	if d.GetSize() == 0 {
		return nil
	}
	tm := (*d).GetFront()
	(*d).DeepSliceDelete(0)
	return tm
}

// PopBack removes the back(last) element of sequence
// and return it at the same time.
func (d *Data) PopBack() interface{} {
	if d.GetSize() == 0 {
		return nil
	}
	tm := (*d).GetBack()
	(*d).DeepSliceDelete((*d).GetSize() - 1)
	return tm
}

// GetElements returns a slice of all valued.
func (d Data) GetElements() []interface{} {
	tm := d.Clone()
	slice := []interface{}{}
	for tm.GetSize() != 0 {
		slice = append(slice, tm.PopFront())
	}
	return slice
}

// CommonPrefix returns the longest common leading components
// among all Data. Python commonPrefix compares the maximal
// Data with the minimal Data, which only takes linear time,
// whereas this compares every possible pair among all Data,
// which makes it slower, but still quadratic, than Python.
// This is to find the common prefix among all Data,
// not just between maximal and minimal Data.
func CommonPrefix(more ...*Data) []interface{} {
	minl := more[0]
	min := more[0].GetSize()
	// to get the Data of the shortest GetSizegth
	for _, val := range more {
		if val.GetSize() < min {
			minl = val
			min = val.GetSize()
		}
	}
	// traverse the minimal Data
	// and compare with other Data
	// elements in the same index
	for key, val := range minl.m {
		// if any value in other Data
		// is different than the one
		// in minimal Data
		for _, other := range more {
			if val != other.m[key] {
				return minl.m[:key]
			}
		}
	}
	return minl.m
}
```

[↑ top](#go-sequence)
<br><br><br><br>
<hr>







#### `slice` vs. `container/list`

Go implements linked list in [`container/list`](http://golang.org/pkg/container/list/) package.
Linked list is more efficient only when we need to do many deletions in the 
'middle' of the list. When ordering of the elements isn't important, most efficient is **slice**.
We can mitigate the deletion problem using this slice trick but there is no way
to mitigate the slowness of traversing linked list:

- [Bjarne Stroustrup: Why you should avoid Linked Lists (C++)](https://groups.google.com/d/msg/golang-nuts/mPKCoYNwsoU/tLefhE7tQjMJ)
- [Why you should never use linked-list](http://www.codeproject.com/Articles/340797/Number-crunching-Why-you-should-never-ever-EVER-us)


```go
func BenchmarkSliceFind(b *testing.B) {
	d := NewData()
	for i := 0; i < 999999; i++ {
		d.PushBack(i)
	}
	for i := 0; i < b.N; i++ {
		d.Contains(999998)
	}
}

func BenchmarkContainerListFind(b *testing.B) {
	l := list.New()
	for i := 0; i < 999999; i++ {
		l.PushBack(i)
	}
	for i := 0; i < b.N; i++ {
		for elem := l.Front(); elem != nil; elem = elem.Next() {
			if reflect.DeepEqual(elem.Value, 999998) {
				break
			}
		}
	}
}
```

<br>
And results are:

```
BenchmarkSliceFind           	      10	 156703074 ns/op	10450659 B/op	  100006 allocs/op
BenchmarkSliceFind-2         	       5	 212367352 ns/op	20901283 B/op	  200010 allocs/op
BenchmarkSliceFind-4         	       5	 223453012 ns/op	20901270 B/op	  200010 allocs/op
BenchmarkSliceFind-8         	       5	 226040441 ns/op	20901270 B/op	  200010 allocs/op
BenchmarkSliceFind-16        	       5	 225683822 ns/op	20901270 B/op	  200010 allocs/op

BenchmarkContainerListFind   	       3	 403878856 ns/op	37333312 B/op	 1666665 allocs/op
BenchmarkContainerListFind-2 	       5	 207013596 ns/op	28799980 B/op	 1399998 allocs/op
BenchmarkContainerListFind-4 	       5	 204088451 ns/op	28799980 B/op	 1399998 allocs/op
BenchmarkContainerListFind-8 	       5	 206244553 ns/op	28799980 B/op	 1399998 allocs/op
BenchmarkContainerListFind-16	       5	 214934224 ns/op	28799980 B/op	 1399998 allocs/op
```

**Use slice!**
<br>

[↑ top](#go-sequence)
<br><br><br><br>
<hr>







#### thread-safe, generic **set**

[Code](http://play.golang.org/p/ddvOu79f9s):

```go
package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	func() {
		d := NewData()
		if reflect.TypeOf(d) != reflect.TypeOf(&Data{}) {
			fmt.Errorf("NewData() should return Data type: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		d.Init()
		if !d.IsEmpty() {
			fmt.Errorf("IsEmpty() should return true: %#v", d)
		}
	}()

	func() {
		d := NewData()
		if d.GetSize() != 0 {
			fmt.Errorf("NewData() should return Data of size 0: %#v", d)
		}
	}()

	func() {
		d := NewData()
		if !d.IsEmpty() {
			fmt.Errorf("IsEmpty() should return true: %#v", d)
		}
	}()

	func() {
		d := NewData()
		d.Insert(1, 2, -.9, "A", 0, 2, 2, 2)
		if d.IsEmpty() {
			fmt.Errorf("IsEmpty() should return false: %#v", d)
		}
		if d.GetSize() != 5 {
			fmt.Errorf("GetSize() should return 5: %#v", d)
		}
		value, exist := d.GetFrequency(2)
		if value != 4 {
			fmt.Errorf("s[2]'s value should be 4: %#v", value)
		}
		if !exist {
			fmt.Errorf("s[2] should exist: %#v", value)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		if d.IsEmpty() {
			fmt.Errorf("IsEmpty() should return false: %#v", d)
		}
		if d.GetSize() != 5 {
			fmt.Errorf("GetSize() should return 5: %#v", d)
		}
		value, exist := d.GetFrequency(2)
		if value != 1 {
			fmt.Errorf("value should be 1: %#v", d)
		}
		if exist != true {
			fmt.Errorf("s[2] should exist: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0, 10, 20)
		sl := d.GetElements()
		if len(sl) != 7 {
			fmt.Errorf("len(sl) should be 7: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		if !d.Contains(-0.9) {
			fmt.Errorf("d.Contains(-0.9) should return true: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		if !d.Delete(-0.9) {
			fmt.Errorf("d.Delete(-0.9) should return true: %#v", d)
		}
		if d.Contains(-.9) || d.Contains(-0.9) {
			fmt.Errorf("d.Contains should return false: %#v", d)
		}
		if !d.Delete("A") {
			fmt.Errorf("s.Delete('A') should return true: %#v", d)
		}
		if d.Delete(10000) {
			fmt.Errorf("d.Delete(10000) should return false: %#v", d)
		}
		if d.GetSize() != 3 {
			fmt.Errorf("d.GetSize() should return 3: %#v", d)
		}
		if d.Delete(100) {
			fmt.Errorf("d.Delete(100) should return false: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := CreateData(1, 2)
		result := d.Intersection(a)
		if len(result) != 2 {
			fmt.Errorf("len(result) should return 2: %#v", d)
		}
		ac := CreateData(2, 1)
		if !a.IsEqual(ac) {
			fmt.Errorf("Should be equal: %#v %#v", a, ac)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := CreateData(100, 200)
		result := d.Union(a)
		if len(result) != 7 {
			fmt.Errorf("len(result) should return 7: %#v", result)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := CreateData(1, 2)
		result := d.Subtract(a)
		if len(result) != 3 {
			fmt.Errorf("len(result) should return 3: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := CreateData(1, 2)
		if d.IsEqual(a) {
			fmt.Errorf("Should be false: %#v %#v", d, a)
		}

		b := CreateData("A", 0, 1, 2, -.9, "A")
		if !d.IsEqual(b) {
			fmt.Errorf("Should be true: %#v %#v", d, b)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := CreateData(1, 2)
		if !d.Subset(a) {
			fmt.Errorf("Should be true: %#v %#v", d, a)
		}

		b := CreateData(1, 2, -.9, "A", 0, 100)
		if d.Subset(b) {
			fmt.Errorf("Should be false: %#v %#v", d, b)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := d.Clone()
		if !d.IsEqual(a) {
			fmt.Errorf("Should be true: %#v %#v", d, a)
		}
	}()
}

// Data is a set of data in map data structure.
// Every element is unique, and it is unordered.
// It maps its value to frequency.
type Data struct {
	// m maps an element to its frequency
	m map[interface{}]int

	// RWMutex is more expensive
	// https://blogs.oracle.com/roch/entry/beware_of_the_performance_of
	// sync.RWMutex
	//
	// to synchronize access to shared state across multiple goroutines.
	//
	sync.Mutex
}

// NewData returns a new Data.
// Map supports the built-in function "make"
// so we do not have to use "new" and
// "make" does not return pointer.
func NewData() *Data {
	nmap := make(map[interface{}]int)
	return &Data{
		m: nmap,
	}
	// return make(Data)
}

// Init initializes the Data.
func (d *Data) Init() {
	// (X) d = NewData()
	// This only updates its pointer
	// , not the Data itself
	//
	*d = *NewData()
}

// GetSize returns the size of set.
func (d Data) GetSize() int {
	return len(d.m)
}

// IsEmpty returns true if the set is empty.
func (d Data) IsEmpty() bool {
	return d.GetSize() == 0
}

// Insert insert values to the set.
func (d *Data) Insert(items ...interface{}) {
	for _, value := range items {
		d.Lock()
		v, ok := d.m[value]
		d.Unlock()
		if ok {
			d.Lock()
			d.m[value] = v + 1
			d.Unlock()
			continue
		}
		d.Lock()
		d.m[value] = 1
		d.Unlock()
	}
}

// CreateData instantiates a set object with initial elements.
func CreateData(items ...interface{}) *Data {
	data := NewData()
	data.Insert(items...)
	return data
}

// GetElements returns the set elements.
func (d Data) GetElements() []interface{} {
	slice := []interface{}{}
	d.Lock()
	for key := range d.m {
		slice = append(slice, key)
	}
	d.Unlock()
	return slice
}

// GetFrequency returns the frequency of an element.
func (d Data) GetFrequency(val interface{}) (int, bool) {
	d.Lock()
	fq, ok := d.m[val]
	d.Unlock()
	return fq, ok
}

// Contains returns true if the value exists in the Data.
func (d Data) Contains(val interface{}) bool {
	d.Lock()
	_, ok := d.m[val]
	d.Unlock()
	if ok {
		return true
	}
	return false
}

// Delete deletes the value, or return false
// if the value does not exist in the Data.
func (d Data) Delete(val interface{}) bool {
	if !d.Contains(val) {
		return false
	}
	d.Lock()
	delete(d.m, val)
	d.Unlock()
	return true
}

// Intersection returns values common in both sets.
func (d *Data) Intersection(a *Data) []interface{} {
	rs := []interface{}{}
	for _, elem := range d.GetElements() {
		a.Lock()
		_, ok := a.m[elem]
		a.Unlock()
		if ok {
			rs = append(rs, elem)
		}
	}
	return rs
}

// Union returns the union of two sets.
func (d *Data) Union(a *Data) []interface{} {
	slice := d.GetElements()
	a.Lock()
	for key := range a.m {
		d.Lock()
		_, ok := d.m[key]
		d.Unlock()
		if !ok {
			slice = append(slice, key)
		}
	}
	a.Unlock()
	return slice
}

// Subtract returns the set `d` - `a`.
func (d *Data) Subtract(a *Data) []interface{} {
	slice := []interface{}{}
	d.Lock()
	for key := range d.m {
		a.Lock()
		_, ok := a.m[key]
		a.Unlock()
		if !ok {
			slice = append(slice, key)
		}
	}
	d.Unlock()
	return slice
}

// IsEqual returns true if the two sets are same,
// regardless of its frequency.
func (d *Data) IsEqual(a *Data) bool {
	if d.GetSize() != a.GetSize() {
		return false
	}
	// for every element of s
	d.Lock()
	for key := range d.m {
		// check if it exists in the Data "a"
		a.Lock()
		_, ok := a.m[key]
		a.Unlock()
		if !ok {
			d.Unlock()
			return false
		}
	}
	d.Unlock()
	return true
}

// Subset returns true if "a" is a subset of "s".
func (d *Data) Subset(a *Data) bool {
	if d.GetSize() < a.GetSize() {
		return false
	}
	a.Lock()
	for key := range a.m {
		d.Lock()
		_, ok := d.m[key]
		d.Unlock()
		if !ok {
			return false
		}
	}
	a.Unlock()
	return true
}

// Clone returns a cloned set
// but does not clone its frequency.
func (d *Data) Clone() *Data {
	return CreateData(d.GetElements()...)

}

// String prints out the Data information.
func (d Data) String() string {
	return fmt.Sprintf("Data: %+v", d.GetElements())
}
```

[↑ top](#go-sequence)
<br><br><br><br>
<hr>
