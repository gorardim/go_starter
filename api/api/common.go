package api

import (
	"context"
	_ "embed"

	"app/api/model"
	"app/pkg/jwtutil"
)

//go:embed openapi.yml
var Openapi string

// AuthUser 授权用户
type AuthUser struct {
	UserId int `json:"user_id"`
}

func UserFormContext(ctx context.Context) (*AuthUser, bool) {
	return jwtutil.FromContext[AuthUser](ctx)
}

type Empty struct{}

const (
	ErrInvalidParam      = "ErrInvalidParam"
	ErrBusiness          = "ErrBusiness"
	ErrPayPasswordNotSet = "ErrPayPasswordNotSet"
	ErrInvalidPayPass    = "ErrInvalidPayPass"
	ErrUserLogOff        = "ErrUserLogOff"
)

type Link struct {
	// link type: ARTICLE, ARTICLE_CATEGORY, Page
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AdItem struct {
	// 图片地址
	ImageUrl string `json:"image_url"`
	// 标题
	Title string `json:"title"`
	// 跳转地址
	Link model.Link `json:"link"`
}
