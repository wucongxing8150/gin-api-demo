package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"hospital-admin-api/config"
)

// Viper 初始化配置文件
func Viper() {
	// 如需增加环境判断 在此处增加
	// 可根据 命令行 > 环境变量 > 默认值 等优先级进行判别读取
	viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("未找到该文件")
		} else {
			panic("读取失败")
		}
	}

	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(&config.C); err != nil {
		panic(fmt.Errorf("解析配置文件失败, err:%s \n", err))
	}

	// 自动监听配置修改
	viper.WatchConfig()

	// 配置文件发生变化后同步到全局变量Conf
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := viper.Unmarshal(&config.C); err != nil {
			panic(fmt.Errorf("重载配置文件失败, err:%s \n", err))
		}
	})
}
