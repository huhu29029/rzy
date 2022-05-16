package v1

import (
	"online-exam/global"
	"online-exam/model"
	"online-exam/model/request"
	"online-exam/model/response"
	"online-exam/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// xAPI 学习行为数据规范

// 行为数据内容 statement = actor+verb+object(+result+context+authority+timestamp)

// actor
// 	uuid 用户唯一标识符
// verb
// 	login 登录
// 	navigate 浏览
// 	practice 做题
// 	examine 考试
// 	logout 登出
// object
// 	module 功能模块
// 	question 习题
// 	wrongQuestion 错题
// result
// context
// 	ip ip地址
// 	device 设备信息
// authority
// timestamp 时间戳

// @Tags xAPI
// @Summary 学习行为数据
// @Security ApiKeyAuth
// @Param data body model.XAPIData true "学习行为数据"
// @Produce application/json
// @Success 200 {object} response.Response "{"code": 0, "data": {}, "msg": "success"}"
// @Router /xapi/new [post]
func NewXAPIData(c *gin.Context) {
	var data request.XAPIData
	err := c.BindJSON(&data)
	if err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.LOG.Error("请求参数错误", zap.Error(err))
		c.Abort()
		return
	}

	if err, user := service.FindUserById(int(getUserID(c))); err == nil {
		xapiData := model.XAPIData{
			Actor: model.XActor{
				UUID: user.UUID.String(),
			},
			Verb:   data.Verb,
			Object: data.Object,
			Result: struct{}{},
			Context: model.XContext{
				IP:     c.ClientIP(),
				Device: c.GetHeader("User-Agent"),
			},
			Authority: struct{}{},
			Timestamp: time.Now(),
		}
		service.NewXAPIData(xapiData)
	} else {
		response.FailWithMessage("获取用户信息失败", c)
		global.LOG.Error("获取用户信息失败", zap.Error(err))
		c.Abort()
		return
	}

	response.OkWithMessage("success", c)
}
