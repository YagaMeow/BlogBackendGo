package system

import (
	"blog-backend/global"
	"blog-backend/model/common/response"
	"blog-backend/model/system"
	systemReq "blog-backend/model/system/request"
	systemRes "blog-backend/model/system/response"
	"blog-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct{}

func (u *UserApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	// key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	userL := &system.User{UserName: l.Username, Password: l.Password}
	user, err := userService.Login(userL)
	if err != nil {
		response.FailWithMessage("用户名不存在或密码错误", c)
		return
	}
	u.TokenNext(c, *user)
	return
}

func (u *UserApi) TokenNext(c *gin.Context, user system.User) {
	token, claims, err := utils.LoginToken(&user)
	if err != nil {
		global.YAGAMI_LOGGER.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	response.OkWithDetailed(systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)
}
