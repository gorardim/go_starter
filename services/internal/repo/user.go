package repo

import (
	"app/api/model"
	"app/pkg/gormx"

	"gorm.io/gorm"
)

// 用户表

type UserRepo struct {
	*gormx.Repo[model.User]
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{Repo: gormx.NewRepo[model.User](db)}
}
