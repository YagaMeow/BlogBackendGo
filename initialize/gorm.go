package initialize

import (
	"blog-backend/global"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {

	mysqlConfig := mysql.Config{
		DSN:                       "root:root@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // 设置默认字符串长度
		SkipInitializeWithVersion: false, // 根据版本自动初始化
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		return nil
	} else {
		return db
	}

}

func RegisterTables() {
	db := global.My_DB
	err := db.AutoMigrate()
	if err != nil {
		global.My_LOG.Error("register tables failed", zap.Error(err))
		os.Exit(0)
	}
}
