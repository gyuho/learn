package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	pathpkg "path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

// https://github.com/golang/go/blob/master/src/go/build/build.go#L320
func envOr(name, def string) string {
	s := os.Getenv(name)
	if s == "" {
		return def
	}
	return s
}

func main() {
	goRoot := pathpkg.Clean(runtime.GOROOT())
	fmt.Println("GOROOT:", goRoot)
	goPath := envOr("GOPATH", "")
	fmt.Println("GOPATH:", goPath)

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	projectPath, err := filepath.Rel(filepath.Join(goPath, "src"), pwd)
	if err != nil {
		panic(err)
	}

	wm, err := walkExt(".", ".go")
	if err != nil {
		panic(err)
	}
	fpaths := []string{}
	for _, v := range wm {
		fpaths = append(fpaths, filepath.Base(v))
	}

	rmap, err := importDeps(pwd, projectPath, fpaths...)
	if err != nil {
		panic(err)
	}

	for k := range rmap {
		fmt.Println(k)
	}
}

func importDeps(dir string, importPath string, fpaths ...string) (map[string]struct{}, error) {
	fSize := len(fpaths)
	if fSize == 0 {
		return nil, nil
	}
	projectPath := importPath
	il := strings.Split(importPath, "/")
	if len(il) > 2 {
		// get github.com/boltdb/bolt
		// if given 'github.com/boltdb/bolt/subpkg'
		projectPath = strings.Join(il[:3], "/")
	}
	var mu sync.Mutex // guards the map
	fmap := make(map[string]struct{})
	done, errCh := make(chan struct{}), make(chan error)
	for _, fs := range fpaths {
		go func(fs string) {
			fset := token.NewFileSet()
			fpath := filepath.Join(dir, fs)
			f, err := parser.ParseFile(fset, fpath, nil, parser.ImportsOnly)
			if err != nil {
				errCh <- err
				return
			}
			for _, elem := range f.Imports {
				pv := strings.TrimSpace(strings.Replace(elem.Path.Value, `"`, "", -1))
				mu.Lock()
				if !strings.HasPrefix(pv, projectPath) {
					fmap[pv] = struct{}{}
				}
				mu.Unlock()
			}
			done <- struct{}{}
		}(fs)
	}
	i := 0
	for {
		select {
		case e := <-errCh:
			return nil, e
		case <-done:
			i++
			if i == fSize {
				close(done)
				return fmap, nil
			}
		}
	}
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

/*
GOROOT: /usr/local/go
GOPATH: /home/gyuho/go
io
net/http
bufio
path
io/ioutil
path/filepath
syscall
go/parser
runtime
sync
strings
log
os/exec
encoding/csv
time
flag
unsafe
encoding/json
os/user
fmt
os
compress/gzip
os/signal
go/token
*/
