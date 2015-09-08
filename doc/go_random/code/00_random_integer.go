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
