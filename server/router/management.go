package router

import (
	v1 "online-exam/api/v1"

	"github.com/gin-gonic/gin"
)

func InitManagementRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("admin")
	{
		ApiRouter.GET("users", v1.GetUsers)
		ApiRouter.GET("users/:id/questions", v1.GetUserQuestionHistory)
		ApiRouter.GET("users/:id/exams", v1.GetUserExamHistory)
		ApiRouter.GET("questions", v1.GetQuestions)
		ApiRouter.POST("questions", v1.AddQuestion)
		ApiRouter.DELETE("questions/:id", v1.DelQuestion)
		ApiRouter.PUT("questions/:id", v1.EditQuestion)
		ApiRouter.GET("exams", v1.ManageExams)
		ApiRouter.POST("exams/exam", v1.AddExam)
		ApiRouter.DELETE("exams/exam/:id", v1.DelExam)
		ApiRouter.PUT("exams/exam/:id", v1.EditExam)
	}
}
