package main

import (
	"blog-backend/core"
	"blog-backend/global"
	"fmt"
)

func main() {
	global.YAGAMI_VIPER = core.InitViper()
	fmt.Println("[Server]APP:", global.YAGAMI_CONFIG.App.AppName)
	core.Run()
}
