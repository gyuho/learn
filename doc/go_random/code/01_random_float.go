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
