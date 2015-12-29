package main

import (
	"fmt"
	"os"
	"path"
)

var (
	path0 = "temporary"
	path1 = "member"
	path2 = "file.txt"

	fpath0 = path0
	fpath1 = path.Join(path0, path1)
	fpath2 = path.Join(path.Join(path0, path1), path2)
)

var isDelete bool

func init() {
	if err := os.RemoveAll(fpath0); err != nil {
		panic(err)
	}
	fmt.Println("fpath0:", fpath0)
	fmt.Println("fpath1:", fpath1)
	fmt.Println("fpath2:", fpath2)
}

func main() {
	defer func() {
		if isDelete {
			os.RemoveAll(fpath0)
		}
	}()

	if err := os.MkdirAll(fpath0, 0700); err != nil {
		panic(err)
	}

	if existDir(fpath1) {
		fmt.Println(fpath1, "already exists... skipping...")
		return
	}

	if err := os.MkdirAll(fpath1, 0700); err != nil {
		panic(err)
	}

	if err := toFileWriteString("hello world!", fpath2); err != nil {
		panic(err)
	}

	fmt.Println("Done")
	isDelete = true
}

/*
fpath0: temporary
fpath1: temporary/member
fpath2: temporary/member/file.txt
Done
*/

// existDir returns true if the specified path points to a directory.
// It returns false and error if the directory does not exist.
func existDir(fpath string) bool {
	st, err := os.Stat(fpath)
	if err != nil {
		return false
	}
	return st.IsDir()
}

func toFileWriteString(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	if _, err := f.WriteString(txt); err != nil {
		return err
	}
	return nil
}
