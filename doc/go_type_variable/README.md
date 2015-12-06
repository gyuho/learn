[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: type, variable

- [Reference](#reference)
- [type](#type)
- [variable](#variable)
- [empty struct](#empty-struct)

[↑ top](#go-type-variable)
<br><br><br><br><hr>


#### Reference

- [Go types](https://golang.org/ref/spec#Types)
- [Go variables](https://golang.org/ref/spec#Variables)

[↑ top](#go-type-variable)
<br><br><br><br><hr>


#### type

All you need to know about Go types can be found [here](https://golang.org/ref/spec#Types).
And try this [code](http://play.golang.org/p/u5Wix-5z2b):

```go
package main

import (
	"fmt"
	"reflect"
)

// type for numeric value
type Int int

// type for struct
type data struct {
	Value string
}

// type for interface
type Interface interface {
	Len() int
}

func (d data) Len() int {
	return len(d.Value)
}

func main() {
	v0 := Int(0)
	fmt.Println(reflect.TypeOf(v0))
	// main.Int

	v1 := data{Value: "A"}
	fmt.Println(reflect.TypeOf(v1))
	// main.data

	v2 := interface{}(v1).(Interface)
	fmt.Println(reflect.TypeOf(v2))
	// main.data
}
```

[↑ top](#go-type-variable)
<br><br><br><br><hr>


#### variable

Try this [code](http://play.golang.org/p/Sd20zM3Lmq):

```go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var val0 string
	fmt.Println(reflect.TypeOf(val0), unsafe.Sizeof(val0), val0)
	// string 16

	var val1 = "A"
	fmt.Println(reflect.TypeOf(val1), unsafe.Sizeof(val1), val1)
	// string 16 A

	val2 := "B"
	fmt.Println(reflect.TypeOf(val2), unsafe.Sizeof(val2), val2)
	// string 16 B

	var data1 = struct{}{}
	fmt.Println(reflect.TypeOf(data1), unsafe.Sizeof(data1), data1)
	// struct {} 0 {}

	var data2 struct{}
	fmt.Println(reflect.TypeOf(data2), unsafe.Sizeof(data2), data2)
	// struct {} 0 {}
}
```

[↑ top](#go-type-variable)
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

[↑ top](#go-type-variable)
<br><br><br><br><hr>

