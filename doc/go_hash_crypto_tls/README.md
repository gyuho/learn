[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: hash, crypto, tls

- [Reference](#reference)
- [hashing, hash table](#hashing-hash-table)
- [hash function](#hash-function)
- [private, public key](#private-public-key)
- [pgp](#pgp)
- [tls](#tls)
- [encrypt, decrypt data](#encrypt-decrypt-data)

[↑ top](#go-hash-crypto-tls)
<br><br><br><br><hr>


#### Reference

- [Journey into cryptography](https://www.khanacademy.org/computing/computer-science/cryptography)
- [Hash function](https://en.wikipedia.org/wiki/Hash_function)
- [Hash table](https://en.wikipedia.org/wiki/Hash_table)
- [Crypto tutorial](https://github.com/joearms/crypto_tutorial)
- [What are Bloom filters?](https://medium.com/the-story/what-are-bloom-filters-1ec2a50c68ff)

[↑ top](#go-hash-crypto-tls)
<br><br><br><br><hr>


#### hashing, hash table

**_Hashing_** is like a **fingerprint** of data. It is very useful when you need to:

1. **Secure data** because a hash function is usually *one-way* operation: you
   cannot reverse the hashed *back to the original*.
2. **Identify, index data** because a hash function always returns
   the **same hash value** for the **same input**.

<br>
There are many [*hash functions*](https://en.wikipedia.org/wiki/Hash_function)
and when you pass a key to a hash function, the function returns a hashed
value. Then you can have 1-to-1 mapping between key and value in the hash
table. Ideally, a hash function returns an unique value to every possible key,
but it is possible that two different keys return the same hashed value: [*hash
collision*](https://en.wikipedia.org/wiki/Collision_(computer_science)).

[↑ top](#go-hash-crypto-tls)
<br><br><br><br><hr>


#### hash function

Cryptographic hash functions are implemented in
[`crypto`](http://golang.org/pkg/crypto/) package:

```go
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

func hashSha1(data []byte) string {
	sum := sha1.Sum(data)
	// convert [20]byte to []byte
	return base64.StdEncoding.EncodeToString(sum[:])
}

func hashSha256(data []byte) string {
	sum := sha256.Sum256(data)
	return base64.StdEncoding.EncodeToString(sum[:])
}

func hashSha512(data []byte) string {
	sum := sha512.Sum512(data)
	return base64.StdEncoding.EncodeToString(sum[:])
}

func hashMd5(data []byte) string {
	sum := md5.Sum(data)
	// convert [20]byte to []byte
	return base64.StdEncoding.EncodeToString(sum[:])
}

func main() {
	data := []byte("Hello World!")
	fmt.Println("hashSha1:", hashSha1(data))
	fmt.Println("hashSha256:", hashSha256(data))
	fmt.Println("hashSha512:", hashSha512(data))
	fmt.Println("hashMd5:", hashMd5(data))
}

/*
hashSha1: Lve95gjOVATpfV8EL5X4nxwjKHE=
hashSha256: f4OxZX/x/FO5LcGBSKHWXfwtSx+j1ncoSt3SABJtkGk=
hashSha512: hhhE1nBOhXP+w02WfiC8/vPUJM9IvgTm3AjyvVjHKXQzcQFerYkcw88cnTS0kmS1EHUbH/nlN5N7xGtdb/TsyA==
hashMd5: 7Qdih1MuhjZehB6Sv8UNjA==
*/

```

[↑ top](#go-hash-crypto-tls)
<br><br><br><br><hr>


#### private, public key

Asymmetric encryption works: a person has private/public lock,
**_private/public key_**. And only shares the **public key** with
others, so that they can encrypt their messages. To receive such
encrypted messages, the **private key** is needed, locked with
password. Only with private key is possible to decrypt the message.

```go
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
	privateKeyPath = "key.pem"
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
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
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

```

[↑ top](#go-hash-crypto-tls)
<br><br><br><br><hr>


#### pgp

```go

```

[↑ top](#go-hash-crypto-tls)
<br><br><br><br><hr>


#### tls

```go
package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	privateKeyPath = "key.pem"
	publicKeyPath  = "public.key"
	certPath       = "cert.pem"
)

func main() {
	defer func() {
		// os.Remove(privateKeyPath)
		// os.Remove(publicKeyPath)
		// os.Remove(certPath)
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
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
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
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		panic(err)
	}
	certTemplate := x509.Certificate{
		SerialNumber: serialNumber,

		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(1000, 1, 1),

		BasicConstraintsValid: true,
		IsCA:        true,
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},

		DNSNames:       []string{"localhost"},
		EmailAddresses: []string{"test@test.com"},
	}

	// func CreateCertificate(rand io.Reader, template, parent *Certificate, pub, priv interface{}) (cert []byte, err error)
	derBytes, err := x509.CreateCertificate(rand.Reader, &certTemplate, &certTemplate, getPublicKey(privateKey), privateKey)
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
			Type:  "CERTIFICATE",
			Bytes: derBytes,
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

	func() {
		f, err := openToRead(privateKeyPath)
		if err != nil {
			panic(err)
		}
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		f.Close()
		block, pemBytes := pem.Decode(bs)
		privateKey, err := x509.ParsePKCS1PrivateKey(append(block.Bytes, pemBytes...))
		if err != nil {
			panic(err)
		}
		fmt.Println("privateKey")
		fmt.Println(privateKey)
	}()

	func() {
		f, err := openToRead(certPath)
		if err != nil {
			panic(err)
		}
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		f.Close()
		block, pemBytes := pem.Decode(bs)
		cert, err := x509.ParseCertificate(append(block.Bytes, pemBytes...))
		if err != nil {
			panic(err)
		}
		fmt.Println("cert")
		fmt.Println(cert)
	}()
}

/*
private.pem
-----BEGIN RSA PRIVATE KEY-----
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
func getPublicKey(priv interface{}) interface{} {
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

```

```go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	privateKeyPath = "key.pem"
	certPath       = "cert.pem"
)

const port = ":8080"

func main() {
	certBytes, err := ioutil.ReadFile(certPath)
	if err != nil {
		panic(err)
	}

	clientCertPool := x509.NewCertPool()
	if ok := clientCertPool.AppendCertsFromPEM(certBytes); !ok {
		panic("can't add certificate to certificate pool!")
	}

	tlsConfigServer := &tls.Config{
		ClientAuth:               tls.RequireAndVerifyClientCert,
		ClientCAs:                clientCertPool,
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
	}
	tlsConfigServer.BuildNameToCertificate()

	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", handler)
	httpServer := &http.Server{
		Addr:      port,
		TLSConfig: tlsConfigServer,
		Handler:   mainRouter,
	}

	go func() {
		fmt.Println("Serving https://localhost" + port)

		// func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
		if err := httpServer.ListenAndServeTLS(certPath, privateKeyPath); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	func() {
		fmt.Println("Sending client requests...")

		tlsCert, err := tls.LoadX509KeyPair(certPath, privateKeyPath)
		if err != nil {
			panic(err)
		}

		certBytes, err := ioutil.ReadFile(certPath)
		if err != nil {
			panic(err)
		}

		clientCertPool := x509.NewCertPool()
		if ok := clientCertPool.AppendCertsFromPEM(certBytes); !ok {
			panic("can't add certificate to certificate pool!")
		}

		tlsConfigClient := &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			RootCAs:      clientCertPool,
		}
		tlsConfigClient.BuildNameToCertificate()

		httpClient := http.DefaultClient
		httpClient.Transport = &http.Transport{TLSClientConfig: tlsConfigClient}
		resp, err := httpClient.Get("https://localhost" + port)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		rb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("response:", string(rb))
	}()
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintln(w, "Hello World!")
		fmt.Fprintf(w, "req.TLS:        %+v\n", req.TLS)
		fmt.Fprintf(w, "DNSNames:       %#q\n", req.TLS.PeerCertificates[0].DNSNames)
		fmt.Fprintf(w, "EmailAddresses: %#q\n", req.TLS.PeerCertificates[0].EmailAddresses)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

/*
Serving https://localhost:8080
Sending client requests...
response: Hello World!
req.TLS:        &{Version:771 HandshakeComplete:true DidResume:false CipherSuite:49199 NegotiatedProtocol: NegotiatedProtocolIsMutual:true ServerName:localhost PeerCertificates:[0xc82017e000] VerifiedChains:[[0xc82017e000 0xc8200a4000]] SignedCertificateTimestamps:[] OCSPResponse:[] TLSUnique:[156 55 205 75 79 27 21 192 84 244 36 226]}
DNSNames:       [`localhost`]
EmailAddresses: [`test@test.com`]
*/

```

[↑ top](#go-hash-crypto-tls)
<br><br><br><br><hr>


#### encrypt, decrypt data

> In cryptography, encryption is the process of encoding messages or
> information in such a way that only authorized parties can read it.
> ... Encryption has long been used by military and governments to facilitate
> secret communication. It is now commonly used in protecting information
> within many kinds of civilian systems.
>
> [*Encryption*](https://en.wikipedia.org/wiki/Encryption) *by Wikipedia*

<br>
Go uses [`crypto/cipher`](http://golang.org/pkg/crypto/cipher/) to encrypt
and decrypt data:

```go
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	data = []byte("Hello World!")
	// AES allows 128, 192 or 256 bit key length. That is 16, 24 or 32 byte.
	// key = "MQCVSYKEY/@#Qwed"
	// key = "MQCVSYKEY/@#Qwedf#d3/434"
	key     = "MQCVSYKEY/@#Qwedf#d3/4345FSAMNPE"
	keyPath = "my_secret.key"
)

func main() {
	fmt.Println()
	encryptedData, err := encrypt(data, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("encryptedData:", string(encryptedData))
	// encryptedData: s��cj����əf1!F���ʶzj��Z����

	fmt.Println()
	decryptedData, err := decrypt(encryptedData, key)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(data, decryptedData) {
		log.Fatalf("%s != %s", data, decryptedData)
	}
	fmt.Println("decryptedData:", string(decryptedData))
	// decryptedData: Hello World!

	fmt.Println()
	defer os.Remove(keyPath)
	if err := encryptToFile(data, key, keyPath); err != nil {
		panic(err)
	}
	decryptedDataFromFile, err := decryptFromFile(keyPath, key)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(data, decryptedDataFromFile) {
		log.Fatalf("%s != %s", data, decryptedDataFromFile)
	}
	fmt.Println("decryptFromFile:", string(decryptedDataFromFile))
	// decryptFromFile: Hello World!
}

// encrypt encrypts data with secret key.
// http://golang.org/pkg/crypto/cipher/#example_NewCFBDecrypter
func encrypt(data []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(data)
	encryptedData := make([]byte, aes.BlockSize+len(b))
	iv := encryptedData[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(encryptedData[aes.BlockSize:], []byte(b))
	return encryptedData, nil
}

func decrypt(encryptedData []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	if len(encryptedData) < aes.BlockSize {
		return nil, errors.New("encryptedData too short")
	}
	iv := encryptedData[:aes.BlockSize]
	encryptedData = encryptedData[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(encryptedData, encryptedData)
	return base64.StdEncoding.DecodeString(string(encryptedData))
}

func read(fpath string) ([]byte, error) {
	bts, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return bts, nil
}

func encryptToFile(data []byte, key, keyPath string) error {
	encryptedData, err := encrypt(data, key)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(keyPath, encryptedData, 0644)
}

func decryptFromFile(keyPath, key string) ([]byte, error) {
	encryptedData, err := read(keyPath)
	if err != nil {
		return nil, err
	}
	return decrypt(encryptedData, key)
}

```

[↑ top](#go-hash-crypto-tls)
<br><br><br><br><hr>
