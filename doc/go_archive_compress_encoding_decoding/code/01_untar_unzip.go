package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// http://golang.org/pkg/archive/tar/
func main() {
	func() {
		fpath := "my.tar.gz"
		defer os.Remove(fpath)
		writeToFiles()
		writeToTar(filePathSlice, fpath)
		deleteFiles()
		untar(fpath)
		deleteFiles()
	}()

	func() {
		fpath := "my.zip"
		defer os.Remove(fpath)
		writeToFiles()
		writeToZip(filePathSlice, fpath)
		deleteFiles()
		unzip(fpath)
		deleteFiles()
	}()
}

func writeToTar(filePathSlice []string, tarPath string) {
	f, err := openToOverwrite(tarPath)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	gw := gzip.NewWriter(f)
	defer gw.Close()
	if err != nil {
		panic(err)
	}
	tw := tar.NewWriter(gw)
	for _, fpath := range filePathSlice {
		sf, err := openToRead(fpath)
		defer sf.Close()
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(sf)
		if err != nil {
			panic(err)
		}
		hdr := &tar.Header{
			Name: fpath,
			Mode: 0600,
			Size: int64(len(body)),

			// Need to set typeflag
			//
			// http://www.gnu.org/software/tar/manual/html_node/Standard.html
			// #define REGTYPE  '0'            /* regular file */
			Typeflag: byte('0'),

			// or use
			// http://golang.org/pkg/archive/tar/#FileInfoHeader
		}
		if err := tw.WriteHeader(hdr); err != nil {
			panic(err)
		}
		if _, err := tw.Write(body); err != nil {
			panic(err)
		}
	}
	if err := tw.Close(); err != nil {
		panic(err)
	}
}

func untar(fpath string) {
	f, err := openToRead(fpath)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	gr, err := gzip.NewReader(f)
	defer gr.Close()
	if err != nil {
		panic(err)
	}
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		path := hdr.Name
		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, os.FileMode(hdr.Mode)); err != nil {
				panic(err)
			}
		case tar.TypeReg:
			ow, err := openToOverwrite(path)
			defer ow.Close()
			if err != nil {
				panic(err)
			}
			if _, err := io.Copy(ow, tr); err != nil {
				panic(err)
			}
		default:
			fmt.Printf("Unable to untar: %c, %s\n", hdr.Typeflag, path)
		}
	}
}

func writeToZip(filePathSlice []string, zipPath string) {
	f, err := openToOverwrite(zipPath)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	zw := zip.NewWriter(f)
	for _, fpath := range filePathSlice {
		fz, err := zw.Create(fpath)
		if err != nil {
			panic(err)
		}
		sf, err := openToRead(fpath)
		defer sf.Close()
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(sf)
		if err != nil {
			panic(err)
		}
		if _, err := fz.Write(body); err != nil {
			panic(err)
		}
	}
	if err := zw.Close(); err != nil {
		panic(err)
	}
}

func unzip(fpath string) {
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
		path := filepath.Join("./", oneFile.Name)
		if oneFile.FileInfo().IsDir() {
			if err := os.MkdirAll(path, oneFile.Mode()); err != nil {
				panic(err)
			}
		} else {
			ow, err := openToOverwrite(path)
			defer ow.Close()
			if err != nil {
				panic(err)
			}
			if _, err := io.Copy(ow, rc); err != nil {
				panic(err)
			}
		}
		rc.Close()
	}
}

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

var files = []struct {
	Name, Body string
}{
	{"readme.txt", "This archive contains some text files."},
	{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
	{"todo.txt", "Get animal handling licence."},
}

var filePathSlice = []string{
	"readme.txt",
	"gopher.txt",
	"todo.txt",
}

func writeToFiles() {
	for _, file := range files {
		f, err := openToOverwrite(file.Name)
		if err != nil {
			panic(err)
		}
		if _, err := f.WriteString(file.Body); err != nil {
			panic(err)
		}
		f.Close()
	}
}

func deleteFiles() {
	for _, file := range files {
		os.Remove(file.Name)
	}
}
