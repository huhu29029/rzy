package service

import (
	"online-exam/global"
	"online-exam/model"
)

func GetUserExamHistory(userId uint, limit int, offset int) []model.ExamHistory {
	// 获取用户考试记录
	var examAnswers []model.ExamAnswers
	global.DB.Where("user_id = ?", userId).Limit(limit).Offset(offset).Find(&examAnswers)

	var examHistory []model.ExamHistory

	// 根据用户考试历史构建返回数据
	for _, examAnswer := range examAnswers {
		var exam model.Exams
		global.DB.Where("id = ?", examAnswer.ExamId).First(&exam)
		examHistory = append(examHistory, model.ExamHistory{
			StartTime: examAnswer.StartTime,
			EndTime:   examAnswer.EndTime,
			ExamId:    examAnswer.ID,
			ExamTitle: exam.Title,
		})
	}
	return examHistory
}
