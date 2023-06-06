package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"hospital-admin-api/config"
	"hospital-admin-api/global"
	"time"
)

// // 定义自己的Writer
// type MyWriter struct {
// 	logger *logrus.Logger
// }
//
// // 实现gorm/logger.Writer接口
// func (m *MyWriter) Printf(format string, v ...interface{}) {
// 	logstr := fmt.Sprintf(format, v...)
// 	// 利用loggus记录日志
// 	m.logger.Info(logstr)
// }
//
// func NewMyWriter() *MyWriter {
// 	l := Logrus()
//
// 	return &MyWriter{logger: l}
// }

func Mysql() {
	var err error

	m := config.C.Mysql
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", m.Username,
		m.Password, m.Host, m.Port, m.DbName, "10s")

	newLogger := logger.New(
		global.Logger,
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	global.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		Logger: newLogger,
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	sqlDB, _ := global.Db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 调试模式
	// Db.LogMode(m.Debug)    // 打印sql
	// Db.SingularTable(true) // 全局禁用表名复数
	fmt.Println("初始化数据库成功......")
}
