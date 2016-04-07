package main

import "fmt"

func main() {
	d := newData()
	d.addLine("1")
	d.addLine("2")
	l1 := d.lines()
	fmt.Println("l1:", l1)

	d.addLine("3")
	l2 := d.lines()
	fmt.Println("l2:", l2)
	fmt.Println("l1:", l1)
}

/*
l1: [1 2]
l2: [1 2 3]
l1: [1 2]
*/

type data struct {
	sl []string
}

func newData() *data {
	return &data{
		sl: make([]string, 0),
	}
}

func (d *data) lines() []string {
	return d.sl
}

func (d *data) addLine(l string) {
	d.sl = append(d.sl, l)
}
