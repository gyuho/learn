package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// exist returns true if the file or directory exists.
func exist(fpath string) bool {
	// Does a directory exist
	st, err := os.Stat(fpath)
	if err != nil {
		return false
	}
	if st.IsDir() {
		return true
	}
	if _, err := os.Stat(fpath); err != nil {
		if strings.Contains(err.Error(), "no such file") {
			return false
		}
	}
	return true
}

// existDir returns true if the specified path points to a directory.
// It returns false and error if the directory does not exist.
func existDir(fpath string) bool {
	st, err := os.Stat(fpath)
	if err != nil {
		return false
	}
	return st.IsDir()
}

// readDir lists files in a directory.
func readDir(fpath string) ([]string, error) {
	dir, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

func main() {
	fmt.Println(exist("00_os.go"))    // true
	fmt.Println(exist("aaaaa.go"))    // false
	fmt.Println(exist("testdata"))    // true
	fmt.Println(existDir("testdata")) // true
	ns, err := readDir("./testdata")
	if err != nil {
		panic(err)
	}
	fmt.Println(ns) // [sample.csv sample.json sample.txt sample_copy.csv sub]
}
