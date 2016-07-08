package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	const (
		fileNameTmpl     = "etcd-{{.Version}}-linux-amd64.tar.gz"
		extractedDirTmpl = "etcd-{{.Version}}-linux-amd64"
		googleTmpl       = "https://storage.googleapis.com/etcd/{{.Version}}/etcd-{{.Version}}-linux-amd64.tar.gz"
		githubTmpl       = "https://github.com/coreos/etcd/releases/download/{{.Version}}/etcd-{{.Version}}-linux-amd64.tar.gz"
	)
	ver := "v3.0.2"
	fileName := insertVersion(fileNameTmpl, ver)
	extractedDir := insertVersion(extractedDirTmpl, ver)

	os.RemoveAll(fileName)
	os.RemoveAll(extractedDir)

	if err := downloadExtractEtcd(fileName, insertVersion(googleTmpl, ver), insertVersion(githubTmpl, ver)); err != nil {
		log.Fatal(err)
	}

	os.RemoveAll(fileName)
	os.RemoveAll(extractedDir)
}

func downloadExtractEtcd(fileName string, urls ...string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, url := range urls {
		var resp *http.Response
		resp, err = http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		return err
	}
	return extractTarGz(fileName)
}

func insertVersion(tmpl, ver string) string {
	buf := new(bytes.Buffer)
	if err := template.Must(template.New("tmpl").Parse(tmpl)).Execute(buf, struct{ Version string }{ver}); err != nil {
		panic(err)
	}
	return buf.String()
}

func extractTarGz(target string) error {
	f, err := os.Open(target)
	if err != nil {
		return err
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if err := untar(tr, header); err != nil {
			return err
		}
	}
	return nil
}

func untar(tr *tar.Reader, header *tar.Header) error {
	switch header.Typeflag {
	case tar.TypeDir:
		return os.MkdirAll(header.Name, 0700)
	case tar.TypeReg, tar.TypeRegA:
		return writeFile(header.Name, tr, header.FileInfo().Mode())
	case tar.TypeSymlink:
		return writeSymlink(header.Name, header.Linkname)
	default:
		return fmt.Errorf("%s has unknown type %v", header.Name, header.Typeflag)
	}
}

func writeFile(fpath string, rd io.Reader, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(fpath), 0700); err != nil {
		return err
	}

	f, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = f.Chmod(mode); err != nil && runtime.GOOS != "windows" {
		return err
	}

	_, err = io.Copy(f, rd)
	return err
}

func writeSymlink(fpath string, target string) error {
	if err := os.MkdirAll(filepath.Dir(fpath), 0700); err != nil {
		return err
	}
	return os.Symlink(target, fpath)
}
