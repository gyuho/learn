[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: hash, crypto, tls

- [Reference](#reference)
- [hashing, hash table](#hashing-hash-table)
- [hash function](#hash-function)
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
