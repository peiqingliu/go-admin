package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

//用户 结构体
type SysUser struct {
	gorm.Model
	UUID        uuid.UUID    	`json:"uuid" gorm:"comment:用户UUID"`
	UserName	string			`json:"userName" gorm:"comment:用户登录名"`
	Password	string 			`json:"password" gorm:"comment:用户登录密码"`
	NickName	string			`json:"nickName" gorm:"comment:用户昵称"`
	HeaderImg	string			`json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Authority   SysAuthority 	`json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       	`json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}
