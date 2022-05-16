package router

import (
	v1 "online-exam/api/v1"

	"github.com/gin-gonic/gin"
)

func InitXAPIRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("xapi")
	{
		ApiRouter.POST("new", v1.NewXAPIData) // 增加学习行为
	}
}
