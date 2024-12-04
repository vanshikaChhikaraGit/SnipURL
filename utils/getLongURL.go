package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func GetLongUrl(ctx *context.Context, rdb *redis.Client, key string)(string,error){

	longURL,err:=rdb.Get(*ctx,key).Result()

	if err==redis.Nil{
		return "", fmt.Errorf("short URL not found")  
	}else if err!=nil{
		return "", fmt.Errorf("short URL not found")
	}
	return longURL,nil
}