package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
	"time"

	"github.com/boltdb/bolt"
	"github.com/renstrom/shortuuid"
)

var (
	dbPath     = ""
	bucketName = "test_bucket"
	writable   = false
)

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	dbPath = filepath.Join(usr.HomeDir, shortuuid.New()+".db")
}

func main() {
	defer os.RemoveAll(dbPath)

	to := time.Now()
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 5 * time.Minute, ReadOnly: true, MmapFlags: syscall.MAP_POPULATE})
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("bolt.Open took", time.Since(to))

	// long-running read

	// long-running write

}
