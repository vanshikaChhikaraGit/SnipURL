package utils

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func RedisClient()(*redis.Client){
	err:=godotenv.Load()
	if err!=nil{
		fmt.Print("failed to load env variables")
		return nil
	}
    rdb:= redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
		Password:os.Getenv("REDIS_PASSWORD"),
		DB:0,
	})

	return rdb
}
