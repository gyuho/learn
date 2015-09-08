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
