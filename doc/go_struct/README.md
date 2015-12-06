[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: struct

- [`struct` for data, `interface` for method](#struct-for-data-interface-for-method)
- [empty `struct`](#empty-struct)

[↑ top](#go-struct)
<br><br><br><br><hr>


#### `struct` for data, `interface` for method

Go **struct** controls the layout of **data**, while Go **interface**:
- is a set of methods.
- is a set of constraints on types.
- is to specify the behavior of an object.

<br>

> Interfaces in Go provide a way to specify the behavior of an object: if
> **something can do this, then it can be used here**.
>
> [**_Effective
> Go_**](https://golang.org/doc/effective_go.html#interfaces_and_types)

<br>

For instance,
[`time.Time`](https://go.googlesource.com/go/+/master/src/time/time.go) is
**_`struct`_** because it contains **data** of your local time:

```go
type Time struct {
    sec  int64
    nsec int32
    loc  *Location
}
```

[**`sort.Interface`**](http://golang.org/pkg/sort/#Interface)
is **_`interface`_** to specify **_behaviors_** or **_requirements_** for
[*`sort`*](http://golang.org/pkg/sort/):

```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

**_`struct`_** is a **type** to **contain data with a set of fields for its
values**. **_`interface`_** is also a **type** to **represent a behavior with a
set of methods**. **_`interface`_** is a set of **constraints on types with
methods**. When a **type implements all the methods in an interface type**, the
**_type_** **implicitly satisfies the interface**. You need not declare that a
type is trying to use the interface. Here’s an
[example](http://play.golang.org/p/WDQWzdlHnu):

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

[↑ top](#go-struct)
<br><br><br><br><hr>


#### empty `struct`

Try [this](http://play.golang.org/p/4B9GnIy-FX)
and read:

- [Empty struct by Dave Cheney](http://dave.cheney.net/2014/03/25/the-empty-struct)


```go
package main

import (
	"fmt"
	"unsafe"
)

func N(n int) []struct{} {
	return make([]struct{}, n)
}

func main() {
	var s0 struct{}
	fmt.Println(unsafe.Sizeof(s0)) // 0

	s1 := struct{}{}
	fmt.Println(unsafe.Sizeof(s1)) // 0

	done := make(chan struct{})
	go func() {
		done <- struct{}{}
	}()
	<-done
	fmt.Println("Done")
	// Done

	for i := range N(10) {
		fmt.Print(i, " ")
	}
	// 0 1 2 3 4 5 6 7 8 9
}

/*
Same in Python:

for i in range(10):
    print i
*/
```

[↑ top](#go-struct)
<br><br><br><br><hr>
