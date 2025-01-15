package system

import (
	"blog-backend/global"
	"blog-backend/model/system"
	"errors"

	"gorm.io/gorm"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (MenuService) GetMenuTreeMap(authorityId uint) (treeMap map[uint][]system.Menu, err error) {
	var allMenus []system.Menu
	var baseMenu []system.BaseMenu
	treeMap = make(map[uint][]system.Menu)

	var AuthorityMenus []system.AuthorityMenu
	err = global.YAGAMI_DB.Where("authority_authority_id = ?", authorityId).Find(&AuthorityMenus).Error
	if err != nil {
		return
	}

	var MenuIds []string

	for i := range AuthorityMenus {
		MenuIds = append(MenuIds, AuthorityMenus[i].MenuId)
	}

	err = global.YAGAMI_DB.Where("id in (?)", MenuIds).Order("sort").Find(&baseMenu).Error
	if err != nil {
		return
	}

	for i := range baseMenu {
		allMenus = append(allMenus, system.Menu{
			BaseMenu:    baseMenu[i],
			Authorityid: authorityId,
			MenuId:      baseMenu[i].ID,
			Parameters:  baseMenu[i].Parameters,
		})
	}

	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}

	return treeMap, err
}

func (MenuService) GetMenuTree(authorityId uint) (menu []system.Menu, err error) {
	menuTree, err := MenuServiceApp.GetMenuTreeMap(authorityId)
	menus := menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = MenuServiceApp.GetChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

func (MenuService) GetChildrenList(menu *system.Menu, treeMap map[uint][]system.Menu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = MenuServiceApp.GetChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (MenuService) GetBaseMenuTreeMap(authorityId uint) (treeMap map[uint][]system.BaseMenu, err error) {
	treeMap = make(map[uint][]system.BaseMenu)
	var authorityMenus []system.AuthorityMenu

	if err = global.YAGAMI_DB.Where("authority_authority_id = ?", authorityId).Find(&authorityMenus).Error; err != nil {
		return
	}

	var menuIds []string
	for i := 0; i < len(authorityMenus); i++ {
		menuIds = append(menuIds, authorityMenus[i].MenuId)
	}

	var menus []system.BaseMenu
	if err = global.YAGAMI_DB.Where("id in (?)", menuIds).Order("sort").Find(&menus).Error; err != nil {
		return
	}

	for _, v := range menus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

func (MenuService) GetBaseMenuTree(authorityId uint) (menuTree []system.BaseMenu, err error) {
	treeMap, err := MenuServiceApp.GetBaseMenuTreeMap(authorityId)
	menus := treeMap[0]
	for i := 0; i < len(menus); i++ {
		err = MenuServiceApp.GetBaseChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

func (MenuService) GetBaseChildrenList(menu *system.BaseMenu, treeMap map[uint][]system.BaseMenu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = MenuServiceApp.GetBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (MenuService) GetInfoList(authorityId uint) (list any, err error) {
	var MenuList []system.BaseMenu
	treeMap, err := MenuServiceApp.GetBaseMenuTreeMap(authorityId)
	if err != nil {
		return
	}
	for i := 0; i < len(MenuList); i++ {
		err = MenuServiceApp.GetBaseChildrenList(&MenuList[i], treeMap)
	}
	return MenuList, err
}

func (MenuService) AddBaseMenu(menu system.BaseMenu) (err error) {
	if !errors.Is(global.YAGAMI_DB.Where("name = ?", menu.Name).First(&system.BaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("当前名称已存在")
	}
	return global.YAGAMI_DB.Create(&menu).Error
}

func (MenuService) AddMenuAuthority(menus []system.BaseMenu, adminAuthorityID uint, authorityId uint) (err error) {
	var auth system.Authority
	auth.AuthorityId = authorityId
	auth.BaseMenus = menus

	// err = AuthorityServiceApp.CreateAuthority()

	return
}

func (MenuService) UserAuthorityDefaultRouter(user *system.User) {
	var menuIds []string
	err := global.YAGAMI_DB.Model(&system.AuthorityMenu{}).Where("authority_authority_id = ?", user.AuthorityId).Pluck("base_menu_id", &menuIds).Error
	if err != nil {
		return
	}
	err = global.YAGAMI_DB.First(&system.BaseMenu{}, "name = ? and id in (?)", user.UserName, menuIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Authority.DefaultRouter = "404"
	}
}
