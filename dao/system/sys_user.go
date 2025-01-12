package system

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	CreateTime string `json:"create_time"`
}

func (u *User) TableName() string {
	return "sys_user"
}
