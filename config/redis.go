package config

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`                // 服务器地址
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`                // 服务器端口
	Password string `mapstructure:"password" json:"password" yaml:"password"`    // 密码
	PoolSize int    `mapstructure:"pool-size" json:"pool-size" yaml:"pool-size"` // 连接池大小
}
