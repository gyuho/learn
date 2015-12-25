[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: random

- [Reference](#reference)
- [random integer](#random-integer)
- [random float](#random-float)
- [random duration](#random-duration)
- [random `bytes`](#random-bytes)

[↑ top](#go-random)
<br><br><br><br><hr>


#### Reference

- [package `math/rand`](http://golang.org/pkg/math/rand/)

[↑ top](#go-random)
<br><br><br><br><hr>


#### random integer

[Code](http://play.golang.org/p/88gzcG-r4v):

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	fmt.Println(src.Int63()) // 8965630292270293660

	random := rand.New(src)
	fmt.Println(random.Int())      // 7742198863449996164
	fmt.Println(random.Int31())    // 1780122247
	fmt.Println(random.Int31n(3))  // 0
	fmt.Println(random.Int63())    // 838216768439018635
	fmt.Println(random.Int63n(10)) // 7
}

```

[↑ top](#go-random)
<br><br><br><br><hr>


#### random float

[Code](http://play.golang.org/p/AyWtUt-W7U):

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	fmt.Println(random.Float32())     // 0.7096111
	fmt.Println(random.Float64())     // 0.7267748269300062
	fmt.Println(random.ExpFloat64())  // 1.4478015992783408
	fmt.Println(random.NormFloat64()) // -1.7676830716730048
}

```

[↑ top](#go-random)
<br><br><br><br><hr>


#### random duration

[Code](http://play.golang.org/p/251bHlcW9S):

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println(duration(5*time.Second, 12*time.Second)) // 5.061306405s
}

func duration(min, max time.Duration) time.Duration {
	if min >= max {
		// return a random duration
		return 7*time.Second + 173*time.Microsecond
	}
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	adt := time.Duration(random.Int63n(int64(max - min)))
	return min + adt
}

```

[↑ top](#go-random)
<br><br><br><br><hr>


#### random `bytes`

```go
package main

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	b := make([]byte, 10)
	if _, err := crand.Read(b); err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// ��.���ms#

	fmt.Println(string(randBytes(10)))
	// IdPDZOxast

	fmt.Println(multiRandBytes(3, 5))
	// [[119 121 67] [114 70 70] [112 90 100] [74 85 77] [84 101 101]]
}

// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func randBytes(bytesN int) []byte {
	const (
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, bytesN)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := bytesN-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}

func multiRandBytes(bytesN, sliceN int) [][]byte {
	m := make(map[string]struct{})
	rs := [][]byte{}
	for len(rs) != sliceN {
		b := randBytes(bytesN)
		if _, ok := m[string(b)]; !ok {
			rs = append(rs, b)
			m[string(b)] = struct{}{}
		} else {
			continue
		}
	}
	return rs
}

```

[↑ top](#go-random)
<br><br><br><br><hr>
