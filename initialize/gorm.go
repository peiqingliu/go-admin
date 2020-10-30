package initialize

import (
	"go-admin/global"
	"go-admin/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB  {

	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB  {
	m := global.GVA_CONFIG.Mysql
	//"root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := m.Username + ":" + m.Password +"@(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db,err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode))
	if err!= nil{
		global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	}else {
		sqlDB,_ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}


// MysqlTables 注册数据库表专用
func MysqlTables(db *gorm.DB)  {
	//自动创建数据表
	err := db.AutoMigrate(
		model.SysUser{},
		model.SysAuthority{},
		model.SysBaseMenu{},
		model.SysBaseMenuParameter{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(mod bool) *gorm.Config {
	if mod {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
}