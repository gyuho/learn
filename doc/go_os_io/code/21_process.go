package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type Process struct {
	Program string
	Socket  string

	Host string
	Port string
	PID  int
}

func main() {
	rs, err := Getpid("bin/etcd", "tcp")
	if err != nil {
		panic(err)
	}
	for _, v := range rs {
		fmt.Printf("Terminating %+v\n", v)
		syscall.Kill(v.PID, syscall.SIGINT)
	}
}

/*
Getpid parses the output of netstat command in linux:

	netstat -tlpn
	(Not all processes could be identified, non-owned process info
	 will not be shown, you would have to be root to see it all.)
	Active Internet connections (only servers)
	Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
	tcp        0      0 127.0.0.1:2379          0.0.0.0:*               LISTEN      21524/bin/etcd
	tcp        0      0 127.0.0.1:22379         0.0.0.0:*               LISTEN      21526/bin/etcd
	tcp        0      0 127.0.0.1:22380         0.0.0.0:*               LISTEN      21526/bin/etcd
	tcp        0      0 127.0.0.1:32379         0.0.0.0:*               LISTEN      21528/bin/etcd
	tcp        0      0 127.0.0.1:12379         0.0.0.0:*               LISTEN      21529/bin/etcd
	tcp        0      0 127.0.0.1:32380         0.0.0.0:*               LISTEN      21528/bin/etcd
	tcp        0      0 127.0.0.1:12380         0.0.0.0:*               LISTEN      21529/bin/etcd
	tcp        0      0 127.0.0.1:53697         0.0.0.0:*               LISTEN      2608/python2
	tcp6       0      0 :::8555                 :::*                    LISTEN      21516/goreman

*/
func Getpid(program, socket string) ([]Process, error) {
	var flag string
	flagFormat := "-%slpn"
	switch socket {
	case "tcp":
		flag = fmt.Sprintf(flagFormat, "t")
	case "tcp6":
		flag = fmt.Sprintf(flagFormat, "t")
	case "udp":
		flag = fmt.Sprintf(flagFormat, "u")
	default:
		return nil, fmt.Errorf("%s is not supported", socket)
	}
	cmd := exec.Command("netstat", flag)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReader(buf)
	lines := [][]string{}
	for {
		l, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		s := string(l)
		if !strings.HasPrefix(s, socket) {
			continue
		}
		ss := strings.Fields(s)
		lines = append(lines, ss)
	}
	addrIdx, pidIdx := 3, 6
	rs := []Process{}
	for _, v := range lines {
		addr := strings.Split(v[addrIdx], ":")
		if len(addr) < 2 {
			continue
		}
		port := ":" + addr[len(addr)-1]
		host := strings.Replace(v[addrIdx], port, "", -1)
		pidProgram := v[pidIdx]
		pp := strings.SplitN(pidProgram, "/", 2)
		if len(pp) != 2 {
			continue
		}
		processID, err := strconv.Atoi(pp[0])
		if err != nil {
			continue
		}
		programName := pp[1]
		if programName != program {
			continue
		}
		p := Process{}
		p.Program = programName
		p.Socket = socket
		p.Host = host
		p.Port = port
		p.PID = processID
		rs = append(rs, p)
	}
	return rs, nil
}
