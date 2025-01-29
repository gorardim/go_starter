package repo

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewSettingRepo,
	NewUserRepo,
	NewUserTokenRepo,
	NewAdminImageRepo,
	NewAdminVideoRepo,
	// center
	NewCenterMenuRepo,
	NewCenterRoleRepo,
	NewCenterUserRepo,
	NewCenterUserRoleRepo,
	NewCenterApiRepo,
	NewCenterMenuApiRepo,
	NewCenterRoleApiRepo,
	NewCenterRoleMenuRepo,
)
