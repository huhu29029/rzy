package source

import (
	"online-exam/global"
	"online-exam/model"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

var authorities = []model.SysAuthority{
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: 1, AuthorityName: "管理员", DefaultRouter: "dashboard"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: 2, AuthorityName: "普通用户", DefaultRouter: "dashboard"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_authorities 表数据初始化
func (a *authority) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "9528"}).Find(&[]model.SysAuthority{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authorities 表初始数据成功!")
		return nil
	})
}
