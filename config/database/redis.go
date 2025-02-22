package database

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

// Global Redis client
var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Change if needed
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("✅ Connected to Redis")
}
