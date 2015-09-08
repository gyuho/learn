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
