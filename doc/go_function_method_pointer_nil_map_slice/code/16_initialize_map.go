package main

import "fmt"

func main() {
	func() {
		m := map[string]bool{"A": true}
		m = make(map[string]bool)
		m["A"] = true
		fmt.Println(m)
		// map[A:true]
	}()

	func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		func() {
			m := map[string]bool{"A": true}
			m = nil
			m["A"] = true
			fmt.Println(m)
		}()
		// panic: assignment to entry in nil map
	}()

	func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		func() {
			m := new(map[string]bool)
			(*m)["A"] = true
			fmt.Println(*m)
		}()
		// panic: assignment to entry in nil map
	}()

	func() {
		m := new(map[string]bool)
		*m = make(map[string]bool)
		(*m)["A"] = true
		fmt.Println(*m)
		// map[A:true]
	}()
}
