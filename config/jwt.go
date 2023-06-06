package config

type Jwt struct {
	SignKey string `mapstructure:"sign-key" json:"sign-key" yaml:"sign-key"` // 私钥
	Ttl     int    `mapstructure:"ttl" json:"ttl" yaml:"ttl"`                // 过期时间 小时
	Algo    string `mapstructure:"algo" json:"algo" yaml:"algo"`             // 加密方式
}
