package Cores

import (
	"gin-api-demo/Cores/mysql"
	"gin-api-demo/Cores/redis"
)

func Init() {

	// 初始化数据库连接
	mysql.Init()

	// 初始化redis连接
	redis.Init()

}
