package system

import (
	"time"
)

type Authority struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time   `sql:"index"`
	AuthorityId   uint        `json:"authorityId" gorm:"not null;unique;primary_key;"`
	AuthorityName string      `json:"authorityName"`
	ParentId      *uint       `json:"parentId"`
	Children      []Authority `json:"children" gorm:"-"`
	BaseMenus     []BaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users         []User      `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter string      `json:"defaultRouter"`
}

func (Authority) TableName() string {
	return "sys_authorities"
}
