[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: spell check

- [Reference](#reference)
- [Spell check algorithm](#spell-check-algorithm)

[↑ top](#go-spell-check)
<br><br><br><br><hr>


#### Reference

- [How to Write a Spelling Corrector by Peter Norvig](http://norvig.com/spell-correct.html)

[↑ top](#go-spell-check)
<br><br><br><br><hr>


#### Spell check algorithm

Try [this](http://play.golang.org/p/Fc2JxOVfTm):

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	fmap := Frequency(strings.NewReader("Francisco, Francisco"))
	fmt.Println(Suggest("Fransisco", fmap))
	// Francisco
}

// Frequency counts the frequency of each word.
func Frequency(reader io.Reader) map[string]int {
	scanner := bufio.NewScanner(reader)
	//
	// This must be called before Scan.
	// The default split function is bufio.ScanLines.
	scanner.Split(bufio.ScanWords)
	//
	fmap := make(map[string]int)
	//
	for scanner.Scan() {
		// Remove all leading and trailing Unicode code points.
		word := strings.Trim(scanner.Text(), ",-!;:\"?.")
		if _, exist := fmap[word]; exist {
			fmap[word]++
		} else {
			fmap[word] = 1
		}
	}
	return fmap
}

// distanceOne sends all possible corrections
// with edit distance 1 to the channel one.
// This is much more probable than the one with 2 edit distance.
func distanceOne(txt string, one chan string) {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	type pair struct {
		front, back string
	}
	pairs := []pair{}
	for i := 0; i <= len(txt); i++ {
		pairs = append(pairs, pair{txt[:i], txt[i:]})
	}
	for _, pair := range pairs {
		// deletion of pair.back[0]
		if len(pair.back) > 0 {
			one <- pair.front + pair.back[1:]
		}
		// transpose of pair.back[0] and pair.back[1]
		if len(pair.back) > 1 {
			one <- pair.front + string(pair.back[1]) + string(pair.back[0]) + pair.back[2:]
		}
		// replace of pair.back[0]
		for _, elem := range alphabet {
			if len(pair.back) > 0 {
				one <- pair.front + string(elem) + pair.back[1:]
			}
		}
		// insertion
		for _, elem := range alphabet {
			one <- pair.front + string(elem) + pair.back
		}
	}
}

// distanceMore sends other possible corrections
// based on the results from distanceOne.
func distanceMore(word string, other chan string) {
	one := make(chan string, 1024*1024)
	go func() {
		distanceOne(word, one)
		close(one)
	}()
	// retrieve from distanceOne results and break when it's done
	for v := range one {
		// run distanceOne in addition to the results from the first distanceOne
		distanceOne(v, other)
	}
}

// known returns the word with maximum frequencies.
func known(txt string, distFunc func(string, chan string), fmap map[string]int) string {
	words := make(chan string, 1024*1024)
	go func() {
		distFunc(txt, words)
		close(words)
	}()
	maxFq := 0
	suggest := ""
	for wd := range words {
		if freq, exist := fmap[wd]; exist && freq > maxFq {
			maxFq, suggest = freq, wd
		}
	}
	return suggest
}

// Suggest suggests the correct spelling based on the sample data.
func Suggest(txt string, fmap map[string]int) string {
	// edit distance 0
	if _, exist := fmap[txt]; exist {
		return txt
	}
	if v := known(txt, distanceOne, fmap); v != "" {
		return v
	}
	if v := known(txt, distanceMore, fmap); v != "" {
		return v
	}
	// edit distance 3, 4, and more ...
	return txt
}

```

[↑ top](#go-spell-check)
<br><br><br><br><hr>
