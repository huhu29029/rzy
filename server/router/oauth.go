package router

import (
	v1 "online-exam/api/v1"

	"github.com/gin-gonic/gin"
)

func InitOAuthRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("oauth")
	{
		ApiRouter.GET("ourspark", v1.Ourspark) // 火花统一登录
	}
}
