package repo

import (
	"app/api/model"
	"app/pkg/gormx"

	"gorm.io/gorm"
)

type CenterUserRepo struct {
	*gormx.Repo[model.CenterUser]
}

func NewCenterUserRepo(db *gorm.DB) *CenterUserRepo {
	return &CenterUserRepo{Repo: gormx.NewRepo[model.CenterUser](db)}
}
