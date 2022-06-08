package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func Init() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 100,
	})
	result := rdb.Ping(context.Background())
	fmt.Println("redis ping:", result.Val())
	if result.Val() != "PONG" {
		// 连接有问题
		return nil
	}
	return rdb
}
