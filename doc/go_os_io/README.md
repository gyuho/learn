[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: os, io

- [package `os`](#package-os)
- [package `os/exec`](#package-osexec)
- [package `flag`](#package-flag)
- [package `io`](#package-io)
- [`io.Pipe`](#iopipe)
- [package `io/ioutil`](#package-ioioutil)
- [stdout, stdin, stderr](#stdout-stdin-stderr)
- [`exist`: files, directories](#exist-files-directories)
- [`create/open/write`: files, directories](#createopenwrite-files-directories)
- [`io/ioutil`, file](#ioioutil-file)
- [temporary file](#temporary-file)
- [`bufio`](#bufios)
- [`copy`: files, directories](#copy-files-directories)
- [`csv`](#csv)
- [`tsv`](#tsv)
- [`compress/gzip`](#compressgzip)
- [walk](#walk)
- [`http.Flusher`](#httpflusher)
- [`os.Signal`](#ossignal)
- [netstat](#netstat)

[↑ top](#go-os-io)
<br><br><br><br>
<hr>







#### package `os`

[Package `os`](http://golang.org/pkg/os) provides a platform-independent
interface to operating system functionality. Note that [`os.File`](http://golang.org/pkg/os/#File) 
implements `Read(b []byte) (n int, err error)` and 
`Write(b []byte) (n int, err error)` methods, therefore satisfying
[`io.Reader`](http://golang.org/pkg/io/#Reader) and 
[`io.Writer`](http://golang.org/pkg/io/#Writer) interface.

Try this code:

```go
package main

import (
	"fmt"
	"os"
	"os/user"
	"time"
)

func main() {
	fmt.Println("TempDir:", os.TempDir())

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// get the current user
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	homePath := usr.HomeDir

	// change the directory
	if err := os.Chdir(homePath); err != nil {
		panic(err)
	}

	// get current working directory
	if twd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		fmt.Println("home pwd:", twd)
	}
	// home pwd: /home/ubuntu

	if err := os.Chdir(pwd); err != nil {
		panic(err)
	}
	// get current working directory
	if twd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		fmt.Println("home pwd:", twd)
	}

	fpath := "temp.txt"

	// this errors for non-existing file.
	if err := os.Remove(fpath); err != nil {
		// panic(err)
		fmt.Println(err)
	}

	// this errors for non-existing file.
	file, err := os.Open(fpath)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		// open temp.txt: no such file or directory

		// THIS DOES NOT GET CALLED
		// BECAUSE the process will be killed
		// in the next lines
		defer func() {
			fmt.Println("Deleting", fpath)
			if err := os.Remove(fpath); err != nil {
				// panic(err)
				fmt.Println(err)
			}
		}()

		fmt.Println("Creating", fpath)
		file, err = os.Create(fpath)
		if err != nil {
			panic(err)
		}
		fmt.Println(file)
	}

	fmt.Println("Deleting", fpath)
	if err := os.Remove(fpath); err != nil {
		// panic(err)
		fmt.Println(err)
	}

	// get process id
	pid := os.Getppid()
	fmt.Println("pid:", pid)

	// find the process
	p, err := os.FindProcess(pid)
	if err != nil {
		panic(err)
	}
	go func() {
		fmt.Println("goroutine: killing the process in 1 second...")
		time.Sleep(time.Second)
		if err := p.Kill(); err != nil {
			panic(err)
		}
	}()
	fmt.Println("Sleeping 1,000 hours in main function...")
	time.Sleep(1000 * time.Hour)
}

/*
remove temp.txt: no such file or directory
open temp.txt: no such file or directory
Creating temp.txt
&{0xc8200164b0}
Deleting temp.txt
pid: 16635
Sleeping 1,000 hours in main function...
goroutine: killing the process in 1 second...
Killed
*/
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>









#### package `os/exec`

[Package `os/exec`](http://golang.org/pkg/os/exec/) runs external commands.

Try this:

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("pwd:", pwd)

	lsCmd := exec.Command("ls", "-a")

	lsOutput1, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(lsOutput1))

	if err := exec.Command("touch", "temp.txt").Run(); err != nil {
		panic(err)
	}

	// lsOutput2, err := lsCmd.Output()
	// if err != nil {
	// 	panic(err)
	// }
	// panic: exec: Stdout already set
	// SHOULD NOT REUSE IT

	lsCmd2 := exec.Command("ls", "-a")
	lsOutput2, err := lsCmd2.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(lsOutput2))

	if err := os.Remove("temp.txt"); err != nil {
		panic(err)
	}

	lsCmd3 := exec.Command("ls", "-a")
	lsOutput3, err := lsCmd3.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(lsOutput3))

	// http://golang.org/pkg/os/exec/#example_Cmd_Start
	slCmd := exec.Command("sleep", "1")
	if err := slCmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("Waiting for command to finish...")
	err = slCmd.Wait()
	fmt.Println("Command finished with error:", err)

	// http://golang.org/pkg/os/exec/#Cmd.Output
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		panic(err)
	}
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
	// Bob is 32 years old
}

/*
pwd: /home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code
.
..
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
10_ioutil_string.go
11_bufio.go
12_copy.go
13_csv.go
14_tsv.go
15_gzip.go
16_walk.go
17_flush.go
stderr.txt
stdin.txt
stdout.txt
testdata

.
..
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
10_ioutil_string.go
11_bufio.go
12_copy.go
13_csv.go
14_tsv.go
15_gzip.go
16_walk.go
17_flush.go
stderr.txt
stdin.txt
stdout.txt
temp.txt
testdata

.
..
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
10_ioutil_string.go
11_bufio.go
12_copy.go
13_csv.go
14_tsv.go
15_gzip.go
16_walk.go
17_flush.go
stderr.txt
stdin.txt
stdout.txt
testdata

Waiting for command to finish...
Command finished with error: <nil>
Bob is 32 years old
*/
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>









#### package `flag`

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	idxPtr := flag.Int(
		"index",
		0,
		"Specify the index.",
	)
	dscPtr := flag.String(
		"description",
		"None",
		"Describe the argument.",
	)
	flag.Parse()
	fmt.Println("index:", *idxPtr)
	fmt.Println("description:", *dscPtr)
}

/*
go run 02_flag.go -h

Usage of /tmp/go-build105642507/command-line-arguments/_obj/exe/02_flag:
  -description string
        Describe the argument. (default "None")
  -index int
        Specify the index.
exit status 2



go run 02_flag.go -index=10 \
-description="Hello World!"
;

index: 10
description: Hello World!
*/

```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>




#### package `io`

[Package `io`](http://golang.org/pkg/io) provides basic interfaces to I/O primitives.
When other packages, other than `io`, implement methods and satisfy interfaces in `io` package,
those different packages can interact based on the shared `io` interfaces. 

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
```

- [`io.Reader`](http://golang.org/pkg/io/#Reader) interface has method `Read`.
- Any type that implements `Read(p []byte) (n int, err error)` method satisfies `io.Reader` interface.
- [`io.Writer`](http://golang.org/pkg/io/#Writer) interface has method `Write`.
- Any type that implements `Write(p []byte) (n int, err error)` method satisfies `io.Writer` interface.
- As long as a type satisfies the interface, it can be used as an function argument.
- As long as a type satisfies the interface, it can contain data.

For example,
- [`os.File`](http://golang.org/pkg/os/#File) implements `Read` and `Write` methods.
- Therefore, it satisfies `io.Reader` and `io.Writer` interfaces.
- And [`json.NewDecoder`](http://golang.org/pkg/encoding/json/#NewDecoder) takes `io.Reader` as an argument.
- Since `os.File` satisfies `io.Reader` interface, we can pass `os.File` as an argument to `json.NewDecoder`.

```go
func (f *File) Read(b []byte) (n int, err error) {
    if f == nil {
        return 0, ErrInvalid
    }
    n, e := f.read(b)
    if n < 0 {
        n = 0
    }
    if n == 0 && len(b) > 0 && e == nil {
        return 0, io.EOF
    }
    if e != nil {
        err = &PathError("read", f.name, e}
    }
    return n, err
}
```

<br>
Try this:

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// And os.File implements Read and Write method
// therefore satisfies io.Reader and io.Writer method

func main() {
	fpath := "testdata/sample.json"

	file, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	tbytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	jsonStream := string(tbytes)
	decodeString(jsonStream)
	// map[Go:Gopher Hello:World]

	decodeFile(file)
	// map[]

	// need to open again
	file2, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	decodeFile(file2)
	// map[Go:Gopher Hello:World]
}

func decodeFile(file *os.File) {
	rmap := map[string]string{}
	dec := json.NewDecoder(file)
	for {
		if err := dec.Decode(&rmap); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%+v\n", rmap)
}

func decodeString(jsonStream string) {
	rmap := map[string]string{}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		if err := dec.Decode(&rmap); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%+v\n", rmap)
}
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>









#### `io.Pipe`

[`io.Pipe`](http://golang.org/pkg/io/#Pipe) creates a synchronous in-memory pipe.

```go
package main

import (
	"fmt"
	"io"
)

func main() {
	done := make(chan struct{})
	r, w := io.Pipe()

	go func() {
		data := []byte("Hello World!")
		n, err := w.Write(data)
		if err != nil {
			panic(err)
		}
		if n != len(data) {
			panic(data)
		}
		done <- struct{}{}
	}()

	buf := make([]byte, 64)
	n, err := r.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("wrote:", n)         // 12
	fmt.Println("buf:", string(buf)) // Hello World!

	<-done

	r.Close()
	w.Close()
}
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>











#### package `io/ioutil`

[package `io/ioutil`](http://golang.org/pkg/io/ioutil) implements some I/O utility functions.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bts))
	// <div id=gbar><nobr><b class=gb1>Search</b>...
}
```

<br>
<br>
[`ioutil.ReadAll`](http://golang.org/pkg/io/ioutil/#ReadAll) takes `io.Reader` as an argument:

```go
func ReadAll(r io.Reader) ([]byte, error)
```

<br>
<br>
[`http.Response`](http://golang.org/pkg/net/http/#Response) struct 
embeds `io.ReadCloser` interface.:

```go
type Response struct {
	...
	Body io.ReadCloser
}
```

<br>
<br>
[`io.ReadCloser`](http://golang.org/pkg/io/#ReadCloser) interface embeds `io.Reader` interface:

```go
type ReadCloser interface {
	Reader
	Closer
}
```

Therefore, `http.Response.Body` of type `io.ReadCloser` is also type `io.Reader`.
Any type that satisfies `io.Reader` interface (*implements* `Read` method)
can be passed to `ioutil.ReadAll`. 

```go
ioutil.ReadAll(resp.Body)
```


<br>
Then does `http.Response.Body` type implement `Read` method? *No.*
`http.Response.Body` is used as `io.Reader` interface argument.
Then you might think `http.Response.Body` type should implement 
`Read` method to satisfy the `io.Reader` interface but it doesn not. 
Neither does `Close`. `http.Response.Body` can hold any type that
implements `Read` and `Close`, but the **actual type** depends on 
the server that you are getting response from, as
[here](https://github.com/golang/go/blob/master/src/net/http/transfer.go):

```go
	// Prepare body reader.  ContentLength < 0 means chunked encoding
	// or close connection when finished, since multipart is not supported yet
	switch {
	case chunked(t.TransferEncoding):
		if noBodyExpected(t.RequestMethod) {
			t.Body = eofReader
		} else {
			t.Body = &body{src: internal.NewChunkedReader(r), hdr: msg, r: r, closing: t.Close}
		}
	case realLength == 0:
		t.Body = eofReader
	case realLength > 0:
		t.Body = &body{src: io.LimitReader(r, realLength), closing: t.Close}
	default:
		// realLength < 0, i.e. "Content-Length" not mentioned in header
		if t.Close {
			// Close semantics (i.e. HTTP/1.0)
			t.Body = &body{src: r, closing: t.Close}
		} else {
			// Persistent connection (i.e. HTTP/1.1)
			t.Body = eofReader
		}
	}
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>










#### stdout, stdin, stderr

To get the input from terminal:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	fmt.Print("input:", input)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("scanner:", scanner.Text())
	}
	// Ctrl + D
}

/*
$ go run 05_stdin.go
Enter text: Hello
input:Hello
scanner:
1
scanner: 1
2
scanner: 2
3
scanner: 3
4
scanner: 4
5
scanner: 5
*/
```

<br>
And note that `log` output goes to standard err (file descriptor `2`):

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("fmt.Println")
	log.Println("log.Println")
	// log.Fatal("log.Fatal")
	panic("panic")
}

/*
go run 06_stdout_stdin_stderr.go Hello 0>>stdin.txt 1>>stdout.txt 2>>stderr.txt


stdin.txt
<empty>

stdout.txt
fmt.Println




log.Println goes to standard err

stderr.txt
2015/08/05 06:09:32 log.Println
2015/08/05 06:09:32 log.Fatal
exit status 1

or

2015/08/05 06:10:42 log.Println
panic: panic

goroutine 1 [running]:
main.main()
	/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/05_stdout_stdin_stderr.go:12 +0x1e4
exit status 2
*/
```

<br>
And here's how you interact with OS standard output, input, error:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello World!")
	// Hello World!

	fmt.Fprintln(os.Stdin, "Input")

	fmt.Fprintln(os.Stderr, "Error")
}

// go run 07_stdout_stdin_stderr_os.go 0>>stdin.txt 1>>stdout.txt 2>>stderr.txt
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>









#### `exist`: files, directories

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

// exist returns true if the file or directory exists.
func exist(fpath string) bool {
	// Does a directory exist
	st, err := os.Stat(fpath)
	if err != nil {
		return false
	}
	if st.IsDir() {
		return true
	}
	if _, err := os.Stat(fpath); err != nil {
		if strings.Contains(err.Error(), "no such file") {
			return false
		}
	}
	return true
}

// existDir returns true if the specified path points to a directory.
// It returns false and error if the directory does not exist.
func existDir(fpath string) bool {
	st, err := os.Stat(fpath)
	if err != nil {
		return false
	}
	return st.IsDir()
}

func main() {
	fmt.Println(exist("00_os.go"))    // true
	fmt.Println(exist("aaaaa.go"))    // false
	fmt.Println(exist("testdata"))    // true
	fmt.Println(existDir("testdata")) // true
}
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>










#### `create/open/write`: files, directories

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// openToRead reads a file.
// Make sure to close the file.
func openToRead(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return f, err
	}
	return f, nil
}

// openToOverwrite creates or opens a file for overwriting.
// Make sure to close the file.
func openToOverwrite(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}

// openToAppend creates a file if it does not exist.
// Otherwise it opens a file.
// Records that are written are to be appended.
// Make sure to close the file.
func openToAppend(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}

func main() {
	fpath := "./testdata/sample.txt"

	func() {
		f, err := openToRead(fpath)
		if err != nil {
			panic(err)
		}
		defer func() {
			fmt.Println("Closing", f.Name())
			f.Close()
		}()
		if f.Name() != fpath {
			panic(f.Name())
		}
		tbytes, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(tbytes))
	}()
	/*
	   Hello World!
	   Closing ./testdata/sample.txt
	*/

	fmt.Println()
	fmt.Println()

	func() {
		fpath := "test.txt"
		for range []int{0, 1} {
			f, err := openToOverwrite(fpath)
			if err != nil {
				panic(err)
			}
			if _, err := f.WriteString("Hello World!"); err != nil {
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
		fmt.Println(string(tbytes))
		if err := os.Remove(fpath); err != nil {
			panic(err)
		}
	}()
	// Hello World!

	fmt.Println()
	fmt.Println()

	func() {
		fpath := "test.txt"
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
		fmt.Println(string(tbytes))
		if err := os.Remove(fpath); err != nil {
			panic(err)
		}
	}()
	/*
	   Hello World! 0
	   Hello World! 1
	*/
}

```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>








#### `io/ioutil`, file

Package [ioutil](http://golang.org/pkg/io/ioutil/) implements some I/O utility functions.

```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func main() {
	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFileWriteString(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenReadAll(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println("fromFileOpenReadAll:", s)
		}
	}()
	// fromFileOpenReadAll: Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFileIO(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenReadAll(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println("fromFileOpenReadAll:", s)
		}
	}()
	// fromFileOpenReadAll: Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFileWrite(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenReadAll(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println("fromFileOpenReadAll:", s)
		}
	}()
	// fromFileOpenReadAll: Hello World!

	func() {
		fpath := "toFileWriteSyscall.txt"
		txt := "Hello World!"
		if err := toFileWriteSyscall(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenReadAll(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println("toFileWriteSyscall and fromFileOpenReadAll:", s)
		}
	}()
	// toFileWriteSyscall and fromFileOpenReadAll: Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFileWriteString(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenFileReadAll(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println("fromFileOpenFileReadAll:", s)
		}
	}()
	// fromFileOpenFileReadAll: Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFileWriteString(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenFileReadFull(fpath, len(txt)); err != nil {
			panic(err)
		} else {
			fmt.Println("fromFileOpenFileReadFull:", s)
		}
	}()
	// fromFileOpenFileReadFull: Hello World!

	func() {
		fpath := "temp.txt"
		txt := strings.Repeat("Hello World!", 10000)
		if err := toFileWriteString(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		isSupported := isDirectIOSupported(fpath)
		fmt.Println("isDirectIOSupported:", isSupported)
		if isSupported {
			if s, err := fromFileDirectIO(fpath); err != nil {
				panic(err)
			} else {
				fmt.Println("fromFileDirectIO:", s[:10], "...")
			}
		}
	}()
	// isDirectIOSupported: true
	// fromFileDirectIO: Hello Worl ...
}

func toFileWriteString(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	if _, err := f.WriteString(txt); err != nil {
		return err
	}
	return nil
}

func toFileIO(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	if _, err := io.WriteString(f, txt); err != nil {
		return err
	}
	return nil
}

func toFileWrite(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	if _, err := f.Write([]byte(txt)); err != nil {
		return err
	}
	return nil
}

func toFileWriteSyscall(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|syscall.MAP_POPULATE, 0777)
	if err != nil {
		// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	if _, err := f.Write([]byte(txt)); err != nil {
		return err
	}
	return nil
}

func fromFileOpenReadAll(fpath string) (string, error) {
	f, err := os.Open(fpath)
	if err != nil {
		// NOT retur nil, err
		// []byte can be null but not string
		return "", err
	}
	defer f.Close()
	// func ReadAll(r io.Reader) ([]byte, error)
	tbytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(tbytes), nil
}

func fromFileOpenFileReadAll(fpath string) (string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0777)
	if err != nil {
		// NOT retur nil, err
		// []byte can be null but not string
		return "", err
	}
	defer f.Close()
	// func ReadAll(r io.Reader) ([]byte, error)
	tbytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(tbytes), nil
}

func fromFileOpenFileReadFull(fpath string, length int) (string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0777)
	if err != nil {
		// NOT retur nil, err
		// []byte can be null but not string
		return "", err
	}
	defer f.Close()
	buf := make([]byte, length)
	if _, err := io.ReadFull(f, buf); err != nil {
		return "", err
	}
	return string(buf), nil
}

func isDirectIOSupported(fpath string) bool {
	f, err := os.OpenFile(fpath, syscall.O_DIRECT, 0)
	defer f.Close()
	return err == nil
}

func fromFileDirectIO(fpath string) (string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY|syscall.O_DIRECT, 0777)
	if err != nil {
		return "", err
	}
	defer f.Close()
	block := AlignedBlock(BlockSize)
	if _, err := io.ReadFull(f, block); err != nil {
		return "", err
	}
	return string(block), nil
}

/*****************************************************/

// Copied from https://github.com/ncw/directio

// alignment returns alignment of the block in memory
// with reference to AlignSize
//
// Can't check alignment of a zero sized block as &block[0] is invalid
func alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}

// AlignedBlock returns []byte of size BlockSize aligned to a multiple
// of AlignSize in memory (must be power of two)
func AlignedBlock(BlockSize int) []byte {
	block := make([]byte, BlockSize+AlignSize)
	if AlignSize == 0 {
		return block
	}
	a := alignment(block, AlignSize)
	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}
	block = block[offset : offset+BlockSize]
	// Can't check alignment of a zero sized block
	if BlockSize != 0 {
		a = alignment(block, AlignSize)
		if a != 0 {
			log.Fatal("Failed to align block")
		}
	}
	return block
}

const (
	// Size to align the buffer to
	AlignSize = 4096

	// Minimum block size
	BlockSize = 4096
)

// OpenFile is a modified version of os.OpenFile which sets O_DIRECT
func OpenFile(name string, flag int, perm os.FileMode) (file *os.File, err error) {
	return os.OpenFile(name, syscall.O_DIRECT|flag, perm)
}

/*****************************************************/

```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>









#### temporary file

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// to create in the current directory.
	// f, err := ioutil.TempFile(".", "temp_prefix_")

	// creates in  TempFile uses the default directory
	// for temporary files (see os.TempDir)
	f, err := ioutil.TempFile("", "temp_prefix_")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	if err := f.Sync(); err != nil {
		panic(err)
	}
	if _, err := f.Seek(0, 0); err != nil {
		panic(err)
	}

	fmt.Println(os.TempDir())            // /tmp
	fmt.Println(f.Name())                // /tmp/temp_prefix_289175735
	fmt.Println(filepath.Base(f.Name())) // temp_prefix_289175735
	fmt.Println(filepath.Dir(f.Name()))  // /tmp
}

```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>








#### `bufio`

Package [`bufio`](http://golang.org/pkg/bufio/) implements buffered I/O.
It wraps an `io.Reader` or `io.Writer` object, creating another object:


```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	func() {
		if err := fromLines([]string{"A", "B", "C"}, "./tmp.txt"); err != nil {
			panic(err)
		}
		defer os.RemoveAll("./tmp.txt")
		lines, err := toLines1("./tmp.txt")
		if err != nil {
			panic(err)
		}
		if len(lines) != 3 {
			panic(fmt.Errorf("expected 3 but %v", lines))
		}
	}()

	func() {
		if err := fromLines([]string{"A", "B", "C"}, "./tmp.txt"); err != nil {
			panic(err)
		}
		defer os.RemoveAll("./tmp.txt")
		lines, err := toLines2("./tmp.txt")
		if err != nil {
			panic(err)
		}
		if len(lines) != 3 {
			panic(fmt.Errorf("expected 3 but %v", lines))
		}
	}()

	func() {
		if err := fromLines([]string{"aaa bbb ccc"}, "./tmp.txt"); err != nil {
			panic(err)
		}
		defer os.RemoveAll("./tmp.txt")
		words, err := toWords("./tmp.txt")
		if err != nil {
			panic(err)
		}
		if len(words) != 3 {
			panic(fmt.Errorf("expected 3 but %v", words))
		}
	}()

	func() {
		fpath := "stdout.txt"
		d, err := toBytes(fpath)
		if err != nil {
			panic(err)
		}
		fmt.Println("toBytes:", string(d))
		/*
		   toBytes: Enter text: input:fmt.Println
		   Enter text: input:fmt.Println
		   Hello World!
		*/
	}()
}

func fromLines(lines []string, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	// func NewWriter(w io.Writer) *Writer
	wr := bufio.NewWriter(f)

	for _, line := range lines {
		// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
		fmt.Fprintln(wr, line)
	}

	if err := wr.Flush(); err != nil {
		return err
	}
	return nil
}

func toLines1(fpath string) ([]string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := []string{}

	// func NewScanner(r io.Reader) *Scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

func toLines2(fpath string) ([]string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rs := []string{}
	br := bufio.NewReader(f)
	for {
		l, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		rs = append(rs, l)
	}
	return rs, nil
}

func toWords(fpath string) ([]string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := []string{}

	// func NewScanner(r io.Reader) *Scanner
	scanner := bufio.NewScanner(f)

	// This must be called before Scan.
	// The default split function is bufio.ScanLines.
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

func toBytes(fpath string) ([]byte, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rs := []byte{}
	br := bufio.NewReader(f)
	for {
		c, err := br.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		rs = append(rs, c)
	}
	return rs, nil
}

```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>








#### `copy`: files, directories

```go
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

```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>








#### `csv`

```go
package main

import (
	"encoding/csv"
	"os"
)

func main() {
	fpath := "test.csv"
	if err := toCSV([][]string{{"A", "B", "C"}, {"D", "E", "F"}}, fpath); err != nil {
		panic(err)
	}
	rows, err := fromCSV(fpath)
	if err != nil {
		panic(err)
	}
	if len(rows) != 2 {
		panic(err)
	}
	os.Remove(fpath)
}

func toCSV(rows [][]string, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	// func NewWriter(w io.Writer) *Writer
	wr := csv.NewWriter(f)

	if err := wr.WriteAll(rows); err != nil {
		return err
	}

	wr.Flush()

	return nil
}

func fromCSV(fpath string) ([][]string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// func NewReader(r io.Reader) *Reader
	rd := csv.NewReader(f)
	// Reading does not require `Flush`

	// in case that rows have different number of fields
	rd.FieldsPerRecord = -1

	// rd.TrailingComma = true
	// rd.TrimLeadingSpace = true
	// rd.LazyQuotes = true

	rows, err := rd.ReadAll()
	if err != nil {
		return rows, err
	}

	return rows, nil
}
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>










#### `tsv`

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fpath := "test.tsv"
	if err := toTSV([][]string{{"A", "B", "C"}, {"D", "E", "F"}}, fpath); err != nil {
		panic(err)
	}
	rows, err := fromTSV(fpath)
	if err != nil {
		panic(err)
	}
	if len(rows) != 2 {
		panic(err)
	}
	os.Remove(fpath)
}

func toTSV(rows [][]string, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	// func NewWriter(w io.Writer) *Writer
	wr := bufio.NewWriter(f)
	for _, row := range rows {
		for idx, elem := range row {

			// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
			fmt.Fprint(wr, elem)

			if len(row)-1 != idx {
				fmt.Fprint(wr, "\t")
			}
		}
		fmt.Fprint(wr, "\n")
	}
	if err := wr.Flush(); err != nil {
		return err
	}
	return nil
}

func fromTSV(fpath string) ([][]string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rows := [][]string{}

	// func NewScanner(r io.Reader) *Scanner
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rows = append(rows, strings.Split(scanner.Text(), "\t"))
	}
	if err := scanner.Err(); err != nil {
		return rows, err
	}
	return rows, f.Close()
}
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>








#### `compress/gzip`

Make sure to be effective when reading from `gzip` as explained:

- [**Crossing Streams: a Love Letter to `io.Reader`** *by Jason Moiron*](http://jmoiron.net/blog/crossing-streams-a-love-letter-to-ioreader/)



```go
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
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>











#### walk

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	func() {
		// the present working directory
		rmap, err := walk(".")
		if err != nil {
			panic(err)
		}
		for _, v := range rmap {
			fmt.Println(v)
		}
	}()

	fmt.Println()
	fmt.Println()

	func() {
		// the present working directory
		rmap, err := walkExt(".", ".go")
		if err != nil {
			panic(err)
		}
		for _, v := range rmap {
			fmt.Println(v)
		}
	}()

	fmt.Println()
	fmt.Println()

	func() {
		rmap, err := walkDir(".")
		if err != nil {
			panic(err)
		}
		for _, v := range rmap {
			fmt.Println(v)
		}
	}()
}

// walk returns all FileInfos with recursive Walk in the target directory.
// It does not include the directories but include the files inside each sub-directories.
// It does not follow the symbolic links. And excludes hidden files.
// It returns the map from os.FileInfo to its absolute path.
func walk(targetDir string) (map[os.FileInfo]string, error) {
	rmap := make(map[os.FileInfo]string)
	visit := func(path string, f os.FileInfo, err error) error {
		if !filepath.HasPrefix(path, ".") && !strings.Contains(path, "/.") {
			if _, ok := rmap[f]; !ok {
				wd, err := os.Getwd()
				if err != nil {
					return err
				}
				rmap[f] = filepath.Join(wd, path)
			}
		}
		return nil
	}
	err := filepath.Walk(targetDir, visit)
	if err != nil {
		return nil, err
	}
	return rmap, nil
}

// walkExt returns all FileInfos with specific extension.
// Make sure to prefix the extension name with dot.
// For example, to find all go files, pass ".go".
func walkExt(targetDir, ext string) (map[os.FileInfo]string, error) {
	rmap := make(map[os.FileInfo]string)
	visit := func(path string, f os.FileInfo, err error) error {
		if f != nil {
			if !f.IsDir() {
				if filepath.Ext(path) == ext {
					if !filepath.HasPrefix(path, ".") && !strings.Contains(path, "/.") {
						if _, ok := rmap[f]; !ok {
							wd, err := os.Getwd()
							if err != nil {
								return err
							}
							thepath := filepath.Join(wd, strings.Replace(path, wd, "", -1))
							rmap[f] = thepath
						}
					}
				}
			}
		}
		return nil
	}
	err := filepath.Walk(targetDir, visit)
	if err != nil {
		return nil, err
	}
	return rmap, nil
}

// walkDir returns all directories.
func walkDir(targetDir string) (map[os.FileInfo]string, error) {
	rmap := make(map[os.FileInfo]string)
	visit := func(path string, f os.FileInfo, err error) error {
		if f != nil {
			if f.IsDir() {
				if !filepath.HasPrefix(path, ".") && !strings.Contains(path, "/.") {
					if _, ok := rmap[f]; !ok {
						rmap[f] = filepath.Join(targetDir, path)
					}
				}
			}
		}
		return nil
	}
	err := filepath.Walk(targetDir, visit)
	if err != nil {
		return nil, err
	}
	return rmap, nil
}

/*
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/01_os_exec.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/06_stdout_stdin_stderr.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/07_stdout_stdin_stderr_os.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample.txt
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample.csv
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/stdout.txt
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample.json
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/12_copy.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/stderr.txt
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/00_os.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/08_exist.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/11_bufio.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample.txt
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/02_io.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/09_open_create.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/15_gzip.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample.csv
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample.json
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub/sample_copy.csv
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/05_stdin.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/13_csv.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/16_walk.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sample_copy.csv
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/03_io_pipe.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/17_flush.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/stdin.txt
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/testdata/sub
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/04_io_ioutil.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/10_ioutil_string.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/14_tsv.go


/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/12_copy.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/13_csv.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/11_bufio.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/01_os_exec.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/06_stdout_stdin_stderr.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/15_gzip.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/00_os.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/04_io_ioutil.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/05_stdin.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/07_stdout_stdin_stderr_os.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/08_exist.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/09_open_create.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/14_tsv.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/16_walk.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/03_io_pipe.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/17_flush.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/10_ioutil_string.go
/home/ubuntu/go/src/github.com/gyuho/learn/doc/go_os_io/code/02_io.go


testdata
testdata/sub
*/
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>













#### `http.Flusher`

```go
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
```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>










#### `os.Signal`

```go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	for _, sig := range []syscall.Signal{syscall.SIGINT, syscall.SIGTERM} {
		c := make(chan os.Signal, 2)
		// Notify causes package signal to relay incoming signals to c. If
		// no signals are provided, all incoming signals will be relayed to c.
		// Otherwise, just the provided signals will.
		signal.Notify(c, sig)

		handleInterrupts()
		p := syscall.Getpid()
		syscall.Kill(p, sig)

		time.Sleep(time.Second)
	}
}

// https://github.com/coreos/etcd/blob/master/pkg/osutil/interrupt_unix.go
func handleInterrupts() {
	notifier := make(chan os.Signal, 1)
	signal.Notify(notifier, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-notifier
		log.Printf("Received %v signal, shutting down...", sig)
		signal.Stop(notifier)
		pid := syscall.Getpid()
		// exit directly if it is the "init" process, since the kernel will not help to kill pid 1.
		if pid == 1 {
			os.Exit(0)
		}
		syscall.Kill(pid, sig.(syscall.Signal))
	}()
}

```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>





#### netstat

How to find and kill processes using Go? How to use `netstat` with Go?

```go
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type Process struct {
	Program string
	Socket  string

	Host string
	Port string
	PID  int
}

func main() {
	rs, err := Getpid("bin/etcd", "tcp")
	if err != nil {
		panic(err)
	}
	for _, v := range rs {
		fmt.Printf("Terminating %+v\n", v)
		syscall.Kill(v.PID, syscall.SIGINT)
	}
}

/*
Getpid parses the output of netstat command in linux:

	netstat -tlpn
	(Not all processes could be identified, non-owned process info
	 will not be shown, you would have to be root to see it all.)
	Active Internet connections (only servers)
	Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
	tcp        0      0 127.0.0.1:2379          0.0.0.0:*               LISTEN      21524/bin/etcd
	tcp        0      0 127.0.0.1:22379         0.0.0.0:*               LISTEN      21526/bin/etcd
	tcp        0      0 127.0.0.1:22380         0.0.0.0:*               LISTEN      21526/bin/etcd
	tcp        0      0 127.0.0.1:32379         0.0.0.0:*               LISTEN      21528/bin/etcd
	tcp        0      0 127.0.0.1:12379         0.0.0.0:*               LISTEN      21529/bin/etcd
	tcp        0      0 127.0.0.1:32380         0.0.0.0:*               LISTEN      21528/bin/etcd
	tcp        0      0 127.0.0.1:12380         0.0.0.0:*               LISTEN      21529/bin/etcd
	tcp        0      0 127.0.0.1:53697         0.0.0.0:*               LISTEN      2608/python2
	tcp6       0      0 :::8555                 :::*                    LISTEN      21516/goreman

*/
func Getpid(program, socket string) ([]Process, error) {
	var flag string
	flagFormat := "-%slpn"
	switch socket {
	case "tcp":
		flag = fmt.Sprintf(flagFormat, "t")
	case "tcp6":
		flag = fmt.Sprintf(flagFormat, "t")
	case "udp":
		flag = fmt.Sprintf(flagFormat, "u")
	default:
		return nil, fmt.Errorf("%s is not supported", socket)
	}
	cmd := exec.Command("netstat", flag)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReader(buf)
	lines := [][]string{}
	for {
		l, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		s := string(l)
		if !strings.HasPrefix(s, socket) {
			continue
		}
		ss := strings.Fields(s)
		lines = append(lines, ss)
	}
	addrIdx, pidIdx := 3, 6
	rs := []Process{}
	for _, v := range lines {
		addr := strings.Split(v[addrIdx], ":")
		if len(addr) < 2 {
			continue
		}
		port := ":" + addr[len(addr)-1]
		host := strings.Replace(v[addrIdx], port, "", -1)
		pidProgram := v[pidIdx]
		pp := strings.SplitN(pidProgram, "/", 2)
		if len(pp) != 2 {
			continue
		}
		processID, err := strconv.Atoi(pp[0])
		if err != nil {
			continue
		}
		programName := pp[1]
		if programName != program {
			continue
		}
		p := Process{}
		p.Program = programName
		p.Socket = socket
		p.Host = host
		p.Port = port
		p.PID = processID
		rs = append(rs, p)
	}
	return rs, nil
}

```

[↑ top](#go-os-io)
<br><br><br><br>
<hr>
