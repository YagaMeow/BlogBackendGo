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

func (AuthorityService) UpdateAuthority(auth system.Authority) (authority system.Authority, err error) {
	var oldAuthority system.Authority
	err = global.YAGAMI_DB.Where("authority_id = ?", auth.AuthorityId).First(&oldAuthority).Error
	if err != nil {
		global.YAGAMI_LOGGER.Debug(err.Error())
		return system.Authority{}, errors.New("该角色不存在")
	}
	err = global.YAGAMI_DB.Model(&oldAuthority).Updates(&auth).Error
	return auth, err
}

func (AuthorityService) DeleteAuthority(auth system.Authority) error {
	if errors.Is(global.YAGAMI_DB.First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("存在该角色用户，禁止删除")
	}
	if !errors.Is(global.YAGAMI_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在该角色用户，禁止删除")
	}
	if (!errors.Is(global.YAGAMI_DB.Where("parent_id = ?", auth.AuthorityId).First(&system.Authority{}).Error, gorm.ErrRecordNotFound)) {
		return errors.New("该角色存在子角色，禁止删除")
	}

	return global.YAGAMI_DB.Transaction(func(tx *gorm.DB) error {
		var err error
		if err = tx.Where("authority_id = ?", auth.AuthorityId).First(auth).Unscoped().Delete(auth).Error; err != nil {
			return err
		}
		return err

		if err = tx.Delete(&system.UserAuthority{}, "authority_authority_id = ?", auth.AuthorityId).Error; err != nil {
			return err
		}

		// authorityId := strconv.Itoa(int(auth.AuthorityId))
		return nil
	})
}

func (AuthorityService) GetAuthorityInfoList(authorityID uint) (list []system.Authority, err error) {
	var authority system.Authority
	err = global.YAGAMI_DB.Where("authority_id = ?", authorityID).First(&authority).Error
	if err != nil {
		return nil, err
	}
	var authorities []system.Authority
	db := global.YAGAMI_DB.Model(&system.Authority{})
	err = db.Where("parent_id = ?", "0").Find(&authorities).Error

	for k := range authorities {
		err = AuthorityServiceApp.FindChildrenAuthority(&authorities[k])
	}

	return authorities, err
}

func (AuthorityService) GetStructAuthorityList(authprityID uint) (list []uint, err error) {
	var auth system.Authority
	_ = global.YAGAMI_DB.Find(&auth, "authority_id = ?", authprityID).Error
	var authorities []system.Authority
	err = global.YAGAMI_DB.Where("parent_id = ?", authprityID).Find(&authorities).Error
	for k := range authorities {
		list = append(list, authorities[k].AuthorityId)
		_, err = AuthorityServiceApp.GetStructAuthorityList(authorities[k].AuthorityId)
	}
	if *auth.ParentId == 0 {
		list = append(list, authprityID)
	}
	return list, err
}

func (AuthorityService) CheckAuthorityIDAuth(authorityID, targetID uint) (err error) {
	authIDS, err := AuthorityServiceApp.GetStructAuthorityList(authorityID)
	var hasAuth bool = false
	for _, v := range authIDS {
		if v == targetID {
			hasAuth = true
			break
		}
	}
	if !hasAuth {
		return errors.New("角色ID不合法")
	}
	return nil
}

func (AuthorityService) FindChildrenAuthority(authority *system.Authority) (err error) {
	var authorities []system.Authority
	if err = global.YAGAMI_DB.Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error; err != nil {
		return err
	}
	for k := range authorities {
		err = AuthorityServiceApp.FindChildrenAuthority(&authorities[k])
	}
	return err
}
