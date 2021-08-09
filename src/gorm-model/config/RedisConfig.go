package config

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v7"
)

// SetupRedis - initializing redis connection
func SetupRedis() *redis.Client {
	redisURL := os.Getenv("REDIS_URL")
	redisPass := os.Getenv("REDIS_PASS")
	redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		panic(err.Error())
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPass,
		DB:       redisDb,
	})

	return redisClient
}
