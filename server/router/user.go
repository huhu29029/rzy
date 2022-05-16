package router

import (
	v1 "online-exam/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("", v1.GetUserInfo)                                                               // 获取用户信息
		UserRouter.POST("changePassword", v1.ChangePassword)                                             // 修改密码
		UserRouter.PUT("setUserInfo", v1.SetUserInfo)                                                    // 设置用户信息
		UserRouter.GET("exams", v1.GetExamHistory)                                                       // 获取用户考试历史
		UserRouter.GET("questions", v1.GetQuestionHistory)                                               // 获取用户做题历史
		UserRouter.GET("WrongQuestion", v1.GetWrongQuestionHistory)                                      //获取用户错题
		UserRouter.DELETE("DeleteQuestionFromWrongQuestionBook", v1.DeleteQuestionFromWrongQuestionBook) //从错题本删除错题
	}
}
