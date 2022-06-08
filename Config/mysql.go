package Config

import "github.com/spf13/viper"

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                           // 服务器地址
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                           // 端口
	DbName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                  // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`               // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`               // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-cons" json:"MaxIdleConns" yaml:"max-idle-cons"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-cons" json:"MaxOpenConns" yaml:"max-open-cons"` // 打开到数据库的最大连接数
	Debug        bool   `mapstructure:"debug" json:"debug" yaml:"debug"`                        // 是否开启Gorm全局日志
}

func GetDevMysql() *Mysql {
	return &Mysql{
		Host:         viper.GetString("dev-database.host"),
		Port:         viper.GetInt("dev-database.port"),
		DbName:       viper.GetString("dev-database.dbName"),
		Username:     viper.GetString("dev-database.userName"),
		Password:     viper.GetString("dev-database.password"),
		MaxIdleConns: viper.GetInt("dev-database.mysql-max-idle-cons"),
		MaxOpenConns: viper.GetInt("dev-database.mysql-max-open-cons"),
		Debug:        viper.GetBool("dev-database.mysql-debug"),
	}
}

func GetMasterMysql() *Mysql {
	return &Mysql{
		Host:         viper.GetString("dev-database.host"),
		Port:         viper.GetInt("dev-database.port"),
		DbName:       viper.GetString("dev-database.dbName"),
		Username:     viper.GetString("dev-database.userName"),
		Password:     viper.GetString("dev-database.password"),
		MaxIdleConns: viper.GetInt("dev-database.mysql-max-idle-cons"),
		MaxOpenConns: viper.GetInt("dev-database.mysql-max-open-cons"),
		Debug:        viper.GetBool("dev-database.mysql-debug"),
	}
}
