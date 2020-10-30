package request

//注册结构体
type RegisterStruct struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	NickName    string `json:"nickName" gorm:"default:'李华'"`
	HeaderImg   string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}