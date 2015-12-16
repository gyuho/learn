package main

import (
	"fmt"
	"math/rand"
	"os/user"
	"path/filepath"
	"time"

	"github.com/boltdb/bolt"
)

var (
	dbPath     = "test.db"
	bucketName = "test_bucket"

	// 5GB
	// numKeys = 1250000

	// 2GB
	numKeys = 500000
	keyLen  = 100
	valLen  = 750
)

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	dbPath = filepath.Join(usr.HomeDir, dbPath)
}

func main() {
	fmt.Println("dbPath:", dbPath)
	fmt.Println("bucketName:", bucketName)

	// Open the dbPath data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Starting writing random data...")
	if err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			return err
		}
		for i := 0; i < numKeys; i++ {
			if i%10000 == 0 {
				fmt.Println("Writing", i+1, "/", numKeys)
			}
			if err := b.Put(randBytes(keyLen), randBytes(valLen)); err != nil {
				return err
			}
			if i+1 == numKeys {
				fmt.Println("Writing", i+1, "/", numKeys)
				fmt.Println("Done with writing random data...")
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}

	fmt.Println("Done! Saved:", db.Path())
}

func randBytes(n int) []byte {
	const (
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
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
