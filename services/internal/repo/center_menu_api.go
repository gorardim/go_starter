package repo

import (
	"app/api/model"
	"app/pkg/gormx"

	"gorm.io/gorm"
)

type CenterMenuApiRepo struct {
	*gormx.Repo[model.CenterMenuApi]
}

func NewCenterMenuApiRepo(db *gorm.DB) *CenterMenuApiRepo {
	return &CenterMenuApiRepo{Repo: gormx.NewRepo[model.CenterMenuApi](db)}
}
