package utils

import (
	
	"encoding/base64"
	"fmt"
	"time"
)

func ShortenURL(inputURL string)(string){
    ts:= time.Now().UnixNano()
	fmt.Println("timestamp is",ts)
	ts_bytes := []byte(fmt.Sprintf("%d",ts))
	key:= base64.StdEncoding.EncodeToString(ts_bytes)
	fmt.Println("key is",key)
	key = key[:len(key)-2];
	return key[16:];
}