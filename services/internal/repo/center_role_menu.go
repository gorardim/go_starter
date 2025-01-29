package repo

import (
	"app/api/model"
	"app/pkg/gormx"
	"gorm.io/gorm"
)

type CenterRoleMenuRepo struct {
	*gormx.Repo[model.CenterRoleMenu]
}

func NewCenterRoleMenuRepo(db *gorm.DB) *CenterRoleMenuRepo {
	return &CenterRoleMenuRepo{Repo: gormx.NewRepo[model.CenterRoleMenu](db)}
}
