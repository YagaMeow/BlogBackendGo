package main

import (
	"blog-backend/global"
	"blog-backend/initialize"
)

func main() {
	global.My_DB = initialize.Gorm()
	if global.My_DB != nil {
		initialize.RegisterTables()
		db, _ := global.My_DB.DB()
		defer db.Close()
	}
}
