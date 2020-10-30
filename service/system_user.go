package service

import (
	"errors"
	"github.com/satori/go.uuid"
	"go-admin/global"
	"go-admin/model"
	"go-admin/utils"
	"gorm.io/gorm"
)

//用户注册
func Register(u model.SysUser) (err error,userInter model.SysUser) {
	var user model.SysUser
	er := global.GVA_DB.Where("username = ?",u.UserName).First(&user).Error
	if !errors.Is(er,gorm.ErrRecordNotFound) {  // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	//保存
	err = global.GVA_DB.Create(&u).Error
	return err, u
}
