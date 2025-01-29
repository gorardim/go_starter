package admin

import (
	"context"
)

// AuthServer 授权服务 (release)
type AuthServer interface {
	//x:api post /admin/auth/login 用户登陆
	AuthLogin(ctx context.Context, req *AuthLoginRequest) (*AuthLoginResponse, error)
	//x:api post /admin/auth/captcha 验证码
	Captcha(ctx context.Context, req *AuthCaptchaRequest) (*AuthCaptchaResponse, error)
	//x:api post /admin/auth/menus 用户菜单
	AuthMenus(ctx context.Context, req *AuthMenusRequest) (*AuthMenusResponse, error)
}

type AuthLoginRequest struct {
	// 用户名
	Username string `json:"username" binding:"required"`
	// 密码
	Password string `json:"password" binding:"required"`
	// 验证码
	VerifyCode string `json:"verify_code"`
	// 验证码ID
	CaptchaId string `json:"captcha_id"`
}

type AuthLoginResponse struct {
	// token
	Token string `json:"token"`
	// 用户id
	UserId string `json:"user_id"`
	// 用户名
	Username string `json:"username"`
	// 昵称
	Nickname string `json:"nickname"`
}

type AuthCaptchaRequest struct{}

type AuthCaptchaResponse struct {
	// 验证码ID
	CaptchaId string `json:"captcha_id"`
	// 验证码图片
	CaptchaImage string `json:"captcha_image"`
}

type AuthMenusRequest struct{}
type AuthMenusResponse struct {
	Items []*CenterMenuInfo `json:"items"`
}
