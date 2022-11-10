package main

import (
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
	for _, val := range r.LRange("111", 0, -1).Val() {
		println(val)
	}

	//keys := r.Keys("[0-9]*").Val()
	//for _, key := range keys {
	//	println(r.LRange("111", 0, -1).Val())
	//	//r.RPop(key)
	//}
}
