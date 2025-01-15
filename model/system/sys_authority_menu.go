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
	MenuId      uint
	AuthorityId uint
}

func (AuthorityMenu) TableName() string {
	return "sys_authority_menus"
}
