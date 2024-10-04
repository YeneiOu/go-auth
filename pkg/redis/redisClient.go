package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func GetNewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return rdb
}
