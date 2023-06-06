package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 全局变量
var (
	Db     *gorm.DB       // 数据库
	Logger *logrus.Logger // 日志
	Redis  *redis.Client  // redis
)
