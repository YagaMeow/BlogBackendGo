package internal

import (
	"blog-backend/global"
	"fmt"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.YAGAMI_CONFIG.App.DbType {
	case "mysql":
		logZap = global.YAGAMI_CONFIG.MySQL.LogZap
	}
	if logZap {
		global.YAGAMI_LOGGER.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
