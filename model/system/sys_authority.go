package system

import (
	"blog-backend/global"
)

type Authority struct {
	global.YAGAMI_MODEL
	AuthorityId   uint        `json:"authorityId" gorm:"not null;unique;primary_key;"`
	AuthorityName string      `json:"authorityName"`
	ParentId      *uint       `json:"parentId"`
	Children      []Authority `json:"children" gorm:"-"`
	Users         []User      `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter string      `json:"defaultRouter"`
}

func (Authority) TableName() string {
	return "sys_authorities"
}
