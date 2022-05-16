package router

import (
	v1 "online-exam/api/v1"

	"github.com/gin-gonic/gin"
)

func InitQuestionRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("questions")
	{
		ApiRouter.GET("random", v1.RandomQuestion)         // 随机题目
		ApiRouter.GET("kt", v1.GetQuestionWithAlgorithm)   // 算法推荐题目
		ApiRouter.GET("wrong", v1.GetWrongQuetion)         // 错题
		ApiRouter.POST("random", v1.GetQuetionsByChapters) // 根据章节获取随机题目
		ApiRouter.POST("submit", v1.SubmitAnswer)          // 提交某道题目的回答
		ApiRouter.GET("question/:id", v1.GetQuestion)      // 某个题目 未遵循 RESTful 规范，原因：https://github.com/gin-gonic/gin/issues/205
		ApiRouter.GET("answer/:id", v1.GetQuestionResult)  // 查询用户答题情况
	}
}
