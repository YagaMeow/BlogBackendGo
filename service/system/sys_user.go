package system

import (
	"blog-backend/dao/system"
	"blog-backend/global"
)

type UserService struct {
}

func (userService *UserService) CreateUser(user *system.User) (err error) {
	if err = global.YAGAMI_DB.Create(user).Error; err != nil {
		return err
	}
	return
}

func (userService *UserService) GetAllUser() (userList []*system.User, err error) {
	if err = global.YAGAMI_DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func (userService *UserService) DeleteUserById(id string) (err error) {
	err = global.YAGAMI_DB.Where("id=?", id).Delete(&system.User{}).Error
	return
}

func (userService *UserService) GetUserById(id string) (user *system.User, err error) {
	if err = global.YAGAMI_DB.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func (userService *UserService) UpdateUser(user *system.User) (err error) {
	err = global.YAGAMI_DB.Save(user).Error
	return
}
