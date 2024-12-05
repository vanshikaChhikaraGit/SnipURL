package utils

import (
	
	"os"

	"github.com/go-redis/redis/v8"
)

func RedisClient()(*redis.Client){
	
    rdb:= redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
		Password:os.Getenv("REDIS_PASSWORD"),
		DB:0,
	})

	return rdb
}
