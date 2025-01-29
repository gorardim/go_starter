package repo

import (
	"app/api/model"
	"app/pkg/gormx"
	"gorm.io/gorm"
)

// 用户token表

type UserTokenRepo struct {
	*gormx.Repo[model.UserToken]
}

func NewUserTokenRepo(db *gorm.DB) *UserTokenRepo {
	return &UserTokenRepo{Repo: gormx.NewRepo[model.UserToken](db)}
}
