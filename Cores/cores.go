package Cores

import (
	"gin-api-demo/Cores/mysql"
	"gin-api-demo/Cores/viper"
)

func Init() {
	// 初始化Viper 加载配置
	viper.Init()

	// 初始化数据库连接
	mysql.Init()

	// 初始化redis连接

}
