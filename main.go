package main

import "C"
import (
	"fmt"
	"hospital-admin-api/api/router"
	"hospital-admin-api/config"
	"hospital-admin-api/core"
	"strconv"
)

func main() {
	// 加载配置文件
	core.Viper()

	// 加载日志
	core.Logrus()

	// 加载数据库
	core.Mysql()

	// 加载redis缓存
	core.Redis()

	// 初始化路由-加载中间件
	r := router.Init()

	if err := r.Run(":" + strconv.Itoa(config.C.Port)); err != nil {
		fmt.Printf("启动失败:%v\n\n", err)
	}
}
