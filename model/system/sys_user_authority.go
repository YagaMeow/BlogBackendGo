package system

type UserAuthority struct {
	UserId      uint
	AuthorityId uint
}

func (UserAuthority) TableName() string {
	return "sys_user_authority"
}
