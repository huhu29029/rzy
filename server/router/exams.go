package router

import (
	v1 "online-exam/api/v1"

	"github.com/gin-gonic/gin"
)

func InitExamRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("exams")
	{
		ApiRouter.GET("random", v1.GetRandomExam)     // 获取随机试卷
		ApiRouter.GET("", v1.GetExams)                // 获取所有试卷
		ApiRouter.GET("exam/:id", v1.GetExam)         // 获取指定试卷
		ApiRouter.POST("submit", v1.SubmitExam)       // 提交用户答卷
		ApiRouter.GET("answer/:id", v1.GetExamResult) // 获取用户答卷
	}
}
