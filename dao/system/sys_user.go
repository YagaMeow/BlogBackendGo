package system

type Login interface {
	GetId() int
	GetName() string
	GetCreateTime() string
}

var _ Login = new(User)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	CreateTime string `json:"create_time"`
}

func (u *User) GetId() int {
	return u.Id
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetCreateTime() string {
	return u.CreateTime
}

func (u *User) TableName() string {
	return "sys_users"
}
