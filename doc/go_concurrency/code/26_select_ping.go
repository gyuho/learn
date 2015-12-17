package main

import (
	"fmt"
	"net/http"
	"time"
)

var sitesToPing = []string{
	"http://www.google.com",
	"http://www.amazon.com",
	"http://nowebsite.net",
}

func main() {
	respChan, errChan := make(chan string), make(chan error)
	for _, target := range sitesToPing {
		go head(target, respChan, errChan)
	}
	for i := 0; i < len(sitesToPing); i++ {
		select {
		case res := <-respChan:
			fmt.Println(res)
		case err := <-errChan:
			fmt.Println(err)
		case <-time.After(time.Second):
			fmt.Println("Timeout!")
		}
	}
	close(respChan)
	close(errChan)
}

/*
200 / http://www.google.com:OK
405 / http://www.amazon.com:Method Not Allowed
Timeout!
*/

func head(
	target string,
	respChan chan string,
	errChan chan error,
) {
	req, err := http.NewRequest("HEAD", target, nil)
	if err != nil {
		errChan <- fmt.Errorf("0 / %s:None with %v", target, err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errChan <- fmt.Errorf("0 / %s:None with %v", target, err)
		return
	}
	defer resp.Body.Close()
	stCode := resp.StatusCode
	stText := http.StatusText(resp.StatusCode)
	respChan <- fmt.Sprintf("%d / %s:%s", stCode, target, stText)
	return
}
