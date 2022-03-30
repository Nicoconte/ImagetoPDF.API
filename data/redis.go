package data

import (
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client = InitClient()

func InitClient() *redis.Client {
	switch Config.Env {
	case "dev":
		return redis.NewClient(&redis.Options{
			Addr:     Config.RedisUrl,
			Password: "",
			DB:       0,
		})

	case "prod":
		opt, err := redis.ParseURL(Config.RedisUrl)

		if err != nil {
			panic(1)
		} else {
			return redis.NewClient(opt)
		}

	default:
		panic(1)
	}
}
