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
