[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: boltdb

- [Reference](#reference)
- [Write random data](#write-random-data)
- [Read all data](#read-all-data)
- [Write, read](#write-read)
- [Long-running write and read](#long-running-write-and-read)

[↑ top](#go-boltdb)
<br><br><br><br><hr>


#### Reference

- [`boltdb/bolt`](https://github.com/boltdb/bolt)

[↑ top](#go-boltdb)
<br><br><br><br><hr>


#### Write random data

```go
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

```

[↑ top](#go-boltdb)
<br><br><br><br><hr>


#### Read all data

```go
package main

import (
	"flag"
	"fmt"
	"os/user"
	"path/filepath"
	"syscall"
	"time"

	"github.com/boltdb/bolt"
)

/*
read with MAP_POPULATE flag...
bolt.Open took 2.879063477s
bolt read took: 51.952703ms

read without MAP_POPULATE flag...
bolt.Open took 1.568715ms
bolt read took: 13.795348869s
*/

var (
	dbPath     = "test.db"
	bucketName = "test_bucket"
	mapPop     = true
	writable   = false
)

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	dbPath = filepath.Join(usr.HomeDir, "test.db")

	mapPt := flag.Bool(
		"populate",
		true,
		"'true' for MAP_POPULATE flag.",
	)
	flag.Parse()
	mapPop = *mapPt
}

func main() {
	read(dbPath, mapPop)
}

func read(dbPath string, mapPop bool) {
	opt := &bolt.Options{Timeout: 5 * time.Minute, ReadOnly: true}
	if mapPop {
		fmt.Println("read with MAP_POPULATE flag...")
		opt = &bolt.Options{Timeout: 5 * time.Minute, ReadOnly: true, MmapFlags: syscall.MAP_POPULATE}
	} else {
		fmt.Println("read without MAP_POPULATE flag...")
	}

	to := time.Now()
	db, err := bolt.Open(dbPath, 0600, opt)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("bolt.Open took", time.Since(to))

	tr := time.Now()
	tx, err := db.Begin(writable)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	bk := tx.Bucket([]byte(bucketName))
	c := bk.Cursor()

	for k, v := c.First(); k != nil; k, v = c.Next() {
		// fmt.Printf("%s ---> %s.\n", k, v)
		_ = k
		_ = v
	}
	fmt.Println("bolt read took:", time.Since(tr))
}

```

[↑ top](#go-boltdb)
<br><br><br><br><hr>


#### Write, read

```go
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

```

[↑ top](#go-boltdb)
<br><br><br><br><hr>


#### Long-running write and read

```go
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

	printIdx     = 2000
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

```

[↑ top](#go-boltdb)
<br><br><br><br><hr>

