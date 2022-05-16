package router

import (
	v1 "online-exam/api/v1"

	"github.com/gin-gonic/gin"
)

func InitTagsRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("tags")
	{
		ApiRouter.GET("", v1.GetTags) // 获取所有标签
	}
}
