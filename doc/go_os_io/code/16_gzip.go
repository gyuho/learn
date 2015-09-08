package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fpath := "test.tar.gz"
	if err := toGzip("Hello World!", fpath); err != nil {
		panic(err)
	}
	if tb, err := toBytes(fpath); err != nil {
		panic(err)
	} else {
		fmt.Println(fpath, ":", string(tb))
		// test.tar.gz : Hello World!
	}
	os.Remove(fpath)
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

func toBytes(fpath string) ([]byte, error) {
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
