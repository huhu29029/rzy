package initialize

import (
	_ "online-exam/docs"
	"online-exam/middleware"
	"online-exam/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	Router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("api")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		router.InitTestRouter(PublicGroup) // just for test
		// router.InitQuestionRouter(PublicGroup) // 答题路由
		// router.InitExamRouter(PublicGroup) // 考试路由
		// router.InitTagsRouter(PublicGroup) // 标签
		router.InitOAuthRouter(PublicGroup) // OAuth
	}
	CORSGroup := Router.Group("api")
	CORSGroup.Use(middleware.Cors()) // 如需跨域可以打开
	{
		router.InitNiepanRouter(CORSGroup) // 涅槃
	}
	PrivateGroup := Router.Group("api")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		router.InitJwtRouter(PrivateGroup)        // jwt相关路由
		router.InitUserRouter(PrivateGroup)       // 注册用户路由
		router.InitQuestionRouter(PrivateGroup)   // 答题路由
		router.InitExamRouter(PrivateGroup)       // 考试路由
		router.InitTagsRouter(PrivateGroup)       // 标签路由
		router.InitChaptersRouter(PrivateGroup)   // 章节路由
		router.InitManagementRouter(PrivateGroup) // 管理员路由
		router.InitXAPIRouter(PrivateGroup)       // 学习行为收集路由
	}
	return Router
}
