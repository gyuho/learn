package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	var (
		writer   = os.Stdout
		doneWait = make(chan struct{})
		errChan  = make(chan error)
		waitSig  = make(chan bool)
		// cmd      = exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin/etcd"))
		cmd = exec.Command("echo", "hello")
	)

	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer cmdOut.Close()
	cmdErr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer cmdErr.Close()

	go func() {
		fmt.Println("cmd.Start:", cmd.Path, cmd.Args)
		if err := cmd.Start(); err != nil {
			errChan <- err
			close(waitSig)
			return
		}
		waitSig <- true
	}()

	go func() {
		scanner := bufio.NewScanner(cmdOut)
		for scanner.Scan() {
			fmt.Fprintln(writer, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			errChan <- err
		}
	}()

	go func() {
		scanner := bufio.NewScanner(cmdErr)
		for scanner.Scan() {
			fmt.Fprintln(writer, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			errChan <- err
		}
	}()

	go func() {
		ready, ok := <-waitSig
		if !ready && !ok {
			log.Fatal("something wrong with cmd.Start!")
		}

		fmt.Println("cmd.Wait")
		if err := cmd.Wait(); err != nil {
			errChan <- err
			return
		}
		doneWait <- struct{}{}
	}()

	select {
	case <-doneWait:
		fmt.Println("cmd done!")

	case err := <-errChan:
		fmt.Println("error:", err)

	case <-time.After(10 * time.Second):
		fmt.Println("timed out and cmd.Process.Kill")
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("error when cmd.Process.Kill:", err)
		}
	}
}

/*
cmd.Start: /bin/echo [echo hello]
cmd.Wait
hello
cmd done!
*/
