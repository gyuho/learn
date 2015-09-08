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
