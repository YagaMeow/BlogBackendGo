package system

import (
	"blog-backend/global"
	"blog-backend/model/system"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type ApiService struct{}

var ApiServiceApp = new(ApiService)

func (ApiService) CreateApi(api system.Api) (err error) {
	if !errors.Is(global.YAGAMI_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.Api{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("api已存在")
	}
	return global.YAGAMI_DB.Create(&api).Error
}

func (ApiService) GetApiGroups() (groups []string, groupApiMap map[string]string, err error) {
	var apis []system.Api
	err = global.YAGAMI_DB.Find(&apis).Error
	if err != nil {
		return
	}
	groupApiMap = make(map[string]string, 0)
	for i := range apis {
		pathArr := strings.Split(apis[i].Path, "/")
		newGroup := true
		for j := range groups {
			if groups[j] == apis[i].ApiGroup {
				newGroup = false
			}
		}
		if newGroup {
			groups = append(groups, apis[i].ApiGroup)
		}
		groupApiMap[pathArr[1]] = apis[i].ApiGroup
	}
	return
}

func (ApiService) SyncApi() (newApis, deleteApis, ignoreApis []system.Api, err error) {
	newApis = make([]system.Api, 0)
	deleteApis = make([]system.Api, 0)
	ignoreApis = make([]system.Api, 0)
	var apis []system.Api
	err = global.YAGAMI_DB.Find(&apis).Error
	if err != nil {
		return
	}
	var ignores []system.IgnoreApi
	err = global.YAGAMI_DB.Find(&ignores).Error
	if err != nil {
		return
	}

	for i := range ignores {
		ignoreApis = append(ignoreApis, system.Api{
			Path:        ignores[i].Path,
			Description: "",
			ApiGroup:    "",
			Method:      ignores[i].Method,
		})
	}

	// var cacheApis []system.Api
	// for i := range global.YAGAMI_ROUTER {
	// 	ignoreFlag := false
	// 	for j := range ignores {
	// 		if ignores[j].Path == global.YAGAMI_ROUTER[i].Path && ignores.method
	// 	}
	// }
	return
}
