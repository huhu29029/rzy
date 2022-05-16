package core

import (
	"fmt"
	"online-exam/global"
	"online-exam/initialize"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 关闭多点登录
	// if global.CONFIG.System.UseMultipoint {
	// 	// 初始化redis服务
	// 	initialize.Redis()
	// }
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	Server is Listening on: http://127.0.0.1%s
	Swagger API documentation:http://127.0.0.1%s/api/swagger/index.html

`, address, address)
	global.LOG.Error(s.ListenAndServe().Error())
}
