package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis"
)

// all database related functions
// mostly just wrappers for ease of use

var (
	redisClient *redis.Client
)

func setupDB() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     appConfig.RedisURL,
		Password: appConfig.RedisPassword,
		DB:       appConfig.RedisDB,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v\n", err)
	}

	log.Printf("pong from redis: %v\n", pong)
}

func hasFullURL(short string) bool {
	err := redisClient.Get(fmt.Sprintf("short:%s", short)).Err()
	if err != nil {
		// Making the assumption that if there's an error trying to get
		// the value that the key doesn't exist.
		// TODO: verify this
		return false
	}
	return true
}

func getFullURL(short string) string {
	full, err := redisClient.Get(fmt.Sprintf("short:%s", short)).Result()
	if err != nil {
		return ""
	}

	if !strings.HasPrefix(full, "http://") && !strings.HasPrefix(full, "https://") {
		return fmt.Sprintf("http://%s", full)
	}
	return full
}

func saveURL(short, long string) bool {
	err := redisClient.Set(fmt.Sprintf("short:%s", short), long, 0).Err()
	if err != nil {
		log.Printf("failed to save short %s: %v\n", short, err)
		return false
	}
	return true
}
