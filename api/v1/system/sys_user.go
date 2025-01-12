package system

import (
	"blog-backend/dao/common/response"
	"blog-backend/dao/system"
	systemReq "blog-backend/dao/system/request"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (u *UserApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	userL := &system.User{Name: l.Username, Password: l.Password}
	user, err := userService.Login(userL)
	if err != nil {
		response.FailWithMessage("用户名不存在或密码错误", c)
		return
	}
	u.TokenNext(c, *user)
	return

}

func (u *UserApi) TokenNext(c *gin.Context, user system.User) {

}

func (u *UserApi) CreateUser(c *gin.Context) {
	var user system.User
	c.BindJSON(&user)

	err := userService.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": user,
		})
	}
}

func (u *UserApi) GetUserList(c *gin.Context) {
	userList, err := userService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    userList,
		})
	}
}

func (u *UserApi) DeleteUser(c *gin.Context) {
	var user system.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = userService.DeleteUserById(user.GetId())
	if err != nil {
		fmt.Print("[DeleteUser] err")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		fmt.Println("[DeleteUser] ok")
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	}
}
