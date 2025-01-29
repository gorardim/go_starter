package httpx

import (
	"net/http"

	"app/pkg/jwtutil"
	"app/services/internal/component/lang"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareOptions struct {
	SkipPaths  []string
	BeforeNext func(c *gin.Context) error
}

//  生成一个gin中间件，用于验证token

func AuthMiddlewareFactory[T any](secret string, ops *AuthMiddlewareOptions) gin.HandlerFunc {
	var skipPathMap = make(map[string]bool)
	for _, path := range ops.SkipPaths {
		skipPathMap[path] = true
	}
	return func(c *gin.Context) {
		if skipPathMap[c.Request.URL.Path] {
			c.Next()
			return
		}
		// token 不能为空
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": lang.T(c.Request.Context(), "LOGIN_EXPIRED"),
				"code":    "ErrUnauthenticated",
				"detail":  "authorization header is empty",
			})
			c.Abort()
			return
		}

		// jwt解析
		claim, err := jwtutil.Parse[T](secret, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": lang.T(c.Request.Context(), "LOGIN_EXPIRED"),
				"code":    "ErrUnauthenticated",
				"detail":  err.Error(),
			})
			c.Abort()
			return
		}
		c.Request = c.Request.WithContext(jwtutil.NewContext(c.Request.Context(), claim.Custom))
		if ops.BeforeNext != nil {
			if err := ops.BeforeNext(c); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": lang.T(c.Request.Context(), "LOGIN_EXPIRED"),
					"code":    "ErrUnauthenticated",
					"detail":  err.Error(),
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
