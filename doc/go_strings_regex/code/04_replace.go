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
