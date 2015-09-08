[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: map

- [`map`](#map)
- [*non-deterministic* `range` `map`](#non-deterministic-range-map)
- [`map` implementation](#map-implementation)

[↑ top](#go-map)
<br><br><br><br>
<hr>










#### `map`

[*Go maps in action by Andrew
Gerrand*](http://blog.golang.org/go-maps-in-action) covers all you need know to
use Go map. This is me trying to understand the internals of Go map. First
here's [how](http://play.golang.org/p/MNOl4o_s3X) I use map:

```go
package main
 
import (
	"fmt"
	"sort"
)
 
// key/value pair of map[string]float64
type MapSF struct {
	key   string
	value float64
}
 
// Sort map pairs implementing sort.Interface
// to sort by value
type MapSFList []MapSF
 
// sort.Interface
// Define our custom sort: Swap, Len, Less
func (p MapSFList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p MapSFList) Len() int      { return len(p) }
func (p MapSFList) Less(i, j int) bool {
	return p[i].value < p[j].value
}
 
// Sort the struct from a map and return a MapSFList
func sortMapByValue(m map[string]float64) MapSFList {
	p := make(MapSFList, len(m))
	i := 0
	for k, v := range m {
		p[i] = MapSF{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
 
func main() {
	// with sort.Interface and struct
	// we can automatically handle the duplicates
	sfmap := map[string]float64{
		"California":    9.9,
		"Japan":         7.23,
		"Korea":         -.3,
		"Hello":         1.5,
		"USA":           8.4,
		"San Francisco": 8.4,
		"Ohio":          -1.10,
		"New York":      1.23,
		"Los Angeles":   23.1,
		"Mountain View": 9.9,
	}
	fmt.Println(sortMapByValue(sfmap), len(sortMapByValue(sfmap)))
	// [{Ohio -1.1} {Korea -0.3} {New York 1.23} {Hello 1.5}
	// {Japan 7.23} {USA 8.4} {San Francisco 8.4} {California 9.9}
	// {Mountain View 9.9} {Los Angeles 23.1}] 10
 
	if v, ok := sfmap["California"]; !ok {
		fmt.Println("California does not exist")
	} else {
		fmt.Println(v, "exists")
	}
	// 9.9 exists
 
	fmt.Println(sfmap["California"]) // 9.9
 
	if v, ok := sfmap["California2"]; !ok {
		fmt.Println("California2 does not exist")
	} else {
		fmt.Println(v, "exists")
	}
	// California2 does not exist
 
	delete(sfmap, "Ohio")
	if v, ok := sfmap["Ohio"]; !ok {
		fmt.Println("Ohio does not exist")
	} else {
		fmt.Println(v, "exists")
	}
	// Ohio does not exist
}
```

And more examples from [Go: function, method,
pointer](https://github.com/gyuho/learn/tree/master/go_function_method_pointer):

```go
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
}
```
```go
package main
 
import "fmt"
 
func main() {
	b := []byte("abc")
	fmt.Println(b, string(b)) // [97 98 99] abc
	b = nil
	fmt.Println(b, string(b)) // []
 
	// str := "abc"
	// str = nil (x) value cannot be nil
 
	mmap := make(map[string]bool)
	mmap["A"] = true
	fmt.Println(mmap) // map[A:true]
	mmap = nil
	fmt.Println(mmap) // map[]
 
	slice := []int{1}
	fmt.Println(slice) // [1]
	slice = nil
	fmt.Println(slice) // []
}
```
```go
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
```

Basically Go map is a [**hash
table**](https://en.wikipedia.org/wiki/Hash_table), like
[below](http://play.golang.org/p/3V2zvcZZ9J):

```go
// A cyclic redundancy check (CRC) is an error-detecting code commonly used in digital
// networks and storage devices to detect accidental changes to raw data.
//
// https://github.com/golang/go/blob/master/src/hash/crc32/crc32.go
// https://github.com/golang/go/blob/master/src/runtime/hashmap.go
 
// A map is just a hash table.  The data is arranged
// into an array of buckets.  Each bucket contains up to
// 8 key/value pairs.  The low-order bits of the hash are
// used to select a bucket.  Each bucket contains a few
// high-order bits of each hash to distinguish the entries
// within a single bucket.
//
// If more than 8 keys hash to a bucket, we chain on
// extra buckets.
//
// When the hashtable grows, we allocate a new array
// of buckets twice as big.  Buckets are incrementally
// copied from the old bucket array to the new bucket array.
 
package main
 
import (
	"bytes"
	"fmt"
	"hash/crc32"
	"strings"
)
 
func main() {
	// hash table using array
	ht := newHashTable()
	for _, elem := range strings.Split("aaaaaaaaaaabbbbcdeftghiklmnopr", "") {
		ht.insert([]byte(elem))
	}
	for _, bucket := range ht.bucketSlice {
		fmt.Println(bucket)
	}
	/*
	   &{true [[102] [98] [101] [97] [100] [116] [103] [99]]}
	   &{true [[109] [105] [110] [112] [111] [107] [108] [104]]}
	   &{false [[] [] [] [] [] [114] [] []]}
	*/
 
	fmt.Println(ht.search([]byte("f"))) // true
	fmt.Println(ht.search([]byte("x"))) // false
}
 
func hashFuncCrc32(val []byte) uint32 {
	// crc64.Checksum(val, crc64.MakeTable(crc64.ISO))
	return crc32.Checksum(val, crc32.MakeTable(crc32.IEEE))
}
 
func hashFunc(val []byte) uint32 {
	return checksum(val, makePolyTable(crc32.IEEE))
}
 
// polyTable is a 256-word table representing the polynomial for efficient processing.
type polyTable [256]uint32
 
func makePolyTable(poly uint32) *polyTable {
	t := new(polyTable)
	for i := 0; i < 256; i++ {
		crc := uint32(i)
		for j := 0; j < 8; j++ {
			if crc&1 == 1 {
				crc = (crc >> 1) ^ poly
			} else {
				crc >>= 1
			}
		}
		t[i] = crc
	}
	return t
}
 
// checksum returns the CRC-32 checksum of data
// using the polynomial represented by the polyTable.
func checksum(data []byte, tab *polyTable) uint32 {
	crc := ^uint32(0)
	for _, v := range data {
		crc = tab[byte(crc)^v] ^ (crc >> 8)
	}
	return ^crc
}
 
const (
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits // 8, Maximum number of key/value pairs a bucket can hold
)
 
type hashTable struct {
	bucketSlice []*bucket
}
 
func newHashTable() *hashTable {
	table := new(hashTable)
	// table.bucketSlice = make([]*bucket, hashTableSize)
	table.bucketSlice = []*bucket{}
	return table
}
 
type bucket struct {
	wrapped bool // already wrapped around from end of bucket array to beginning
	data    [bucketCnt][]byte
	// type byteData []byte
	// []byte == []uint8
}
 
func newBucket() *bucket {
	newBucket := &bucket{}
	newBucket.wrapped = false
	newBucket.data = [bucketCnt][]byte{}
	return newBucket
}
 
func (h *hashTable) search(val []byte) bool {
	if len(h.bucketSlice) == 0 {
		return false
	}
	probeIdx := hashFunc(val) % uint32(bucketCnt)
	for _, bucket := range h.bucketSlice {
		// check the probeIdx
		if bucket.data[probeIdx] != nil {
			if bytes.Equal(bucket.data[probeIdx], val) {
				return true
			}
		}
		// linear probe
		for idx, elem := range bucket.data {
			if uint32(idx) == probeIdx {
				continue
			}
			if bytes.Equal(elem, val) {
				return true
			}
		}
	}
	return false
}
 
// hashFunc -> probeIdx -> linear probe to fill up bucket
func (h *hashTable) insert(val []byte) {
	if h.search(val) {
		return
	}
	if len(h.bucketSlice) == 0 {
		h.bucketSlice = append(h.bucketSlice, newBucket())
	}
	probeIdx := hashFunc(val) % uint32(bucketCnt)
	isInserted := false
Loop:
	for _, bucket := range h.bucketSlice {
		// if the bucket is already full, skip it
		if bucket.wrapped {
			continue
		}
		// if the index is not taken yet, map it
		if bucket.data[probeIdx] == nil {
			bucket.data[probeIdx] = val
			isInserted = true
			break
		}
		// linear probe
		for idx, elem := range bucket.data {
			if uint32(idx) == probeIdx {
				continue
			}
			if elem == nil {
				bucket.data[idx] = val
				isInserted = true
				break Loop
			}
		}
		bucket.wrapped = true
	}
	if !isInserted {
		nb := newBucket()
		nb.data[probeIdx] = val
		h.bucketSlice = append(h.bucketSlice, nb)
	}
}
```

[↑ top](#go-map)
<br><br><br><br>
<hr>













#### *non-deterministic* `range` `map`

Try [this](http://play.golang.org/p/CgDbiFTJDX):

```go
package main

import (
	"fmt"
	"strings"
)

func nonDeterministicMapUpdateV1() {
	for i := 0; i < 10; i++ {
		fmt.Println("nonDeterministicMapUpdateV1 TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		length := len(mmap)
		for k, v := range mmap {
			mmap[strings.ToUpper(k)] = v * v
			delete(mmap, k)
		}
		if length == len(mmap) {
			fmt.Println("Luckily, Deterministic with nonDeterministicMapUpdateV1:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with nonDeterministicMapUpdateV1:", length, len(mmap))
	}
}

func nonDeterministicMapUpdateV2() {
	for i := 0; i < 10; i++ {
		fmt.Println("nonDeterministicMapUpdateV2 TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		ks := []string{}
		length := len(mmap)
		for k, v := range mmap {
			mmap[strings.ToUpper(k)] = v * v
			ks = append(ks, k)
		}
		for _, k := range ks {
			delete(mmap, k)
		}
		if length == len(mmap) {
			fmt.Println("Luckily, Deterministic with nonDeterministicMapUpdateV2:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with nonDeterministicMapUpdateV2:", length, len(mmap))
	}
}

func nonDeterministicMapUpdateV3() {
	for i := 0; i < 10; i++ {
		fmt.Println("nonDeterministicMapUpdateV3 TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		length := len(mmap)
		for k := range mmap {
			v := mmap[k]
			mmap[strings.ToUpper(k)] = v * v
			delete(mmap, k)
		}
		if length == len(mmap) {
			fmt.Println("Luckily, Deterministic with nonDeterministicMapUpdateV3:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with nonDeterministicMapUpdateV3:", length, len(mmap))
	}
}

func deterministicMapSet() {
	for i := 0; i < 10000; i++ {
		mmap := make(map[int]bool)
		for i := 0; i < 10000; i++ {
			mmap[i] = true
		}
		length := len(mmap)
		for k := range mmap {
			delete(mmap, k)
		}
		if len(mmap) == 0 {
			fmt.Println("Deterministic with deterministicMapSet:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with deterministicMapSet:", length, len(mmap))
	}
}

func deterministicMapDelete() {
	for i := 0; i < 10000; i++ {
		fmt.Println("deterministicMapDelete TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		length := len(mmap)
		for k := range mmap {
			delete(mmap, k)
		}
		if len(mmap) == 0 {
			fmt.Println("Deterministic with deterministicMapDelete:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with deterministicMapDelete:", length, len(mmap))
	}
}

func deterministicMapUpdate() {
	for i := 0; i < 10000; i++ {
		fmt.Println("deterministicMapUpdate TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		mmapCopy := make(map[string]int)
		length := len(mmap)
		for k, v := range mmap {
			mmapCopy[strings.ToUpper(k)] = v * v
		}
		for k := range mmap {
			delete(mmap, k)
		}
		if length == len(mmapCopy) || len(mmap) != 0 {
			fmt.Println("Deterministic with deterministicMapUpdate:", length, len(mmapCopy))
			return
		} else {
			mmapCopy = make(map[string]int) // to initialize(empty)
			//
			// (X)
			// mmapCopy = nil
		}
		fmt.Println("Non-Deterministic with deterministicMapUpdate:", length, len(mmap))
	}
}

func main() {
	nonDeterministicMapUpdateV1()
	fmt.Println()
	nonDeterministicMapUpdateV2()
	fmt.Println()
	nonDeterministicMapUpdateV3()

	fmt.Println()

	deterministicMapSet()
	fmt.Println()
	deterministicMapDelete()
	fmt.Println()
	deterministicMapUpdate()
}

/*
These are all non-deterministic.
If you are lucky, the map gets updated inside range.

nonDeterministicMapUpdateV1 TRY = 0
Non-Deterministic with nonDeterministicMapUpdateV1: 5 4
nonDeterministicMapUpdateV1 TRY = 1
Non-Deterministic with nonDeterministicMapUpdateV1: 5 4
nonDeterministicMapUpdateV1 TRY = 2
Luckily, Deterministic with nonDeterministicMapUpdateV1: 5 5

nonDeterministicMapUpdateV2 TRY = 0
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 1
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 2
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 3
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 4
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 5
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 6
Non-Deterministic with nonDeterministicMapUpdateV2: 5 3
nonDeterministicMapUpdateV2 TRY = 7
Non-Deterministic with nonDeterministicMapUpdateV2: 5 4
nonDeterministicMapUpdateV2 TRY = 8
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 9
Non-Deterministic with nonDeterministicMapUpdateV2: 5 4

nonDeterministicMapUpdateV3 TRY = 0
Non-Deterministic with nonDeterministicMapUpdateV3: 5 4
nonDeterministicMapUpdateV3 TRY = 1
Non-Deterministic with nonDeterministicMapUpdateV3: 5 4
nonDeterministicMapUpdateV3 TRY = 2
Non-Deterministic with nonDeterministicMapUpdateV3: 5 4
nonDeterministicMapUpdateV3 TRY = 3
Luckily, Deterministic with nonDeterministicMapUpdateV3: 5 5

Deterministic with deterministicMapSet: 10000 0

deterministicMapDelete TRY = 0
Deterministic with deterministicMapDelete: 5 0

deterministicMapUpdate TRY = 0
Deterministic with deterministicMapUpdate: 5 5
*/
```

<br>
This tells that `map` is:
- **_Non-deterministic_** on **update** when it *updates* with `for range` of the map. 
- **_Deterministic_** on **`update`** when it *deletes* with `for range`, **NOT** on the map. 
	- `for i := 0; i < 10000; i++ {mmap[i] = true}`
- **_Deterministic_** on **`delete`** when it *deletes* with `for range` of the map. 
- **_Deterministic_** on **`update`** when it *updates* with `for range` of the **COPIED** map. 


<br>
[Go FAQ](http://golang.org/doc/faq#atomic_maps) explains:

> Why are map operations not defined to be atomic?
>
> After long discussion it was decided that the typical use of maps did not
> require safe access from multiple goroutines, and in those cases where it did,
> the map was probably part of some larger data structure or computation that was
> already synchronized. Therefore requiring that all map operations grab a mutex
> would slow down most programs and add safety to few. This was not an easy
> decision, however, since it means uncontrolled map access can crash the
> program.
>
> The language does not preclude atomic map updates. When required, such as when
> hosting an untrusted program, the implementation could interlock map access.
>
> [*Go FAQ*](http://golang.org/doc/faq#atomic_maps)


And about `for` loop:

> The iteration order over maps is not specified and is not guaranteed to be
> the same from one iteration to the next. If map entries that have not yet
> been reached are removed during iteration, the corresponding iteration values
> will not be produced. If map entries are **created during iteration**, that entry
> may be produced during the iteration or **may be skipped**. The choice may vary
> for each entry created and from one iteration to the next. If the map is nil,
> the number of iterations is 0.
>
> [Go Spec](https://golang.org/ref/spec#For_statements)

[↑ top](#go-map)
<br><br><br><br>
<hr>




















#### `map` implementation

The actual implementation is much more complicated, which I can only cover the
fraction of. Source can be found here
[**_`/master/src/runtime/hashmap.go`_**](https://github.com/golang/go/blob/master/src/runtime/hashmap.go):

> **A map is just a hash table. The data is arranged
> into an array of buckets. Each bucket contains up to
> 8 key/value pairs. The low-order bits of the hash are
> used to select a bucket. Each bucket contains a few
> high-order bits of each hash to distinguish the entries
> within a single bucket.**
>
> If more than 8 keys hash to a bucket, we chain on
> extra buckets.
>
> When the hashtable grows, we allocate a new array
> of buckets twice as big. Buckets are incrementally
> copied from the old bucket array to the new bucket array.
> 
> Map iterators walk through the array of buckets and
> return the keys in walk order (bucket #, then overflow
> chain order, then bucket index). To maintain iteration
> semantics, we never move keys within their bucket (if
> we did, keys might be returned 0 or 2 times). When
> growing the table, iterators remain iterating through the
> old table and must check the new table if the bucket
> they are iterating through has been moved (“evacuated”)
> to the new table.

First take a look at [**_runtime type
represenation_**](https://github.com/golang/go/blob/master/src/runtime/hashmap.go):

```go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
 
// Runtime type representation.
 
package runtime
 
type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	_unused    uint8
	align      uint8
	fieldalign uint8
	kind       uint8
	alg        *typeAlg
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata  *byte
	_string *string
	x       *uncommontype
	ptrto   *_type
	zero    *byte // ptr to the zero value for this type
}
 
type maptype struct {
	typ           _type
	key           *_type
	elem          *_type
	bucket        *_type // internal type representing a hash bucket
	hmap          *_type // internal type representing a hmap
	keysize       uint8  // size of key slot
	indirectkey   bool   // store ptr to key instead of key itself
	valuesize     uint8  // size of value slot
	indirectvalue bool   // store ptr to value instead of value itself
	bucketsize    uint16 // size of bucket
	reflexivekey  bool   // true if k==k for all keys
}
```

And also [type
algorithms](https://github.com/golang/go/blob/master/src/runtime/alg.go) for
compiler—[alg.go](https://github.com/golang/go/blob/master/src/runtime/alg.go)
contains the hash functions that are used in Go map implementation:

```go
package runtime
 
// typeAlg is also copied/used in reflect/type.go.
// keep them in sync.
type typeAlg struct {
	// function for hashing objects of this type
	// (ptr to object, seed) -> hash
	hash func(unsafe.Pointer, uintptr) uintptr
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}
```

If you look at
[**_/master/src/runtime/hashmap.go_**](https://github.com/golang/go/blob/master/src/runtime/hashmap.go),
Go map has two parts:
[**_hmap_**](https://github.com/golang/go/blob/master/src/runtime/hashmap.go#L102)
as a header for a Go map, and
[**_bmap_**](https://github.com/golang/go/blob/master/src/runtime/hashmap.go#L127)
as a bucket in a Go map. And
[**_makemap_**](https://github.com/golang/go/blob/master/src/runtime/hashmap.go#L187)
function **initializes a map** and returns **_hmap_** **pointer**:

```go
func makemap(
	t *maptype,
	hint int64,
	h *hmap,
	bucket unsafe.Pointer,
) *hmap {
	...
```

```go
package runtime
 
 
const (
	bucketCntBits = 3
	bucketCnt = 1 << bucketCntBits // 8
)
 
// A header for a Go map.
type hmap struct {
	// Note: the format of the Hmap is encoded in ../../cmd/internal/gc/reflect.go and
	// ../reflect/type.go.  Don't change this structure without also changing that code!
	count int // # live cells == size of map.  Must be first (used by len() builtin)
	flags uint8
	B     uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	hash0 uint32 // hash seed
 
	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)
 
	// If both key and value do not contain pointers and are inline, then we mark bucket
	// type as containing no pointers. This avoids scanning such maps.
	// However, bmap.overflow is a pointer. In order to keep overflow buckets
	// alive, we store pointers to all overflow buckets in hmap.overflow.
	// Overflow is used only if key and value do not contain pointers.
	// overflow[0] contains overflow buckets for hmap.buckets.
	// overflow[1] contains overflow buckets for hmap.oldbuckets.
	// The first indirection allows us to reduce static size of hmap.
	// The second indirection allows to store a pointer to the slice in hiter.
	overflow *[2]*[]*bmap
}
 
// A bucket for a Go map.
type bmap struct {
	tophash [bucketCnt]uint8
	// Followed by bucketCnt keys and then bucketCnt values.
	// NOTE: packing all the keys together and then all the values together makes the
	// code a bit more complicated than alternating key/value/key/value/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
}
```

Note that **bucket stores 8 key/value pairs.**

> *Followed by bucketCnt(8) keys and then bucketCnt values.*

Bucket consists of *key1/key2/key3/value1/value2/value3/…* and internally it
calculates the hash values and fills up the bucket as
[here](https://github.com/golang/go/blob/master/src/runtime/hashmap.go#L411):

```go
hash := alg.hash(key, uintptr(h.hash0))
top := uint(hash >> (ptrSize*8 - 8))
```

To understand this snippet, you need to look at
[stubs.go](https://github.com/golang/go/blob/master/src/runtime/stubs.go):

```go
// https://github.com/golang/go/blob/master/src/runtime/stubs.go
package runtime
 
import "unsafe"
 
// Declarations for runtime services implemented in C or assembly.
 
const ptrSize = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const
const regSize = 4 << (^uintreg(0) >> 63) // unsafe.Sizeof(uintreg(0)) but an ideal const
 
// Should be a built-in for unsafe.Pointer?
//go:nosplit
func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}
```

And here's
[code](https://github.com/golang/go/blob/master/src/runtime/hashmap.go) to
access the map value by key:

```go
// https://github.com/golang/go/blob/master/src/runtime/hashmap.go
package runtime
 
// mapaccess1 returns a pointer to h[key].  Never returns nil, instead
// it will return a reference to the zero object for the value type if
// the key is not in the map.
// NOTE: The returned pointer may keep the whole map live, so don't
// hold onto it for very long.
func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	if raceenabled && h != nil {
		callerpc := getcallerpc(unsafe.Pointer(&t))
		pc := funcPC(mapaccess1)
		racereadpc(unsafe.Pointer(h), callerpc, pc)
		raceReadObjectPC(t.key, key, callerpc, pc)
	}
	if h == nil || h.count == 0 {
		return unsafe.Pointer(t.elem.zero)
	}
	alg := t.key.alg
	hash := alg.hash(key, uintptr(h.hash0))
	m := uintptr(1)<<h.B - 1
	b := (*bmap)(add(h.buckets, (hash&m)*uintptr(t.bucketsize)))
	if c := h.oldbuckets; c != nil {
		oldb := (*bmap)(add(c, (hash&(m>>1))*uintptr(t.bucketsize)))
		if !evacuated(oldb) {
			b = oldb
		}
	}
	top := uint8(hash >> (ptrSize*8 - 8))
	if top < minTopHash {
		top += minTopHash
	}
	for {
		for i := uintptr(0); i < bucketCnt; i++ {
			if b.tophash[i] != top {
				continue
			}
			k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
			if t.indirectkey {
				k = *((*unsafe.Pointer)(k))
			}
			if alg.equal(key, k) {
				v := add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.valuesize))
				if t.indirectvalue {
					v = *((*unsafe.Pointer)(v))
				}
				return v
			}
		}
		b = b.overflow(t)
		if b == nil {
			return unsafe.Pointer(t.elem.zero)
		}
	}
}
```

What this does:

1. `alg := t.key.alg` loads the *type algorithm(or hash function)* from the
   `maptype`.
2. `alg.hash(key, uintptr(h.hash0)` calculates the hash value by its key.
3. `h.hash0` is just a random integer generated by a assembly code.
4. **_for-loop_** is process of **iterating buckets inside map** *until it
   finds the key from the function arguments.*
5. If the key is not found, `unsafe.Pointer(t.elem.zero)` returns a reference
   to the zero object for the value type—it never returns `nil`.

[↑ top](#go-map)
<br><br><br><br>
<hr>
