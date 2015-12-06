[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: leveldb

- [leveldb](#leveldb)

[↑ top](#go-leveldb)
<br><br><br><br><hr>


#### leveldb

> LevelDB is a light-weight, single-purpose library for persistence with
> bindings to many platforms.
>
> [*LevelDB*](http://leveldb.org/)


```go
package main
 
import (
	"log"
	"runtime"
	"strings"
	"time"
 
	"github.com/syndtr/goleveldb/leveldb"
)
 
func init() {
	maxCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("Concurrent execution with", maxCPU, "CPUs.")
}
 
func main() {
	start := time.Now()
 
	levelDBpath := "./db"
	ldb, err := leveldb.OpenFile(levelDBpath, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ldb.Close()
 
	maxSet := 450 // import maxSet * 10 messages
	decodedEvents, receiptHandlersToDelete := getMessages(maxSet)
 
	foundEvent := make(map[string]bool)
	idsToPut := []string{}
	for _, msg := range decodedEvents {
 
		// check id collision within a batch
		if _, ok := foundEvent[msg.EventHash]; ok {
			continue
		} else {
			foundEvent[msg.EventHash] = true
 
			data, err := ldb.Get([]byte(msg.EventHash), nil)
			if err != nil {
				if !strings.Contains(err.Error(), "not found") {
					log.Fatal(err)
				}
			}
			if data != nil {
				log.Printf("Found Duplicate Event: %s", string(data))
				continue
			}
 
			idsToPut = append(idsToPut, msg.EventHash)
		}
	}
 
	// do your data import job here
 
	log.Println("maintain the lookup table")
	for _, id := range idsToPut {
		// maintain the lookup table
		if err := ldb.Put([]byte(id), []byte("true"), nil); err != nil {
			log.Fatal(err)
		}
	}
}

```

[↑ top](#go-leveldb)
<br><br><br><br><hr>
