package core

import (
	"blog-backend/global"
	"blog-backend/initialize"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func Run() {

	Router := initialize.Routers()

	address := fmt.Sprintf(":%s", global.YAGAMI_CONFIG.App.Port)
	s := InitServer(address, Router)

	global.YAGAMI_LOGGER.Info("server run success on", zap.String("address", address))

	time.Sleep(10 * time.Millisecond)

	global.YAGAMI_LOGGER.Error(s.ListenAndServe().Error())

}
