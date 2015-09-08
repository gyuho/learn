package main

import (
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	func() {
		fpath := "test.tar.gz"
		if err := toGzip("Hello World!", fpath); err != nil {
			panic(err)
		}
		if tb, err := gZipToBytes(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println(fpath, ":", string(tb))
			// test.tar.gz : Hello World!
		}
		os.Remove(fpath)
	}()

	func() {
		fpath := "test.tar.zlib"
		if err := toZlib("Hello World!", fpath); err != nil {
			panic(err)
		}
		if tb, err := zLibToBytes(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println(fpath, ":", string(tb))
			// test.tar.zlib : Hello World!
		}
		os.Remove(fpath)
	}()
}

// exec.Command("gzip", "-f", fpath).Run()
func toGzip(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	gw := gzip.NewWriter(f)
	if _, err := gw.Write([]byte(txt)); err != nil {
		return err
	}
	gw.Close()
	gw.Flush()
	return nil
}

func toZlib(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	zw := zlib.NewWriter(f)
	if _, err := zw.Write([]byte(txt)); err != nil {
		return err
	}
	zw.Close()
	zw.Flush()
	return nil
}

func gZipToBytes(fpath string) ([]byte, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer fz.Close()

	// or JSON
	// http://jmoiron.net/blog/crossing-streams-a-love-letter-to-ioreader/
	s, err := ioutil.ReadAll(fz)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func zLibToBytes(fpath string) ([]byte, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	z, err := zlib.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer z.Close()
	s, err := ioutil.ReadAll(z)
	if err != nil {
		return nil, err
	}
	return s, nil
}
