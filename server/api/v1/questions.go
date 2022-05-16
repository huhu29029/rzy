package v1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"online-exam/global"
	"online-exam/model"
	"online-exam/model/request"
	"online-exam/model/response"
	"online-exam/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 算法推荐题目，根据用户最近练习历史
func GetQuestionWithAlgorithm(c *gin.Context) {
	ALG_LEN := 50 // 算法所需的用户最大练习长度
	uuid := getUserUuid(c)
	uid := getUserID(c)
	REDIS_KEY := uuid + "_" + "kt"

	userLen, err := global.REDIS.LLen(REDIS_KEY).Result()
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("获取算法推荐题目出错", c)
		return
	}

	// 获取用户练习历史
	// [[q1, q2, ...], [a1, a2, ...]]
	exercises := make([][]int, 2)

	if userLen == 0 {
		var answers []model.Answers
		fmt.Println("uid: ", uid)
		global.DB.Select("id", "question_id").Where("user_id = ?", uid).Order("id desc").Limit(ALG_LEN).Find(&answers)
		for _, answer := range answers {
			result := service.GetQuestionResult(int(answer.ID))
			var aQuestion model.Questions
			global.DB.Where("id = ?", answer.QuestionId).Preload("Options").Preload("Tags").First(&aQuestion)
			exercises[0] = append(exercises[0], int(aQuestion.Tags[0].ID))
			if result.IsRight {
				exercises[1] = append(exercises[1], 1)
			} else {
				exercises[1] = append(exercises[1], 0)
			}
		}

		// 获取算法推荐内容
		questionIDs, err := service.GetKTRecommendations(exercises)
		if err != nil {
			global.LOG.Error(err.Error())
			response.FailWithMessage("获取算法推荐题目出错", c)
			return
		}

		for _, id := range questionIDs {
			err = global.REDIS.RPush(REDIS_KEY, id).Err()
		}
	}

	questionID, err := global.REDIS.LPop(REDIS_KEY).Result()

	var question model.Questions
	global.DB.Preload("Options").Preload("Tags").First(&question, questionID)

	response.OkWithData(response.RandomQuestionResponse{
		Question: question,
	}, c)
}

// 错题重练
func GetWrongQuetion(c *gin.Context) {
	uuid := getUserUuid(c)
	uid := getUserID(c)
	REDIS_KEY := uuid + "_" + "wrong"
	WINDOW_SIZE := 10

	sugar := global.LOG.Sugar()

	userLen, err := global.REDIS.LLen(REDIS_KEY).Result()
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("获取错题出错", c)
		return
	}

	if userLen == 0 {
		var answers []model.Answers
		// 选取不同题目最新的提交记录 -> 选择做错的
		err := global.DB.Raw("select * from exam.answers where id in (SELECT max(id) as id FROM exam.answers where user_id = ? group by question_id) and is_right = 0 limit ?", uid, WINDOW_SIZE).Scan(&answers).Error
		if err != nil {
			sugar.Error(err)
			response.Fail(c)
			return
		}
		if len(answers) == 0 {
			response.Result(1000, struct{}{}, "暂无错题", c)
			return
		}
		for _, answer := range answers {
			sugar.Info(answer.ID, answer.QuestionId, answer.IsRight, uid, WINDOW_SIZE)
			err = global.REDIS.RPush(REDIS_KEY, answer.QuestionId).Err()
			if err != nil {
				sugar.Info(err)
				response.Fail(c)
				return
			}
		}
	}

	questionID, err := global.REDIS.LPop(REDIS_KEY).Result()

	var question model.Questions
	global.DB.Preload("Options").Preload("Tags").First(&question, questionID)

	response.OkWithData(response.RandomQuestionResponse{
		Question: question,
	}, c)
}

// 随机推荐问题
func RandomQuestion(c *gin.Context) {
	QUESTION_WINDOW := 10

	var question model.Questions
	uuid := getUserUuid(c)

	userLen, err := global.REDIS.LLen(uuid).Result()
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("获取随机题目出错", c)
	}
	if userLen == 0 {
		var questions []model.Questions
		global.DB.Select("id").Order("rand()").Limit(QUESTION_WINDOW).Find(&questions)
		for _, question := range questions {
			err = global.REDIS.RPush(uuid, question.ID).Err()
		}
	}

	questionID, err := global.REDIS.LPop(uuid).Result()
	global.DB.Preload("Options").Preload("Tags").First(&question, questionID)

	response.OkWithData(response.RandomQuestionResponse{
		Question: question,
	}, c)
}

// @Tags Questions
// @Summary 根据章节随机推荐题目
// @Security ApiKeyAuth
// @Param data body request.GetQuestionByChaptersRequest true "章节id列表"
// @Produce application/json
// @Success 200 {object} response.RandomQuestionResponse "response.Response{data: response.RandomQuestionResponse}"
// @Router /questions/random [post]
func GetQuetionsByChapters(c *gin.Context) {
	var data request.GetQuestionByChaptersRequest
	c.BindJSON(&data)

	var question model.Questions
	result := global.DB.Where("chapter_id IN ?", data.ChapterIds).Preload("Options").Preload("Tags").Order("rand()").First(&question)

	if result.Error != nil {
		response.FailWithMessage(result.Error.Error(), c)
		return
	}
	response.OkWithData(response.RandomQuestionResponse{
		Question: question,
	}, c)
}

