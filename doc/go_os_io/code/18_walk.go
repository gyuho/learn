package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	func() {
		// recursively walk
		rmap, err := walk(".")
		if err != nil {
			panic(err)
		}
		for _, v := range rmap {
			fmt.Println("walk:", v)
		}
	}()

	fmt.Println()
	fmt.Println()

	func() {
		// recursively walk
		rmap, err := walkExt(".", ".txt")
		if err != nil {
			panic(err)
		}
		for _, v := range rmap {
			fmt.Println("walkExt:", v)
		}
	}()

	fmt.Println()
	fmt.Println()

	func() {
		// only the present working directory
		rmap, err := walkExtCurrentDir(".", ".txt")
		if err != nil {
			panic(err)
		}
		for _, v := range rmap {
			fmt.Println("walkExtCurrentDir:", v)
		}
	}()

	fmt.Println()
	fmt.Println()

	func() {
		// walk only directories
		rmap, err := walkDir(".")
		if err != nil {
			panic(err)
		}
		for _, v := range rmap {
			fmt.Println("walkDir:", v)
		}
	}()
}

// walk returns all FileInfos with recursive Walk in the target directory.
// It does not include the directories but include the files inside each sub-directories.
// It does not follow the symbolic links. And excludes hidden files.
// It returns the map from os.FileInfo to its absolute path.
func walk(targetDir string) (map[os.FileInfo]string, error) {
	rmap := make(map[os.FileInfo]string)
	visit := func(path string, f os.FileInfo, err error) error {
		if !filepath.HasPrefix(path, ".") && !strings.Contains(path, "/.") {
			if _, ok := rmap[f]; !ok {
				wd, err := os.Getwd()
				if err != nil {
					return err
				}
				rmap[f] = filepath.Join(wd, path)
			}
		}
		return nil
	}
	err := filepath.Walk(targetDir, visit)
	if err != nil {
		return nil, err
	}
	return rmap, nil
}

// walkExt returns all FileInfos with specific extension.
// Make sure to prefix the extension name with dot.
// For example, to find all go files, pass ".go".
func walkExt(targetDir, ext string) (map[os.FileInfo]string, error) {
	rmap := make(map[os.FileInfo]string)
	visit := func(path string, f os.FileInfo, err error) error {
		if f != nil {
			if !f.IsDir() {
				if filepath.Ext(path) == ext {
					if !filepath.HasPrefix(path, ".") && !strings.Contains(path, "/.") {
						if _, ok := rmap[f]; !ok {
							wd, err := os.Getwd()
							if err != nil {
								return err
							}
							thepath := filepath.Join(wd, strings.Replace(path, wd, "", -1))
							rmap[f] = thepath
						}
					}
				}
			}
		}
		return nil
	}
	err := filepath.Walk(targetDir, visit)
	if err != nil {
		return nil, err
	}
	return rmap, nil
}

// walkExtCurrentDir only walks the current directory, not sub-directories.
func walkExtCurrentDir(targetDir, ext string) (map[os.FileInfo]string, error) {
	rmap := make(map[os.FileInfo]string)
	visit := func(path string, f os.FileInfo, err error) error {
		if f != nil {
			if !f.IsDir() {
				if filepath.Ext(path) == ext {
					if !filepath.HasPrefix(path, ".") && !strings.Contains(path, "/.") {
						if _, ok := rmap[f]; !ok {
							wd, err := os.Getwd()
							if err != nil {
								return err
							}
							thepath := filepath.Join(wd, strings.Replace(path, wd, "", -1))
							if wd == filepath.Dir(thepath) {
								rmap[f] = thepath
							}
						}
					}
				}
			}
		}
		return nil
	}
	err := filepath.Walk(targetDir, visit)
	if err != nil {
		return nil, err
	}
	return rmap, nil
}

// walkDir returns all directories.
func walkDir(targetDir string) (map[os.FileInfo]string, error) {
	rmap := make(map[os.FileInfo]string)
	visit := func(path string, f os.FileInfo, err error) error {
		if f != nil {
			if f.IsDir() {
				if !filepath.HasPrefix(path, ".") && !strings.Contains(path, "/.") {
					if _, ok := rmap[f]; !ok {
						rmap[f] = filepath.Join(targetDir, path)
					}
				}
			}
		}
		return nil
	}
	if err := filepath.Walk(targetDir, visit); err != nil {
		return nil, err
	}
	return rmap, nil
}

/*
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/02_flag.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/07_stdout_stdin_stderr.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/18_walk.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample.csv
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample_copy.csv
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/09_exist.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/22_importdeps.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample_copy.csv
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample.txt
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample.csv
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/03_io.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/13_bufio.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/15_csv.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/21_temp_file.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/stderr.txt
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/stdin.txt
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/17_gzip.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/01_os_exec.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/05_io_ioutil.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/06_stdin.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/11_io_ioutil_file.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/14_copy.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/12_temp_file.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/16_tsv.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample.txt
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample.json
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/00_os.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/04_io_pipe.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/08_stdout_stdin_stderr_os.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/10_open_create.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/19_flush.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/20_signal.go
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/stdout.txt
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata
walk: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample.json


walkExt: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/stderr.txt
walkExt: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/stdin.txt
walkExt: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/stdout.txt
walkExt: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample.txt
walkExt: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample.txt


walkExtCurrentDir: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/stderr.txt
walkExtCurrentDir: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/stdin.txt
walkExtCurrentDir: /home/gyuho/go/src/github.com/gyuho/learn/doc/go_os_io/code/stdout.txt


walkDir: testdata
walkDir: testdata/sub

*/
