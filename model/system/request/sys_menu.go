package request

import (
	"blog-backend/global"
	"blog-backend/model/system"
)

type AddMenuAuthorityInfo struct {
	Menus       []system.BaseMenu `json:"menus"`
	AuthorityId uint              `json:"authorityId"`
}

func DefaultMenu() []system.BaseMenu {
	return []system.BaseMenu{
		{
			YAGAMI_MODEL: global.YAGAMI_MODEL{ID: 1},
			ParentId:     0,
			Path:         "",
			Name:         "",
			Component:    "",
			Sort:         1,
			Meta: system.Meta{
				Title: "仪表盘",
				Icon:  "setting",
			},
		},
	}
}
