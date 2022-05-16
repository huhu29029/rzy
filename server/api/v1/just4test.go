package v1

import (
	"online-exam/global"
	"online-exam/model/response"

	"github.com/gin-gonic/gin"
)

// @Tags Test
// @Summary 测试专用
// @Produce application/json
// @Router /test/just4test [get]
// @Success 200 {object} response.Response "{"code": 0, "data": {}, "msg": "just4test"}"
func Just4test(c *gin.Context) {
	sugar := global.LOG.Sugar()

	sugar.Info(c.Request.Header)
	response.OkWithMessage("just4test", c)
}

// @Tags Home
// @Summary HelloWorld
// @Produce application/json
// @Router /test/just4test [get]
// @Success 200 {object} response.Response "{"code": 0, "data": {}, "msg": "just4test"}"
func HelloWorld(c *gin.Context) {
	response.OkWithMessage("helloworld, I'm online-exam sever", c)
}
