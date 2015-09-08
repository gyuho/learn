package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	func() {
		d := NewData()
		if reflect.TypeOf(d) != reflect.TypeOf(&Data{}) {
			fmt.Errorf("NewData() should return Data type: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		d.Init()
		if !d.IsEmpty() {
			fmt.Errorf("IsEmpty() should return true: %#v", d)
		}
	}()

	func() {
		d := NewData()
		if d.GetSize() != 0 {
			fmt.Errorf("NewData() should return Data of size 0: %#v", d)
		}
	}()

	func() {
		d := NewData()
		if !d.IsEmpty() {
			fmt.Errorf("IsEmpty() should return true: %#v", d)
		}
	}()

	func() {
		d := NewData()
		d.Insert(1, 2, -.9, "A", 0, 2, 2, 2)
		if d.IsEmpty() {
			fmt.Errorf("IsEmpty() should return false: %#v", d)
		}
		if d.GetSize() != 5 {
			fmt.Errorf("GetSize() should return 5: %#v", d)
		}
		value, exist := d.GetFrequency(2)
		if value != 4 {
			fmt.Errorf("s[2]'s value should be 4: %#v", value)
		}
		if !exist {
			fmt.Errorf("s[2] should exist: %#v", value)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		if d.IsEmpty() {
			fmt.Errorf("IsEmpty() should return false: %#v", d)
		}
		if d.GetSize() != 5 {
			fmt.Errorf("GetSize() should return 5: %#v", d)
		}
		value, exist := d.GetFrequency(2)
		if value != 1 {
			fmt.Errorf("value should be 1: %#v", d)
		}
		if exist != true {
			fmt.Errorf("s[2] should exist: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0, 10, 20)
		sl := d.GetElements()
		if len(sl) != 7 {
			fmt.Errorf("len(sl) should be 7: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		if !d.Contains(-0.9) {
			fmt.Errorf("d.Contains(-0.9) should return true: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		if !d.Delete(-0.9) {
			fmt.Errorf("d.Delete(-0.9) should return true: %#v", d)
		}
		if d.Contains(-.9) || d.Contains(-0.9) {
			fmt.Errorf("d.Contains should return false: %#v", d)
		}
		if !d.Delete("A") {
			fmt.Errorf("s.Delete('A') should return true: %#v", d)
		}
		if d.Delete(10000) {
			fmt.Errorf("d.Delete(10000) should return false: %#v", d)
		}
		if d.GetSize() != 3 {
			fmt.Errorf("d.GetSize() should return 3: %#v", d)
		}
		if d.Delete(100) {
			fmt.Errorf("d.Delete(100) should return false: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := CreateData(1, 2)
		result := d.Intersection(a)
		if len(result) != 2 {
			fmt.Errorf("len(result) should return 2: %#v", d)
		}
		ac := CreateData(2, 1)
		if !a.IsEqual(ac) {
			fmt.Errorf("Should be equal: %#v %#v", a, ac)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := CreateData(100, 200)
		result := d.Union(a)
		if len(result) != 7 {
			fmt.Errorf("len(result) should return 7: %#v", result)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := CreateData(1, 2)
		result := d.Subtract(a)
		if len(result) != 3 {
			fmt.Errorf("len(result) should return 3: %#v", d)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := CreateData(1, 2)
		if d.IsEqual(a) {
			fmt.Errorf("Should be false: %#v %#v", d, a)
		}

		b := CreateData("A", 0, 1, 2, -.9, "A")
		if !d.IsEqual(b) {
			fmt.Errorf("Should be true: %#v %#v", d, b)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := CreateData(1, 2)
		if !d.Subset(a) {
			fmt.Errorf("Should be true: %#v %#v", d, a)
		}

		b := CreateData(1, 2, -.9, "A", 0, 100)
		if d.Subset(b) {
			fmt.Errorf("Should be false: %#v %#v", d, b)
		}
	}()

	func() {
		d := CreateData(1, 2, -.9, "A", 0)
		a := d.Clone()
		if !d.IsEqual(a) {
			fmt.Errorf("Should be true: %#v %#v", d, a)
		}
	}()
}

// Data is a set of data in map data structure.
// Every element is unique, and it is unordered.
// It maps its value to frequency.
type Data struct {
	// m maps an element to its frequency
	m map[interface{}]int

	// RWMutex is more expensive
	// https://blogs.oracle.com/roch/entry/beware_of_the_performance_of
	// sync.RWMutex
	//
	// to synchronize access to shared state across multiple goroutines.
	//
	sync.Mutex
}

// NewData returns a new Data.
// Map supports the built-in function "make"
// so we do not have to use "new" and
// "make" does not return pointer.
func NewData() *Data {
	nmap := make(map[interface{}]int)
	return &Data{
		m: nmap,
	}
	// return make(Data)
}

// Init initializes the Data.
func (d *Data) Init() {
	// (X) d = NewData()
	// This only updates its pointer
	// , not the Data itself
	//
	*d = *NewData()
}

// GetSize returns the size of set.
func (d Data) GetSize() int {
	return len(d.m)
}

// IsEmpty returns true if the set is empty.
func (d Data) IsEmpty() bool {
	return d.GetSize() == 0
}

// Insert insert values to the set.
func (d *Data) Insert(items ...interface{}) {
	for _, value := range items {
		d.Lock()
		v, ok := d.m[value]
		d.Unlock()
		if ok {
			d.Lock()
			d.m[value] = v + 1
			d.Unlock()
			continue
		}
		d.Lock()
		d.m[value] = 1
		d.Unlock()
	}
}

// CreateData instantiates a set object with initial elements.
func CreateData(items ...interface{}) *Data {
	data := NewData()
	data.Insert(items...)
	return data
}

// GetElements returns the set elements.
func (d Data) GetElements() []interface{} {
	slice := []interface{}{}
	d.Lock()
	for key := range d.m {
		slice = append(slice, key)
	}
	d.Unlock()
	return slice
}

// GetFrequency returns the frequency of an element.
func (d Data) GetFrequency(val interface{}) (int, bool) {
	d.Lock()
	fq, ok := d.m[val]
	d.Unlock()
	return fq, ok
}

// Contains returns true if the value exists in the Data.
func (d Data) Contains(val interface{}) bool {
	d.Lock()
	_, ok := d.m[val]
	d.Unlock()
	if ok {
		return true
	}
	return false
}

// Delete deletes the value, or return false
// if the value does not exist in the Data.
func (d Data) Delete(val interface{}) bool {
	if !d.Contains(val) {
		return false
	}
	d.Lock()
	delete(d.m, val)
	d.Unlock()
	return true
}

// Intersection returns values common in both sets.
func (d *Data) Intersection(a *Data) []interface{} {
	rs := []interface{}{}
	for _, elem := range d.GetElements() {
		a.Lock()
		_, ok := a.m[elem]
		a.Unlock()
		if ok {
			rs = append(rs, elem)
		}
	}
	return rs
}

// Union returns the union of two sets.
func (d *Data) Union(a *Data) []interface{} {
	slice := d.GetElements()
	a.Lock()
	for key := range a.m {
		d.Lock()
		_, ok := d.m[key]
		d.Unlock()
		if !ok {
			slice = append(slice, key)
		}
	}
	a.Unlock()
	return slice
}

// Subtract returns the set `d` - `a`.
func (d *Data) Subtract(a *Data) []interface{} {
	slice := []interface{}{}
	d.Lock()
	for key := range d.m {
		a.Lock()
		_, ok := a.m[key]
		a.Unlock()
		if !ok {
			slice = append(slice, key)
		}
	}
	d.Unlock()
	return slice
}

// IsEqual returns true if the two sets are same,
// regardless of its frequency.
func (d *Data) IsEqual(a *Data) bool {
	if d.GetSize() != a.GetSize() {
		return false
	}
	// for every element of s
	d.Lock()
	for key := range d.m {
		// check if it exists in the Data "a"
		a.Lock()
		_, ok := a.m[key]
		a.Unlock()
		if !ok {
			d.Unlock()
			return false
		}
	}
	d.Unlock()
	return true
}

// Subset returns true if "a" is a subset of "s".
func (d *Data) Subset(a *Data) bool {
	if d.GetSize() < a.GetSize() {
		return false
	}
	a.Lock()
	for key := range a.m {
		d.Lock()
		_, ok := d.m[key]
		d.Unlock()
		if !ok {
			return false
		}
	}
	a.Unlock()
	return true
}

// Clone returns a cloned set
// but does not clone its frequency.
func (d *Data) Clone() *Data {
	return CreateData(d.GetElements()...)

}

// String prints out the Data information.
func (d Data) String() string {
	return fmt.Sprintf("Data: %+v", d.GetElements())
}
