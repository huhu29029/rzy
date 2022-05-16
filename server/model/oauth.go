package model

import "online-exam/global"

type OauthOursparkUsers struct {
	global.MODEL
	Name       string  `json:"name"`
	Role       int     `json:"role"`
	University uint    `json:"university"`
	Stuno      string  `json:"stuno"`
	SysUserId  uint    `json:"sysUserId"`
	Oid        string  `json:"oid"`
	SysUser    SysUser `json:"-" gorm:"foreignKey:SysUserId;references:ID"`
}
