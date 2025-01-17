package system

type UserAuthority struct {
	UserId      uint `gorm:"column:user_id"`
	AuthorityId uint `gorm:"column:authority_autohrity_id"`
}

func (UserAuthority) TableName() string {
	return "sys_user_authority"
}
