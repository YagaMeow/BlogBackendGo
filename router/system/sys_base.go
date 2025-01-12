package system

import (
	v1 "blog-backend/api/v1"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) gin.IRoutes {
	baseRouter := Router.Group("base")

	// baseRouter.POST("register", func(context *gin.Context) {
	// 	var form system.Register
	// 	if err := context.ShouldBindJSON(&form); err != nil {
	// 		context.JSON(http.StatusOK, gin.H{
	// 			"error": utils.GetErrorMsg(form, err),
	// 		})
	// 		return
	// 	}
	// 	context.JSON(http.StatusOK, gin.H{
	// 		"message": "success",
	// 	})
	// })

	baseRouter.POST("users", v1.ApiGroupApp.SystemApiGroup.CreateUser)
	baseRouter.GET("users", v1.ApiGroupApp.SystemApiGroup.GetUserList)

	return baseRouter
}
