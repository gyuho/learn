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
