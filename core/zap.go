package core

import (
	"blog-backend/core/internal"
	"blog-backend/global"
	"blog-backend/utils"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.YAGAMI_CONFIG.Zap.Director); !ok {
		fmt.Println("[Zap]创建日志文件夹", global.YAGAMI_CONFIG.Zap.Director)
		_ = os.Mkdir(global.YAGAMI_CONFIG.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.YAGAMI_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	fmt.Println("[Zap]日志初始化完成")
	return logger
}
