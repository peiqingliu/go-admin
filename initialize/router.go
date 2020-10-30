package initialize

import (
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/middleware"
	"go-admin/router"
	"net/http"
)

//初始化路由
func Routers() *gin.Engine{
	var Router = gin.Default()
	//1、 为用户头像和文件提供静态地址
	Router.StaticFS(global.GVA_CONFIG.Local.Path,http.Dir(global.GVA_CONFIG.Local.Path))
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use Routers logger")
	//2、设置跨域
	Router.Use(middleware.Cors())
	global.GVA_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup)   // 注册用户路由
	global.GVA_LOG.Info("router register success")
	return Router
}