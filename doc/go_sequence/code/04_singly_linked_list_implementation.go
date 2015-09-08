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
