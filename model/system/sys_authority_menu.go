package system

type Menu struct {
	BaseMenu
	MenuId      uint
	Authorityid uint
	Children    []Menu
	Parameters  []BaseMenuParameter
	Btns        map[string]int
}

type AuthorityMenu struct {
	MenuId      string
	AuthorityId string
}

func (AuthorityMenu) TableName() string {
	return "sys_authority_menus"
}
