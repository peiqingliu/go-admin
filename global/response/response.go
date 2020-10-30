package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//全局同意方法返回

type Response struct {
	Code 	int				`json:"code"`
	Data	interface{}		`json:"data"`
	Msg 	string			`json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int,msg string, data interface{}, c *gin.Context)  {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS,"操作成功", map[string]interface{}{}, c)
}


func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS,  message,map[string]interface{}{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS,"操作成功", data,  c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS,message, data,  c)
}

func Fail(c *gin.Context) {
	Result(ERROR, "操作失败",map[string]interface{}{},  c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, message, map[string]interface{}{}, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code,message, data,  c)
}