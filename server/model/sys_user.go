package model

import (
	"online-exam/global"

	uuid "github.com/satori/go.uuid"
)

// AuthorityId: 1 管理员 2 普通用户
type SysUser struct {
	global.MODEL
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`
	Username    string       `json:"userName" gorm:"comment:用户登录名"`
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`
	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg   string       `json:"headerImg" gorm:"default:https://i.loli.net/2021/03/26/NfFzdA6YQKjV9Ga.jpg;comment:用户头像"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId uint         `json:"authorityId" gorm:"default:2;comment:用户角色ID"`

}
