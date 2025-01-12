package system

import (
	"blog-backend/service"
)

type ApiGroup struct {
	UserApi
}

var (
	userService    = service.ServiceGroupApp.SystemServiceGroup.UserService
	articleService = service.ServiceGroupApp.SystemServiceGroup.ArticleService
)
