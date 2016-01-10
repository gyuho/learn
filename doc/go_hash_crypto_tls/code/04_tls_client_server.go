package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	privateKeyPath = "private.pem"
	publicKeyPath  = "public.key"
	certPath       = "cert.pem"
)

func main() {
	defer func() {
		// os.Remove(privateKeyPath)
		// os.Remove(publicKeyPath)
		// os.Remove(certPath)
	}()

	f0, err := openToRead(privateKeyPath)
	if err != nil {
		panic(err)
	}
	bs0, err := ioutil.ReadAll(f0)
	if err != nil {
		panic(err)
	}
	f0.Close()
	block0, pemBytes0 := pem.Decode(bs0)
	privateKey, err := x509.ParsePKCS1PrivateKey(append(block0.Bytes, pemBytes0...))
	if err != nil {
		panic(err)
	}
	fmt.Println("privateKey")
	fmt.Println(privateKey)

	f1, err := openToRead(certPath)
	if err != nil {
		panic(err)
	}
	bs1, err := ioutil.ReadAll(f1)
	if err != nil {
		panic(err)
	}
	f1.Close()
	block1, pemBytes1 := pem.Decode(bs1)
	cert, err := x509.ParseCertificate(append(block1.Bytes, pemBytes1...))
	if err != nil {
		panic(err)
	}
	fmt.Println("cert")
	fmt.Println(cert)
}

func openToRead(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return f, err
	}
	return f, nil
}

func openToOverwrite(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}
