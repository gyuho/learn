package main

import (
	"fmt"
	"math/rand"
)

// You can either pass the pointer of map or just map to update.
// But if you want to initialize with assignment, you have to pass pointer.

func updateMap1(m map[int]bool) {
	for {
		num := rand.Intn(150)
		if _, ok := m[num]; !ok {
			m[num] = true
		}
		if len(m) == 5 {
			return
		}
	}
}

func initializeMap1(m map[int]bool) {
	m = nil
	m = make(map[int]bool)
}

type mapType map[int]bool

func (m mapType) updateMap1() {
	m[0] = false
	m[1] = false
}

func (m mapType) initializeMap1() {
	m = nil
	m = make(map[int]bool)
}

func updateMap2(m *map[int]bool) {
	for {
		num := rand.Intn(150)
		if _, ok := (*m)[num]; !ok {
			(*m)[num] = true
		}
		if len(*m) == 5 {
			return
		}
	}
}

func initializeMap2(m *map[int]bool) {
	// *m = nil
	*m = make(map[int]bool)
}

func (m *mapType) updateMap2() {
	(*m)[0] = false
	(*m)[1] = false
}

func (m *mapType) initializeMap2() {
	// *m = nil
	*m = make(map[int]bool)
}

func main() {
	m0 := make(map[int]bool)
	m0[1] = true
	m0[2] = true
	fmt.Println("Done:", m0) // Done: map[1:true 2:true]

	m0 = make(map[int]bool)
	fmt.Println("After:", m0) // After: map[]

	m1 := make(map[int]bool)
	updateMap1(m1)
	fmt.Println("updateMap1:", m1)
	// (o) change
	// updateMap1: map[131:true 87:true 47:true 59:true 31:true]

	initializeMap1(m1)
	fmt.Println("initializeMap1:", m1)
	// (X) no change
	// initializeMap1: map[59:true 31:true 131:true 87:true 47:true]

	mapType(m1).updateMap1()
	fmt.Println("mapType(m1).updateMap1():", m1)
	// (o) change
	// mapType(m1).updateMap1(): map[87:true 47:true 59:true 31:true 0:false 1:false 131:true]

	mapType(m1).initializeMap1()
	fmt.Println("mapType(m1).initializeMap1():", m1)
	// (X) no change
	// mapType(m1).initializeMap1(): map[59:true 31:true 0:false 1:false 131:true 87:true 47:true]

	m2 := make(map[int]bool)
	updateMap2(&m2)
	fmt.Println("updateMap2:", m2)
	// (o) change
	// updateMap2: map[140:true 106:true 0:true 18:true 25:true]

	initializeMap2(&m2)
	fmt.Println("initializeMap2:", m2)
	// (o) change
	// initializeMap2: map[]

	(*mapType)(&m2).updateMap2()
	fmt.Println("(*mapType)(&m2).updateMap2:", m2)
	// (o) change
	// (*mapType)(&m2).updateMap2: map[0:false 1:false]

	(*mapType)(&m2).initializeMap2()
	fmt.Println("(*mapType)(&m2).initializeMap2:", m2)
	// (o) change
	// (*mapType)(&m2).initializeMap2: map[]
}
