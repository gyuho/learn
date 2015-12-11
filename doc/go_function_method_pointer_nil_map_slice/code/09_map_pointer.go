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

	mmap := make(map[string]map[string]struct{})
	mmap["A"] = make(map[string]struct{})
	mmap["A"]["B"] = struct{}{}
	mmap["A"]["C"] = struct{}{}
	tm := mmap["A"]
	tm["D"] = struct{}{}
	delete(tm, "B")
	fmt.Println(mmap) // map[A:map[C:{} D:{}]]
}
