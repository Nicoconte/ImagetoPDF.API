package data

import (
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client = InitClient()

func InitClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     Config.RedisConnection,
		Password: "",
		DB:       0,
	})
}
