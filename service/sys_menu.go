package service

import "blog-backend/model/system"

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (MenuService) UserAuthorityDefaultRouter(user *system.User) {
}
