package system

import (
	"blog-backend/global"
	"blog-backend/model/common"

	"github.com/gofrs/uuid/v5"
)

type Login interface {
	GetUsername() string
	GetNickname() string
	GetUUID() uuid.UUID
	GetUserId() uint
	GetAuthorityId() uint
	GetUserInfo() any
}

var _ Login = new(User)

type User struct {
	global.YAGAMI_MODEL
	UUID          uuid.UUID      `json:"uuid" gorm:"index"`
	UserName      string         `json:"userName" gorm:"index"`
	Password      string         `json:"-"`
	NickName      string         `json:"nickName"`
	AuthorityId   uint           `json:"authorityId" gorm:"default: 888;"`
	Authority     Authority      `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;"`
	Authorities   []Authority    `json:"authorities" gorm:"many2many:sys_user_authority"`
	Phone         string         `json:"phone"`
	Email         string         `json:"email"`
	Enable        string         `json:"enable"`
	OriginSetting common.JSONMap `json:"originSetting" form:"originSetting" gorm:"type:text;default:null;column:origin_setting;"`
}

func (u *User) GetUsername() string {
	return u.UserName
}
func (u *User) GetNickname() string {
	return u.NickName
}
func (u *User) GetUUID() uuid.UUID {
	return u.UUID
}
func (u *User) GetUserId() uint {
	return u.ID
}
func (u *User) GetAuthorityId() uint {
	return u.AuthorityId
}
func (u *User) GetUserInfo() any {
	return *u
}

func (u *User) TableName() string {
	return "sys_users"
}
