package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"hospital-admin-api/config"
	"hospital-admin-api/global"
	"os"
	"path"
)

// Logrus 日志记录到文件
func Logrus() *logrus.Logger {
	// 日志文件
	fileName := path.Join(config.C.Log.FilePath, config.C.Log.FileName)

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("err", err)
	}

	global.Logger = logrus.New()

	// 设置输出
	global.Logger.Out = src

	// 设置日志级别
	global.Logger.SetLevel(logrus.DebugLevel)

	// 设置日志格式
	global.Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return global.Logger
}
