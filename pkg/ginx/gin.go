package ginx

import (
	"app/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Default() *gin.Engine {
	engine := gin.New()
	engine.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowOrigins:     []string{"*"},
		MaxAge:           12 * time.Hour,
	}))
	engine.Use(gin.Logger())
	engine.Use(Context())
	engine.Use(func(ctx *gin.Context) {
		ctx.Request = ctx.Request.WithContext(logger.NewLoggerContextWithTraceId(ctx.Request.Context(), ""))
		ctx.Next()
	})
	return engine
}
