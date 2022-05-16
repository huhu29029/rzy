package router

import (
	v1 "online-exam/api/v1"
	"online-exam/middleware"

	"github.com/gin-gonic/gin"
)

func InitNiepanRouter(Router *gin.RouterGroup) {
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	ApiRouter := Router.Group("niepan")
	{
		ApiRouter.GET("questions", v1.GetQuestions2Niepan) // 涅槃
	}
}
