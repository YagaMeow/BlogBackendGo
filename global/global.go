package global

import (
	"blog-backend/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	YAGAMI_CONFIG config.Config
	YAGAMI_VIPER  *viper.Viper
	YAGAMI_LOGGER *zap.Logger
)
