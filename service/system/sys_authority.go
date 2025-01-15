package system

import (
	"blog-backend/global"
	"blog-backend/model/system"
	systemReq "blog-backend/model/system/request"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

var ErrRoleExistence = errors.New("角色已存在")

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (AuthorityService) CreateAuthority(auth system.Authority) (authority system.Authority, err error) {
	if !errors.Is(global.YAGAMI_DB.Where("authority = id", auth.AuthorityId).First(&system.Authority{}).Error, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}
	e := global.YAGAMI_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&auth).Error; err != nil {
			return err
		}
		auth.BaseMenus = systemReq.DefaultMenu()
		if err = tx.Model(&auth).Association("BaseMenus").Replace(&auth.BaseMenus); err != nil {
			return err
		}

		casbinInfos := systemReq.DefaultCasbin()
		authorityId := strconv.Itoa(int(auth.AuthorityId))
		rules := [][]string{}
		for _, v := range casbinInfos {
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
		return nil
	})
	return auth, e
}
