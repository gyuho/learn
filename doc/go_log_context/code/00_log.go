// go run 03_log.go 1>>stdout.log 2>>stderr.log;
package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("First log message!")
	ss := []int{1, 2, 3}
	fmt.Println(ss)
	fmt.Println(ss[5])
}
