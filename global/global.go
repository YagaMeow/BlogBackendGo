package global

import (
	"blog-backend/config"

	"github.com/spf13/viper"
)

var (
	YAGAMI_CONFIG config.Config
	YAGAMI_VIPER  *viper.Viper
)
