package cache

import (
	"me-english/utils/config"

	"github.com/go-redis/redis"
)

func redisConnect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDR,
		Password: "",
		DB:       0,
	})
}
