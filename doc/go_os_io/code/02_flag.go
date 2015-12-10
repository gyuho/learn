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
You can:

-description aaa
-description 'aaa'
-description "aaa"
-description=aaa
-description='aaa'
-description="aaa"

--description aaa
--description 'aaa'
--description "aaa"
--description=aaa
--description='aaa'
--description="aaa"
*/
