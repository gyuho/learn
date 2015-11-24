package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	func() {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		cf, err := ioutil.TempFile(wd, "hello")
		if err != nil {
			panic(err)
		}
		fmt.Println(cf.Name())
		os.Remove(cf.Name())
	}()

	func() {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		cf, err := ioutil.TempFile(wd, "hello")
		if err != nil {
			panic(err)
		}
		op := cf.Name()
		os.Rename(op, "new_name")
		fmt.Println(op, "to new_name")
		os.Remove("new_name")
	}()

	func() {
		tmp := os.TempDir()
		f, err := ioutil.TempFile(tmp, "hello")
		if err != nil {
			panic(err)
		}
		fpath, err := filepath.Abs(f.Name())
		if err != nil {
			panic(err)
		}
		fmt.Println(fpath)

		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		npath := filepath.Join(wd, "hello")
		if err := copy(fpath, npath); err != nil {
			panic(err)
		}

		os.Remove(fpath)
		os.Remove(npath)
	}()
}

/*
0777	full access for everyone
0700	only private access
0755	private read/write access, others only read access
0750	private read/write access, group read access, others no access
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
