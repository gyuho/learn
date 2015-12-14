package main

import (
	"fmt"
	"os/user"
	"path/filepath"
	"syscall"
	"time"

	"github.com/boltdb/bolt"
)

const (
	bucketName = "test_bucket"
	writable   = false
)

func main() {
	fmt.Println("read")
	read()

	fmt.Println()
	fmt.Println("readMapPopulate")
	readMapPopulate()
}

/*
read
bolt.Open took 47.419µs
bolt Read took: 75.532225ms

readMapPopulate
bolt.Open took 254.063181ms
bolt Read took: 51.274025ms
*/

func read() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	dbPath := filepath.Join(usr.HomeDir, "test.db")

	to := time.Now()
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 5 * time.Minute, ReadOnly: true})
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
	fmt.Println("bolt Read took:", time.Since(tr))
}

func readMapPopulate() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	dbPath := filepath.Join(usr.HomeDir, "test.db")

	to := time.Now()
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 5 * time.Minute, ReadOnly: true, MmapFlags: syscall.MAP_POPULATE})
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
	fmt.Println("bolt Read took:", time.Since(tr))
}
