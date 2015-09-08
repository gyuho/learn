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
