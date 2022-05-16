package model

import (
	"online-exam/global"
)

type JwtBlacklist struct {
	global.MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
