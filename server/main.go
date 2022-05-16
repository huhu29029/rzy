package main

import (
	"context"
	"online-exam/core"
	"online-exam/global"
	"online-exam/initialize"
)

// @title online-exam API
// @version 0.0.1
// @description online exam APIs
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /api
func main() {
	global.VP = core.Viper()      // 初始化Viper
	global.LOG = core.Zap()       // 初始化zap日志库
	global.DB = initialize.Gorm() // gorm连接数据库
	global.MongoDB = initialize.MongoDB()
	initialize.Redis() // 连接Redis
	if global.DB != nil {
		initialize.MysqlTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	if global.MongoDB != nil {
		defer func() {
			if err := global.MongoDB.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()
	}
	core.RunWindowsServer()
}
