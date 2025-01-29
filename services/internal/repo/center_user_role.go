package repo

import (
	"app/api/model"
	"app/pkg/gormx"
	"gorm.io/gorm"
)

type CenterUserRoleRepo struct {
	*gormx.Repo[model.CenterUserRole]
}

func NewCenterUserRoleRepo(db *gorm.DB) *CenterUserRoleRepo {
	return &CenterUserRoleRepo{Repo: gormx.NewRepo[model.CenterUserRole](db)}
}
