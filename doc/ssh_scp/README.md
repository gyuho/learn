[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# ssh, scp

- [Reference](#reference)
- [`ssh`, `scp`](#ssh-scp)
- [implement `ssh`](#implement-ssh)
- [`scp` from a local file, directory](#scp-from-a-local-file-directory)
- [`scp` from a remote file](#scp-from-a-remote-file)

[↑ top](#ssh-scp)
<br><br><br><br>
<hr>






#### Reference

- [Secure Shell](https://en.wikipedia.org/wiki/Secure_Shell)
- [Secure copy](https://en.wikipedia.org/wiki/Secure_copy)
- [How the SCP protocol works](https://blogs.oracle.com/janp/entry/how_the_scp_protocol_works)
- [Why aren’t we using SSH for everything?](https://medium.com/@shazow/ssh-how-does-it-even-9e43586e4ffc)

[↑ top](#ssh-scp)
<br><br><br><br>
<hr>






#### `ssh`, `scp`

Application layer refers to shared protocols and interface
methods between hosts, such as HTTP, [SSH](https://en.wikipedia.org/wiki/Secure_Shell)
, SMTP.

> Secure Shell, or SSH, is a cryptographic (encrypted) network protocol for
> initiating text-based shell sessions on remote machines
> in a secure way.
>
> This allows a user to run commands on a machine's command prompt without them
> being physically present near the machine. It also allows a user to establish
> a secure channel over an insecure network in a client-server architecture,
> connecting an SSH client application with an SSH server. Common
> applications include remote command-line login and remote command execution,
> but any network service can be secured with SSH. The protocol specification
> distinguishes between two major versions, referred to as SSH-1 and SSH-2.
>
> [*Secure Shell*](https://en.wikipedia.org/wiki/Secure_Shell) *by Wikipedia*

<br>
> Secure copy or SCP is a means of securely transferring computer files between
> a local host and a remote host or between two remote hosts. It is based on
> the Secure Shell (SSH) protocol.
>
> The SCP protocol is a network protocol, based on the BSD RCP protocol,
> which supports file transfers between hosts on a network. SCP uses Secure
> Shell (SSH) for data transfer and uses the same mechanisms for
> authentication, thereby ensuring the authenticity and confidentiality of the
> data in transit. A client can send (upload) files to a server, optionally
> including their basic attributes (permissions, timestamps). Clients can also
> request files or directories from a server (download). SCP runs over TCP port
> 22 by default. Like RCP, there is no RFC that defines the specifics of the
> protocol.
> 
> [Secure copy](https://en.wikipedia.org/wiki/Secure_copy) *by Wikipedia*

<br>
Here's how you use `ssh` and `scp` in Ubuntu:

```
# ssh into a remote machine:
ssh -i KEY_PATH \
USER@HOST
;

# upload from local to remote:
scp -i KEY_PATH \
-r SOURCE_PATH_IN_LOCAL \
USER@HOST:DESTINATION_PATH_IN_REMOTE
;

# download from remote to local:
scp -i KEY_PATH \
USER@HOST:SOURCE_PATH_IN_REMOTE \
DESTINATION_PATH_IN_LOCAL
;
```

[↑ top](#ssh-scp)
<br><br><br><br>
<hr>








#### implement `ssh`

[golang.org/x/crypto/ssh](https://godoc.org/golang.org/x/crypto/ssh)
implements `SSH` client and server.
We use `SSH` to run commands in a remote machine.
And here's how you would do it in Go:

```go
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

```

[↑ top](#ssh-scp)
<br><br><br><br>
<hr>











#### `scp` from a local file, directory

```go
package main

import (
	"bufio"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	func() {
		var (
			keyPath     = "/home/ubuntu/my.pem"
			user        = "ubuntu"
			host        = "YOUR_HOST"
			port        = "22"
			dialTimeout = 5 * time.Second
			execTimeout = 15 * time.Second

			fromPath = "testdata/hello.txt"
			toPath   = "/home/ubuntu/hello_copy.txt"
		)
		if err := scpToRemoteFile(keyPath, user, host, port, dialTimeout, fromPath, toPath, execTimeout); err != nil {
			panic(err)
		}
	}()

	func() {
		var (
			keyPath     = "/home/ubuntu/my.pem"
			user        = "ubuntu"
			host        = "YOUR_HOST"
			port        = "22"
			dialTimeout = 5 * time.Second
			execTimeout = 15 * time.Second

			fromPath = "testdata/world.txt"
			toPath   = "/home/ubuntu/world_copy.txt"
		)
		if err := scpToRemoteFile(keyPath, user, host, port, dialTimeout, fromPath, toPath, execTimeout); err != nil {
			panic(err)
		}
	}()

	func() {
		var (
			keyPath     = "/home/ubuntu/my.pem"
			user        = "ubuntu"
			host        = "YOUR_HOST"
			port        = "22"
			dialTimeout = 5 * time.Second
			execTimeout = 15 * time.Second

			fromDirPath = "testdata"
			toDirPath   = "/home/ubuntu/testdata_copy"
		)
		if err := scpToRemoteDir(keyPath, user, host, port, dialTimeout, fromDirPath, toDirPath, execTimeout); err != nil {
			panic(err)
		}
	}()
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
	return ssh.Dial("tcp", addr, &clientConfig)
}

func check(stdoutBufioReader *bufio.Reader) error {
	code, err := stdoutBufioReader.ReadByte()
	if err != nil {
		return err
	}
	//  0 (OK), 1 (warning) or 2 (fatal error; will end the connection)
	if code != 0 {
		msg, _, err := stdoutBufioReader.ReadLine()
		if err != nil {
			return fmt.Errorf("stdoutBufioReader.ReadLine error: %+v", err)
		}
		return fmt.Errorf("stdoutBufioReader.ReadByte error: %+v / %s", err, string(msg))
	}
	return nil
}

func uploadFile(
	fromPath string,
	toPath string,
	stdinPipe io.Writer,
	stdoutBufioReader *bufio.Reader,
) error {
	/////////////////////////////
	// copy to a temporary file.
	r, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer r.Close()
	w, err := ioutil.TempFile("", "temp_prefix_")
	if err != nil {
		return err
	}
	defer w.Close()
	if _, err = io.Copy(w, r); err != nil {
		return err
	}
	if err := w.Sync(); err != nil {
		return err
	}
	if _, err := w.Seek(0, 0); err != nil {
		return err
	}
	wi, err := w.Stat()
	if err != nil {
		return err
	}
	fsize := wi.Size()
	/////////////////////////////

	/////////////////////////////
	// start scp protocol.
	toPathFileName := filepath.Base(toPath)
	fmt.Fprintln(stdinPipe, "C0644", fsize, toPathFileName)
	if err := check(stdoutBufioReader); err != nil {
		return err
	}
	/////////////////////////////

	/////////////////////////////
	// start writing.
	if _, err := io.Copy(stdinPipe, w); err != nil {
		return err
	}
	fmt.Fprint(stdinPipe, "\x00")
	if err := check(stdoutBufioReader); err != nil {
		return err
	}
	/////////////////////////////

	return nil
}

func scpToRemoteFile(
	keyPath string,
	user string,
	host string,
	port string,
	dialTimeout time.Duration,
	fromPath string,
	toPath string,
	execTimeout time.Duration,
) error {
	/////////////////////////////
	f, err := openToRead(keyPath)
	if err != nil {
		return err
	}
	sshSigner, err := getSSHSigner(f)
	if err != nil {
		return err
	}
	sshClient, err := getSSHClient(sshSigner, user, host, port, dialTimeout)
	if err != nil {
		return err
	}
	session, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	stdoutPipe, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	stdoutBufioReader := bufio.NewReader(stdoutPipe)
	stdinPipe, err := session.StdinPipe()
	if err != nil {
		return err
	}
	if stdinPipe == nil {
		return fmt.Errorf("stdinPipe is nil")
	}
	// defer stdinPipe.Close()
	// make sure to close this before session.Wait()
	//
	// https://blogs.oracle.com/janp/entry/how_the_scp_protocol_works
	// In all cases aside from remote-to-remote scenario the scp command
	// processes command line options and then starts an SSH connection
	// to the remote host. Another scp command is run on the remote side
	// through that connection in either source or sink mode. Source mode
	// reads files and sends them over to the other side, sink mode accepts them.
	// Source and sink modes are triggered using -f (from) and -t (to) options, respectively.
	if err := session.Start(fmt.Sprintf("scp -vt %s", toPath)); err != nil {
		return err
	}
	/////////////////////////////

	/////////////////////////////
	if err := uploadFile(fromPath, toPath, stdinPipe, stdoutBufioReader); err != nil {
		return err
	}
	// make sure to close this before session.Wait()
	stdinPipe.Close()
	/////////////////////////////

	/////////////////////////////
	// wait for session to finish.
	doneChan := make(chan struct{})
	errChan := make(chan error)
	go func() {
		if err := session.Wait(); err != nil {
			fmt.Println("wait returns err", err)
			if exitErr, ok := err.(*ssh.ExitError); ok {
				fmt.Printf("non-zero exit status: %d", exitErr.ExitStatus())
				// If we exited with status 127, it means SCP isn't available in remote server.
				// Return a more descriptive error for that.
				if exitErr.ExitStatus() == 127 {
					errChan <- errors.New("SCP is not installed in the remote server: `apt-get install openssh-client`")
					return
				}
			}
			errChan <- err
			return
		}
		doneChan <- struct{}{}
	}()
	select {
	case <-doneChan:
		fmt.Println("done with scpToRemoteFile.")
		return nil

	case err := <-errChan:
		return err

	case <-time.After(execTimeout):
		return fmt.Errorf("execution timeout.")
	}
	/////////////////////////////
}

func writeDirProtocal(
	dirInfo os.FileInfo,
	toDirPath string,
	stdinPipe io.Writer,
	stdoutBufioReader *bufio.Reader,
	uploadFunc func() error,
) error {
	fmt.Println("writeDirProtocal from", dirInfo.Name(), "to", toDirPath)
	fmt.Fprintln(stdinPipe, fmt.Sprintf("D%04o", dirInfo.Mode().Perm()), 0, toDirPath)
	if err := check(stdoutBufioReader); err != nil {
		return err
	}
	if err := uploadFunc(); err != nil {
		return err
	}
	fmt.Fprintln(stdinPipe, "E")
	if err := check(stdoutBufioReader); err != nil {
		return err
	}
	return nil
}

func recursiveUploadDir(
	fromDirPath string,
	fileInfos []os.FileInfo,
	toDirPath string,
	stdinPipe io.Writer,
	stdoutBufioReader *bufio.Reader,
) error {
	for _, fi := range fileInfos {
		localFilePath := filepath.Join(fromDirPath, fi.Name())
		fmt.Println("recursiveUploadDir from", localFilePath, "to", filepath.Join(toDirPath, fi.Name()))

		// check if this is actually a symlink to a directory. If it is
		// a symlink to a file we don't do any special behavior because uploading
		// a file just works. If it is a directory, we need to know so we
		// treat it differently.
		isSymlinkToDir := false
		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			symPath, err := filepath.EvalSymlinks(localFilePath)
			if err != nil {
				return err
			}
			symFi, err := os.Lstat(symPath)
			if err != nil {
				return err
			}
			isSymlinkToDir = symFi.IsDir()
		}
		// neither directory, nor symlink
		// then it is just a regular file
		if !fi.IsDir() && !isSymlinkToDir {
			if err := uploadFile(
				localFilePath,
				fi.Name(),
				stdinPipe,
				stdoutBufioReader,
			); err != nil {
				return err
			}
			continue
		}
		// to create the directory
		uploadFunc := func() error {
			remotePath := filepath.Join(fromDirPath, fi.Name())
			f, err := os.Open(remotePath)
			if err != nil {
				return err
			}
			defer f.Close()
			rfs, err := f.Readdir(-1)
			if err != nil {
				return err
			}
			// recursively call
			return recursiveUploadDir(
				remotePath,
				rfs,
				toDirPath,
				stdinPipe,
				stdoutBufioReader,
			)
		}
		if err := writeDirProtocal(
			fi,
			fi.Name(),
			stdinPipe,
			stdoutBufioReader,
			uploadFunc,
		); err != nil {
			return err
		}
	}
	return nil
}

func startUploadingDir(
	keyPath string,
	user string,
	host string,
	port string,
	dialTimeout time.Duration,
	fromDirPath string,
	toDirPath string,
	execTimeout time.Duration,
	stdinPipe io.Writer,
	stdoutBufioReader *bufio.Reader,
) error {
	r, err := os.Open(fromDirPath)
	if err != nil {
		return err
	}
	defer r.Close()
	ri, err := r.Stat()
	if err != nil {
		return err
	}
	if !ri.IsDir() {
		fmt.Println(ri.Name(), "is not a directory. Running scpToRemoteFile")
		return scpToRemoteFile(keyPath, user, host, port, dialTimeout, fromDirPath, toDirPath, execTimeout)
	}
	// now we need to recursively scp directories and files
	//
	// to create the directory
	uploadFunc := func() error {
		f, err := os.Open(fromDirPath)
		if err != nil {
			return err
		}
		defer f.Close()
		rfs, err := f.Readdir(-1)
		if err != nil {
			return err
		}
		return recursiveUploadDir(
			fromDirPath,
			rfs,
			toDirPath,
			stdinPipe,
			stdoutBufioReader,
		)
	}
	return writeDirProtocal(
		ri,
		ri.Name(),
		stdinPipe,
		stdoutBufioReader,
		uploadFunc,
	)
}

func scpToRemoteDir(
	keyPath string,
	user string,
	host string,
	port string,
	dialTimeout time.Duration,
	fromDirPath string,
	toDirPath string,
	execTimeout time.Duration,
) error {
	/////////////////////////////
	f, err := openToRead(keyPath)
	if err != nil {
		return err
	}
	sshSigner, err := getSSHSigner(f)
	if err != nil {
		return err
	}
	sshClient, err := getSSHClient(sshSigner, user, host, port, dialTimeout)
	if err != nil {
		return err
	}
	session, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	stdoutPipe, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	stdoutBufioReader := bufio.NewReader(stdoutPipe)
	stdinPipe, err := session.StdinPipe()
	if err != nil {
		return err
	}
	if stdinPipe == nil {
		return fmt.Errorf("stdinPipe is nil")
	}
	// defer stdinPipe.Close()
	// make sure to close this before session.Wait()
	//
	// https://blogs.oracle.com/janp/entry/how_the_scp_protocol_works
	// In all cases aside from remote-to-remote scenario the scp command
	// processes command line options and then starts an SSH connection
	// to the remote host. Another scp command is run on the remote side
	// through that connection in either source or sink mode. Source mode
	// reads files and sends them over to the other side, sink mode accepts them.
	// Source and sink modes are triggered using -f (from) and -t (to) options, respectively.
	if err := session.Start(fmt.Sprintf("scp -rvt %s", toDirPath)); err != nil {
		return err
	}
	/////////////////////////////

	/////////////////////////////
	if err := startUploadingDir(keyPath, user, host, port, dialTimeout, fromDirPath, toDirPath, execTimeout, stdinPipe, stdoutBufioReader); err != nil {
		return err
	}
	// make sure to close this before session.Wait()
	stdinPipe.Close()
	/////////////////////////////

	/////////////////////////////
	// wait for session to finish.
	doneChan := make(chan struct{})
	errChan := make(chan error)
	go func() {
		if err := session.Wait(); err != nil {
			fmt.Println("wait returns err", err)
			if exitErr, ok := err.(*ssh.ExitError); ok {
				fmt.Printf("non-zero exit status: %d", exitErr.ExitStatus())
				// If we exited with status 127, it means SCP isn't available in remote server.
				// Return a more descriptive error for that.
				if exitErr.ExitStatus() == 127 {
					errChan <- errors.New("SCP is not installed in the remote server: `apt-get install openssh-client`")
					return
				}
			}
			errChan <- err
			return
		}
		doneChan <- struct{}{}
	}()
	select {
	case <-doneChan:
		fmt.Println("done with scpToRemoteDir.")
		return nil

	case err := <-errChan:
		return err

	case <-time.After(execTimeout):
		return fmt.Errorf("execution timeout.")
	}
	/////////////////////////////
}

```

[↑ top](#ssh-scp)
<br><br><br><br>
<hr>










#### `scp` from a remote file

```go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	if err := scpFromRemoteFile(
		"/home/ubuntu/my.pem",
		"ubuntu",
		"YOUR_HOST",
		"/home/ubuntu/hello.txt",
		"/home/ubuntu/hello_copy.txt",
	); err != nil {
		panic(err)
	}
}

func scpFromRemoteFile(
	keyPath string,
	user string,
	host string,
	fromPath string,
	toPath string,
) error {
	target := fmt.Sprintf(
		"%s@%s:%s",
		user,
		host,
		fromPath,
	)
	if err := exec.Command("scp", "-i", keyPath, target, toPath).Run(); err != nil {
		return err
	}
	return nil
}

```

[↑ top](#ssh-scp)
<br><br><br><br>
<hr>
