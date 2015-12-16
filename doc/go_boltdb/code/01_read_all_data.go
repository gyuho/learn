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
