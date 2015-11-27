package main

import (
	"fmt"
	"go/build"
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

	func() {
		rmap, err := importDeps(pwd)
		if err != nil {
			panic(err)
		}
		for k := range rmap {
			fmt.Println("importDeps:", k)
		}
	}()

	func() {
		rmap, err := importDepsWithProjectPath(pwd, projectPath, fpaths...)
		if err != nil {
			panic(err)
		}
		for k := range rmap {
			fmt.Println("importDepsWithProjectPath:", k)
		}
	}()
}

// https://github.com/golang/go/blob/master/src/go/build/syslist.go#L7
const goosList = "android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows "
const goarchList = "386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64 "

func importDeps(dir string) (map[string]struct{}, error) {
	tm, err := walkExt(dir, ".go")
	if err != nil {
		return nil, err
	}
	wm := make(map[string]struct{})
	for _, v := range tm {
		wm[v] = struct{}{}
	}
	fSize := len(wm)
	if fSize == 0 {
		return nil, nil
	}
	var mu sync.Mutex // guards the map
	fmap := make(map[string]struct{})
	done, errCh := make(chan struct{}), make(chan error)
	for fpath := range wm {
		go func(fpath string) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, fpath, nil, parser.ImportsOnly|parser.ParseComments)
			if err != nil {
				errCh <- err
				return
			}
			ignore := false
			for _, cc := range f.Comments {
				for _, v := range cc.List {
					if strings.HasPrefix(v.Text, "// +build ignore") {
						ignore = true
						break
					}
					if strings.HasPrefix(v.Text, "// +build") {
						p := strings.Replace(v.Text, "// +build ", "", -1)
						if !strings.Contains(goosList, p) && !strings.Contains(goarchList, p) {
							ignore = true
							break
						}
					}
				}
				if ignore {
					break
				}
			}
			if !ignore {
				for _, elem := range f.Imports {
					pv := strings.TrimSpace(strings.Replace(elem.Path.Value, `"`, "", -1))
					if pv == "C" || build.IsLocalImport(pv) || strings.HasPrefix(pv, ".") {
						continue
					}
					mu.Lock()
					fmap[pv] = struct{}{}
					mu.Unlock()
				}
			}
			done <- struct{}{}
		}(fpath)
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

func importDepsWithProjectPath(dir string, importPath string, fpaths ...string) (map[string]struct{}, error) {
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
importDeps: log
importDeps: net/http
importDeps: time
importDeps: runtime
importDeps: sync
importDeps: io
importDeps: os/exec
importDeps: os/user
importDeps: go/token
importDeps: os/signal
importDeps: flag
importDeps: go/parser
importDeps: fmt
importDeps: path/filepath
importDeps: encoding/csv
importDeps: encoding/json
importDeps: bufio
importDeps: syscall
importDeps: path
importDeps: os
importDeps: io/ioutil
importDeps: strings
importDeps: compress/gzip
importDeps: unsafe
importDeps: go/build
importDepsWithProjectPath: flag
importDepsWithProjectPath: os/exec
importDepsWithProjectPath: go/build
importDepsWithProjectPath: io/ioutil
importDepsWithProjectPath: syscall
importDepsWithProjectPath: io
importDepsWithProjectPath: path/filepath
importDepsWithProjectPath: go/parser
importDepsWithProjectPath: path
importDepsWithProjectPath: runtime
importDepsWithProjectPath: encoding/json
importDepsWithProjectPath: fmt
importDepsWithProjectPath: os
importDepsWithProjectPath: os/signal
importDepsWithProjectPath: unsafe
importDepsWithProjectPath: sync
importDepsWithProjectPath: os/user
importDepsWithProjectPath: net/http
importDepsWithProjectPath: strings
importDepsWithProjectPath: time
importDepsWithProjectPath: compress/gzip
importDepsWithProjectPath: bufio
importDepsWithProjectPath: log
importDepsWithProjectPath: encoding/csv
importDepsWithProjectPath: go/token

*/
