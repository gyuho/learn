[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: sort, algorithm

- [sort `int`](#sort-int)
- [sort `float`](#sort-float)
- [sort `string`](#sort-string)
- [sort `string` `slice` by length](#sort-string-slice-by-length)
- [sort `struct`](#sort-struct)
- [sort `map`](#sort-map)
- [sort table](#sort-table)
- [bubble sort](#bubble-sort)
- [insertion sort](#insertion-sort)
- [selection sort](#selection-sort)
- [counting sort](#counting-sort)
- [radix sort](#radix-sort)
- [sort search](#sort-search)

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### sort `int`

[Code](http://play.golang.org/p/OAFF1GGYD4):

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	fmt.Println(s)
	// [1 2 3 4 5 6]
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### sort `float`

[Code](http://play.golang.org/p/X9OgpX4eWg):

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []float64{5.4, 2.1, 3.5, 6.1, -10.5} // unsorted
	sort.Float64s(s)
	fmt.Println(s)
	// [-10.5 2.1 3.5 5.4 6.1]
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### sort `string`

[Code](http://play.golang.org/p/8gqNtcYSvK):

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"X", "x", "a", "A", "G"} // unsorted
	sort.Strings(s)
	fmt.Println(s)
	// [A G X a x]
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### sort `string` `slice` by length

[Code](http://play.golang.org/p/WDQWzdlHnu):

```go
package main

import (
	"fmt"
	"sort"
)

var words = []string{
	"adasdasd", "d", "aaasdasdasd", "qqqq", "kkkk",
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j]) // ascending order
}

func main() {
	sort.Sort(sort.StringSlice(words))
	// sort.Strings(words)
	fmt.Printf("%q\n", words)
	// ["aaasdasdasd" "adasdasd" "d" "kkkk" "qqqq"]

	sort.Sort(byLength(words))
	fmt.Printf("%q\n", words)
	// ["d" "kkkk" "qqqq" "adasdasd" "aaasdasdasd"]
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### sort `struct`

[Code](http://play.golang.org/p/Ss48HqtqMo):

```go
package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// byAge implements sort.Interface for []person based on
// the Age field.
type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people) // [Bob: 31 John: 42 Michael: 17 Jenny: 26]
	sort.Sort(byAge(people))
	fmt.Println(people) // [Michael: 17 Jenny: 26 Bob: 31 John: 42]
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### sort `map`

[Code](http://play.golang.org/p/MNOl4o_s3X):

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

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### sort table

Try [this](http://play.golang.org/p/b3iZfgsGe5) and [this](http://play.golang.org/p/h2OLcglgjq):

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	rows := [][]string{
		[]string{"1", "a", "1", "10"},
		[]string{"1", "b", "1", "9"},
		[]string{"1", "c", "1", "8"},
		[]string{"1", "d", "1", "7"},
		[]string{"1", "e", "1", "6"},
		[]string{"1", "f", "1", "5"},
		[]string{"1", "g", "1", "4"},
		[]string{"1", "h", "1", "3"},
		[]string{"1", "i", "1", "2"},
		[]string{"1", "j", "1", "1"},
	}
	rs1 := stringsAscending(rows, 1)
	if fmt.Sprintf("%v", rs1) != "[[1 a 1 10] [1 b 1 9] [1 c 1 8] [1 d 1 7] [1 e 1 6] [1 f 1 5] [1 g 1 4] [1 h 1 3] [1 i 1 2] [1 j 1 1]]" {
		fmt.Errorf("rs1 %v", rs1)
	}
	rs2 := stringsAscending(rows, 3)
	if fmt.Sprintf("%v", rs2) != "[[1 j 1 1] [1 a 1 10] [1 i 1 2] [1 h 1 3] [1 g 1 4] [1 f 1 5] [1 e 1 6] [1 d 1 7] [1 c 1 8] [1 b 1 9]]" {
		fmt.Errorf("rs2 %v", rs2)
	}
	rs3 := stringsDescending(rows, 1)
	if fmt.Sprintf("%v", rs3) != "[[1 j 1 1] [1 i 1 2] [1 h 1 3] [1 g 1 4] [1 f 1 5] [1 e 1 6] [1 d 1 7] [1 c 1 8] [1 b 1 9] [1 a 1 10]]" {
		fmt.Errorf("rs3 %v", rs3)
	}
	rs4 := stringsDescending(rows, 3)
	if fmt.Sprintf("%v", rs4) != "[[1 b 1 9] [1 c 1 8] [1 d 1 7] [1 e 1 6] [1 f 1 5] [1 g 1 4] [1 h 1 3] [1 i 1 2] [1 a 1 10] [1 j 1 1]]" {
		fmt.Errorf("rs4 %v", rs4)
	}
}

var sortColumnIndex int

// sortByIndexAscending sorts two-dimensional strings in an ascending order, at a specified index.
type sortByIndexAscending [][]string

func (s sortByIndexAscending) Len() int {
	return len(s)
}

func (s sortByIndexAscending) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByIndexAscending) Less(i, j int) bool {
	return s[i][sortColumnIndex] < s[j][sortColumnIndex]
}

// stringsAscending sorts two dimensional strings in an ascending order.
func stringsAscending(rows [][]string, idx int) [][]string {
	sortColumnIndex = idx
	sort.Sort(sortByIndexAscending(rows))
	return rows
}

// sortByIndexDescending sorts two-dimensional strings in an Descending order, at a specified index.
type sortByIndexDescending [][]string

func (s sortByIndexDescending) Len() int {
	return len(s)
}

func (s sortByIndexDescending) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByIndexDescending) Less(i, j int) bool {
	return s[i][sortColumnIndex] > s[j][sortColumnIndex]
}

// stringsDescending sorts two dimensional strings in a descending order.
func stringsDescending(rows [][]string, idx int) [][]string {
	sortColumnIndex = idx
	sort.Sort(sortByIndexDescending(rows))
	return rows
}
```

```go
package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func strToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func main() {
	rows := [][]string{
		[]string{"cdomain.com", "3", "-5.02", "aaa", "aaa"},
		[]string{"cdomain.com", "2", "133.02", "aaa", "aaa"},
		[]string{"cdomain.com", "1", "1.02", "aaa", "aaa"},
		[]string{"bdomain.com", "2", "23.02", "aaa", "aaa"},
		[]string{"bdomain.com", "1", "12.02", "aaa", "aaa"},
		[]string{"bdomain.com", "3", "53.02", "aaa", "aaa"},
		[]string{"adomain.com", "5", "32.1232", "aaa", "aaa"},
		[]string{"adomain.com", "3", "2.02202", "aaa", "aaa"},
		[]string{"adomain.com", "1", "511.02", "aaa", "aaa"},
	}
	ascendingName0 := func(row1, row2 *[]string) bool {
		return (*row1)[0] < (*row2)[0]
	}
	descendingVal := func(row1, row2 *[]string) bool {
		return strToFloat64((*row1)[2]) > strToFloat64((*row2)[2])
	}
	ascendingName1 := func(row1, row2 *[]string) bool {
		return (*row1)[1] < (*row2)[1]
	}
	by(rows, ascendingName0, descendingVal, ascendingName1).Sort(rows)
	rs := fmt.Sprintf("%v", rows)
	if rs != "[[adomain.com 1 511.02 aaa aaa] [adomain.com 5 32.1232 aaa aaa] [adomain.com 3 2.02202 aaa aaa] [bdomain.com 3 53.02 aaa aaa] [bdomain.com 2 23.02 aaa aaa] [bdomain.com 1 12.02 aaa aaa] [cdomain.com 2 133.02 aaa aaa] [cdomain.com 1 1.02 aaa aaa] [cdomain.com 3 -5.02 aaa aaa]]" {
		fmt.Errorf("%v", rows)
	}
}

// by returns a multiSorter that sorts using the less functions
func by(rows [][]string, lesses ...lessFunc) *multiSorter {
	return &multiSorter{
		data: rows,
		less: lesses,
	}
}

// lessFunc compares between two string slices.
type lessFunc func(p1, p2 *[]string) bool

func makeAscendingFunc(idx int) func(row1, row2 *[]string) bool {
	return func(row1, row2 *[]string) bool {
		return (*row1)[idx] < (*row2)[idx]
	}
}

// multiSorter implements the Sort interface
// , sorting the two dimensional string slices within.
type multiSorter struct {
	data [][]string
	less []lessFunc
}

// Sort sorts the rows according to lessFunc.
func (ms *multiSorter) Sort(rows [][]string) {
	sort.Sort(ms)
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.data)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.data[i], ms.data[j] = ms.data[j], ms.data[i]
}

// Less is part of sort.Interface.
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.data[i], &ms.data[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q
			return true
		case less(q, p):
			// p > q
			return false
		}
		// p == q; try next comparison
	}
	return ms.less[k](p, q)
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### bubble sort

Try [this](http://play.golang.org/p/l1P4dAaKhd):

```go
package main

import "fmt"

func main() {
	nums := []int{1, -1, 23, -2, 23, 123, 12, 1}
	bubbleSort(nums)
	fmt.Println(nums)
	// [-2 -1 1 1 12 23 23 123]
}

/*
O (n^2)

bubbleSort(A)
for i = 1 to A.length - 1
	for j = A.length downto i + 1
		if A[j] < A[j-1]
			exchange A[j] with A[j-1]
*/
func bubbleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := len(nums) - 1; j != i-1; j-- {
			// the bigger value 'bubbles up' to the last position
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### insertion sort

Try [this](http://play.golang.org/p/pB6ecWnjGV):

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, -1, 23, -2, 23, 123, 12, 1}
	insertionSort(nums1)
	fmt.Println(nums1)
	// [-2 -1 1 1 12 23 23 123]

	nums2 := []int{1, -1, 23, -2, 23, 123, 12, 1}
	insertionSortInterface(sort.IntSlice(nums2), 0, len(nums2))
	fmt.Println(nums2)
	// [-2 -1 1 1 12 23 23 123]
}

// O (n^2)
func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; (j > 0) && (nums[j] < nums[j-1]); j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

func insertionSortInterface(data sort.Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### selection sort

Try [this](http://play.golang.org/p/VXu4DSRl5D):

```go
package main

import "fmt"

func main() {
	nums := []int{1, -1, 23, -2, 23, 123, 12, 1}
	selectionSort(nums)
	fmt.Println(nums)
}

// O (n^2)
func selectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		// Min is the index of the minimum element.
		// Swap it with the current position
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### counting sort

Try [this](http://play.golang.org/p/Wo5EZ_MKOC):

```go
package main

import "fmt"

func main() {
	nums := []int{20, 370, 45, 75, 410, 1802, 24, 2, 66}
	fmt.Println(countingSort(nums))
	// [0 2 20 24 45 66 75 370 410 1802]
}

/*
Counting Sort is O(n).

It does not do any comparison.
Instead, counting sort uses the actual values
of the elements to index into an array.
It only works for positive integers.
The running time depends on the largest element.
Therefore, if the maximum value is very large, the sorting takes long time.

range 0 to k, for some integer k:

1. Create an array(slice) of the size of the maximum value + 1.
2. Count each element.
3. Add up the elements.
4. Put them back to result.
*/

func countingSort(nums []int) []int {

	// 1. Create an array(nums) of the size of the maximum value + 1
	k := max(nums)
	count := make([]int, k+1)

	// 2. Count each element
	for i := 0; i < len(nums); i++ {
		count[nums[i]] = count[nums[i]] + 1
	}

	// 3. Add up the elements
	for i := 1; i < k+1; i++ {
		count[i] = count[i] + count[i-1]
	}

	// 4. Put them back to result
	rs := make([]int, len(nums)+1)
	for j := 0; j < len(nums); j++ {
		rs[count[nums[j]]] = nums[j]
		count[nums[j]] = count[nums[j]] - 1
	}

	return rs
}

func max(nums []int) int {
	max := nums[0]
	for _, elem := range nums {
		if max < elem {
			max = elem
		}
	}
	return max
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### radix sort

Try [this](http://play.golang.org/p/Xmy0RPVXAv):

```go
package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	nums := []int{732, 23, 1, 55, 7130, 321, 223, 5}
	radixSort(nums)
	fmt.Println(nums)
	// [1 5 23 55 223 321 732 7130]
}

/*
Radix Sort
1. Set up an array of initially empty "buckets"
2. Take the smallest of each element
3. Group elements from the smallest
4. Repeat the process
*/
func radixSort(nums []int) {

	// 1. Set up an array of initially empty "buckets"
	// create 10 buckets of which is a list
	var bucketList [10]*list.List
	for i := 0; i < 10; i++ {
		// initialize each bucket
		bucketList[i] = list.New()
	}

	max := max(nums)
	maxdigit := 0
	for max > 0 {
		// 2/10 == 0, 2%10 == 2
		max /= 10

		// if max is 812, maxdigit is 3
		maxdigit++
	}

	/*
		2. Take the smallest of each element
		3. Group elements from the smallest
		4. Repeat the process
	*/
	// if i is 2, then it means 3rd digit
	// if i is 2, in 321, i is 1
	for i := 0; i < maxdigit; i++ {

		// Pow10 returns 10**e, the base-10 exponential of e
		// math.Pow10(2) is 100
		p := int(math.Pow10(i + 1))
		q := int(math.Pow10(i))

		for j := 0; j < len(nums); j++ {
			/*
				x is the i-th digit

				if nums[0] is 123, and i is 0
				then 123 % 10 / 1 ---> x is 3

				if nums[0] is 123, and i is 1
				then 123 % 100 / 10 ---> x is 2
			*/
			x := nums[j] % p / q

			// add nums[j] to x th bucket
			// group by the digit
			bucketList[x].PushBack(nums[j])
		}

		count := 0
		for k := 0; k < 10; k++ {
			for elem := bucketList[k].Front(); elem != nil; elem = elem.Next() {
				nums[count] = elem.Value.(int)
				count++
			}
			bucketList[k].Init()
		}
	}
}

func max(nums []int) int {
	max := nums[0]
	for _, elem := range nums {
		if max < elem {
			max = elem
		}
	}
	return max
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>


#### sort search

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	{
		// given a slice data sorted in ascending order
		names := []string{"a", "b", "c", "d", "e"}

		idx := sort.Search(
			len(names),
			func(i int) bool {
				fmt.Println("searching at", i)
				return names[i] >= "d"
			})

		if idx == len(names) {
			fmt.Println("d is not found")
		} else {
			fmt.Println(idx, names[idx])
		}
	}
	/*
	   searching at 2
	   searching at 4
	   searching at 3
	   3 d
	*/

	{
		// Searching data sorted in descending order would use
		// the <= operator instead of the >= operator.
		names := []string{"e", "d", "c", "b", "a"}

		idx := sort.Search(
			len(names),
			func(i int) bool {
				fmt.Println("searching at", i)
				return names[i] <= "d"
			})

		if idx == len(names) {
			fmt.Println("d is not found")
		} else {
			fmt.Println(idx, names[idx])
		}
	}
	/*
		searching at 2
		searching at 1
		searching at 0
		1 d
	*/
}

```

[↑ top](#go-sort-algorithm)
<br><br><br><br><hr>
