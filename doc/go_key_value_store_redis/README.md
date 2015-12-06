[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: key/value store, redis

- [key/value store](#keyvalue-store)
- [redis](#redis)

[↑ top](#go-keyvalue-store-redis)
<br><br><br><br><hr>


#### key/value store

A **key-value store** stores data with key-value pair using data structures
like *dictionary*, *map*, or *hash*. It work much different than traditional
*SQL database*, or *relational database*, so it's often called
[*NoSQL*](https://en.wikipedia.org/wiki/NoSQL) database. None is better than
the other. If you have relation-less data and need fast-retrieval by providing
the key, the key-value database is a better choice. If you deal with data sets,
each of which relates to another in complicated ways, the relational-database
is the right choice.

The simplest way is **to use your file systems as a key-value storage**:

```go
package main
 
import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
 
func exist(fpath string) bool {
	// Does a directory exist
	st, err := os.Stat(fpath)
	if err != nil {
		return false
	}
	if st.IsDir() {
		return true
	}
	if _, err := os.Stat(fpath); err != nil {
		if strings.Contains(err.Error(), "no such file") {
			return false
		}
	}
	return true
}
 
func getHash(bt []byte) string {
	if bt == nil {
		return ""
	}
	h := sha512.New()
	h.Write(bt)
	sha512Hash := hex.EncodeToString(h.Sum(nil))
	return sha512Hash
}
 
func main() {
	var events = []string{
		"Hello World!",
		"Hello World!",
		"different",
	}
	for _, event := range events {
		eventHash := getHash([]byte(event))
		if !exist(eventHash) {
			if err := ioutil.WriteFile(eventHash, nil, 0644); err != nil {
				log.Fatal(err)
			}
			fmt.Println("Saved", event, "with", eventHash)
		} else {
			fmt.Println("Found duplicate events:", event)
		}
	}
}
 
/*
Saved Hello World! with 861844d6704e8573fec34d967e20bcfef3d424cf48be04e6dc08f2bd58c729743371015ead891cc3cf1c9d34b49264b510751b1ff9e537937bc46b5d6ff4ecc8
Found duplicate events: Hello World!
Saved different with 49d5b8799558e22d3890d03b56a6c7a46faa1a7d216c2df22507396242ab3540e2317b870882b2384d707254333a8439fd3ca191e93293f745786ff78ef069f8
*/
```

[↑ top](#go-keyvalue-store-redis)
<br><br><br><br><hr>


#### redis

> Redis is an open source, BSD licensed, advanced **key-value cache** and
> **store**. It is often referred to as a **data structure server** since keys
> can contain [strings](http://redis.io/topics/data-types-intro#strings),
> [hashes](http://redis.io/topics/data-types-intro#hashes),
> [lists](http://redis.io/topics/data-types-intro#lists),
> [sets](http://redis.io/topics/data-types-intro#sets), [sorted
> sets](http://redis.io/topics/data-types-intro#sorted-sets),
> [bitmaps](http://redis.io/topics/data-types-intro#bitmaps),
> [hyperloglogs](http://redis.io/topics/data-types-intro#hyperloglogs).
>
> [*Redis*](http://redis.io/)

```go
package main
 
import (
	"fmt"
	"log"
 
	"github.com/garyburd/redigo/redis"
)
 
func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
 
	// SET command only lets you have one value for the key
	c.Do("SET", "MY_KEY", "MY_VALUE")
 
	// HSET lets you have multiple fields and values
	c.Do("HSET", "myhash", "field1", "Hello")
	// myhash = { field1 : "Hello" }
 
	val1, err := redis.String(c.Do("GET", "MY_KEY"))
	if err != nil {
		fmt.Println("key not found")
	}
	fmt.Println(val1)
	// MY_VALUE
 
	val2, err := redis.String(c.Do("HGET", "myhash", "field1"))
	if err != nil {
		fmt.Println("key not found")
	}
	fmt.Println(val2)
	// Hello
}
```

[↑ top](#go-keyvalue-store-redis)
<br><br><br><br><hr>
