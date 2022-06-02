package main

import (
	"fmt"
	"gin-api-demo/App/Api/Router"
	"gin-api-demo/Cores"
	"github.com/spf13/viper"
	"strconv"
)

func main() {
	// 初始化路由
	r := Router.Init()

	// 加载核心方法
	Cores.Init()

	if err := r.Run(":" + strconv.Itoa(viper.GetInt("port"))); err != nil {
		fmt.Println("启动失败:%v\n", err)
	}
}
