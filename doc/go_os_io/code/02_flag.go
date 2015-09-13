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
