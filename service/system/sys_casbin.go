package system

import (
	"blog-backend/model/system/request"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

func (CasbinService) UpdateCasbin(adminAuthorityId, authorityId uint, casbinInfos []request.CasbinInfo) error {
	err := AuthorityServiceApp.CheckAuthorityIDAuth(adminAuthorityId, authorityId)
	if err != nil {
		return err
	}

	return ErrRoleExistence
}
