package v1

import (
	"online-exam/global"
	"online-exam/model"
	"online-exam/model/response"
	"online-exam/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func JsonInBlacklist(c *gin.Context) {

	err, user := service.FindUserById(int(getUserID(c)))
	if err != nil {
		response.FailWithMessage("获取用户信息失败", c)
		return
	} else {
		// xapi: 用户登出
		service.NewXAPIData(model.XAPIData{
			Actor: model.XActor{
				UUID: (*user).UUID.String(),
			},
			Verb: "logout",
			Object: model.XObject{
				Info: "System",
			},
			Result: struct{}{},
			Context: model.XContext{
				IP:     c.ClientIP(),
				Device: c.GetHeader("User-Agent"),
			},
			Authority: struct{}{},
			Timestamp: time.Now(),
		})
	}

	token := c.Request.Header.Get("x-token")
	jwt := model.JwtBlacklist{Jwt: token}
	if err := service.JsonInBlacklist(jwt); err != nil {
		global.LOG.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
