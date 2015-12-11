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
