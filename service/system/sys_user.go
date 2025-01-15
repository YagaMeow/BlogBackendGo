package system

import (
	"blog-backend/global"
	"blog-backend/model/system"
	"blog-backend/utils"
	"errors"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type UserService struct {
}

var UserServiceApp = new(UserService)

func (UserService) Register(u system.User) (userInter system.User, err error) {
	var user system.User
	if !errors.Is(global.YAGAMI_DB.Where("username = ?", u.UserName).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}

	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err = global.YAGAMI_DB.Create(&u).Error
	return u, err
}

func (UserService) Login(u *system.User) (userInter *system.User, err error) {
	if nil == global.YAGAMI_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.User
	err = global.YAGAMI_DB.Where("username = ?", u.UserName).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

func (UserService) ChangePassword(u *system.User, newPassWord string) (userInter *system.User, err error) {
	var user system.User
	if err = global.YAGAMI_DB.Where("username = ?", u.UserName).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassWord)
	err = global.YAGAMI_DB.Save(&user).Error
	return &user, err
}

func (UserService) GetUserInfo(uuid uuid.UUID) (user system.User, err error) {
	var reqUser system.User
	return reqUser, err
}
