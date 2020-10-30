package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/global/response"
	"go-admin/model"
	"go-admin/model/request"
	resp "go-admin/model/response"
	"go-admin/service"
	"go-admin/utils"
)

//注册用户
func Register(c *gin.Context)  {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	UserVerify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"NickName":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
		"AuthorityId": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R,UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(),c)
		return
	}

	user := &model.SysUser{
		UserName:R.UserName,
		NickName: R.NickName,
		Password: R.Password,
		HeaderImg: R.HeaderImg,
		AuthorityId: R.AuthorityId,
	}
	err,userReturn := service.Register(*user)
	if err != nil {
		response.FailWithDetailed(response.ERROR, resp.SysUserResponse{User: userReturn}, fmt.Sprintf("%v", err), c)
	} else {
		response.OkDetailed(resp.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}