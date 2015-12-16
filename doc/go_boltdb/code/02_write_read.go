package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/renstrom/shortuuid"
)

var (
	dbPath     = shortuuid.New() + ".db"
	bucketName = shortuuid.New()

	numKeys = 10
	keyLen  = 3
	valLen  = 7

	keys = make([][]byte, numKeys)
	vals = make([][]byte, numKeys)
)

func init() {
	fmt.Println("Starting writing random data...")
	for i := range keys {
		keys[i] = randBytes(keyLen)
		vals[i] = randBytes(valLen)
	}
	fmt.Println("Done with writing random data...")
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

	fmt.Println("Starting writing...")
	if err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			return err
		}
		for i := range keys {
			if err := b.Put(keys[i], vals[i]); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}
	fmt.Println("Done with writing...")

	fmt.Println("Starting reading...")
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		for i := range keys {
			k := keys[i]
			v := b.Get(k)
			fmt.Printf("Read: %s ---> %s\n", k, v)
		}
		return nil
	}); err != nil {
		panic(err)
	}
	fmt.Println("Done with reading...")
}

/*
Starting writing random data...
Done with writing random data...
dbPath: MAD96XbjYFenxMQBSnMxvY.db
bucketName: zeXeJqUhpMhYD5kPVMRVBi
Starting writing...
Done with writing...
Starting reading...
Read: uWN ---> qpjxtgO
Read: vPH ---> kgBMOcz
Read: unX ---> XgGmeOJ
Read: gLd ---> ekPAyFk
Read: dVa ---> yjgXook
Read: otW ---> nyiOWvj
Read: rBV ---> viufTKO
Read: wto ---> NVFceGz
Read: Dtg ---> RlPULZS
Read: PLC ---> NqLporc
Done with reading...
*/

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
