package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	var (
		w = os.Stdout

		sudo = true

		// socket = "tcp"
		socket = "tcp6"

		program = ""
		// program = "bin/etcd"

		// port = ""
		port = ":8080"
	)

	ps, err := NetStat(w, sudo, socket, program, port)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	Kill(w, ps...)
	/*
	   [NetStat] socket: 'tcp6' / program: '00_hello_world' / host: '::' / port: ':8080' / pid: '9768'
	   [Kill] syscall.Kill -> socket: 'tcp6' / program: '00_hello_world' / host: '::' / port: ':8080' / pid: '9768'
	   [Kill] Done!
	*/
}

// Process describes OS processes.
type Process struct {
	Socket  string
	Program string
	Host    string
	Port    string
	PID     int
}

func (p Process) String() string {
	return fmt.Sprintf("socket: '%s' / program: '%s' / host: '%s' / port: '%s' / pid: '%d'",
		p.Socket,
		p.Program,
		p.Host,
		p.Port,
		p.PID,
	)
}

// Kill kills all processes in arguments.
func Kill(w io.Writer, ps ...Process) {
	defer func() {
		recover()
	}()
	for _, v := range ps {
		fmt.Fprintf(w, "[Kill] syscall.Kill -> %s\n", v)
		if err := syscall.Kill(v.PID, syscall.SIGINT); err != nil {
			fmt.Fprintln(w, "[Kill - error]", err)
		}
	}
	fmt.Fprintln(w, "[Kill] Done!")
}

/*
NetStat parses the output of netstat command in linux.
Pass '' or '*' to match all. For example, call Kill("tcp", "bin/etcd", "*")
to kill all processes that are running "bin/etcd":

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

Otherwise, you would have to run something like the following:

	sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:12379/gio');
	sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:22379/gio');
	sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:32379/gio');
	sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:2379/gio');
	sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:8080/gio');

*/
func NetStat(w io.Writer, sudo bool, socket, program, port string) ([]Process, error) {
	socket = strings.TrimSpace(socket)
	program = strings.TrimSpace(program)
	if program == "" {
		program = "*"
	}
	port = strings.TrimSpace(port)
	if port == "" {
		port = "*"
	}
	if port != "*" {
		if !strings.HasPrefix(port, ":") {
			port = ":" + port
		}
	}
	if program == "*" && port == "*" {
		fmt.Fprintln(w, "[NetStat - warning] grepping all programs.")
	}

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
		return nil, fmt.Errorf("socket '%s' is unknown", socket)
	}

	cmd := exec.Command("netstat", flag)
	if sudo {
		cmd = exec.Command("sudo", "netstat", flag)
	}
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf
	if err := cmd.Run(); err != nil {
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
		sl := strings.Fields(s)
		lines = append(lines, sl)
	}

	socketIdx := 0
	portIdx := 3
	programIdx := 6

	ps := []Process{}

	fmt.Fprintf(w, "[netstat] 'netstat %s' returned %d lines.\n", flag, len(lines))
	for _, sl := range lines {

		theSocket := sl[socketIdx]
		if theSocket != socket {
			fmt.Fprintln(w, "[NetStat] different socket. Skipping", sl)
			continue
		}

		asl := strings.Split(sl[portIdx], ":")
		if len(asl) < 2 {
			fmt.Fprintln(w, "[NetStat] skipping", sl)
			continue
		}
		thePort := ":" + asl[len(asl)-1]
		if port != "*" {
			if thePort != port {
				fmt.Fprintln(w, "[NetStat] different port. Skipping", sl)
				continue
			}
		}

		theHost := strings.TrimSpace(strings.Replace(sl[portIdx], thePort, "", -1))

		psl := strings.SplitN(sl[programIdx], "/", 2)
		if len(psl) != 2 {
			continue
		}
		theProgram := strings.TrimSpace(psl[1])
		if program != "*" {
			if theProgram != program {
				fmt.Fprintln(w, "[NetStat] different program. Skipping", sl)
				continue
			}
		}

		thePID := 0
		if d, err := strconv.Atoi(psl[0]); err != nil {
			fmt.Fprintln(w, "[NetStat - error] %v / Skipping %+v\n", err, sl)
			continue
		} else {
			thePID = d
		}

		p := Process{}
		p.Socket = theSocket
		p.Program = theProgram
		p.Host = theHost
		p.Port = thePort
		p.PID = thePID
		ps = append(ps, p)
	}

	for _, v := range ps {
		fmt.Fprintf(w, "[NetStat] %s\n", v)
	}
	return ps, nil
}
