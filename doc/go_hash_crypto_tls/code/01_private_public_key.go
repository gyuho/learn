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
	privateKeyPath = "private-key.pem"
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

	// "PEM" format is a method to encode binary data into text
	if err := pem.Encode(
		privateKeyFile,
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	); err != nil {
		panic(err)
	}
	privateKeyFile.Close()

	// write public key
	pbf, err := openToOverwrite(publicKeyPath)
	if err != nil {
		panic(err)
	}

	// if err := gob.NewEncoder(pbf).Encode(&privateKey.PublicKey); err != nil {
	// 	panic(err)
	// }

	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}
	if _, err := pbf.Write(ssh.MarshalAuthorizedKey(publicKey)); err != nil {
		panic(err)
	}
	pbf.Close()

	pbf, err = openToAppend(publicKeyPath)
	if err != nil {
		panic(err)
	}
	// add comment
	pbf.WriteString(" gyuhox@gmail.com")
	pbf.Close()

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
private-key.pem
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC9XMz3R6pSYgGwxPX91UJlgdd8/VGLmZgk2GVkWLFa+GdcdWdg
EIC4N2h53K2APbo/0i2KDJtIQX4P4BRDlzxI78oDhfEmssrS5TLPoHTT5dzhQkdn
JFKKmkpJ1qo/o1rnO4Hv4rTDw0vK0gy/ep5i2OY6oeAyAFc7ENLJDmPk5QIDAQAB
AoGAKbkQ0EtSE+TUSoabTNp4TrVVLY0DMqcdBsFHVdzU9x5UZ+LWbCw2sGBE/NTK
xb7UEsvUjN5KOJl1lTniPSJNfcNt8uUybH5D4sW7ea1vYF2BPm8aGqNEGg4WcxvR
rZs8lfMH9E32MkrTNDSjgCV1/NsmRqddqRcFmfRv6FLXzQECQQDQuEN4BqpwsPc/
Ox+8aa2qB1+gihVknGO0WIFmUv3asby8B5xDl5w5Pz4qTMibc4+U3mFm+D2zWpcn
zijE23IZAkEA6EIAY3xljbBRwzcYQp/TCSw8RYqn4p+mFFA4KMZIgSP33z/Da01f
Hil/5L3AiNOzni2K+5L9uF/Cyure7TNarQJBALpg8I6LlUNQI1jpWOuMirFcKD5Z
T8UqCbaPme1fiqPxNxHI0fdhuPU9zitDqZd21+4drmien6o66ON4qtsvAnECQQDC
WwjsN5rb2KJzE9WvWwNEd8nv/7nBwQs/kGmOZW8i8jBol3k2f8aK/PtTNR664T07
rqzRHQ5IjYn6OFVYdVL5AkB87/GGL9XFfyKP9ZYQygJRrsvN8z9ZF4YfM38CAhgd
aBSYmk20KkbND74dV8iNqT3mwh9SnkHC6fn5AiZcE6cR
-----END RSA PRIVATE KEY-----

public.key
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC9XMz3R6pSYgGwxPX91UJlgdd8/VGLmZgk2GVkWLFa+GdcdWdgEIC4N2h53K2APbo/0i2KDJtIQX4P4BRDlzxI78oDhfEmssrS5TLPoHTT5dzhQkdnJFKKmkpJ1qo/o1rnO4Hv4rTDw0vK0gy/ep5i2OY6oeAyAFc7ENLJDmPk5Q==
 gyuhox@gmail.com

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

func openToAppend(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}
