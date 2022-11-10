package main

import (
	. "fmt"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	var r = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := r.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	Println(r.Keys("[0-9]*"))
}
