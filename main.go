package main

import (
	"blog-backend/core"
	"blog-backend/global"
	"blog-backend/initialize"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	global.YAGAMI_VIPER = core.InitViper()
	global.YAGAMI_LOGGER = core.InitZap()
	zap.ReplaceGlobals(global.YAGAMI_LOGGER)
	global.YAGAMI_LOGGER.Info("server run success on ", zap.String("zap_log", "zap_log"))
	global.YAGAMI_DB = initialize.Gorm()

	fmt.Println("[Server]APP:", global.YAGAMI_CONFIG.App.AppName)
	core.Run()
}
