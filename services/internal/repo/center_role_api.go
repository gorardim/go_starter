package repo

import (
	"app/api/model"
	"app/pkg/gormx"
	"gorm.io/gorm"
)

type CenterRoleApiRepo struct {
	*gormx.Repo[model.CenterRoleApi]
}

func NewCenterRoleApiRepo(db *gorm.DB) *CenterRoleApiRepo {
	return &CenterRoleApiRepo{Repo: gormx.NewRepo[model.CenterRoleApi](db)}
}
