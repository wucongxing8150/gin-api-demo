package core

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"hospital-admin-api/config"
	"hospital-admin-api/global"
	"strconv"
)

// Redis redis缓存
func Redis() {
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Host + ":" + strconv.Itoa(config.C.Redis.Port),
		Password: config.C.Redis.Password, // no password set
		DB:       0,                       // use default DB
		PoolSize: config.C.Redis.PoolSize,
	})
	_, err := global.Redis.Ping(context.Background()).Result()
	if err != nil {
		panic("redis初始化失败! " + err.Error())
	}
	fmt.Println("初始化redis成功......")
}
