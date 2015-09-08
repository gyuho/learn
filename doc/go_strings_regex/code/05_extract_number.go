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
