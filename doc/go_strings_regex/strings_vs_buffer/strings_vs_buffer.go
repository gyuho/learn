package strings_vs_buffer

import (
	"bytes"
	"fmt"
	"strings"
)

var data = []string{}

func init() {
	for i := 0; i < 3000; i++ {
		data = append(data, fmt.Sprintf("%d", i))
	}
}

func stringsJoin(d []string) string {
	slice := make([]string, 0, len(d))
	for i := range data {
		slice = append(slice, data[i])
	}
	return strings.Join(slice, ",")
}

func bufferJoin(d []string) string {
	buf := &bytes.Buffer{}
	for i := range data {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(data[i])
	}
	return buf.String()
}

/*
sh bench.sh
testing: warning: no tests to run
PASS
BenchmarkStringsJoin  	   30000	     59646 ns/op	   77312 B/op	       3 allocs/op
BenchmarkStringsJoin-2	   20000	     63908 ns/op	   77312 B/op	       3 allocs/op
BenchmarkStringsJoin-4	   20000	     66558 ns/op	   77313 B/op	       3 allocs/op
BenchmarkStringsJoin-8	   20000	    103958 ns/op	   77313 B/op	       3 allocs/op
BenchmarkBufferJoin   	   30000	     59510 ns/op	   77312 B/op	       3 allocs/op
BenchmarkBufferJoin-2 	   20000	     62430 ns/op	   77312 B/op	       3 allocs/op
BenchmarkBufferJoin-4 	   20000	     66373 ns/op	   77313 B/op	       3 allocs/op
BenchmarkBufferJoin-8 	   20000	     77346 ns/op	   77313 B/op	       3 allocs/op
ok  	github.com/gyuho/learn/doc/go_strings_regex/strings_vs_buffer	17.823s
*/
