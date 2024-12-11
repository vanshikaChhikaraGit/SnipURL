package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func SetKey(ctx *context.Context, rdb *redis.Client, key string, value string){

	fmt.Println("Setting key", key, "to", value, "in Redis")
	rdb.Set(*ctx,key,value,0)
	fmt.Println("The key", key, "has been set to", value, "in Redis")

}