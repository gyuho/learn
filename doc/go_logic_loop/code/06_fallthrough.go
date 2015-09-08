package main

import "fmt"

func main() {

	// "fallthrough" statement transfers control
	// to the first statement of the next case
	// clause in a expression "switch" statement.
	// It may be used only as the final non-empty
	// statement in such a clause.

	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		// To fall through to a subsequent case
		fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 1")
	}
	// 1 > 10

	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		// fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 2")
	}
	// [no output]

	bf := 2
	switch bf {
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("None 3")
	}
	// 3

	cf := 5
	switch cf {
	case 1:
		fmt.Println("1")
	case 2:
		// fallthrough
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("None 4")
	}
	// None 4

	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		break
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 5")
	}
	// [no output]
}
