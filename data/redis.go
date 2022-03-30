package data

import (
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client = InitClient()

func InitClient() *redis.Client {
	opt, err := redis.ParseURL(Config.RedisUrl)

	if err != nil {
		panic(1)
	}
	return redis.NewClient(opt)
}
