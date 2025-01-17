package system

import "blog-backend/global"

type Api struct {
	global.YAGAMI_MODEL
	Path        string
	Description string
	ApiGroup    string
	Method      string
}

func (Api) TableName() string {
	return "sys_apis"
}

type IgnoreApi struct {
	global.YAGAMI_MODEL
	Path   string
	Method string
	Flag   bool `gorm:"-"`
}

func (IgnoreApi) TableName() string {
	return "sys_ignore_apis"
}
