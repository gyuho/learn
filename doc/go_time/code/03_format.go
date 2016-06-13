package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	{
		st := time.Now()
		for i := 0; i < 10000; i++ {
			buf := bytes.NewBuffer(nil)
			now := time.Now()
			ts := now.Format("2006-01-02 15:04:05")
			buf.WriteString(ts)
			ms := now.Nanosecond() / 1000
			buf.WriteString(fmt.Sprintf(".%06d", ms))
			_ = buf.String()
		}
		fmt.Println("took", time.Since(st))
	}

	{
		st := time.Now()
		for i := 0; i < 10000; i++ {
			_ = time.Now().String()[:26]
		}
		fmt.Println("took", time.Since(st))
	}
}

/*
took 8.934964ms
took 5.888227ms
*/
