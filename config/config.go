package config

var C Config

type Config struct {
	Port  int    `mapstructure:"port" json:"port" yaml:"port"`
	Env   string `mapstructure:"env" json:"Env" yaml:"Env"`
	Mysql Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Log   Log    `mapstructure:"log" json:"log" yaml:"log"`
	Redis Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Jwt   Jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
