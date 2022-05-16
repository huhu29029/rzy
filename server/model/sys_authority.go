package model

import (
	"time"
)

type SysAuthority struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
	AuthorityId   uint       `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName string     `json:"authorityName" gorm:"comment:角色名"`
	DefaultRouter string     `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`
}
