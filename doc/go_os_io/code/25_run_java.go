package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	cmd := exec.Command("bash", "-c", "java HelloWorld")
	err := cmd.Start()
	fmt.Printf("PID: %d\n", cmd.Process.Pid)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Printf("Got signal to %d\n", cmd.Process.Pid)
		syscall.Kill(cmd.Process.Pid, syscall.SIGHUP)
		done <- true
	}()
	<-done
}
