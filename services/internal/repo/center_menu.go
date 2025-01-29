package repo

import (
	"app/api/model"
	"app/pkg/gormx"
	"gorm.io/gorm"
)

type CenterMenuRepo struct {
	*gormx.Repo[model.CenterMenu]
}

func NewCenterMenuRepo(db *gorm.DB) *CenterMenuRepo {
	return &CenterMenuRepo{Repo: gormx.NewRepo[model.CenterMenu](db)}
}
