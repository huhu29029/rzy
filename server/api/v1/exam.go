package v1

import (
	"fmt"
	"online-exam/global"
	"online-exam/model"
	"online-exam/model/request"
	"online-exam/model/response"
	"online-exam/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Tags Exam
// @Summary 获取随机试卷
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} request.Exam "response.Response{data: request.Exam}"
// @Router /exams/random [get]
func GetRandomExam(c *gin.Context) {
	var exam model.Exams

	global.DB.Order("rand()").First(&exam)

	global.DB.Model(&exam).Preload("Options").Association("Questions").Find(&exam.Questions)

	var questions []request.Question
	for _, question := range exam.Questions {
		var options []request.QuestionOption
		for _, option := range question.Options {
			options = append(options, request.QuestionOption{
				MODEL: option.MODEL,
				Title: option.Title,
			})
		}
		questions = append(questions, request.Question{
			MODEL:   question.MODEL,
			Title:   question.Title,
			Type:    question.Type,
			Options: options,
		})
	}

	finalExam := request.Exam{
		MODEL:     exam.MODEL,
		Title:     exam.Title,
		Questions: questions,
	}

	response.OkWithData(response.ExamResponse{
		Exam: finalExam,
	}, c)
}

// @Tags Exam
// @Summary 查询所有试卷
// @Security ApiKeyAuth
// @Param page query string false "所在页面 默认值：1"
// @Param size query string false "分页大小 默认值：10"
// @Param userId query string false "所属用户 id"
// @Produce application/json
// @Success 200 {object} response.GetExamsResponse "response.Response{data: response.GetExamsResponse}"
// @Router /exams [get]
func GetExams(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	userId := c.DefaultQuery("userId", "")

	limit, err := strconv.Atoi(size)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	offset, err := strconv.Atoi(page)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	offset = (offset - 1) * limit

	var exams []model.Exams
	var count int64
	if userId == "" {
		global.DB.Limit(limit).Offset(offset).Find(&exams)
		global.DB.Model(&model.Exams{}).Count(&count)
	} else {
		global.DB.Limit(limit).Offset(offset).Where("user_id = ?", userId).Find(&exams)
		global.DB.Model(&model.Exams{}).Where("user_id = ?", userId).Count(&count)
	}

	response.OkWithData(response.GetExamsResponse{
		Total: uint(count),
		Exams: exams,
	}, c)
}

// @Tags Exam
// @Summary 查询某个试卷
// @Security ApiKeyAuth
// @Param id path string true "试卷id"
// @Produce application/json
// @Success 200 {object} response.ExamResponse "response.Response{data: response.ExamResponse}"
// @Router /exams/exam/:id [get]
func GetExam(c *gin.Context) {
	examId := c.Param("id")
	userId := getUserID(c)

	// 考试限制: 提交次数，提交间隔、考试前练习次数
	totalCount := 3
	submitDuration := 30 * 60 * time.Second // 考试提交间隔30分钟

	var count int64

	global.DB.Model(&model.ExamAnswers{}).Where("user_id = ? and exam_id = ?", userId, examId).Count(&count)
	fmt.Println(count)

	if int(count) >= totalCount {
		response.Result(2000, struct{}{}, fmt.Sprintf("考试次数已用完，总次数%d", totalCount), c)
		return
	}

	var examAnswer model.ExamAnswers
	global.DB.Select("end_time").Where("user_id = ? and exam_id= ?", userId, examId).Last(&examAnswer)

	// location, _ := time.LoadLocation("Asia/Shanghai")
	layout := "2006-01-02 15:04:05"
	// lastSubmitTime, _ := time.ParseInLocation(layout, examAnswer.EndTime, location)
	lastSubmitTime := examAnswer.EndTime
	nextSubmitTime := lastSubmitTime.Add(submitDuration)
	fmt.Println(lastSubmitTime, nextSubmitTime, time.Now())

	if time.Now().Before(nextSubmitTime) {
		response.Result(2001, struct{}{}, fmt.Sprintf("未满足考试提交间隔30分钟，下次提交时间为%s", nextSubmitTime.Format(layout)), c)
		return
	}

	global.DB.Model(&model.Answers{}).Where("user_id = ?", userId).Count(&count)
	practiceCount := 50
	if int(count) < practiceCount {
		response.Result(2000, struct{}{}, fmt.Sprintf("请在练习后在参加考试，至少练习%d道，当前%d道", practiceCount, count), c)
		return
	}

	var exam model.Exams

	result := global.DB.First(&exam, examId)

	if result.Error != nil {
		response.FailWithMessage("get exam error", c)
		return
	}

	global.DB.Model(&exam).Preload("Options").Association("Questions").Find(&exam.Questions)

	var questions []request.Question
	for _, question := range exam.Questions {
		var options []request.QuestionOption
		for _, option := range question.Options {
			options = append(options, request.QuestionOption{
				MODEL: option.MODEL,
				Title: option.Title,
			})
		}
		questions = append(questions, request.Question{
			MODEL:   question.MODEL,
			Title:   question.Title,
			Type:    question.Type,
			Options: options,
		})
	}

	finalExam := request.Exam{
		MODEL:     exam.MODEL,
		Title:     exam.Title,
		Questions: questions,
	}

	response.OkWithData(response.ExamResponse{
		Exam: finalExam,
	}, c)
}

