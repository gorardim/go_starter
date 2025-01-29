package repo

import (
	"app/api/model"
	"app/pkg/gormx"
	"gorm.io/gorm"
)

type CenterRoleRepo struct {
	*gormx.Repo[model.CenterRole]
}

func NewCenterRoleRepo(db *gorm.DB) *CenterRoleRepo {
	return &CenterRoleRepo{Repo: gormx.NewRepo[model.CenterRole](db)}
}
