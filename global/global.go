package global

import (
	"github.com/spf13/viper"
	"go-admin/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//存放全局变量


var (
	GVA_DB 				*gorm.DB
	GVA_CONFIG 			config.Server
	GVA_LOG 			*zap.Logger
	GVA_VP				*viper.Viper
)