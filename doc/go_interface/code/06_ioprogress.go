package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mitchellh/ioprogress"
)

func main() {
	downloadTarget := "http://releases.ubuntu.com/14.04.2/ubuntu-14.04.2-desktop-amd64.iso"
	resp, err := http.Get(downloadTarget)
	if err != nil {
		log.Fatal(err)
	}

	size := resp.ContentLength
	progressR := &ioprogress.Reader{
		Reader:       resp.Body,
		Size:         size,
		DrawInterval: time.Millisecond,
		DrawFunc: func(progress, total int64) error {
			if progress == total {
				fmt.Printf("\rDownloading: %s%10s", ioprogress.DrawTextFormatBytes(size, size), "")
				return nil
			}
			fmt.Printf("\rDownloading: %s%10s", ioprogress.DrawTextFormatBytes(progress, total), "")
			return nil
		},
	}

	file, err := open("ubuntu.iso")
	if err != nil {
		log.Fatal(err)
	}
	if _, err = io.Copy(file, progressR); err != nil {
		log.Fatal(err)
	}
}

func open(fpath string) (*os.File, error) {
	file, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		file, err = os.Create(fpath)
		if err != nil {
			return file, err
		}
	}
	return file, nil
}
