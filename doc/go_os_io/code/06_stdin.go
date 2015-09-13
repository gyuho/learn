package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	fmt.Print("input:", input)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("scanner:", scanner.Text())
	}
	// Ctrl + D
}

/*
$ go run 05_stdin_0.go
Enter text: Hello
input:Hello
scanner:
1
scanner: 1
2
scanner: 2
3
scanner: 3
4
scanner: 4
5
scanner: 5
*/
