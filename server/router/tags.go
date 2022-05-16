package router

import (
	v1 "online-exam/api/v1"

	"github.com/gin-gonic/gin"
)

func InitChaptersRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("chapters")
	{
		ApiRouter.GET("", v1.GetChapters) // 获取所有章节
	}
}
