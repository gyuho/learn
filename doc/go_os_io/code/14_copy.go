package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

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

func copyToTempFile(src, tempPrefix string) (string, error) {
	r, err := os.Open(src)
	if err != nil {
		return "", fmt.Errorf("copy: open(%q): %v", src, err)
	}
	defer r.Close()

	w, err := ioutil.TempFile("", tempPrefix)
	if err != nil {
		return "", fmt.Errorf("ioutil.TempFile error: %+v", err)
	}
	defer w.Close()

	if _, err = io.Copy(w, r); err != nil {
		return "", err
	}
	if err := w.Sync(); err != nil {
		return "", err
	}
	if _, err := w.Seek(0, 0); err != nil {
		return "", err
	}
	return w.Name(), nil
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

func main() {
	func() {
		fpath := "test.txt"
		defer os.Remove(fpath)
		for _, k := range []int{0, 1} {
			f, err := openToAppend(fpath)
			if err != nil {
				panic(err)
			}
			if _, err := f.WriteString(fmt.Sprintf("Hello World! %d\n", k)); err != nil {
				panic(err)
			}
			f.Close()
		}
		f, err := openToRead(fpath)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		tbytes, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println("fpath:", string(tbytes))
		/*
		   fpath: Hello World! 0
		   Hello World! 1
		*/

		fpathCopy := "test_copy.txt"
		defer os.Remove(fpathCopy)
		if err := copy(fpath, fpathCopy); err != nil {
			panic(err)
		}
		fc, err := openToRead(fpathCopy)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		tbc, err := ioutil.ReadAll(fc)
		if err != nil {
			panic(err)
		}
		fmt.Println("fpathCopy:", string(tbc))
		/*
		   fpathCopy: Hello World! 0
		   Hello World! 1
		*/
	}()

	func() {
		fpath := "test.txt"
		defer os.Remove(fpath)
		for _, k := range []int{0, 1} {
			f, err := openToAppend(fpath)
			if err != nil {
				panic(err)
			}
			if _, err := f.WriteString(fmt.Sprintf("Hello World! %d\n", k)); err != nil {
				panic(err)
			}
			f.Close()
		}
		tempPath, err := copyToTempFile(fpath, "temp_prefix_")
		if err != nil {
			panic(err)
		}
		fmt.Println("tempPath:", tempPath)
	}()

	func() {
		if err := copyDir("testdata", "testdata2"); err != nil {
			panic(err)
		}
		os.RemoveAll("testdata2")
	}()
}

func openToRead(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return f, err
	}
	return f, nil
}

func openToAppend(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}
