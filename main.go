package main

import (
	"fmt"
	"go-admin/core"
	"go-admin/global"
	"go-admin/initialize"
)

func main()  {
	fmt.Println("hello,go")
	global.GVA_VP = core.Viper()          	//初始化Viper
	global.GVA_LOG = core.Zap()				//初始化日志
	global.GVA_DB = initialize.Gorm() 		//初始化数据库
	initialize.MysqlTables(global.GVA_DB)	//初始化表结构
	// 程序结束前关闭数据库链接
	db,_ := global.GVA_DB.DB()
	defer db.Close()
	core.RunWindowsServer()
}
