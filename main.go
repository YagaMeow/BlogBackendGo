package main

import (
	"blog-backend/core"
	"blog-backend/global"
	"blog-backend/initialize"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const AppMode = "debug"

func main() {
	gin.SetMode(AppMode)

	global.YAGAMI_VIPER = core.InitViper()

	global.YAGAMI_LOGGER = core.InitZap()
	zap.ReplaceGlobals(global.YAGAMI_LOGGER)
	global.YAGAMI_LOGGER.Info("server run success on ", zap.String("zap_log", global.YAGAMI_CONFIG.Zap.Director))

	global.YAGAMI_DB = initialize.Gorm()

	initialize.OtherInit()

	fmt.Println("[Server]APP:", global.YAGAMI_CONFIG.App.AppName)
	core.Run()
}
