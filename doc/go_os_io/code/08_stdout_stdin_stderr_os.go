package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello World!")
	// Hello World!

	fmt.Fprintln(os.Stdin, "Input")

	fmt.Fprintln(os.Stderr, "Error")
}

// go run 07_stdout_stdin_stderr_os.go 0>>stdin.txt 1>>stdout.txt 2>>stderr.txt
