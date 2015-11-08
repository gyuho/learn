package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	for _, sig := range []syscall.Signal{syscall.SIGINT, syscall.SIGTERM} {
		c := make(chan os.Signal, 2)
		// Notify causes package signal to relay incoming signals to c. If
		// no signals are provided, all incoming signals will be relayed to c.
		// Otherwise, just the provided signals will.
		signal.Notify(c, sig)

		handleInterrupts()
		p := syscall.Getpid()
		syscall.Kill(p, sig)

		time.Sleep(time.Second)
	}
}

// https://github.com/coreos/etcd/blob/master/pkg/osutil/interrupt_unix.go
func handleInterrupts() {
	notifier := make(chan os.Signal, 1)
	signal.Notify(notifier, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-notifier
		log.Printf("Received %v signal, shutting down...", sig)
		signal.Stop(notifier)
		pid := syscall.Getpid()
		// exit directly if it is the "init" process, since the kernel will not help to kill pid 1.
		if pid == 1 {
			os.Exit(0)
		}
		syscall.Kill(pid, sig.(syscall.Signal))
	}()
}
