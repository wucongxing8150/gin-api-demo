package redis

import (
	"context"
	"gin-api-demo/Config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"strconv"
)

var Rdb *redis.Client

var r *Config.Redis

func Init() {
	if viper.GetString("env") == "product" {
		r = Config.GetMasterRedis()
	} else {
		r = Config.GetDevRedis()
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + strconv.Itoa(r.Port),
		Password: r.Password, // no password set
		DB:       0,          // use default DB
		PoolSize: r.PoolSize,
	})
	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("redis初始化失败! " + err.Error())
	}
}
