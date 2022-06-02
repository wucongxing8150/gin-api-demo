// @Description: viper
// @Author: wucongxing
// @Date:2021/12/23 13:52

package viper

import (
	"github.com/spf13/viper"
)

// Init
// @Description:初始化viper 读取配置文件
func Init() {
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

	// // 将读取的配置信息保存至全局变量Conf
	// if err := viper.Unmarshal(&Config.C); err != nil {
	// 	panic(fmt.Errorf("解析配置文件失败, err:%s \n", err))
	// }
	//
	// // 注意！！！配置文件发生变化后要同步到全局变量Conf
	// viper.OnConfigChange(func(in fsnotify.Event) {
	// 	if err := viper.Unmarshal(&Config.C); err != nil {
	// 		panic(fmt.Errorf("解析配置文件失败, err:%s \n", err))
	// 	}
	// })

	// 自动监听配置修改
	viper.WatchConfig()
}
