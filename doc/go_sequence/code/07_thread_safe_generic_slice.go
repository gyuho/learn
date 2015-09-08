package main

import (
	"container/list"
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func main() {
	func() {
		dt1 := NewData()
		dt1.Init()
		dt2 := NewData()
		if !reflect.DeepEqual(dt1, dt2) {
			fmt.Errorf("%#v %#v", dt1, dt2)
		}
	}()

	func() {
		dt := NewData()
		dt.Insert([]interface{}{1, "A", 3, -.9, "B"}...)
		if dt.GetSize() != 5 {
			fmt.Errorf("Should be '5': %#v", dt)
		}
	}()

	func() {
		d := NewData()
		d.Insert([]interface{}{1, "A", 3, -.9, "B"}...)
		d.Init()
		if d.GetSize() != 0 {
			fmt.Errorf("Should be '0': %#v", d)
		}
		dt := NewData()
		dt.PushBack(1)
		dt.PushBack(2)
		dt.PushBack(3)
		dt.Insert(3, 1, 7, 11)
		isempty1 := dt.IsEmpty()
		dt.Init()
		isempty2 := dt.IsEmpty()
		if isempty1 != false && isempty2 != true {
			fmt.Errorf("Should return 'false' and 'true': %v, %v", isempty1, isempty2)
		}
	}()

	func() {
		dt := NewData()
		if !dt.IsEmpty() {
			fmt.Errorf("Should return 'true': %#v", dt)
		}
	}()

	func() {
		dt := NewData()
		dt.Insert(1, 3, 4, 5, "A", "7", 100)
		if dt.GetSize() != 7 {
			fmt.Errorf("Should return 7: %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, 3, 4, 5, "A", "7", 100)
		if dt.GetSize() != 7 {
			fmt.Errorf("Should return 7: %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		c := dt.Clone()
		if dt.GetSize() != c.GetSize() {
			fmt.Errorf("Should return true but %#v / %#v", dt, c)
		}
		if !dt.IsDeepEqual(*c) || !dt.IsSemiEqual(*c) {
			fmt.Errorf("Should return true but %+v %+v", dt, c)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		idx, ok := dt.Contains("3")
		if ok {
			fmt.Errorf("Should return false but %+v %+v", idx, ok)
		}
		idx, ok = dt.Contains(3)
		if !ok {
			fmt.Errorf("Should return true but %+v %+v", idx, ok)
		}
		a, b := dt.Contains("A")
		if a != 1 && b != true {
			fmt.Errorf("Should return '1, true': %#v", dt)
		}
		c, d := dt.Contains(-.8)
		if c != 0 && d != false {
			fmt.Errorf("Should return '0, false': %#v", dt)
		}
	}()

	func() {
		s1 := CreateData(1, "A", 3, -.9, "B")
		s1c := CreateData()
		s1c.Insert(1, "A", 3, -.9, "B")
		s2 := CreateData(3, -.9, "B", 1, "A")
		s3 := CreateData()
		s3.Insert(-.9, "B", 1, 3, "A")
		if !s1.IsDeepEqual(*s1c) {
			fmt.Errorf("Should return true but %+v %+v", s1, s1c)
		}
		if s1.IsDeepEqual(*s2) {
			fmt.Errorf("Should return false but %+v %+v", s1, s2)
		}
		if !s1.IsSemiEqual(*s3) {
			fmt.Errorf("Should return true but %+v %+v", s1, s3)
		}
		if !s2.IsSemiEqual(*s3) {
			fmt.Errorf("Should return true but %+v %+v", s2, s3)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		dt.PushFront("Front")
		if dt.GetFront() != "Front" {
			fmt.Errorf("dt.GetFront() should be 'Front': %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		dt.PushBack("Back")
		if dt.GetBack() != "Back" {
			fmt.Errorf("dt.GetBack() should be 'Back': %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		dt.DeepSliceDelete(2)
		if dt.m[2] != -0.9 {
			fmt.Errorf("dt.m[2] should be '-0.9': %#v", dt)
		}
		for _ = range dt.m {
			dt.DeepSliceDelete(0)
			// Don't do dt.DeepSliceDelete(k)
			// the slice length decreases at the same time
		}
		if dt.GetSize() != 0 {
			fmt.Errorf("Should be empty but: %#v", dt.GetSize())
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		ok := dt.FindAndDelete(3)
		if !ok || dt.m[2] != -0.9 {
			fmt.Errorf("Should return true, but %#v, and dt.m[2] should be '-0.9': %#v", ok, dt)
		}

		// list := dt
		// this does deep copy
		// (they are the same)
		// so we need to use Copy
		list := dt.Clone()
		for _, v := range list.m {
			dt.FindAndDelete(v)
		}
		l := dt.GetSize()
		if l != 0 {
			fmt.Errorf("Should be empty but: %#v / %#v / %#v", l, dt, list)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		dt.DeepSliceCut(2, 3)
		if dt.m[2] != "B" {
			fmt.Errorf("dt[2] should be 'B': %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		tm := dt.GetFront()
		if tm != 1 {
			fmt.Errorf("dt.GetFront() should be 1: %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		tm := dt.GetBack()
		if tm != "B" {
			fmt.Errorf("dt.GetBack() should be 'B': %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		tm := dt.PopFront()
		if tm != 1 {
			fmt.Errorf("dt.PopFront() should return 1: %#v", dt)
		}
		if dt.m[0] != "A" {
			fmt.Errorf("dt[0] should be 'A': %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		tm := dt.PopBack()
		if tm != "B" {
			fmt.Errorf("dt.PopBack() should return 'B': %#v", dt)
		}
		if dt.m[3] != -0.9 {
			fmt.Errorf("dt[3] should be -0.9: %#v", dt)
		}
	}()

	func() {
		dt := CreateData(1, "A", 3, -.9, "B")
		slice := dt.GetElements()
		if len(slice) != 5 {
			fmt.Errorf("len(slice) should return 5: %#v", dt)
		}
		if slice[3] != -0.9 {
			fmt.Errorf("slice[3] should be -0.9: %#v", dt)
		}
	}()

	func() {
		s1 := CreateData(1, "A", 3, -.9, "B", "e", "f", "G")
		s2 := CreateData(1, "A", 3, -.9, "B")
		s3 := CreateData(1, "A", 3, -.9, "H", 2, 3, 4)
		s4 := CreateData(1, "A", 3, -.9, "H", 2, 3, 4)
		s5 := CreateData(1, "A", 3, -.9, "B", "e", "f")
		result := CommonPrefix(s1, s2, s3, s4, s5)
		if len(result) != 4 {
			fmt.Errorf("len(result) should return 4: %#v", result)
		}
		if result[3] != -0.9 {
			fmt.Errorf("result[3] should be -0.9: %#v", result)
		}
	}()
}

// Data can contain any type of values,
// because its data is a slice of interface{} type.
// It is an empty interface, which means that it can
// be satisfied by any type of value.
type Data struct {
	m []interface{}

	// RWMutex is more expensive
	// https://blogs.oracle.com/roch/entry/beware_of_the_performance_of
	// sync.RWMutex
	//
	// to synchronize access to shared state across multiple goroutines.
	//
	sync.Mutex
}

// NewData returns a new object of Data.
func NewData() *Data {
	nslice := []interface{}{}
	return &Data{
		m: nslice,
	}
}

// GetSize returns the GetSizegth of sequence.
// If the method needs to mutate the receiver,
// the receiver must be a pointer.
// (http://golang.org/doc/faq#methods_on_values_or_pointers)
func (d Data) GetSize() int {
	d.Lock()
	size := len(d.m)
	d.Unlock()
	return size
}

// Init initializes the Data.
func (d *Data) Init() {
	// (X) d = NewData()
	// This only updates its pointer
	// , not the Data itself
	//
	*d = *NewData()
}

// IsEmpty returns true if the sequence is empty.
func (d Data) IsEmpty() bool {
	return d.GetSize() == 0
}

// Insert appends values to Data.
func (d *Data) Insert(vals ...interface{}) {
	d.Lock()
	for _, elem := range vals {
		d.m = append(d.m, elem)
	}
	d.Unlock()
}

// CreateData instantiates a set object with initial elements.
func CreateData(vals ...interface{}) *Data {
	data := NewData()
	data.Insert(vals...)
	return data
}

// Clone returns a copy of the sequence.
// This is useful because ":=" operator
// does deep copy and when we manipulate
// either one, then the other one also changed.
func (d Data) Clone() *Data {
	td := NewData()
	for _, v := range d.m {
		td.PushBack(v)
	}
	return td
}

// Contains returns true if the elem exists
// in the Data.
func (d Data) Contains(elem interface{}) (int, bool) {
	d.Lock()
	defer d.Unlock()
	for idx, val := range d.m {
		if reflect.DeepEqual(val, elem) {
			return idx, true
		}
	}
	return 0, false
}

// IsSemiEqual returns true if s1 is equal to s2
// regardless of its ordering of elementd.
func (d Data) IsSemiEqual(a Data) bool {
	if d.GetSize() != a.GetSize() {
		return false
	}
	for _, val := range a.m {
		_, ok := d.Contains(val)
		if !ok {
			return false
		}
	}
	return true
}

// IsDeepEqual returns true if s1 is equal to s2.
func (d Data) IsDeepEqual(a Data) bool {
	if d.GetSize() != a.GetSize() {
		return false
	}
	return reflect.DeepEqual(d.m, a.m)
}

// PushFront adds an element to the front of sequence.
func (d *Data) PushFront(val interface{}) {
	size := d.GetSize()
	ts := make([]interface{}, size+1)
	ts[0] = val
	d.Lock()
	copy(ts[1:], d.m)
	d.m = ts
	d.Unlock()
}

// PushBack adds an element to the back of sequence.
func (d *Data) PushBack(val interface{}) {
	d.Lock()
	d.m = append(d.m, val)
	d.Unlock()
}

// DeepSliceDelete deletes the element in the index.
func (d *Data) DeepSliceDelete(idx int) {
	size := d.GetSize()
	d.Lock()
	copy(d.m[idx:], d.m[idx+1:])
	d.m[size-1] = nil // zero value of type or nil
	d.m = d.m[:size-1 : size-1]
	d.Unlock()
}

// FindAndDelete finds the element and delete it.
func (d *Data) FindAndDelete(val interface{}) bool {
	idx, ok := d.Contains(val)
	if !ok {
		return false
	}
	d.DeepSliceDelete(idx)
	return true
}

// DeepSliceCut deletes the elements from indices a to b.
func (d *Data) DeepSliceCut(a, b int) {
	if b > d.GetSize()-1 || a < 0 || a > b {
		panic("Index out of range! You can cut only inside slice.")
	}
	diff := b - a + 1
	idx := a
	i := 0
	for i < diff {
		d.DeepSliceDelete(idx)
		i++
	}
}

// GetFront returns the first(front) element of sequence.
func (d Data) GetFront() interface{} {
	if d.GetSize() == 0 {
		return nil
	}
	return d.m[0]
}

// GetBack returns the last(back) element of sequence.
func (d Data) GetBack() interface{} {
	if d.GetSize() == 0 {
		return nil
	}
	return d.m[d.GetSize()-1]
}

// PopFront removes the front(first) element of sequence
// and return it at the same time.
func (d *Data) PopFront() interface{} {
	if d.GetSize() == 0 {
		return nil
	}
	tm := (*d).GetFront()
	(*d).DeepSliceDelete(0)
	return tm
}

// PopBack removes the back(last) element of sequence
// and return it at the same time.
func (d *Data) PopBack() interface{} {
	if d.GetSize() == 0 {
		return nil
	}
	tm := (*d).GetBack()
	(*d).DeepSliceDelete((*d).GetSize() - 1)
	return tm
}

// GetElements returns a slice of all valued.
func (d Data) GetElements() []interface{} {
	tm := d.Clone()
	slice := []interface{}{}
	for tm.GetSize() != 0 {
		slice = append(slice, tm.PopFront())
	}
	return slice
}

// CommonPrefix returns the longest common leading components
// among all Data. Python commonPrefix compares the maximal
// Data with the minimal Data, which only takes linear time,
// whereas this compares every possible pair among all Data,
// which makes it slower, but still quadratic, than Python.
// This is to find the common prefix among all Data,
// not just between maximal and minimal Data.
func CommonPrefix(more ...*Data) []interface{} {
	minl := more[0]
	min := more[0].GetSize()
	// to get the Data of the shortest GetSizegth
	for _, val := range more {
		if val.GetSize() < min {
			minl = val
			min = val.GetSize()
		}
	}
	// traverse the minimal Data
	// and compare with other Data
	// elements in the same index
	for key, val := range minl.m {
		// if any value in other Data
		// is different than the one
		// in minimal Data
		for _, other := range more {
			if val != other.m[key] {
				return minl.m[:key]
			}
		}
	}
	return minl.m
}

func BenchmarkSliceFind(b *testing.B) {
	d := NewData()
	for i := 0; i < 999999; i++ {
		d.PushBack(i)
	}
	for i := 0; i < b.N; i++ {
		d.Contains(999998)
	}
}

func BenchmarkContainerListFind(b *testing.B) {
	l := list.New()
	for i := 0; i < 999999; i++ {
		l.PushBack(i)
	}
	for i := 0; i < b.N; i++ {
		for elem := l.Front(); elem != nil; elem = elem.Next() {
			if reflect.DeepEqual(elem.Value, 999998) {
				break
			}
		}
	}
}

/*
go test -bench 01_slice_vs_linked_list_test.go -benchmem -cpu 1,2,4,8,16;

BenchmarkSliceFind           	      10	 156703074 ns/op	10450659 B/op	  100006 allocs/op
BenchmarkSliceFind-2         	       5	 212367352 ns/op	20901283 B/op	  200010 allocs/op
BenchmarkSliceFind-4         	       5	 223453012 ns/op	20901270 B/op	  200010 allocs/op
BenchmarkSliceFind-8         	       5	 226040441 ns/op	20901270 B/op	  200010 allocs/op
BenchmarkSliceFind-16        	       5	 225683822 ns/op	20901270 B/op	  200010 allocs/o

BenchmarkContainerListFind   	       3	 403878856 ns/op	37333312 B/op	 1666665 allocs/op
BenchmarkContainerListFind-2 	       5	 207013596 ns/op	28799980 B/op	 1399998 allocs/op
BenchmarkContainerListFind-4 	       5	 204088451 ns/op	28799980 B/op	 1399998 allocs/op
BenchmarkContainerListFind-8 	       5	 206244553 ns/op	28799980 B/op	 1399998 allocs/op
BenchmarkContainerListFind-16	       5	 214934224 ns/op	28799980 B/op	 1399998 allocs/op
*/
