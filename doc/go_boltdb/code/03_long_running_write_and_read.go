package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/renstrom/shortuuid"
)

var (
	dbPath     = shortuuid.New() + ".db"
	bucketName = shortuuid.New()

	// 2GB
	numKeys = 500000
	keyLen  = 100
	valLen  = 750

	keys = make([][]byte, numKeys)
	vals = make([][]byte, numKeys)

	keysForRead = make([][]byte, numKeys)

	printIdx     = 20000
	readStartIdx = 300000
	timeout      = time.Second
)

func init() {
	fmt.Println("Starting generating random data...")
	for i := range keys {
		if i%printIdx == 0 {
			fmt.Println("Generating", i+1, "/", numKeys)
		}
		keys[i] = randBytes(keyLen)
		vals[i] = randBytes(valLen)
		if i+1 == numKeys {
			fmt.Println("Generating", i+1, "/", numKeys)
			fmt.Println("Done with generating random data...")
		}
	}

	fmt.Println("Copying 'keys' to 'keysForRead'...")
	copy(keysForRead, keys)
	fmt.Println("Done with copying...")
}

func main() {
	fmt.Println("dbPath:", dbPath)
	fmt.Println("bucketName:", bucketName)

	defer os.Remove(dbPath)

	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Creating bucket:", bucketName)
	if err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucket([]byte(bucketName)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		panic(err)
	}
	fmt.Println("Done with creating bucket:", bucketName)

	readyToRead := make(chan struct{})
	doneWithWrite := make(chan struct{})
	doneWithRead := make(chan struct{})

	go func() {
		fmt.Println("Starting long-running writing...")
		for i := range keys {
			if i == readStartIdx {
				readyToRead <- struct{}{}
			}
			st := time.Now()
			ch := make(chan struct{})
			go func() {
				if err := db.Update(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte(bucketName))
					if err := b.Put(keys[i], vals[i]); err != nil {
						return err
					}
					return nil
				}); err != nil {
					panic(err)
				}
				ch <- struct{}{}
			}()
			select {
			case <-ch:
				if i%printIdx == 0 {
					fmt.Printf("Write took: %v (%d/%d)\n", time.Since(st), i+1, numKeys)
				}
				continue
			case <-time.After(timeout):
				log.Fatalf("Write timeout: %v (%d/%d)\n", time.Since(st), i+1, numKeys)
			}
		}
		doneWithWrite <- struct{}{}
		fmt.Println("Done with long-running writing...")
	}()

	go func() {
		<-readyToRead
		fmt.Println("Starting long-running reading...")
		for i := range keysForRead {
			st := time.Now()
			ch := make(chan struct{})
			go func() {
				if err := db.View(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte(bucketName))
					k := keys[i]
					v := b.Get(k)
					// fmt.Printf("Read: %s ---> %s\n", k, v)
					_ = k
					_ = v
					return nil
				}); err != nil {
					panic(err)
				}
				ch <- struct{}{}
			}()
			select {
			case <-ch:
				if i%printIdx == 0 {
					fmt.Printf("Read took: %v (%d/%d)\n", time.Since(st), i+1, numKeys)
				}
				continue
			case <-time.After(timeout):
				log.Fatalf("Read timeout: %v (%d/%d)\n", time.Since(st), i+1, numKeys)
			}
		}
		doneWithRead <- struct{}{}
		fmt.Println("Done with long-running reading...")
	}()

	<-doneWithWrite
	<-doneWithRead
	fmt.Println("Done!")
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
