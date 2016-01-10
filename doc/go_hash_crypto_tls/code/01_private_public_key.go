package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

var (
	privateKeyPath = "private.pem"
	publicKeyPath  = "public.key"
)

func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}

func main() {
	defer func() {
		os.Remove(privateKeyPath)
		os.Remove(publicKeyPath)
	}()

	// GenerateKey generates an RSA keypair
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	// write private key
	privateKeyFile, err := openToOverwrite(privateKeyPath)
	if err != nil {
		panic(err)
	}
	// if err := gob.NewEncoder(privateKeyFile).Encode(privateKey); err != nil {
	// 	panic(err)
	// }
	//
	// "PEM" format is a method to encode binary data into text
	if err := pem.Encode(
		privateKeyFile,
		&pem.Block{
			Type:    "RSA PRIVATE KEY",
			Headers: map[string]string{"TEST_KEY": "TEST_VALUE"},
			Bytes:   x509.MarshalPKCS1PrivateKey(privateKey),
		},
	); err != nil {
		panic(err)
	}
	privateKeyFile.Close()

	// write public key
	publicKeyFile, err := openToOverwrite(publicKeyPath)
	if err != nil {
		panic(err)
	}
	// if err := gob.NewEncoder(publicKeyFile).Encode(&privateKey.PublicKey); err != nil {
	// 	panic(err)
	// }
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}
	if _, err := publicKeyFile.Write(ssh.MarshalAuthorizedKey(publicKey)); err != nil {
		panic(err)
	}
	publicKeyFile.Close()

	// read private key
	func() {
		f, err := openToRead(privateKeyPath)
		if err != nil {
			panic(err)
		}
		defer func() {
			f.Close()
		}()
		tbytes, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(f.Name())
		fmt.Println(string(tbytes))
	}()

	// read public key
	func() {
		f, err := openToRead(publicKeyPath)
		if err != nil {
			panic(err)
		}
		defer func() {
			f.Close()
		}()
		tbytes, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(f.Name())
		fmt.Println(string(tbytes))
	}()
}

/*
private.key
-----BEGIN RSA PRIVATE KEY-----
TEST_KEY: TEST_VALUE

MIICXQIBAAKBgQCgMRqaLHKOdf9HPElaav3G8yxVWAeb0eB4Wvy7QDsHgOroJlmg
+LylCoUCHdd0Ly/rBJH6AGMcDjBh/jEE5YK8kj/wV+UGh7g+3n9I4ez3rNrYiytM
6c3baTxDYN73RSw6tjjGIzg/tmuUrk6i0eslL83g5INMjKs31LCsoitMsQIDAQAB
AoGAcyLX+/f2Xm5xDMJH9rTvsg8VzkF3Noeizt6Wx/9ibgI61KC7yvb8n6Lv9pV8
RgWka0bdpNKiaYfJPqV0lhBf5gX0PImKn4mNiIwS5rMjps8Ymeth+sNXJh8n9OX1
PYOIVQ5oERqyHQiqG+AJ+rk6tls4NGNFFdN39aJfkuDfgAECQQDD5FYBiAceCkB2
5J+kaS7ZhbLgOMge9MM/gmn16uJrOxVEOUeUwP74fqu/tuYfotMFNLYmGnpXzavo
Adt1k6W1AkEA0Vh1etiZvOm4s0fzS18fKI2esIjri8tNx+kLbdrYXm6T7jAEzWel
SsGUs/PD+7sxWf9Xz9YTduCgpqeSNxnojQJBAJXx0Reo/PG0nTWkuMJLtQ3R9mMF
c8GmT1HszJjtq1SzTAsF4VHvDPw/Uc4U/T9YDjjc6VRvThipmR2lVkxAsUUCQQCh
htaGpe/hgpjfxAlmQ4vgF321CsBsCb8HG7qU1cITAtEjfGuILYutJbZeLx0t8569
qTaRB8XW+LUcQbmgyF3VAkBWCQf3Z7wyPM+1qeXzIa4ZUZcpMOcaQ98TuSQhVVcD
pZnPsO5Ni1wjKhKWHP3lOXx9N+e9NtCunfjyv5C2SIs3
-----END RSA PRIVATE KEY-----

public.key
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQCgMRqaLHKOdf9HPElaav3G8yxVWAeb0eB4Wvy7QDsHgOroJlmg+LylCoUCHdd0Ly/rBJH6AGMcDjBh/jEE5YK8kj/wV+UGh7g+3n9I4ez3rNrYiytM6c3baTxDYN73RSw6tjjGIzg/tmuUrk6i0eslL83g5INMjKs31LCsoitMsQ==

*/

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
