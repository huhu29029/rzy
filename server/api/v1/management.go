package v1

import (
	"online-exam/global"
	"online-exam/model"
	"online-exam/model/request"
	"online-exam/model/response"
	"online-exam/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Tags Admin
// @Summary manage users
// @Security ApiKeyAuth
// @Param page query string false "所在页面 默认值：1"
// @Param size query string false "分页大小 默认值：10"
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /admin/users [get]
func GetUsers(c *gin.Context) {
	userId := getUserID(c)
	if !service.IsAdmin(userId) {
		response.FailWithMessage("用户权限不足", c)
		return
	}

	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")

	limit, err := strconv.Atoi(size)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	offset, err := strconv.Atoi(page)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	offset = (offset - 1) * limit

	var count int64
	var users []model.SysUser
	global.DB.Model(&model.SysUser{}).Count(&count)
	global.DB.Limit(limit).Offset(offset).Find(&users)

	response.OkWithData(response.ManagementUsersResp{
		Total: uint(count),
		Users: users,
	}, c)
}

// @Tags Admin
// @Summary user question history
// @Security ApiKeyAuth
// @Produce  application/json
// @Param id path string true "user id"
// @Param page query string false "所在页面 默认值：1"
// @Param size query string false "分页大小 默认值：10"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /admin/users/:id/questions [get]
func GetUserQuestionHistory(c *gin.Context) {
	if !service.IsAdmin(getUserID(c)) {
		response.FailWithMessage("用户权限不足", c)
		return
	}

	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	userIdStr := c.Param("id")
	userId, _ := strconv.Atoi(userIdStr)

	limit, err := strconv.Atoi(size)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	offset, err := strconv.Atoi(page)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	offset = (offset - 1) * limit

	var count int64
	global.DB.Model(&model.Answers{}).Where("user_id = ?", userId).Count(&count)

	questionHistory := service.GetUserQuestionHistory(uint(userId), limit, offset)

	response.OkWithData(response.GetQuestionHistoryResponse{
		Total:           uint(count),
		QuestionHistory: questionHistory,
	}, c)
}

// @Tags Admin
// @Summary user exam history
// @Security ApiKeyAuth
// @Param id path string true "user id"
// @Param page query string false "所在页面 默认值：1"
// @Param size query string false "分页大小 默认值：10"
// @Produce application/json
// @Success 200 {object} response.GetExamHistoryResponse "model.Response{data: response.GetExamHistoryResponse}"
// @Router /admin/users/:id/exams [get]
func GetUserExamHistory(c *gin.Context) {
	if !service.IsAdmin(getUserID(c)) {
		response.FailWithMessage("用户权限不足", c)
		return
	}

	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	userIdStr := c.Param("id")
	userId, _ := strconv.Atoi(userIdStr)

	limit, err := strconv.Atoi(size)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	offset, err := strconv.Atoi(page)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	offset = (offset - 1) * limit

	var count int64
	global.DB.Model(&model.ExamAnswers{}).Where("user_id = ?", userId).Count(&count)
	examHistory := service.GetUserExamHistory(uint(userId), limit, offset)

	response.OkWithData(response.GetExamHistoryResponse{
		Total:       uint(count),
		ExamHistory: examHistory,
	}, c)
}

// @Tags Admin
// @Summary 查询所有题目
// @Security ApiKeyAuth
// @Param page query string false "所在页面 默认值：1"
// @Param size query string false "分页大小 默认值：10"
// @Produce application/json
// @Success 200 {object} response.GetQuestionsResponse "response.Response{data: response.GetQuestionsResponse}"
// @Router /admin/questions [get]
func GetQuestions(c *gin.Context) {
	userId := getUserID(c)
	if !service.IsAdmin(userId) && !service.IsPreacher(userId) {
		response.FailWithMessage("用户权限不足", c)
		return
	}

	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")

	limit, err := strconv.Atoi(size)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	offset, err := strconv.Atoi(page)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	offset = (offset - 1) * limit

	var questions []model.Questions
	var count int64
	if service.IsAdmin(userId) {
		global.DB.Limit(limit).Offset(offset).Preload("Options").Preload("Tags").Find(&questions)
		global.DB.Model(&model.Questions{}).Count(&count)
	} else if service.IsPreacher(userId) {
		global.DB.Limit(limit).Offset(offset).Preload("Options").Preload("Tags").Where("user_id = ?", userId).Find(&questions)
		global.DB.Model(&model.Questions{}).Where("user_id = ?", userId).Count(&count)
	}

	response.OkWithData(response.GetQuestionsResponse{
		Total:     uint(count),
		Questions: questions,
	}, c)
}

// @Tags Admin
// @Summary 添加新问题
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.AddQuestionRequest true "问题"
// @Success 200 {object} response.AddQuestionResponse true "response.Response{data: response.AddQuestionResponse}"
// @Router /admin/questions [post]
func AddQuestion(c *gin.Context) {
	userId := getUserID(c)
	if !service.IsAdmin(userId) && !service.IsPreacher(userId) {
		response.FailWithMessage("用户权限不足", c)
		return
	}

	var question request.AddQuestionRequest
	err := c.ShouldBindJSON(&question)
	if err != nil {
		response.FailWithMessage("parameters error", c)
		return
	}

	var options []model.QuestionOptions
	for _, option := range question.Options {
		options = append(options, model.QuestionOptions{
			Title:   option.Title,
			IsRight: option.IsRight,
		})
	}

	var tags []model.Tags
	if len(question.Tags) != 0 {
		global.DB.Find(&tags, question.Tags)
	}

	createdQuestion := model.Questions{
		MODEL: global.MODEL{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:     question.Title,
		Analysis:  question.Analysis,
		Type:      question.Type,
		ChapterId: question.ChapterId,
		Options:   options,
		Tags:      tags,
		UserId:    userId,
	}

	result := global.DB.Create(&createdQuestion)

	if result.Error != nil {
		response.FailWithMessage("create question error", c)
		return
	}

	response.OkWithData(response.AddQuestionResponse{
		Question: createdQuestion,
	}, c)
}

// @Tags Admin
// @Summary 删除问题
// @Security ApiKeyAuth
// @Param id path string true "问题id"
// @Produce application/json
// @Success 200 {object} response.DelQuestionResponse "response.Response{data: response.DelQuestionResponse}"
// @Router /admin/questions/:id [delete]
func DelQuestion(c *gin.Context) {
	userId := getUserID(c)
	if !service.IsAdmin(userId) && !service.IsPreacher(userId) {
		response.FailWithMessage("用户权限不足", c)
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if service.IsPreacher(userId) {
		var count int64
		global.DB.Model(&model.Questions{}).Where(&model.Questions{
			MODEL: global.MODEL{
				ID: uint(id),
			},
			UserId: userId,
		}).Count(&count)

		if count == 0 {
			response.FailWithMessage("用户权限不足", c)
			return
		}
	}

	question := model.Questions{
		MODEL: global.MODEL{
			ID: uint(id),
		},
	}

	result := global.DB.Unscoped().Delete(&question)
	global.DB.Model(&question).Association("Tags").Clear()

	if result.Error != nil {
		response.FailWithMessage("del question error", c)
		return
	}

	response.OkWithData(response.DelQuestionResponse{
		Question: struct{}{},
	}, c)
}

// @Tags Admin
// @Summary 编辑问题
// @Security ApiKeyAuth
// @Param id path string true "问题id"
// @Param data body request.EditQuestionRequest true "问题"
// @Produce application/json
// @Success 200 {object} response.AddQuestionResponse "response.Response{data: response.AddQuestionResponse}"
// @Router /admin/questions/:id [put]
func EditQuestion(c *gin.Context) {
	userId := getUserID(c)
	if !service.IsAdmin(userId) && !service.IsPreacher(userId) {
		response.FailWithMessage("用户权限不足", c)
		return
	}

	questionIdStr := c.Param("id")

	questionId, err := strconv.Atoi(questionIdStr)
	if err != nil {
		response.FailWithMessage("parameters error", c)
		return
	}

	if service.IsPreacher(userId) {
		var count int64
		global.DB.Model(&model.Questions{}).Where(&model.Questions{
			MODEL: global.MODEL{
				ID: uint(questionId),
			},
			UserId: userId,
		}).Count(&count)

		if count == 0 {
			response.FailWithMessage("用户权限不足", c)
			return
		}
	}

	var question request.EditQuestionRequest
	c.ShouldBindJSON(&question)

	// 删除 Options
	var options []model.QuestionOptions
	for _, option := range question.Options {
		options = append(options, model.QuestionOptions{
			MODEL: global.MODEL{
				ID:        option.ID,
				UpdatedAt: time.Now(),
			},
			Title:      option.Title,
			IsRight:    option.IsRight,
			QuestionId: option.QuestionId,
		})
	}

	var prevOptions []model.QuestionOptions

	global.DB.Where("question_id = ?", questionId).Find(&prevOptions)

	for _, prevOption := range prevOptions {
		delFlag := true

		for _, option := range question.Options {
			if prevOption.ID == option.ID {
				delFlag = false
			}
		}

		if delFlag {
			global.DB.Delete(&model.QuestionOptions{}, prevOption.ID)
		}
	}

	global.DB.Model(&model.Questions{
		MODEL: global.MODEL{
			ID: uint(questionId),
		},
	}).Association("Tags").Clear()

	var tags []model.Tags
	global.DB.Find(&tags, question.Tags)

	// options 单独更新
	for _, option := range options {
		global.DB.Updates(&option)
	}

	updatedQuestion := model.Questions{
		MODEL: global.MODEL{
			ID:        uint(questionId),
			UpdatedAt: time.Now(),
		},
		Title:     question.Title,
		Analysis:  question.Analysis,
		Type:      question.Type,
		ChapterId: question.ChapterId,
		Options:   options,
		Tags:      tags,
		UserId:    userId,
	}

	result := global.DB.Updates(&updatedQuestion)

	if result.Error != nil {
		response.FailWithMessage("update question error", c)
		return
	}

	response.OkWithData(response.AddQuestionResponse{
		Question: updatedQuestion,
	}, c)
}

// @Tags Admin
// @Summary 查询管理员所属试卷
// @Security ApiKeyAuth
// @Param page query string false "所在页面 默认值：1"
// @Param size query string false "分页大小 默认值：10"
// @Produce application/json
// @Success 200 {object} response.GetExamsResponse "response.Response{data: response.GetExamsResponse}"
// @Router /admin/exams [get]
func ManageExams(c *gin.Context) {
	userId := getUserID(c)
	if !service.IsAdmin(userId) && !service.IsPreacher(userId) {
		response.FailWithMessage("用户权限不足", c)
		return
	}

	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")

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
	if service.IsAdmin(userId) {
		global.DB.Limit(limit).Offset(offset).Find(&exams)
		global.DB.Model(&model.Exams{}).Count(&count)
	} else if service.IsPreacher(userId) {
		global.DB.Limit(limit).Offset(offset).Where("user_id = ?", userId).Find(&exams)
		global.DB.Model(&model.Exams{}).Where("user_id = ?", userId).Count(&count)
	}

	response.OkWithData(response.GetExamsResponse{
		Total: uint(count),
		Exams: exams,
	}, c)
}

// @Tags Admin
// @Summary 创建试卷
// @Security ApiKeyAuth
// @Param data body request.ExamReq true "题目id列表"
// @Produce application/json
// @Success 200 {object} response.ExamResponse "response.Response{data: response.ExamResponse}"
// @Router /admin/exams/exam [post]
func AddExam(c *gin.Context) {
	userId := getUserID(c)
	if !service.IsAdmin(userId) && !service.IsPreacher(userId) {
		response.FailWithMessage("用户权限不足", c)
		return
	}

	var Req request.ExamReq
	c.BindJSON(&Req)

	var questions []model.Questions
	for _, id := range Req.QuestionIDs {
		questions = append(questions, model.Questions{
			MODEL: global.MODEL{
				ID: id,
			},
		})
	}
	exam := model.Exams{
		Title:     Req.Title,
		Questions: questions,
		UserId:    userId,
	}

	global.DB.Create(&exam)

	global.DB.First(&exam, exam.ID)
	global.DB.Model(&exam).Preload("Options").Association("Questions").Find(&exam.Questions)

	var finalQuestions []request.Question
	for _, question := range exam.Questions {
		var options []request.QuestionOption
		for _, option := range question.Options {
			options = append(options, request.QuestionOption{
				MODEL: option.MODEL,
				Title: option.Title,
			})
		}
		finalQuestions = append(finalQuestions, request.Question{
			MODEL:   question.MODEL,
			Title:   question.Title,
			Type:    question.Type,
			Options: options,
		})
	}

	finalExam := request.Exam{
		MODEL:     exam.MODEL,
		Title:     exam.Title,
		Questions: finalQuestions,
	}

	response.OkWithData(response.ExamResponse{
		Exam: finalExam,
	}, c)
}

// @Tags Admin
// @Summary 删除试卷
// @Security ApiKeyAuth
// @Param id path string true "试卷id"
// @Produce application/json
// @Success 200 {object} response.ExamResponse "response.Response{data: response.ExamResponse}"
// @Router /admin/exams/exam/:id [delete]
func DelExam(c *gin.Context) {
	userId := getUserID(c)
	if !service.IsAdmin(userId) && !service.IsPreacher(userId) {
		response.FailWithMessage("用户权限不足", c)
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	exam := model.Exams{
		MODEL: global.MODEL{
			ID: uint(id),
		},
	}

	if service.IsPreacher(userId) {
		var count int64
		global.DB.Model(&model.Exams{}).Where(&model.Exams{
			MODEL: global.MODEL{
				ID: uint(id),
			},
			UserId: userId,
		}).Count(&count)

		if count == 0 {
			response.FailWithMessage("用户权限不足", c)
			return
		}
	}

	result := global.DB.Unscoped().Delete(&exam)

	if result.Error != nil {
		response.FailWithMessage("del exam error", c)
		return
	}

	response.OkWithData(response.ExamResponse{
		Exam: request.Exam{},
	}, c)
}

// @Tags Admin
// @Summary 编辑试卷
// @Security ApiKeyAuth
// @Param id path string true "试卷id"
// @Param data body request.ExamReq true "试卷id"
// @Produce application/json
// @Success 200 {object} response.ExamResponse "response.Response{data: response.ExamResponse}"
// @Router /admin/exams/exam/:id [put]
func EditExam(c *gin.Context) {
	userId := getUserID(c)
	if !service.IsAdmin(userId) && !service.IsPreacher(userId) {
		response.FailWithMessage("用户权限不足", c)
		return
	}

	id := c.Param("id")
	var Req request.ExamReq
	c.BindJSON(&Req)

	examId, err := strconv.Atoi(id)
	if err != nil {
		response.FailWithMessage("parameters error", c)
		return
	}

	if service.IsPreacher(userId) {
		var count int64
		global.DB.Model(&model.Exams{}).Where(&model.Exams{
			MODEL: global.MODEL{
				ID: uint(examId),
			},
			UserId: userId,
		}).Count(&count)

		if count == 0 {
			response.FailWithMessage("用户权限不足", c)
			return
		}
	}

	var prevQuestions []model.ExamsQuestions
	global.DB.Where("exams_id = ?", examId).Find(&prevQuestions)
	for _, prevQuestion := range prevQuestions {
		delFlag := true

		for _, questionId := range Req.QuestionIDs {
			if questionId == prevQuestion.QuestionsId {
				delFlag = false
				break
			}
		}

		if delFlag {
			global.DB.Where("questions_id = ? AND exams_id = ?", prevQuestion.QuestionsId, id).Delete(&prevQuestion)
		}
	}

	var questions []model.Questions
	for _, id := range Req.QuestionIDs {
		questions = append(questions, model.Questions{
			MODEL: global.MODEL{
				ID: id,
			},
		})
	}

	updatedExam := model.Exams{
		MODEL: global.MODEL{
			ID: uint(examId),
		},
		Title:     Req.Title,
		Questions: questions,
	}
	c.ShouldBindJSON(&Req)

	global.DB.Updates(&updatedExam)

	var exam model.Exams
	global.DB.First(&exam, updatedExam.ID)
	global.DB.Model(&exam).Preload("Options").Association("Questions").Find(&exam.Questions)

	var finalQuestions []request.Question
	for _, question := range exam.Questions {
		var options []request.QuestionOption
		for _, option := range question.Options {
			options = append(options, request.QuestionOption{
				MODEL: option.MODEL,
				Title: option.Title,
			})
		}
		finalQuestions = append(finalQuestions, request.Question{
			MODEL:   question.MODEL,
			Title:   question.Title,
			Type:    question.Type,
			Options: options,
		})
	}

	finalExam := request.Exam{
		MODEL:     exam.MODEL,
		Title:     exam.Title,
		Questions: finalQuestions,
	}

	response.OkWithData(response.ExamResponse{
		Exam: finalExam,
	}, c)
}
