package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/boltdb/bolt"
)

const (
	// these will create 2GB database.
	numKeys = 500000
	keyLen  = 100
	valLen  = 750

	bucketName = "test_bucket"
	writable   = true
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	dbPath := filepath.Join(usr.HomeDir, "test.db")

	if err := os.RemoveAll(dbPath); err != nil {
		panic(err)
	}

	// Open the dbPath data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Start a writable transaction.
	tx, err := db.Begin(writable)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	// Use the transaction
	b, err := tx.CreateBucket([]byte(bucketName))
	if err != nil {
		panic(err)
	}

	// Write to database
	for i := 0; i < numKeys; i++ {
		k := randBytes(keyLen)
		v := randBytes(valLen)
		fmt.Println("Writing", i, "/", numKeys)
		if err := b.Put(k, v); err != nil {
			panic(err)
		}
	}

	// Commit the transaction
	fmt.Println("Committing...")
	if err := tx.Commit(); err != nil {
		panic(err)
	}

	fmt.Println("Done")
}

const (
	// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randBytes(n int) []byte {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}
