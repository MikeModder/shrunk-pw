package main

import (
	"log"
	"github.com/go-redis/redis"
)

// all database related functions
// mostly just wrappers for ease of use

var (
	redisClient *redis.Client
)

func setupDB() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v\n", err)
	}
}

func getFullURL(short string) string {

}

func saveURL(short, long string) bool {

}
