package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
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

	if err := copyDir("testdata", "copy_test"); err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered:", err)
		}
		fmt.Println("deleting...")
		os.RemoveAll("copy_test")
	}()
	panic(111)
}

func copyDir(src, dst string) error {
	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dst, si.Mode()); err != nil {
		return err
	}

	dir, err := os.Open(src)
	if err != nil {
		return err
	}
	defer dir.Close()

	fis, err := dir.Readdir(-1)
	for _, fi := range fis {
		sp := src + "/" + fi.Name()
		dp := dst + "/" + fi.Name()
		if fi.IsDir() {
			if err := copyDir(sp, dp); err != nil {
				// create sub-directories - recursively
				return err
			}
		} else {
			if err := copy(sp, dp); err != nil {
				return err
			}
		}
	}

	return nil
}

/*
0777    full access for everyone
0700    only private access
0755    private read/write access, others only read access
0750    private read/write access, group read access, others no access
*/
func copy(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return fmt.Errorf("copy: mkdirall: %v", err)
	}

	r, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("copy: open(%q): %v", src, err)
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("copy: create(%q): %v", dst, err)
	}
	defer w.Close()

	// func Copy(dst Writer, src Reader) (written int64, err error)
	if _, err = io.Copy(w, r); err != nil {
		return err
	}
	if err := w.Sync(); err != nil {
		return err
	}
	if _, err := w.Seek(0, 0); err != nil {
		return err
	}
	return nil
}
