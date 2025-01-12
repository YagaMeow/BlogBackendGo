package system

import (
	"blog-backend/dao/system"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

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
