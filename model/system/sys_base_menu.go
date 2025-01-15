package system

import "blog-backend/global"

type BaseMenu struct {
	global.YAGAMI_MODEL
	MenuLevel uint
	ParentId  uint
	Path      string
	Name      string
	Hidden    bool
	Component string
	Sort      int
	Meta
	Authoritys []Authority `gorm:"many2many:sys_authority_menus"`
	Children   []BaseMenu  `gorm:"-"`
	Parameters []BaseMenuParameter
}

type Meta struct {
	ActiveName  string
	KeepAlive   bool
	DefaultMenu bool
	Icon        string
	Title       string
	CloseTab    bool
}

type BaseMenuParameter struct {
	global.YAGAMI_MODEL
	BaseMenuId uint
	Type       string
	Key        string
	Value      string
}

func (BaseMenu) TableName() string {
	return "sys_menus"
}
