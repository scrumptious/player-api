package redis

import (
	redis "github.com/redis/go-redis/v9"
	"log"
)

var R *redis.Client

func InitRedis(options *redis.Options) {
	R = redis.NewClient(options)
	if R == nil {
		log.Fatalln("Failed to initialize Redis")
	}
}
