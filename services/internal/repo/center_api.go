package repo

import (
	"app/api/model"
	"app/pkg/gormx"

	"gorm.io/gorm"
)

type CenterApiRepo struct {
	*gormx.Repo[model.CenterApi]
}

func NewCenterApiRepo(db *gorm.DB) *CenterApiRepo {
	return &CenterApiRepo{Repo: gormx.NewRepo[model.CenterApi](db)}
}