// @Tags Niepan
// @Summary 根据知识点获取题目 - 用于涅槃
// @Param point query string true "对应涅槃知识点，使用 , 分割。示例：JavaScript,js简介,简介,JS有什么作用"
// @Param page query string false "所在页面 默认值：1"
// @Param size query string false "分页大小 默认值：1"
// @Produce application/json
// @Success 200 {object} response.GetQuestions2NiepanResponse "response.Response{data: response.GetQuestions2NiepanResponse}"
// @Router /niepan/questions [get]
func GetQuestions2Niepan(c *gin.Context) {
	// 参数解析
	pointStr := c.Query("point")

	points := strings.Split(pointStr, ",")

	if len(points) != 4 {
		response.FailWithMessage("知识点层级错误", c)
		return
	}

	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "1")

	limit, err := strconv.Atoi(size)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	offset, err := strconv.Atoi(page)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	offset = (offset - 1) * limit

	// 数据查询
	tag := model.Tags{
		Layer1: points[0],
		Layer2: points[1],
		Layer3: points[2],
		Layer4: points[3],
	}
	result := global.DB.Where(&tag).First(&tag)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage("标签不存在", c)
		return
	}

	var total int64
	global.DB.Where("tags_id = ?", tag.ID).Model(&model.QuestionsTags{}).Count(&total)
	var questionsTags []model.QuestionsTags
	global.DB.Where("tags_id = ?", tag.ID).Limit(limit).Offset(offset).Find(&questionsTags)

	var questions []model.Questions
	for _, el := range questionsTags {
		var question model.Questions
		global.DB.Preload("Options").Preload("Tags").First(&question, el.QuestionsId)
		questions = append(questions, question)
	}

	response.OkWithData(response.GetQuestions2NiepanResponse{
		Total:     uint(total),
		Questions: questions,
	}, c)
}

// @Tags Questions
// @Summary 提交某个问题答案
// @Security ApiKeyAuth
// @accept application/json
// @Param data body request.Answer true "答案"
// @Produce application/json
// @Success 200 {object} response.Response ""
// @Router /questions/submit [post]
func SubmitAnswer(c *gin.Context) {
	var data request.Answer
	c.BindJSON(&data)

	uid := getUserID(c)

	sugar := global.LOG.Sugar()
	db := global.DB

	// 记录用户答题原始数据
	isRight := true

	var question model.Questions
	err := db.Preload("Options").First(&question, data.QuestionId).Error
	if err != nil {
		sugar.Error(err)
		response.Fail(c)
		return
	}

	for _, option := range question.Options {
		isChoosed := false

		for _, selectedOptionId := range data.Options {
			if option.ID == selectedOptionId {
				isChoosed = true
				break
			}
		}

		if (isChoosed && !option.IsRight) || (!isChoosed && option.IsRight) {
			isRight = false
			break
		}
	}

	sugar.Info("回答问题", uid, data.QuestionId, isRight)

	tx := db.Begin()

	answer := model.Answers{
		StartTime:  data.StartTime,
		EndTime:    data.EndTime,
		UserId:     uid,
		QuestionId: data.QuestionId,
		IsRight:    isRight,
	}
	err = tx.Create(&answer).Error
	if err != nil {
		sugar.Error(err)
		tx.Rollback()
		response.Fail(c)
	}

	var answerOptions []model.AnswerOptions
	for _, val := range data.Options {
		answerOptions = append(answerOptions, model.AnswerOptions{
			AnswerId:         answer.ID,
			QuestionOptionId: val,
		})
	}
	err = tx.Create(&answerOptions).Error
	if err != nil {
		sugar.Error(err)
		tx.Rollback()
		response.Fail(c)
		return
	}

	err = tx.Commit().Error
	if err != nil {
		sugar.Error(err)
		response.Fail(c)
		return
	}

	response.Ok(c)
}

// @Tags Questions
// @Summary 查询某个问题
// @Security ApiKeyAuth
// @Param id path string true "问题id"
// @Produce application/json
// @Success 200 {object} response.AddQuestionResponse "response.Response{data: response.AddQuestionResponse}"
// @Router /questions/question/:id [get]
func GetQuestion(c *gin.Context) {
	id := c.Param("id")

	var question model.Questions
	result := global.DB.Preload("Options").Preload("Tags").First(&question, id)

	if result.Error != nil {
		response.FailWithMessage("get question error", c)
		return
	}

	response.OkWithData(response.AddQuestionResponse{
		Question: question,
	}, c)
}

// @Tags Questions
// @Summary 查询问题答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Param id path string true "用户答题id"
// @Produce application/json
// @Success 200 {object} response.QuestionResultResponse "response.Response{data: response.QuestionResultResponse}"
// @Router /questions/answer/:id [get]
func GetQuestionResult(c *gin.Context) {
	paramId := c.Param("id")
	answerId, err := strconv.Atoi(paramId)

	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	result := service.GetQuestionResult(answerId)

	response.OkWithData(response.QuestionResultResponse{
		Result: result,
	}, c)
}
