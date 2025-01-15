package initialize

import (
	"blog-backend/global"
	"blog-backend/model/system"
	"fmt"

	"go.uber.org/zap"
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

func RegisterTables() {
	fmt.Println("[SQL]初始化表单...")

	db := global.YAGAMI_DB
	err := db.AutoMigrate(
		system.User{},
		system.BaseMenu{},
		system.Authority{},
		system.Article{},
	)

	if err != nil {
		global.YAGAMI_LOGGER.Error("[SQL]注册表失败 ", zap.Error(err))
		return
	}

	err = bizModel()

	if err != nil {
		global.YAGAMI_LOGGER.Error("[SQL]注册表失败 ", zap.Error(err))
	}

	fmt.Println("[SQL]注册表成功")
}
