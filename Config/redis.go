package Config

import "github.com/spf13/viper"

type Redis struct {
	Host     string `mapstructure:"host" json:"host"`           // 服务器地址
	Port     int    `mapstructure:"port" json:"port"`           // 服务器端口
	Password string `mapstructure:"password" json:"password"`   // 密码
	PoolSize int    `mapstructure:"pool-size" json:"pool-size"` // 连接池大小
}

func GetDevRedis() *Redis {
	return &Redis{
		Host:     viper.GetString("dev-redis.host"),
		Port:     viper.GetInt("dev-redis.port"),
		Password: viper.GetString("dev-redis.password"),
		PoolSize: viper.GetInt("dev-redis.pool-size"),
	}
}

func GetMasterRedis() *Redis {
	return &Redis{
		Host:     viper.GetString("master-redis.host"),
		Port:     viper.GetInt("master-redis.port"),
		Password: viper.GetString("master-redis.password"),
		PoolSize: viper.GetInt("master-redis.pool-size"),
	}
}
