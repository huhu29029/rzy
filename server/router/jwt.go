package router

import (
	v1 "online-exam/api/v1"

	"github.com/gin-gonic/gin"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("jwt")
	{
		ApiRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
