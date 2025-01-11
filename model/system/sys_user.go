package system

import "blog-backend/utils"

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (register Register) GetMessages() utils.ValidatorMessages {
	return utils.ValidatorMessages{
		"Name.required":     "用户名不能为空",
		"Mobile.required":   "手机号码不能为空",
		"Password.required": "用户密码不能为空",
	}
}
