package repo

import (
	"app/api/model"
	"app/pkg/gormx"

	"gorm.io/gorm"
)

// 图片表

type AdminImageRepo struct {
	*gormx.Repo[model.AdminImage]
}

func NewAdminImageRepo(db *gorm.DB) *AdminImageRepo {
	return &AdminImageRepo{Repo: gormx.NewRepo[model.AdminImage](db)}
}
