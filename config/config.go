package config

type Server struct {
	//全局环境变量
	System System 		`mapstructure:"system" json:"system" yaml:"system"`
	//gorm
	Mysql  Mysql      	`mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//日志
	Zap    Zap     		`mapstructure:"zap" json:"zap" yaml:"zap"`
	//oss 文件对象存储
	Local Local 		`mapstructure:"local" json:"local" yaml:"local"`
}
