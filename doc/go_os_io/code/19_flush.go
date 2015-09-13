package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

type flusherWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flusherWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	fw := flusherWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}
	cmd := exec.Command("ls")
	cmd.Stdout = &fw
	cmd.Stderr = &fw
	cmd.Run()
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
00_os.go
01_os_exec.go
02_io.go
03_io_pipe.go
04_io_ioutil.go
05_stdin.go
06_stdout_stdin_stderr.go
07_stdout_stdin_stderr_os.go
08_exist.go
09_open_create.go
10_bufio.go
11_copy.go
12_flush.go
stderr.txt
stdin.txt
stdout.txt
testdata
*/
