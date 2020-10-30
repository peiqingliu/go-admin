package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/api/v1"
)

//初始化用户路由
func InitUserRouter(Router *gin.RouterGroup)  {
	UserRouter := Router.Group("user")

	UserRouter.POST("register",v1.Register)
}