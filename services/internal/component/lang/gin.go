package lang

import (
	"context"

	"github.com/gin-gonic/gin"
)

type langContextKey struct{}

func FromContext(ctx context.Context) string {
	return ctx.Value(langContextKey{}).(string)
}

func NewContext(ctx context.Context, l string) context.Context {
	return context.WithValue(ctx, langContextKey{}, l)
}

var Middleware gin.HandlerFunc = func(c *gin.Context) {
	l := c.GetHeader("Accept-Language")
	c.Request = c.Request.WithContext(NewContext(c.Request.Context(), l))
	c.Next()
}
