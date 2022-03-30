package data

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client = InitClient()

func InitClient() *redis.Client {
	switch os.Getenv("APP_ENV") {
	case "dev":
		return redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", Config.RedisUrl, Config.RedisPort),
			Password: "",
			DB:       0,
		})

	case "prod":
		opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))

		if err != nil {
			panic(1)
		} else {
			return redis.NewClient(opt)
		}

	default:
		panic(1)
	}
}
