package router

import (
	v1 "online-exam/api/v1"

	"github.com/gin-gonic/gin"
)

func InitTestRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("test")
	{
		ApiRouter.GET("just4test", v1.Just4test) // just for test
	}

	HomeRouter := Router.Group("")
	{
		HomeRouter.GET("", v1.HelloWorld)
	}
}
