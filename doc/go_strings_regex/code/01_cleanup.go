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
