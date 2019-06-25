package redis

import (
	"github.com/go-redis/redis"
	"github.com/gomsa/tools/env"
	"github.com/micro/go-micro/util/log"
)

// redis 连接
var (
	// redis 连接
	Client *redis.Client
)

func init() {
	Client := redis.NewClient(&redis.Options{
		Addr:     env.Getenv("REDIS_HOST", "127.0.0.1"),
		Password: env.Getenv("REDIS_PASSWORD", ""),
		DB:       0, // use default DB
	})
	pong, err := Client.Ping().Result()
	if err != nil {
		log.Log(pong, err)
	}
}
