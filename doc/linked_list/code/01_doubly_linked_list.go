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
