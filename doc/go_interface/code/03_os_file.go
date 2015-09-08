package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var fpath = "a.txt"

func main() {
	// func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", lines)
	// a.txt contents ...
}