// @Tags Exam
// @Summary 提交试卷
// @Security ApiKeyAuth
// @accept application/json
// @Param data body request.ExamAnswer true "试卷答案"
// @Produce application/json
// @Success 200 {object} response.SubmitExamResponse ""
// @Router /exams/submit [post]
func SubmitExam(c *gin.Context) {
	var data request.ExamAnswer
	c.BindJSON(&data)

	var examAnswerQuestions []model.ExamAnswerQuestions
	for _, question := range data.Questions {
		var examAnswerQuestionOptions []model.ExamAnswerQuestionOptions
		for _, option := range question.Options {
			examAnswerQuestionOptions = append(examAnswerQuestionOptions, model.ExamAnswerQuestionOptions{
				QuestionOptionId: option,
			})
		}
		examAnswerQuestions = append(examAnswerQuestions, model.ExamAnswerQuestions{
			QuestionId:      question.ID,
			QuestionOptions: examAnswerQuestionOptions,
		})
	}

	examAnswer := model.ExamAnswers{
		ExamId:    data.ExamId,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
		UserId:    getUserID(c),
		Questions: examAnswerQuestions,
	}
	result := global.DB.Create(&examAnswer)

	if result.Error != nil {
		response.FailWithMessage("提交试卷失败", c)
		return
	}
	response.OkWithData(response.SubmitExamResponse{
		ID: examAnswer.ID,
	}, c)
}

// @Tags Exam
// @Summary 查询试卷答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Param id path string true "答卷id"
// @Produce application/json
// @Success 200 {object} response.ExamResultResponse "response.Response{data: response.ExamResultResponse}"
// @Router /exams/answer/:id [get]
func GetExamResult(c *gin.Context) {
	// 获取用户答卷
	paramId := c.Param("id")
	examAnswerId, err := strconv.Atoi(paramId)

	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var examAnswer model.ExamAnswers

	userId := getUserID(c)
	var examAnswerRes *gorm.DB
	if service.IsAdmin(userId) {
		examAnswerRes = global.DB.Where(&model.ExamAnswers{
			MODEL: global.MODEL{
				ID: uint(examAnswerId),
			},
		}).First(&examAnswer)
	} else {
		examAnswerRes = global.DB.Where(&model.ExamAnswers{
			MODEL: global.MODEL{
				ID: uint(examAnswerId),
			},
			UserId: getUserID(c),
		}).First(&examAnswer)
	}

	if examAnswerRes.Error != nil {
		response.FailWithMessage(examAnswerRes.Error.Error(), c)
		return
	}

	// 获取试卷信息
	var exam model.Exams

	result := global.DB.First(&exam, examAnswer.ExamId)

	if result.Error != nil {
		response.FailWithMessage("get exam error", c)
		return
	}

	global.DB.Preload("Questions.Options").Preload("Questions.Tags").Preload("Questions").Find(&exam)

	// 试卷信息 + 用户答卷 => 批改后试卷
	var examResult response.ExamResult

	examResult.ID = uint(examAnswerId)
	examResult.Title = exam.Title
	examResult.TotalScore = 0
	examResult.Score = 0
	examResult.StartTime = examAnswer.StartTime
	examResult.EndTime = examAnswer.EndTime

	exercises := make([][]int, 2) // 答题情况

	for _, question := range exam.Questions {

		isRight := true
		var options []model.QuestionOptionResult

		var examAnswerQuestion model.ExamAnswerQuestions

		global.DB.Where(&model.ExamAnswerQuestions{
			ExamAnswerId: examAnswer.ID,
			QuestionId:   question.ID,
		}).Preload("QuestionOptions").First(&examAnswerQuestion)

		for _, option := range question.Options {
			isChoosed := false

			for _, userOption := range examAnswerQuestion.QuestionOptions {
				if option.ID == userOption.QuestionOptionId {
					isChoosed = true
				}
			}

			if (isChoosed && !option.IsRight) || (!isChoosed && option.IsRight) {
				isRight = false
			}

			options = append(options, model.QuestionOptionResult{
				Title: option.Title,
				// IsRight:   option.IsRight,
				IsChoosed: isChoosed,
			})
		}

		examResult.TotalScore++
		if isRight {
			examResult.Score++
		} else {
			// TODO: write in wrongbook
		}

		examResult.Questions = append(examResult.Questions, model.QuestionResult{
			Title: question.Title,
			// Analysis: question.Analysis,
			Type: question.Type,
			// IsRight:  isRight,
			Options: options,
		})

		exercises[0] = append(exercises[0], int(question.Tags[0].ID))
		if isRight {
			exercises[1] = append(exercises[1], 1)
		} else {
			exercises[1] = append(exercises[1], 0)
		}
	}

	questionIDs, err := service.GetKTRecommendations(exercises)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("获取算法推荐题目出错", c)
	}

	response.OkWithData(response.ExamResultResponse{
		Result: examResult,
		Recommendations: response.ExamRecommendations{
			QuestionIds: questionIDs,
		},
	}, c)
}
