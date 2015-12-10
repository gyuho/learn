package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["A"] = 1
	m["B"] = 2
	for k, v := range m {
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err, "at", k, v)
				}
			}()
			panic("panic")
		}()
	}
}

/*
panic at A 1
panic at B 2
*/
