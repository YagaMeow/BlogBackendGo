package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	My_DB  *gorm.DB
	My_LOG *zap.Logger
	My_VP  *viper.Viper
)
