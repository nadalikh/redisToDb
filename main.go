package main

import (
	"context"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	//Redis connection
	var r = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := r.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	//Creating mongodb connection
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.8.115:27017"))
	if err != nil {
		log.Fatal(err)
	}
	//Start the connection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = db.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var imeiDetail []interface{}
	for {
		keys := r.Keys("[0-9]*").Val()
		for _, key := range keys {
			numberOfDetailsForAnImei := len(r.LRange(key, 0, -1).Val())
			for i := 0; i < numberOfDetailsForAnImei; i++ {
				imeiDetail = append(imeiDetail, bson.D{{key, r.RPop(key).Val()}})
			}
		}
		collection := db.Database("GPS").Collection("statuses")
		res, err := collection.InsertMany(ctx, imeiDetail)
		if err != nil {
			return
		}
		print(res.InsertedIDs)
		imeiDetail = nil
		time.Sleep(time.Second * 20)
	}
}
