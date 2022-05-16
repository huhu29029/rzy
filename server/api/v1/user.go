package v1

import (
	"online-exam/global"
	"online-exam/middleware"
	"online-exam/model"
	"online-exam/model/request"
	response "online-exam/model/response"
	"online-exam/service"
	"online-exam/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	if err := utils.Verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	U := &model.SysUser{Username: L.Username, Password: L.Password}
	if err, user := service.Login(U); err != nil {
		global.LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		// xapi: 用户登录
		service.NewXAPIData(model.XAPIData{
			Actor: model.XActor{
				UUID: (*user).UUID.String(),
			},
			Verb: "login",
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

		token, err := GenToken(*user)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		}
		response.OkWithDetailed(response.TokenResponse{
			Token: token,
		}, "登录成功", c)
	}
}

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	var R request.Register
	_ = c.ShouldBindJSON(&R)
	if err := utils.Verify(R, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := model.SysUser{Username: R.Username, NickName: R.Username, Password: R.Password, AuthorityId: 2}
	err, userReturn := service.Register(user)
	if err != nil {
		global.LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		token, err := GenToken(userReturn)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		}
		response.OkWithDetailed(response.TokenResponse{
			Token: token,
		}, "登录成功", c)
	}
}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func ChangePassword(c *gin.Context) {
	var user request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	U := &model.SysUser{Username: user.Username, Password: user.Password}
	if err, _ := service.ChangePassword(U, user.NewPassword); err != nil {
		global.LOG.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /user/setUserInfo [put]
func SetUserInfo(c *gin.Context) {
	var user model.SysUser
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, ReqUser := service.SetUserInfo(user); err != nil {
		global.LOG.Error("设置失败", zap.Any("err", err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
	}
}

// @Tags SysUser
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "response.Response{data: response.SysUserResponse}"
// @Router /user [get]
func GetUserInfo(c *gin.Context) {
	err, user := service.FindUserById(int(getUserID(c)))
	if err != nil {
		response.FailWithMessage("获取用户信息失败", c)
		return
	}

	response.OkWithData(response.SysUserResponse{
		User: *user,
	}, c)
}

// @Tags SysUser
// @Summary 查询用户考试历史
// @Security ApiKeyAuth
// @Param page query string false "所在页面 默认值：1"
// @Param size query string false "分页大小 默认值：10"
// @Produce application/json
// @Success 200 {object} response.GetExamHistoryResponse "model.Response{data: response.GetExamHistoryResponse}"
// @Router /user/exams [get]
func GetExamHistory(c *gin.Context) {
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
	global.DB.Model(&model.ExamAnswers{}).Where("user_id = ?", getUserID(c)).Count(&count)
	examHistory := service.GetUserExamHistory(getUserID(c), limit, offset)

	response.OkWithData(response.GetExamHistoryResponse{
		Total:       uint(count),
		ExamHistory: examHistory,
	}, c)
}

// @Tags SysUser
// @Summary 查询用户做题历史
// @Security ApiKeyAuth
// @Param page query string false "所在页面 默认值：1"
// @Param size query string false "分页大小 默认值：10"
// @Produce application/json
// @Success 200 {object} response.GetQuestionHistoryResponse "model.Response{data: response.GetQuestionHistoryResponse}"
// @Router /user/questions [get]
func GetQuestionHistory(c *gin.Context) {
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
	global.DB.Model(&model.Answers{}).Where("user_id = ?", getUserID(c)).Count(&count)

	questionHistory := service.GetUserQuestionHistory(getUserID(c), limit, offset)

	response.OkWithData(response.GetQuestionHistoryResponse{
		Total:           uint(count),
		QuestionHistory: questionHistory,
	}, c)
}

// @Tags SysUser
// @Summary 查询用户错题历史
// @Security ApiKeyAuth
// @Param page query string false "所在页面 默认值：1"
// @Param size query string false "分页大小 默认值：10"
// @Produce application/json
// @Success 200 {object} response.GetQuestionHistoryResponse "model.Response{data: response.GetQuestionHistoryResponse}"
// @Router /user/WrongQuestion [get]
func GetWrongQuestionHistory(c *gin.Context) {
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
	global.DB.Model(&model.Answers{}).Where("user_id = ?", getUserID(c)).Count(&count)
	count = service.NewCount(getUserID(c))
	questionHistory := service.GetWrongQuestionHistory(getUserID(c), limit, offset)
	response.OkWithData(response.GetQuestionHistoryResponse{
		Total:           uint(count),
		QuestionHistory: questionHistory,
	}, c)
}

//@Tags SysUser
//@Summary 删除错题
//@Security ApiKeyAuth
//@Param id path string true "问题id"
//@Produce application/json
//@Success 200 {object} response.DelQuestionFromQuestionBook "response.Response{data: response.DelQuestionFromQuestionBook}"
//@Router /user/DeleteQuestionFromWrongQuestionBook [delete]
func DeleteQuestionFromWrongQuestionBook(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	WrongQustion := model.WrongQuestionBook{
		MODEL: global.MODEL{
			ID: uint(id),
		},
	}
	result := global.DB.Unscoped().Delete(&WrongQustion)
	global.DB.Model(&WrongQustion).Association("tags").Clear()
	if result.Error != nil {
		response.FailWithMessage("del question error", c)
		return
	}
	response.OkWithData(response.DelQuestionFromQuestionBook{
		WrongQuestionBook: struct{}{},
	},
		c)
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func getUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件")
		return 0
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.ID
	}
}

// 从Gin的Context中获取从jwt解析出来的用户UUID
func getUserUuid(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UUID.String()
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func getUserAuthorityId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		global.LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件")
		return 0
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.AuthorityId
	}
}

// @description: 根据用户信息生成token
func GenToken(user model.SysUser) (string, error) {
	j := &middleware.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		BufferTime:  global.CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    global.CONFIG.JWT.SigningKey,                      // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败", zap.Any("err", err))
	}

	return token, err
}
