package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	func() {
		// the present working directory
		rmap, err := walk(".")
		if err != nil {
			panic(err)
		}
		for _, v := range rmap {
			fmt.Println(v)
		}
	}()

	fmt.Println()
	fmt.Println()

	func() {
		// the present working directory
		rmap, err := walkExt(".", ".go")
		if err != nil {
			panic(err)
		}
		for _, v := range rmap {
			fmt.Println(v)
		}
	}()

	fmt.Println()
	fmt.Println()

	func() {
		rmap, err := walkDir(".")
		if err != nil {
			panic(err)
		}
		for _, v := range rmap {
			fmt.Println(v)
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
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/01_os_exec.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/06_stdout_stdin_stderr.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/07_stdout_stdin_stderr_os.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample.txt
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample.csv
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/stdout.txt
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample.json
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/12_copy.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/stderr.txt
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/00_os.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/08_exist.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/11_bufio.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample.txt
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/02_io.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/09_open_create.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/15_gzip.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample.csv
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample.json
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample_copy.csv
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/05_stdin.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/13_csv.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/16_walk.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample_copy.csv
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/03_io_pipe.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/17_flush.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/stdin.txt
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/04_io_ioutil.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/10_ioutil_string.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/14_tsv.go


/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/12_copy.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/13_csv.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/11_bufio.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/01_os_exec.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/06_stdout_stdin_stderr.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/15_gzip.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/00_os.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/04_io_ioutil.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/05_stdin.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/07_stdout_stdin_stderr_os.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/08_exist.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/09_open_create.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/14_tsv.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/16_walk.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/03_io_pipe.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/17_flush.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/10_ioutil_string.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/02_io.go


testdata
testdata/sub
*/
