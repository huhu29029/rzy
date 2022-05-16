package v1

import (
	"net/http"
	"online-exam/global"
	"online-exam/model"
	"online-exam/service"
	"online-exam/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	FRONTEND_URL = "/#/oauth/ourspark"
)

// @Tags OAuth
// @Summary OAuth 登录
// @Param code query string true "code"
// @Router /oauth/ourspark [get]
func Ourspark(c *gin.Context) {
	code := c.Query("code")

	oid, err := utils.GetOursparkOID(code)
	if err != nil {
		global.LOG.Error(err.Error())
		c.Redirect(http.StatusMovedPermanently, FRONTEND_URL)
		return
	}

	detail, err := utils.GetOursaprkDetail(oid)
	if err != nil {
		global.LOG.Error(err.Error())
		c.Redirect(http.StatusMovedPermanently, FRONTEND_URL)
		return
	}

	global.LOG.Info("oauth user detail", zap.Any("user detail", detail))

	user, err := findOrCreateSysUser(oid, detail)

	if err != nil {
		global.LOG.Error(err.Error())
		c.Redirect(http.StatusMovedPermanently, FRONTEND_URL)
		return
	}

	// xapi: 用户登录
	service.NewXAPIData(model.XAPIData{
		Actor: model.XActor{
			UUID: user.UUID.String(),
		},
		Verb: "login-oauth-ourspark",
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

	token, err := GenToken(user)

	if err != nil {
		global.LOG.Error(err.Error())
		c.Redirect(http.StatusMovedPermanently, FRONTEND_URL)
		return
	}
	global.LOG.Info(token)

	c.Redirect(http.StatusMovedPermanently, FRONTEND_URL+"?key="+token)
}

func findOrCreateSysUser(oid string, oursparkUser utils.OursparkDetail) (model.SysUser, error) {
	var oauthOursparkUser model.OauthOursparkUsers
	result := global.DB.Where(&model.OauthOursparkUsers{
		Oid: oid,
	}).First(&oauthOursparkUser)

	if result.Error == gorm.ErrRecordNotFound {
		var err error
		var sysUser model.SysUser
		if oauthOursparkUser.Role == 5 { // 校外教师
			err, sysUser = service.Register(model.SysUser{
				Username:    "ourspark_user_" + oid + "_" + oursparkUser.Name,
				NickName:    oursparkUser.Name,
				AuthorityId: 3, // 布道师身份
			})
		} else {
			err, sysUser = service.Register(model.SysUser{
				Username:    "ourspark_user_" + oid + "_" + oursparkUser.Name,
				NickName:    oursparkUser.Name,
				AuthorityId: 2, // 默认普通用户
			})
		}

		if err != nil {
			global.LOG.Error(err.Error())
			return model.SysUser{}, err
		}
		global.DB.Create(&model.OauthOursparkUsers{
			Name:       oursparkUser.Name,
			University: uint(oursparkUser.Univ),
			Stuno:      oursparkUser.Stuno,
			Role:       oursparkUser.Role,
			Oid:        oid,
			SysUserId:  sysUser.ID,
		})
		return sysUser, err
	} else { // 更新 oauth User 信息
		err, sysUser := service.FindUserById(int(oauthOursparkUser.SysUserId))
		if err != nil {
			global.LOG.Error(err.Error())
			return model.SysUser{}, err
		}
		if oauthOursparkUser.Role == 5 { // 校外教师
			global.DB.Model(&sysUser).Updates(model.SysUser{
				Username:    "ourspark_user_" + oid + "_" + oursparkUser.Name,
				NickName:    oursparkUser.Name,
				AuthorityId: 3, // 布道师身份
			})
		} else {
			global.DB.Model(&sysUser).Updates(model.SysUser{
				Username:    "ourspark_user_" + oid + "_" + oursparkUser.Name,
				NickName:    oursparkUser.Name,
				AuthorityId: 2, // 默认普通用户
			})
		}
		global.DB.Model(&oauthOursparkUser).Updates(model.OauthOursparkUsers{
			Name:       oursparkUser.Name,
			University: uint(oursparkUser.Univ),
			Stuno:      oursparkUser.Stuno,
			Role:       oursparkUser.Role,
		})
		err, sysUser = service.FindUserById(int(oauthOursparkUser.SysUserId))
		return *sysUser, err
	}
}
