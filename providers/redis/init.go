package redis

import (
	"github.com/go-redis/redis"
	"github.com/gomsa/tools/env"
	"github.com/micro/go-micro/util/log"
)

// NewClient 创建新的 redis 连接
func NewClient() (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     env.Getenv("REDIS_HOST", "127.0.0.1:6379"),
		Password: env.Getenv("REDIS_PASSWORD", ""),
		DB:       0, // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Log(pong, err)
	}
	return
}
