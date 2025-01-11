package initialize

import (
	"blog-backend/global"

	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.YAGAMI_CONFIG.App.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
