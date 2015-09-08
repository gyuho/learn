[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: argument, flag

- [flag](#flag)

[↑ top](#go-argument-flag)
<br><br><br><br>
<hr>








#### flag

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
go run 00_flag.go -h

Usage of /tmp/go-build105642507/command-line-arguments/_obj/exe/00_flag:
  -description string
    	Describe the argument. (default "None")
  -index int
    	Specify the index.
exit status 2



go run 00_flag.go -index=10 \
-description="Hello World!"
;

index: 10
description: Hello World!
*/
```

[↑ top](#go-argument-flag)
<br><br><br><br>
<hr>
