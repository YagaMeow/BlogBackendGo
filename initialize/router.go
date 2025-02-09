package initialize

import (
	"blog-backend/global"
	"blog-backend/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System

	PublicGroup := Router.Group("")

	PublicGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	systemRouter.InitBaseRouter(PublicGroup)

	global.YAGAMI_ROUTER = Router.Routes()

	return Router
}
