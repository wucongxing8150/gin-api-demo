package main

import (
	"fmt"
	"gin-api-demo/App/Api/Router"
	"gin-api-demo/Cores"
	vip "gin-api-demo/Cores/viper"
	"github.com/spf13/viper"
	"strconv"
)

func main() {
	// 初始化Viper 加载配置
	vip.Init()

	// 初始化路由
	r := Router.Init()

	// 加载核心方法
	Cores.Init()

	if err := r.Run(":" + strconv.Itoa(viper.GetInt("port"))); err != nil {
		fmt.Printf("启动失败:%v\n\n", err)
	}
}
