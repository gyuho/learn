package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// openToOverwrite creates or opens a file for overwriting.
// Make sure to close the file.
func openToOverwrite(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}

// openToRead reads a file.
// Make sure to close the file.
func openToRead(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return f, err
	}
	return f, nil
}

// http://golang.org/pkg/archive/tar/
func main() {
	func() {
		fpath := "my.tar"
		defer os.Remove(fpath)

		func() {
			// Open to write a tar file.
			f, err := openToOverwrite(fpath)
			defer f.Close()
			if err != nil {
				panic(err)
			}

			// buf := new(bytes.Buffer) // Create a buffer to write our archive to.
			// tw := tar.NewWriter(buf) // Create a new tar archive.
			tw := tar.NewWriter(f)

			// Add some files to the archive.
			var files = []struct {
				Name, Body string
			}{
				{"readme.txt", "This archive contains some text files."},
				{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
				{"todo.txt", "Get animal handling licence."},
			}
			for _, file := range files {
				hdr := &tar.Header{
					Name: file.Name,
					Mode: 0600,
					Size: int64(len(file.Body)),
				}
				if err := tw.WriteHeader(hdr); err != nil {
					panic(err)
				}
				if _, err := tw.Write([]byte(file.Body)); err != nil {
					panic(err)
				}
			}
			// Make sure to check the error on Close.
			if err := tw.Close(); err != nil {
				panic(err)
			}
		}()

		fmt.Println()

		func() {
			// Open the tar archive for reading.
			//
			// r := bytes.NewReader(buf.Bytes())
			// tr := tar.NewReader(r)
			f, err := openToRead(fpath)
			defer f.Close()
			if err != nil {
				panic(err)
			}
			tr := tar.NewReader(f)
			for {
				hdr, err := tr.Next()
				if err == io.EOF {
					// end of tar archive
					break
				}
				if err != nil {
					panic(err)
				}
				fmt.Println()
				fmt.Printf("Contents of %s:\n", hdr.Name)
				if _, err := io.Copy(os.Stdout, tr); err != nil {
					panic(err)
				}
				fmt.Println()
			}
		}()
	}()

	func() {
		fpath := "my.zip"
		defer os.Remove(fpath)

		func() {
			f, err := openToOverwrite(fpath)
			defer f.Close()
			if err != nil {
				panic(err)
			}
			zw := zip.NewWriter(f)
			var files = []struct {
				Name, Body string
			}{
				{"readme.txt", "This archive contains some text files."},
				{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
				{"todo.txt", "Get animal handling licence."},
			}
			for _, file := range files {
				fz, err := zw.Create(file.Name)
				if err != nil {
					panic(err)
				}
				if _, err := fz.Write([]byte(file.Body)); err != nil {
					panic(err)
				}
			}
			// Make sure to check the error on Close.
			if err := zw.Close(); err != nil {
				panic(err)
			}
		}()

		fmt.Println()

		func() {
			fz, err := zip.OpenReader(fpath)
			defer fz.Close()
			if err != nil {
				panic(err)
			}
			for _, oneFile := range fz.File {
				rc, err := oneFile.Open()
				if err != nil {
					panic(err)
				}
				fmt.Println()
				fmt.Printf("Contents of %s:\n", oneFile.Name)
				if _, err := io.Copy(os.Stdout, rc); err != nil {
					panic(err)
				}
				fmt.Println()

				// defer runs when the for loop ends
				rc.Close()
			}
		}()
	}()
}
