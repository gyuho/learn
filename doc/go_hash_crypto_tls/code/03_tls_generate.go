package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	privateKeyPath = "private.pem"
	publicKeyPath  = "public.key"
	certPath       = "cert.pem"
)

func main() {
	defer func() {
		os.Remove(privateKeyPath)
		os.Remove(publicKeyPath)
		os.Remove(certPath)
	}()

	// write private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	privateKeyFile, err := openToOverwrite(privateKeyPath)
	if err != nil {
		panic(err)
	}
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
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}
	if _, err := publicKeyFile.Write(ssh.MarshalAuthorizedKey(publicKey)); err != nil {
		panic(err)
	}
	publicKeyFile.Close()

	// write cert
	certTemplate := x509.Certificate{
		Signature:          []byte("test"),
		SignatureAlgorithm: x509.SHA512WithRSA,

		PublicKeyAlgorithm: x509.RSA,
		PublicKey:          convertPublicKey(privateKey),

		Version: 0,

		Issuer: pkix.Name{
			Country:            []string{"Issuer"},
			Organization:       []string{"Issuer"},
			OrganizationalUnit: []string{"Issuer"},
		},

		Subject: pkix.Name{
			Country:            []string{"Subject"},
			Organization:       []string{"Subject"},
			OrganizationalUnit: []string{"Subject"},
		},

		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Minute),

		SubjectKeyId:          []byte("TEST"),
		BasicConstraintsValid: true,
		IsCA: true,

		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		panic(err)
	}
	certTemplate.SerialNumber = serialNumber
	certTemplate.IsCA = true
	certTemplate.KeyUsage |= x509.KeyUsageCertSign
	//
	// func CreateCertificate(rand io.Reader, template, parent *Certificate, pub, priv interface{}) (cert []byte, err error)
	derBytes, err := x509.CreateCertificate(rand.Reader, &certTemplate, &certTemplate, convertPublicKey(privateKey), privateKey)
	if err != nil {
		panic(err)
	}
	certKeyFile, err := openToOverwrite(certPath)
	if err != nil {
		panic(err)
	}
	if err := pem.Encode(
		certKeyFile,
		&pem.Block{
			Type:    "CERTIFICATE",
			Headers: map[string]string{"TEST_KEY": "TEST_VALUE"},
			Bytes:   derBytes,
		},
	); err != nil {
		panic(err)
	}
	certKeyFile.Close()

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

	// read public key
	func() {
		f, err := openToRead(certPath)
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
private.pem
-----BEGIN RSA PRIVATE KEY-----
TEST_KEY: TEST_VALUE

MIICXgIBAAKBgQDVgp1RWPs80m3gAJBrbhCjL21FyXGn9AEnINDg5W5D2aaKnE0D
26lW6BqR0h1iRI9F60HSMxeCqDDMou4To4/BIHZpePP/X7S+zQJPdeRuB07/0989
oig/UYkA2fbhKCbpBGKLf15BjW8sOtYzFscFhoByYWZ+w0zFblmNfr3GMwIDAQAB
AoGAFcuJh55Ptzu735vvIihQJnhW7ULNCVoNLBNbfzmscdyr9YZTDkvEE40J+Uy7
lyZsgbSsOWrhwYKtyJXxO6v8poiJc5TF4EYiL2zRu0mZAzZWIkcj87Qf9Z4QMHXa
iFP26NwJNgimX5+sHbP8FaSdgcrfrYHattao47bNkfyhjEECQQDh+A6nkiIJKpC0
gKSShQ8+A5Nv4glmsTgEJo3hCltdq3WHvuk2WatYEF6l01lA2yN1CE8ocrFd1jho
Co/oWTIRAkEA8eKvEjnMWPiiAOPuQxpOqbY9HxnypCx+/2fc5CVNcbvvqY+UcC/R
Iu+gMdVnsw4nZEwd9awQsSntix6D+F8wAwJBANJP0TPdKphlaXDWGlXUSa9qHJsR
Qba2UnBqgbplrUus/SJuaRgQtQytj6m+318hlgqixSncNYAklTMgQXf7LEECQQDt
JnC7D8vX1z0OXmp1g89n+PKIEaqhZ7bDthMN47zAK6BXwBuqulbzR7jp4u8e0Fuy
rCYbfa2H5TGuWibNVpX9AkEAivN5+lgFSzBgz6QLdTahz5d49yOhROqd6V5aPES6
T5f99iFzRFmr0LKsDHC+kolFxlws1hn0XE0QOTBryLMycw==
-----END RSA PRIVATE KEY-----

public.key
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDVgp1RWPs80m3gAJBrbhCjL21FyXGn9AEnINDg5W5D2aaKnE0D26lW6BqR0h1iRI9F60HSMxeCqDDMou4To4/BIHZpePP/X7S+zQJPdeRuB07/0989oig/UYkA2fbhKCbpBGKLf15BjW8sOtYzFscFhoByYWZ+w0zFblmNfr3GMw==

cert.pem
-----BEGIN CERTIFICATE-----
TEST_KEY: TEST_VALUE

MIICSTCCAbKgAwIBAgIQTt/O2lGhYG1VwaCDFj6zGDANBgkqhkiG9w0BAQ0FADA2
MRAwDgYDVQQGEwdTdWJqZWN0MRAwDgYDVQQKEwdTdWJqZWN0MRAwDgYDVQQLEwdT
dWJqZWN0MB4XDTE2MDEwOTIzNTgwMFoXDTE2MDEwOTIzNTkwMFowNjEQMA4GA1UE
BhMHU3ViamVjdDEQMA4GA1UEChMHU3ViamVjdDEQMA4GA1UECxMHU3ViamVjdDCB
nzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA1YKdUVj7PNJt4ACQa24Qoy9tRclx
p/QBJyDQ4OVuQ9mmipxNA9upVugakdIdYkSPRetB0jMXgqgwzKLuE6OPwSB2aXjz
/1+0vs0CT3XkbgdO/9PfPaIoP1GJANn24Sgm6QRii39eQY1vLDrWMxbHBYaAcmFm
fsNMxW5ZjX69xjMCAwEAAaNYMFYwDgYDVR0PAQH/BAQDAgIEMBMGA1UdJQQMMAoG
CCsGAQUFBwMBMA8GA1UdEwEB/wQFMAMBAf8wDQYDVR0OBAYEBFRFU1QwDwYDVR0j
BAgwBoAEVEVTVDANBgkqhkiG9w0BAQ0FAAOBgQB5pxTJjUwkfwO+Lxjc1Joyx98R
gwiRB/25EXucguL7UtmD/GeIfxW7zxdoTChcOX2c2qgKcwKvH9V1TKBLc1/co4Ui
NkqelHYczPWIXueqJxJa7NfEjhc2ehMNcuPOJbESGVvegY8jcsKckrc8H7RG2bMj
XHBkArYcakTskF5A4Q==
-----END CERTIFICATE-----

*/

// https://golang.org/src/crypto/tls/generate_cert.go
func convertPublicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
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
