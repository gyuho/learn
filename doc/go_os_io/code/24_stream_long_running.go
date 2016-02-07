package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/coreos/etcd/Godeps/_workspace/src/github.com/spf13/cobra"
)

type Flag struct {
	EtcdOld string
	EtcdNew string
}

var (
	rootCommand = &cobra.Command{
		Use:        "migration",
		Short:      "migration handles etcd migration.",
		SuggestFor: []string{"migration", "miation", "miration"},
	}
)

func init() {
	cobra.EnablePrefixMatching = true
}

func init() {
	rootCommand.AddCommand(releaseCommand)
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

var (
	releaseCommand = &cobra.Command{
		Use:   "release",
		Short: "release checks etcd migration between two releases.",
		Run:   CommandFunc,
	}

	cmdFlag = Flag{}
)

func init() {
	cobra.EnablePrefixMatching = true
}

func init() {
	releaseCommand.PersistentFlags().StringVarP(&cmdFlag.EtcdOld, "etcd-binary-old", "a", "bin/etcd", "Path of executable etcd binary to migrate from.")
	releaseCommand.PersistentFlags().StringVarP(&cmdFlag.EtcdNew, "etcd-binary-new", "b", "bin/etcd", "Path of executable etcd binary to migrate to.")
}

var (
	defaultFlags1 = []string{
		"--name", "infra1",
		"--listen-client-urls", "http://localhost:12379",
		"--advertise-client-urls", "http://localhost:12379",
		"--listen-peer-urls", "http://localhost:12380",
		"--initial-advertise-peer-urls", "http://localhost:12380",
		"--initial-cluster-token", "etcd-cluster-1",
		"--initial-cluster", "infra1=http://localhost:12380,infra2=http://localhost:22380,infra3=http://localhost:32380",
		"--initial-cluster-state", "new",
	}
	defaultFlags2 = []string{
		"--name", "infra2",
		"--listen-client-urls", "http://localhost:22379",
		"--advertise-client-urls", "http://localhost:22379",
		"--listen-peer-urls", "http://localhost:22380",
		"--initial-advertise-peer-urls", "http://localhost:22380",
		"--initial-cluster-token", "etcd-cluster-1",
		"--initial-cluster", "infra1=http://localhost:12380,infra2=http://localhost:22380,infra3=http://localhost:32380",
		"--initial-cluster-state", "new",
	}
	defaultFlags3 = []string{
		"--name", "infra3",
		"--listen-client-urls", "http://localhost:32379",
		"--advertise-client-urls", "http://localhost:32379",
		"--listen-peer-urls", "http://localhost:32380",
		"--initial-advertise-peer-urls", "http://localhost:32380",
		"--initial-cluster-token", "etcd-cluster-1",
		"--initial-cluster", "infra1=http://localhost:12380,infra2=http://localhost:22380,infra3=http://localhost:32380",
		"--initial-cluster-state", "new",
	}
	memberStartReadyString   = "etcdserver: set the initial cluster version to "
	memberReStartReadySuffix = " became active"
)

func getInfraFlags(i int) []string {
	switch i {
	case 1:
		return defaultFlags1
	case 2:
		return defaultFlags2
	case 3:
		return defaultFlags3
	default:
		panic(fmt.Sprintf("%d is not defined", i))
	}
}

func CommandFunc(cmd *cobra.Command, args []string) {
	defer func() {
		fmt.Println("deleting...")
		os.RemoveAll("infra1.etcd")
		os.RemoveAll("infra2.etcd")
		os.RemoveAll("infra3.etcd")
	}()

	oldCmds := make([]*exec.Cmd, 3)
	oldOutputs := make([]io.ReadCloser, 3)
	newCmds := make([]*exec.Cmd, 3)
	newOutputs := make([]io.ReadCloser, 3)
	for i := range oldCmds {
		oldCmd := exec.Command(cmdFlag.EtcdOld, getInfraFlags(i+1)...)
		oldCmds[i] = oldCmd
		oldOutput, err := oldCmd.StderrPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		oldOutputs[i] = oldOutput

		newCmd := exec.Command(cmdFlag.EtcdNew, getInfraFlags(i+1)...)
		newCmds[i] = newCmd
		newOutput, err := newCmd.StderrPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		newOutputs[i] = newOutput
	}

	errChan := make(chan error)
	done := make(chan struct{})
	for i := range oldCmds {
		cmd := oldCmds[i]
		go func(i int, cmd *exec.Cmd) {
			if err := cmd.Start(); err != nil {
				errChan <- err
				return
			}
			done <- struct{}{}
		}(i, cmd)
	}
	cn := 0
	for cn != 3 {
		cn++
		select {
		case err := <-errChan:
			fmt.Fprintln(os.Stderr, err)
			return
		case <-done:
		}
	}

	becameActiveCount := 0
	for i, o := range oldOutputs {
		go func(i int, reader io.ReadCloser) {
			scanner := bufio.NewScanner(reader)
			for {
				for scanner.Scan() {
					txt := scanner.Text()
					fmt.Printf("[old infra%d] %s\n", i+1, txt)
					if strings.Contains(txt, memberStartReadyString) {
						done <- struct{}{}
					}
					if strings.HasSuffix(txt, memberReStartReadySuffix) {
						fmt.Printf("[old infra%d] reconnected!\n", i+1)
						becameActiveCount++
					}
				}
			}
			if err := scanner.Err(); err != nil {
				errChan <- err
				return
			}
		}(i, o)
	}
	for i, o := range newOutputs {
		go func(i int, reader io.ReadCloser) {
			scanner := bufio.NewScanner(reader)
			for {
				for scanner.Scan() {
					txt := scanner.Text()
					fmt.Printf("[new infra%d] %s\n", i+1, txt)
					if strings.HasSuffix(txt, memberReStartReadySuffix) {
						fmt.Printf("[new infra%d] reconnected!\n", i+1)
						becameActiveCount++
					}
				}
			}
			if err := scanner.Err(); err != nil {
				errChan <- err
				return
			}
		}(i, o)
	}
	cn = 0
	for cn != 3 {
		cn++
		select {
		case err := <-errChan:
			fmt.Fprintln(os.Stderr, err)
			return
		case <-done:
		}
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("[old infra%d] killing...\n", i+1)
		if err := syscall.Kill(oldCmds[i].Process.Pid, syscall.SIGKILL); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Printf("[old infra%d] killed!\n", i+1)
		time.Sleep(10 * time.Second)

		fmt.Printf("[new infra%d] restarting...\n", i+1)
		if err := newCmds[i].Start(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Printf("[new infra%d] restarted!\n", i+1)
		time.Sleep(10 * time.Second)
	}

	// 6(2 per node) at the beginning of cluster + 12(4 per kill) during migration = 18
	if becameActiveCount == 18 {
		fmt.Printf("migration successful from %s to %s\n", cmdFlag.EtcdOld, cmdFlag.EtcdNew)
	} else {
		fmt.Printf("migration failed from %s to %s (becameActiveCount %d)\n", cmdFlag.EtcdOld, cmdFlag.EtcdNew, becameActiveCount)
	}
}
