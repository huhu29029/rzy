package service

import (
	"online-exam/global"
	"online-exam/model"
	"os/exec"
	"strconv"
)

// 获取问题答题情况
func GetQuestionResult(answerId int) model.QuestionResult {
	// 获取用户答题信息
	var answer model.Answers
	// TODO 用户鉴权
	global.DB.Preload("Options").First(&answer, answerId)

	// 获取题目信息
	var question model.Questions
	global.DB.Preload("Options").First(&question, answer.QuestionId)

	// 题目 + 用户答题 => 批改题目
	var result model.QuestionResult

	result.Title = question.Title
	result.Analysis = question.Analysis
	result.Type = question.Type

	isRight := true

	for _, option := range question.Options {
		isChoosed := false

		for _, answerOption := range answer.Options {
			if option.ID == answerOption.QuestionOptionId {
				isChoosed = true
			}
		}

		if (isChoosed && !option.IsRight) || (!isChoosed && option.IsRight) {
			isRight = false
		}

		result.Options = append(result.Options, model.QuestionOptionResult{
			Title:     option.Title,
			IsRight:   option.IsRight,
			IsChoosed: isChoosed,
		})
	}

	result.IsRight = isRight

	return result
}

func GetUserQuestionHistory(userId uint, limit int, offset int) []model.QuestionHistory {
	// 获取用户做题记录
	var answers []model.Answers
	global.DB.Where("user_id = ?", userId).Limit(limit).Offset(offset).Find(&answers)

	var questionHistory []model.QuestionHistory

	for _, question := range answers {
		questionResult := GetQuestionResult(int(question.ID))

		questionHistory = append(questionHistory, model.QuestionHistory{
			ID:             question.ID,
			StartTime:      question.StartTime,
			EndTime:        question.EndTime,
			QuestionId:     question.QuestionId,
			QuestionResult: questionResult,
		})
	}
	return questionHistory
}

func NewCount(userId uint) int64 {
	var answers []model.Answers
	var count int64
	global.DB.Model(&model.Answers{}).Where("user_id = ?", userId).Find(&answers).Count(&count)
	exec.Command("echo " + strconv.Itoa(int(count)))
	for _, question := range answers {
		questionResult := GetQuestionResult(int(question.ID))
		if questionResult.IsRight == true {
			count--
		}
	}
	return count
}

func GetWrongQuestionHistory(userId uint, limit int, offset int) []model.QuestionHistory {
	// 获取用户做题记录
	var answers []model.Answers
	global.DB.Where("user_id = ?", userId).Limit(limit).Offset(offset).Find(&answers)

	var questionHistory []model.QuestionHistory

	for _, question := range answers {
		questionResult := GetQuestionResult(int(question.ID))
		if questionResult.IsRight == false {
			questionHistory = append(questionHistory, model.QuestionHistory{
				ID:             question.ID,
				StartTime:      question.StartTime,
				EndTime:        question.EndTime,
				QuestionId:     question.QuestionId,
				QuestionResult: questionResult,
			})
		}
	}
	for _, question := range questionHistory {
		type WrongQuestion struct {
			userID       uint
			questionID   uint
			title        string
			analysis     string
			questionType uint
			isRight      bool
		}

		global.DB.Table("Wrong_question_book").Select("userID", "questionID", "title",
			"analysis", "type", "is_right").Create(&WrongQuestion{
			userID:       userId,
			questionID:   question.QuestionId,
			title:        question.QuestionResult.Title,
			analysis:     question.QuestionResult.Analysis,
			questionType: question.QuestionResult.Type,
			isRight:      question.QuestionResult.IsRight,
		})
	}
	return questionHistory
}
