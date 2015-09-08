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
