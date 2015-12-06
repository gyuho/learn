package main

import (
	"bytes"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/testdata"
)

func main() {
	func() {
		var (
			keyPath     = "/home/ubuntu/my.pem"
			user        = "ubuntu"
			host        = "YOUR_HOST"
			port        = "22"
			dialTimeout = 5 * time.Second
			execTimeout = 5 * time.Second
			cmd1        = "cd $HOME; pwd; ls;"
			cmd2        = "cd $HOME; pwd; ls;"
		)
		f, err := openToRead(keyPath)
		if err != nil {
			panic(err)
		}
		sshSigner, err := getSSHSigner(f)
		if err != nil {
			panic(err)
		}
		fmt.Printf("sshSigner: %+v\n", sshSigner)
		//
		sshClient, err := getSSHClient(sshSigner, user, host, port, dialTimeout)
		if err != nil {
			panic(err)
		}
		fmt.Printf("sshClient: %+v\n", sshClient)
		//
		output1, err := runCommand(sshClient, cmd1, execTimeout)
		if err != nil {
			panic(err)
		}
		fmt.Printf("output1:\n%+v\n", output1)
		//
		output2, err := runCommand(sshClient, cmd2, execTimeout)
		if err != nil {
			panic(err)
		}
		fmt.Printf("output2:\n%+v\n", output2)
	}()

	runTestServer()
}

func openToRead(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return f, err
	}
	return f, nil
}

func getSSHSigner(rd io.Reader) (ssh.Signer, error) {
	// ioutil.ReadAll can take `os.File` as a `io.Reader` or `io.Writer`
	// Make sure to get the fresh reader for every GetSSHSigner call
	bts, err := ioutil.ReadAll(rd)
	if err != nil {
		return nil, err
	}
	// parse the private key to check if the private key has a password.
	block, _ := pem.Decode(bts)
	if block == nil {
		return nil, fmt.Errorf("no PEM data is found")
	}
	if block.Headers["Proc-Type"] == "4,ENCRYPTED" {
		return nil, fmt.Errorf("Password protected key is not supported. Please decrypt the key prior to use.")
	}
	sg, err := ssh.ParsePrivateKey(bts)
	if err != nil {
		return nil, err
	}
	if t, ok := rd.(*os.File); ok {
		t.Close()
	}
	return sg, nil
}

func getSSHClient(
	sshSigner ssh.Signer,
	user string,
	host string,
	port string,
	dialTimeout time.Duration,
) (*ssh.Client, error) {

	clientConfig := ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(sshSigner)},
		// ssh.Password("password"),
	}
	addr := host + ":" + port

	// if we need to set up dial timeout.
	//
	c, err := net.DialTimeout("tcp", addr, dialTimeout)
	if err != nil {
		return nil, err
	}
	if tc, ok := c.(*net.TCPConn); ok {
		// if c is tcp connection, set these:
		tc.SetKeepAlive(true)
		tc.SetKeepAlivePeriod(5 * time.Second)
	}
	// func NewClientConn(c net.Conn, addr string, config *ClientConfig)
	// (Conn, <-chan NewChannel, <-chan *Request, error)
	conn, newChan, reqChan, err := ssh.NewClientConn(
		c,
		addr,
		&clientConfig,
	)
	if err != nil {
		return nil, err
	}
	if conn == nil {
		return nil, errors.New("Can't establish SSH")
	}
	// func NewClient(c Conn, chans <-chan NewChannel, reqs <-chan *Request) *Client
	return ssh.NewClient(conn, newChan, reqChan), nil

	// or
	//
	// return ssh.Dial("tcp", addr, &clientConfig)
}

func runCommand(
	sshClient *ssh.Client,
	cmd string,
	execTimeout time.Duration,
) (string, error) {

	session, err := sshClient.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	termModes := ssh.TerminalModes{
		ssh.ECHO:          0,     // do not echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err := session.RequestPty("xterm", 80, 40, termModes); err != nil {
		return "", err
	}

	stdinBuf, stdoutBuf, stderrBuf := new(bytes.Buffer), new(bytes.Buffer), new(bytes.Buffer)
	session.Stdin = stdinBuf
	session.Stdout = stdoutBuf
	session.Stderr = stderrBuf

	outputChan := make(chan string)
	errChan := make(chan error)
	go func() {
		if err := session.Run(cmd); err != nil {
			errChan <- fmt.Errorf("%v, %s", err, stderrBuf.String())
			return
		}
		outputChan <- stdoutBuf.String()
	}()
	select {
	case output := <-outputChan:
		return output, nil

	case err := <-errChan:
		return "", err

	case <-time.After(execTimeout):
		return "", fmt.Errorf("execution timeout.")
	}
}

func runTestServer() {
	// Parse and set the private key of the server, required to accept connections
	//
	// testKeyByte is RSA sample private key from
	// https://github.com/golang/crypto/blob/master/ssh/testdata/keys.go
	sshSigner, err := getSSHSigner(bytes.NewReader(testdata.PEMBytes["rsa"]))
	// sshSigner, err := ssh.ParsePrivateKey(testdata.PEMBytes["rsa"])
	if err != nil {
		panic(err)
	}
	serverConfig := &ssh.ServerConfig{
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			log.Println(conn.RemoteAddr(), "is authenticated with", key.Type())
			return nil, nil
		},
	}
	serverConfig.AddHostKey(sshSigner)
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	go func() {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		defer c.Close()
		conn, chans, _, err := ssh.NewServerConn(c, serverConfig)
		if err != nil {
			fmt.Printf("Handshaking error: %v", err)
		}
		fmt.Println("Accepted SSH connection")
		for newChannel := range chans {
			channel, _, err := newChannel.Accept()
			if err != nil {
				panic("Unable to accept channel.")
			}
			fmt.Println("Accepted channel")
			go func() {
				defer channel.Close()
				conn.OpenChannel(newChannel.ChannelType(), nil)
			}()
		}
		conn.Close()
	}()
	addr := l.Addr().String()
	fmt.Println("Returning address:", addr)
	clientConfig := ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{ssh.PublicKeys(sshSigner)},
	}
	sshClient, err := ssh.Dial("tcp", addr, &clientConfig)
	if err != nil {
		panic(err)
	}
	session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	fmt.Printf("session: %+v\n", session)
}
