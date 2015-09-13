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
