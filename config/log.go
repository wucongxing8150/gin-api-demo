package config

type Log struct {
	FilePath string `mapstructure:"file-path" json:"file-path" yaml:"file-path"` // 日志目录
	FileName string `mapstructure:"file-name" json:"file-name" yaml:"file-name"` // 日志名称
}
