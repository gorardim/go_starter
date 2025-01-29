package provider

import (
	"github.com/mojocn/base64Captcha"
)

func NewCaptcha() *base64Captcha.Captcha {
	return base64Captcha.NewCaptcha(
		base64Captcha.NewDriverString(
			80, 160, 0, 0, 4,
			"23456789wertyupkjhgfdsazxcvbnm",
			nil,
			nil,
			[]string{"actionj.ttf"},
		),
		base64Captcha.DefaultMemStore,
	)
}
