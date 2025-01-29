package repo

import (
	"app/api/model"
	"app/pkg/gormx"

	"gorm.io/gorm"
)

type AdminVideoRepo struct {
	*gormx.Repo[model.AdminVideo]
}

func NewAdminVideoRepo(db *gorm.DB) *AdminVideoRepo {
	return &AdminVideoRepo{Repo: gormx.NewRepo[model.AdminVideo](db)}
}
