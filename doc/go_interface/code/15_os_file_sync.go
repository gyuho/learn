package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := openToAppend("test.log")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	var wc WriteSyncer
	wc = &osFile{file: f}
	_, err = wc.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	if err = wc.Sync(); err != nil {
		log.Fatal(err)
	}
	// calling *osFile.Write
	// calling *osFile.Sync
}

// A WriteSyncer is an io.Writer that can also flush any buffered data. Note
// that *os.File (and thus, os.Stderr and os.Stdout) implement WriteSyncer.
type WriteSyncer interface {
	io.Writer
	Sync() error
}

type writerWrapper struct {
	io.Writer
}

// in case io.Writer does not implement Sync
// (if io.Writer implements Sync (like *os.File), this does not get called)
func (w writerWrapper) Sync() error {
	fmt.Println("calling writerWrapper.Sync")
	return nil
}

type osFile struct {
	file *os.File
}

func (f *osFile) Write(p []byte) (n int, err error) {
	fmt.Println("calling *osFile.Write")
	return f.file.Write(p)
}

func (f *osFile) Sync() error {
	fmt.Println("calling *osFile.Sync")
	return f.file.Sync()
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
