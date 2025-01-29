package repo

import (
	"app/api/model"
	"app/pkg/gormx"
	"gorm.io/gorm"
)

type SettingRepo struct {
	*gormx.Repo[model.Setting]
}

func NewSettingRepo(db *gorm.DB) *SettingRepo {
	return &SettingRepo{Repo: gormx.NewRepo[model.Setting](db)}
}
