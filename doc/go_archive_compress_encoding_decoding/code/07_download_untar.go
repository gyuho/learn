package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	dir, err := Download(os.TempDir(), "v3.0.2", "linux")
	fmt.Println("Downloaded at", dir)
	fmt.Println(err)

	os.RemoveAll(dir)
}

const (
	fileNamelLinux    = "etcd-{{.Version}}-linux-amd64.tar.gz"
	extractedDirLinux = "etcd-{{.Version}}-linux-amd64"

	fileNameDarwin     = "etcd-{{.Version}}-darwin-amd64.zip"
	extractedDirDarwin = "etcd-{{.Version}}-darwin-amd64"

	fileNameWindows     = "etcd-{{.Version}}-windows-amd64.zip"
	extractedDirWindows = "etcd-{{.Version}}-windows-amd64"

	googleTmpl = "https://storage.googleapis.com/etcd/{{.Version}}/{{.FileName}}"
	githubTmpl = "https://github.com/coreos/etcd/releases/download/{{.Version}}/{{.FileName}}"
)

// Download downloads and extracts etcd release binaries.
//
//   dir, _ = Download(os.TempDir(), "v3.0.2", "linux")
//   dir, _ = Download(os.TempDir(), "v3.0.2", "darwin")
//   dir, _ = Download(os.TempDir(), "v3.0.2", "windows")
//
// downloads the etcd v3.0.2 and extracts to 'dir' directory.
func Download(targetDir, ver, hostOS string) (string, error) {
	var (
		fileNameTmpl     string
		extractedDirTmpl string
		extractFunc      func(string, string) error
	)
	switch hostOS {
	case "linux":
		fileNameTmpl, extractedDirTmpl = fileNamelLinux, extractedDirLinux
		extractFunc = extractTarGz
	case "osx", "darwin":
		fileNameTmpl, extractedDirTmpl = fileNameDarwin, extractedDirDarwin
		extractFunc = extractZip
	case "windows":
		fileNameTmpl, extractedDirTmpl = fileNameWindows, extractedDirWindows
		extractFunc = extractZip
	default:
		return "", fmt.Errorf("unknown OS %q", hostOS)
	}

	var (
		fileName     = insertVersionFileName(fileNameTmpl, ver, "")
		extractedDir = insertVersionFileName(extractedDirTmpl, ver, "")
		urls         = []string{insertVersionFileName(googleTmpl, ver, fileName), insertVersionFileName(githubTmpl, ver, fileName)}
	)
	fileName = filepath.Join(targetDir, fileName)
	extractedDir = filepath.Join(targetDir, extractedDir)

	f, err := os.Create(fileName)
	if err != nil {
		return extractedDir, err
	}
	defer f.Close()

	for _, url := range urls {
		var resp *http.Response
		resp, err = http.Get(url)
		if err != nil {
			continue
		}
		defer gracefulClose(resp)

		if _, err = io.Copy(f, resp.Body); err != nil {
			continue
		}
		break
	}
	if err != nil {
		return extractedDir, err
	}

	return extractedDir, extractFunc(fileName, targetDir)
}

func gracefulClose(resp *http.Response) {
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
}

func extractTarGz(fileToExtract, dir string) error {
	f, err := os.Open(fileToExtract)
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

		if err := untar(tr, header, dir); err != nil {
			return err
		}
	}
	return nil
}

func extractZip(fileToExtract, dir string) error {
	zr, err := zip.OpenReader(fileToExtract)
	if err != nil {
		return err
	}
	defer zr.Close()

	for _, zf := range zr.File {
		if err := unzip(zf, dir); err != nil {
			return err
		}
	}

	return nil
}

func untar(tr *tar.Reader, header *tar.Header, dir string) error {
	switch header.Typeflag {
	case tar.TypeDir:
		return os.MkdirAll(filepath.Join(dir, header.Name), 0700)
	case tar.TypeReg, tar.TypeRegA:
		return writeFile(filepath.Join(dir, header.Name), tr, header.FileInfo().Mode())
	case tar.TypeSymlink:
		return writeSymlink(filepath.Join(dir, header.Name), header.Linkname)
	default:
		return fmt.Errorf("%s has unknown type %v", header.Name, header.Typeflag)
	}
}

func unzip(zf *zip.File, dir string) error {
	if strings.HasSuffix(zf.Name, "/") {
		return os.MkdirAll(filepath.Join(dir, zf.Name), 0700)
	}

	rc, err := zf.Open()
	if err != nil {
		return fmt.Errorf("%s: open compressed file: %v", zf.Name, err)
	}
	defer rc.Close()

	return writeFile(filepath.Join(dir, zf.Name), rc, zf.FileInfo().Mode())
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

func insertVersionFileName(tmpl, ver, fileName string) string {
	buf := new(bytes.Buffer)
	if err := template.Must(template.New("tmpl").Parse(tmpl)).Execute(buf, struct{ Version, FileName string }{ver, fileName}); err != nil {
		panic(err)
	}
	return buf.String()
}
