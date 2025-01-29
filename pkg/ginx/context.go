package ginx

import (
	"context"
	"github.com/gin-gonic/gin"
)

type ctxKey struct{}

func FromContext(ctx context.Context) *gin.Context {
	return ctx.Value(ctxKey{}).(*gin.Context)
}

func NewContext(ctx context.Context, c *gin.Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, c)
}

var Context = func() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := NewContext(c.Request.Context(), c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
