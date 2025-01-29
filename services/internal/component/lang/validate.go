package lang

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/vi"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"

	enTranslations "github.com/go-playground/validator/v10/translations/en"
	viTranslations "github.com/go-playground/validator/v10/translations/vi"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	zhLocal, enLocal, viLocal = zh.New(), en.New(), vi.New()
)

type transKey struct{}

func TransContext(ctx context.Context) (ut.Translator, bool) {
	v, ok := ctx.Value(transKey{}).(ut.Translator)
	return v, ok
}

func NewTransContext(ctx context.Context, trans locales.Translator) context.Context {
	return context.WithValue(ctx, transKey{}, trans)
}

func Trans() gin.HandlerFunc {
	uni := NewUniversalTranslator()
	return func(c *gin.Context) {
		locale := c.GetHeader("Accept-Language")
		if !lo.Contains([]string{Mongolian, English}, locale) {
			locale = Mongolian
		}
		if trans, found := uni.GetTranslator(locale); found {
			c.Request = c.Request.WithContext(NewTransContext(c.Request.Context(), trans))
		}
		c.Next()
	}
}

func NewUniversalTranslator() *ut.UniversalTranslator {
	uni := ut.New(zhLocal, zhLocal, enLocal, viLocal)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if trans, found := uni.GetTranslator("en"); found {
			_ = enTranslations.RegisterDefaultTranslations(v, trans)
		}

		if trans, found := uni.GetTranslator("zh"); found {
			_ = zhTranslations.RegisterDefaultTranslations(v, trans)
			trans.Add("required", "{0}不能为空", true)
		}

		if trans, found := uni.GetTranslator("vi"); found {
			_ = viTranslations.RegisterDefaultTranslations(v, trans)
		}
	}

	return uni
}
