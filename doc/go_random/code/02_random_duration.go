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
