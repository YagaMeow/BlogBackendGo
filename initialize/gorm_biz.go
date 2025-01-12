package initialize

import "blog-backend/global"

func bizModel() error {
	db := global.YAGAMI_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
