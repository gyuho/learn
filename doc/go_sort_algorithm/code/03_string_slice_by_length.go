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
