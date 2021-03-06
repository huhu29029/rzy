package router

import (
	v1 "online-exam/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("register", v1.Register)
		BaseRouter.POST("captcha", v1.Captcha)
	}
	return BaseRouter
}
