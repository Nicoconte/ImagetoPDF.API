package data

import (
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client = InitClient()

func InitClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
