package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func Init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "Wucongxing1..", // no password set
		DB:       0,               // use default DB
		PoolSize: 100,
	})
	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("redis初始化失败! " + err.Error())
	}
}
