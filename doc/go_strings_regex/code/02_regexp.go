package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("(.{0,5})")
	fmt.Printf("%s\n", re.FindString("12312312321"))    // 12312
	fmt.Printf("%s\n", re.FindString("abcdadfasfsddf")) // abcda
}
