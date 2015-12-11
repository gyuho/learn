package main

import (
	"fmt"
	"io"
)

// - interface
// - error interface
// - pointer of a type
// - map
// - slice (not array)
// - byte slice
// - channel

func main() {
	func() {
		var v io.Reader
		v = nil
		fmt.Println("interface became nil:", v)
	}()

	func() {
		v := fmt.Errorf("error")
		v = nil
		fmt.Println("error interface became nil:", v)
	}()

	func() {
		type t struct {
			a string
		}
		v := &t{}
		v = nil
		fmt.Println("pointer of a type became nil:", v)
	}()

	func() {
		v := make(map[string]bool)
		v = nil
		fmt.Println("map became nil:", v)
	}()

	func() {
		v := []int{}
		v = nil
		fmt.Println("slice became nil:", v)

		// v := [3]int{}
		// v = nil
		// cannot use nil as type [3]int in assignment
	}()

	func() {
		v := []byte("Hello")
		v = nil
		fmt.Println("byte slice became nil:", v)
	}()

	func() {
		v := make(chan int)
		v = nil
		fmt.Println("channel became nil:", v)
	}()
}

/*
interface became nil: <nil>
error interface became nil: <nil>
pointer of a type became nil: <nil>
map became nil: map[]
slice became nil: []
byte slice became nil: []
channel became nil: <nil>
*/
