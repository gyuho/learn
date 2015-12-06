[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: strings, regex

- [`string`](#string)
- [clean up](#clean-up)
- [regex](#regex)
- [comma](#comma)
- [replace](#replace)
- [extract numbers](#extract-numbers)
- [space, tab](#space-tab)

[↑ top](#go-strings-regex)
<br><br><br><br><hr>


#### `string`

[Code](http://play.golang.org/p/HW7yvOBYwK):

```go
package main

import "fmt"

func main() {
	str := "Hello World!"
	fmt.Println(str[1]) // 101

	fmt.Println()

	for _, c := range "Hello World!" {
		fmt.Println(c, string(c))
	}
	/*
	   72 H
	   101 e
	   108 l
	   108 l
	   111 o
	   32
	   87 W
	   111 o
	   114 r
	   108 l
	   100 d
	   33 !
	*/

	fmt.Println()

	for _, c := range []byte("Hello World!") {
		fmt.Println(c, string(c))
	}
	/*
	   72 H
	   101 e
	   108 l
	   108 l
	   111 o
	   32
	   87 W
	   111 o
	   114 r
	   108 l
	   100 d
	   33 !
	*/
}
```

[↑ top](#go-strings-regex)
<br><br><br><br><hr>


#### clean up

[Here](http://play.golang.org/p/fwqycU9gqQ) are some useful functions to preprocess Go strings:

```go
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	func() {
		str1 := "   Hello,    World! 124  2 This 23is Go project,		It is amazing  . I am excited!    Are you too?    "
		r1 := cleanUp(str1)
		if r1 != "Hello, World! 124 2 This 23is Go project, It is amazing . I am excited! Are you too?" {
			fmt.Errorf("Expected \"Hello, World! 124 2 This 23is Go project, It is amazing . I am excited! Are you too?\": %s", r1)
		}
		str2 := "		Hello World!	This is Go 		"
		r2 := cleanUp(str2)
		if r2 != "Hello World! This is Go" {
			fmt.Errorf("Expected \"Hello World! This is Go\": %s", r2)
		}
		str3 := "	\n\n\n	Hello World! \n\n	This is Go 		"
		r3 := cleanUp(str3)
		if r3 != "Hello World! This is Go" {
			fmt.Errorf("Expected \"Hello World! This is Go\": %s", r3)
		}
	}()

	func() {
		str := "was't weren't aren't isn't I'd you'd I'll he doesn't Don't I won't I'm You're i've what's up it's I didn't"
		r := expand(strings.ToLower(str))
		rc := "was not were not are not is not i would you would i will he does not do not i will not i am you are i have what is up it is i did not"
		if r != rc {
			fmt.Errorf("Expected\n%v\nbut\n%v", rc, r)
		}
	}()

	func() {
		slice1 := toWords("sadf \nsdafsdf a  s  \nsadfsdf")
		if len(slice1) != 5 {
			fmt.Errorf("Expected 5 but %#v", slice1)
		}
		str2 := "Hello World!  It is Good to See You, Them, 10. This is Go. How are you? He said \"I am x1000 good.\""
		slice2 := toWords(strings.ToLower(str2))
		cslice2 := []string{"hello", "world!", "it", "is", "good", "to", "see", "you,", "them,", "10.", "this", "is", "go.", "how", "are", "you?", "he", "said", "i", "am", "x1000", "good."}
		if len(slice2) != len(cslice2) {
			fmt.Errorf("%#v\n%#v", slice2, cslice2)
		}
		for k, v := range slice2 {
			if cslice2[k] != v {
				fmt.Errorf("%#v != %#v", cslice2[k], v)
			}
		}
	}()

	func() {
		str1 := "Hello World!  It is Good to See You, Them, 10. This is Go. How are you? He said \"I am x1000 good.\""
		slice1 := toSentences(str1)
		cslice1 := []string{"Hello World!", "It is Good to See You, Them, 10.", "This is Go.", "How are you?", "He said I am x1000 good."}
		if len(slice1) != len(cslice1) {
			fmt.Errorf("%#v\n%%#v", slice1, cslice1)
		}
		for k, v := range slice1 {
			if cslice1[k] != v {
				fmt.Errorf("%#v != %#v", cslice1[k], v)
			}
		}
		str2 := "Hello World!  It is Good to See You, Them, 10. This is Go. How are you? He said \"I am x1000 good.\""
		slice2 := toSentences(str2)
		cslice2 := []string{"Hello World!", "It is Good to See You, Them, 10.", "This is Go.", "How are you?", "He said I am x1000 good."}
		if len(slice2) != len(cslice2) {
			fmt.Errorf("%#v\n%%#v", slice2, cslice2)
		}
		for k, v := range slice2 {
			if cslice2[k] != v {
				fmt.Errorf("%#v != %#v", cslice2[k], v)
			}
		}
	}()
}

// cleanUp cleans up unnecessary characters in string.
// It cleans up the blank characters that carry no meaning in context,
// converts all white spaces into single whitespace.
// String is immutable, which means the original string would not change.
func cleanUp(str string) string {
	// validID := regexp.MustCompile(`\s{2,}`)
	// func TrimSpace(s string) string
	// slicing off all "leading" and
	// "trailing" white space, as defined by Unicode.
	str = strings.TrimSpace(str)

	// func Fields(s string) []string
	// Fields splits the slice s around each instance
	// of "one or more consecutive white space"
	slice := strings.Fields(str)

	// now join them with a single white space character
	return strings.Join(slice, " ")
}

func expand(str string) string {
	str = strings.Replace(str, "'d", " would", -1)
	str = strings.Replace(str, "isn't", "is not", -1)
	str = strings.Replace(str, "aren't", "are not", -1)
	str = strings.Replace(str, "was't", "was not", -1)
	str = strings.Replace(str, "weren't", "were not", -1)

	str = strings.Replace(str, "'ve", " have", -1)
	str = strings.Replace(str, "'re", " are", -1)
	str = strings.Replace(str, "'m", " am", -1)
	str = strings.Replace(str, "it's", "it is", -1)
	str = strings.Replace(str, "what's", "what is", -1)
	str = strings.Replace(str, "'ll", " will", -1)

	str = strings.Replace(str, "won't", "will not", -1)
	str = strings.Replace(str, "can't", "can not", -1)
	str = strings.Replace(str, "mustn't", "must not", -1)

	str = strings.Replace(str, "haven't", "have not", -1)
	str = strings.Replace(str, "hasn't", "has not", -1)

	str = strings.Replace(str, "dn't", "d not", -1)
	str = strings.Replace(str, "don't", "do not", -1)
	str = strings.Replace(str, "doesn't", "does not", -1)
	str = strings.Replace(str, "didn't", "did not", -1)

	return str
}

func toWords(str string) []string {
	str = expand(str)
	validID := regexp.MustCompile(`\"`)
	str = validID.ReplaceAllString(str, "")
	return strings.Split(cleanUp(str), " ")
}

func toSentences(str string) []string {
	validID1 := regexp.MustCompile(`\"`)
	str = validID1.ReplaceAllString(str, "")

	validID2 := regexp.MustCompile(`[.]`)
	str = validID2.ReplaceAllString(str, ".___")

	validID3 := regexp.MustCompile(`[?]`)
	str = validID3.ReplaceAllString(str, "?___")

	validID4 := regexp.MustCompile(`[!]`)
	str = validID4.ReplaceAllString(str, "!___")

	slice := strings.Split(str, "___")
	// to clean up the empty strings
	result := []string{}
	for _, value := range slice {
		value = cleanUp(value)
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}
```

[↑ top](#go-strings-regex)
<br><br><br><br><hr>


#### regex

In Python, you can use regex to truncate the string like here:

```python
import re
r = re.match("(.{0,5})", "12312312321")
print r.groups()[0]
# 12312
 
r = re.match("(.{0,5})", "abcdadfasfsddf")
print r.groups()[0]
# abcda
```

In Go, you can do [this](http://play.golang.org/p/IHSIwffAU2):

```go
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
```

[↑ top](#go-strings-regex)
<br><br><br><br><hr>


#### comma

Try this [code](http://play.golang.org/p/6-4ukzzn6o):

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num1 int64 = 21391239213
	r1 := insertCommaInt64(num1)

	if r1 != "21,391,239,213" {
		fmt.Errorf("InsertCommaInt(num1) should return \"21,391,239,213\" : %#v", r1)
	}

	// var num2 float64 = 21391239213.5122123
	num2 := 21391239213.5122123
	r2 := insertCommaFloat64(num2)

	if r2 != "21,391,239,213.51221" {
		fmt.Errorf("InsertCommaFloat(num2) should return \"21,391,239,213.51221\": %#v", r2)
	}
}

// insertCommaInt64 inserts comma in every three digit.
// It returns the new version of input integer, in string format.
func insertCommaInt64(num int64) string {
	// func FormatUint(i uint64, base int) string
	str := strconv.FormatInt(num, 10)
	result := []string{}
	i := len(str) % 3
	if i == 0 {
		i = 3
	}
	for index, elem := range strings.Split(str, "") {
		if i == index {
			result = append(result, ",")
			i += 3
		}
		result = append(result, elem)
	}
	return strings.Join(result, "")
}

// insertCommaFloat64 inserts comma in every three digit in integer part.
// It returns the new version of input float number, in string format.
func insertCommaFloat64(num float64) string {
	// FormatFloat(num, 'f', 6, 64) with precision 6
	// for arbitrary precision, put -1
	str := strconv.FormatFloat(num, 'f', -1, 64)
	slice := strings.Split(str, ".")
	intpart := slice[0]
	floatpart := ""
	if len(slice) > 1 {
		floatpart = slice[1]
	}
	result := []string{}
	i := len(intpart) % 3
	if i == 0 {
		i = 3
	}
	for index, elem := range strings.Split(intpart, "") {
		if i == index {
			result = append(result, ",")
			i += 3
		}
		result = append(result, elem)
	}

	intpart = strings.Join(result, "")
	return intpart + "." + floatpart
}
```

[↑ top](#go-strings-regex)
<br><br><br><br><hr>


#### replace

[Here](http://play.golang.org/p/cBAz-F8RtP)'s how to replace strings
with [regexp](http://golang.org/pkg/regexp):

```go
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	func() {
		str1 := "Hello 1231thi21s123 is 12G33o25!!!!"
		r1 := replaceNonAlpha(strings.ToLower(str1), " ")
		if r1 != "hello     thi  s    is   g  o      " {
			fmt.Errorf("Expected \"hello     thi  s    is   g  o      \": %s", r1)
		}
		str2 := "Hello 1231thi21s123 is 12G33o25!!!!"
		r2 := replaceNonAlpha(strings.ToLower(str2), "")
		if r2 != "hello this is go" {
			fmt.Errorf("Expected \"hello this is go\": %s", r2)
		}
	}()

	func() {
		str1 := "It's me! Hello !@##%W@!#orl!@#!@#!@#d!"
		r1 := replaceNonAlnum(expand(strings.ToLower(str1)), " ")
		if r1 != "it is me  hello      w   orl         d " {
			fmt.Errorf("Expected \"it is me  hello      w   orl         d \": %s", r1)
		}
		str2 := "Hello !@##%W@!#orl!@#!@#!@#d!"
		r2 := replaceNonAlnum(expand(strings.ToLower(str2)), "")
		if r2 != "hello world" {
			fmt.Errorf("Expected \"hello world\": %s", r2)
		}
	}()

	func() {
		slice1 := []string{
			strings.ToLower(strings.Replace(replaceNonAlnumDash("List(Linked List) vs. Slice(Array)", ""), " ", "-", -1)),
			strings.ToLower(strings.Replace(replaceNonAlnumDash("What is Graph? (YouTube Clips)", ""), " ", "-", -1)),
			strings.ToLower(strings.Replace(replaceNonAlnumDash("Adjacency List vs. Adjacency Matrix", ""), " ", "-", -1)),
			strings.ToLower(strings.Replace(replaceNonAlnumDash("C++ Version", ""), " ", "-", -1)),
			strings.ToLower(strings.Replace(replaceNonAlnumDash("what is go-learn", ""), " ", "-", -1)),
		}
		slice2 := []string{
			"listlinked-list-vs-slicearray",
			"what-is-graph-youtube-clips",
			"adjacency-list-vs-adjacency-matrix",
			"c-version",
			"what-is-go-learn",
		}
		for idx, elem := range slice1 {
			if elem != slice2[idx] {
				fmt.Errorf("Should be same:\n%s\n%s\n", elem, slice2[idx])
			}
		}

	}()
}

// replaceNonAlpha replaces all non-alphabetic characters.
func replaceNonAlpha(str, replace string) string {
	// alphabetic (== [A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alpha:]\s]`)
	return validID.ReplaceAllString(str, replace)
}

// replaceNonAlnum replaces all alphanumeric characters.
func replaceNonAlnum(str, replace string) string {
	// alphanumeric (== [0-9A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alnum:]\s]`)
	return validID.ReplaceAllString(str, replace)
}

// replaceNonAlnumDash replaces all alphanumeric characters or dash(-).
func replaceNonAlnumDash(str, replace string) string {
	// alphanumeric (== [0-9A-Za-z])
	// \s is a white space character

	validID := regexp.MustCompile(`[^A-Za-z0-9-\s]`)
	// validID := regexp.MustCompile(`[^A-Za-z0-9\s-]`)

	return validID.ReplaceAllString(str, replace)
}

func expand(str string) string {
	str = strings.Replace(str, "'d", " would", -1)
	str = strings.Replace(str, "isn't", "is not", -1)
	str = strings.Replace(str, "aren't", "are not", -1)
	str = strings.Replace(str, "was't", "was not", -1)
	str = strings.Replace(str, "weren't", "were not", -1)

	str = strings.Replace(str, "'ve", " have", -1)
	str = strings.Replace(str, "'re", " are", -1)
	str = strings.Replace(str, "'m", " am", -1)
	str = strings.Replace(str, "it's", "it is", -1)
	str = strings.Replace(str, "what's", "what is", -1)
	str = strings.Replace(str, "'ll", " will", -1)

	str = strings.Replace(str, "won't", "will not", -1)
	str = strings.Replace(str, "can't", "can not", -1)
	str = strings.Replace(str, "mustn't", "must not", -1)

	str = strings.Replace(str, "haven't", "have not", -1)
	str = strings.Replace(str, "hasn't", "has not", -1)

	str = strings.Replace(str, "dn't", "d not", -1)
	str = strings.Replace(str, "don't", "do not", -1)
	str = strings.Replace(str, "doesn't", "does not", -1)
	str = strings.Replace(str, "didn't", "did not", -1)

	return str
}
```

[↑ top](#go-strings-regex)
<br><br><br><br><hr>


#### extract numbers

[Here](http://play.golang.org/p/cBAz-F8RtP)'s how to extract numbers from a string:

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "I have $1,500,000. The car is $79,000. But I pay only 20 dollars, and the tax was $10.57. The final price is $80,000.35. BRC Burrito is $0.50 right now. I've got change of $.20."
	r := extractNumber(str)
	s := []string{"1500000.", "79000.", "10.57", "80000.35", "0.50", ".20"}
	for idx, elem := range s {
		if r[idx] != elem {
			fmt.Errorf("Not same: \n%#v\n%#v", r, s)
		}
	}
}

func extractNumber(str string) []string {
	ns := regexp.MustCompile(`,`).ReplaceAllString(str, "")
	validID := regexp.MustCompile(`\d{0,}[.]\d{0,}`)
	numslice := validID.FindAllString(ns, -1)
	rs := []string{}
	for _, elem := range numslice {
		if elem != "." {
			rs = append(rs, elem)
		}
	}
	return rs
}
```

[↑ top](#go-strings-regex)
<br><br><br><br><hr>


#### space, tab

Try this [code](http://play.golang.org/p/f0lLG_pKik):

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	func() {
		str := "Hello	World	!	Hello"
		r := allTabIntoSingleSpace(str)
		if r != "Hello World ! Hello" {
			fmt.Errorf("AllTabIntoSingleSpace(str) should return \"Hello World ! Hello\": %#v", r)
		}
	}()

	func() {
		str := "Hello World! Hello"
		r := allSpaceIntoSingleTab(str)

		if r != "Hello	World!	Hello" {
			fmt.Errorf("AllSpaceIntoSingleTab(str) should return \"Hello	World!	Hello\": %#v", r)
		}

	}()

	func() {
		str := "Hello	World	Hello"
		r := tabToSpace(str)

		if r != "Hello World Hello" {
			fmt.Errorf("TabToSpace(str) should return \"Hello World Hello\": %#v", r)
		}
	}()

	func() {
		str := "Hello World Hello"
		r := spaceToTab(str)

		if r != "Hello	World	Hello" {
			fmt.Errorf("SpaceToTab(str) should return \"Hello	World	Hello\": %#v", r)
		}
	}()
}

// allTabIntoSingleSpace converts all tab characters into single whitespace character.
func allTabIntoSingleSpace(str string) string {
	// to take any tab chracters: single tab, double tabs, ...
	validID := regexp.MustCompile(`\t{1,}`)
	return validID.ReplaceAllString(str, " ")
}

// allSpaceIntoSingleTab converts all whitespace characters into single tab character.
func allSpaceIntoSingleTab(str string) string {
	// to take any whitespace characters: single whitespace, doulbe _, ...
	validID := regexp.MustCompile(`\s{1,}`)
	return validID.ReplaceAllString(str, "	")
}

// tabToSpace converts all tab characters into whitespace characters.
func tabToSpace(str string) string {
	validID := regexp.MustCompile(`\t`)
	return validID.ReplaceAllString(str, " ")
}

// spaceToTab converts all whitespace characters into tab characters.
func spaceToTab(str string) string {
	validID := regexp.MustCompile(`\s`)
	return validID.ReplaceAllString(str, "	")
}
```

[↑ top](#go-strings-regex)
<br><br><br><br><hr>
