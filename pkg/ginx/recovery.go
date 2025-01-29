package ginx

import (
	"app/pkg/alert"
	"app/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"time"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var buf [1024]byte
				_ = runtime.Stack(buf[:], false)
				c.AbortWithStatusJSON(500, gin.H{
					"code": "SystemError",
					"msg":  fmt.Sprintf("系统错误，请重试[%5d]", time.Now().UnixMilli()%100000),
				})
				alert.Alert(c.Request.Context(), "panic: "+fmt.Sprint(err), []string{
					"path: " + string(c.Request.RequestURI),
					"stack: " + string(buf[:]),
				})
				logger.Printf(c.Request.Context(), "panic: %s\n%s", fmt.Sprint(err), string(buf[:]))
			}
		}()
		c.Next()
	}
}
