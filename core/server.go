package core

import (
	"fmt"
	"go-admin/global"
	"go-admin/initialize"
	"go.uber.org/zap"
	"time"
)

func RunWindowsServer()  {

	if global.GVA_CONFIG.System.UseMultipoint {  //
		// 初始化redis服务
	}
	Router := initialize.Routers()  //初始化路由
	Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Println("欢迎使用", address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}

type server interface {
	ListenAndServe() error
}