package router

import (
	"fmt"
	"net/http"
	"time"

	"app/api/api"
	"app/pkg/errx"
	"app/pkg/ginx"
	"app/pkg/swagger"
	"app/services/internal/component/apiuser"
	"app/services/internal/component/httpx"
	"app/services/internal/component/lang"
	"app/services/internal/config"
	"app/services/internal/repo"
	svcapi "app/services/internal/service/api"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	Config        *config.Config
	AuthServer    api.AuthServer
	UserRepo      *repo.UserRepo
	UserTokenRepo *repo.UserTokenRepo
	Upload        *svcapi.Upload
}

func (r *ApiRouter) Run(addr string) error {
	router := ginx.Default()

	if r.Config.Env != "local" {
		router.Use(ginx.AccessLog([]string{
			"/api/_swagger",
		}, ginx.Desensitize{
			"/api/auth/login":          []string{"password"},
			"/api/auth/register":       []string{"password", "confirm_password"},
			"/api/auth/password/reset": []string{"confirm_password", "new_password"},
		}))
	}

	if r.Config.Env == "local" {
		router.Use(gin.Recovery())
	} else {
		router.Use(ginx.Recovery())
	}

	middlewares := []gin.HandlerFunc{
		ginx.Context(),
		lang.Middleware,
		httpx.AuthMiddlewareFactory[api.AuthUser](r.Config.Jwt.OpenApiSecret, &httpx.AuthMiddlewareOptions{
			SkipPaths: []string{
				"/api/auth/login",
				"/api/auth/register",
				"/api/auth/email/send",
				"/api/auth/password/reset",
				"/api/auth/two_fa_login",
				"/api/auth/check_username",
			},
			BeforeNext: func(c *gin.Context) error {
				authUser, ok := api.UserFormContext(c.Request.Context())
				if !ok {
					return fmt.Errorf("invalid token")
				}
				token := c.GetHeader("Authorization")
				// find token
				has, err := r.UserTokenRepo.Exists(c.Request.Context(), "user_id = ? and token = ?", authUser.UserId, token)
				if err != nil {
					return err
				}
				if !has {
					return fmt.Errorf("invalid token")
				}
				return nil
			},
		}),
		apiuser.BindUser(r.UserRepo),
	}

	api.RegisterAuthServer(router, r.AuthServer, apiHandleConvert, middlewares...)

	router.POST("/api/upload/image", append(middlewares, apiHandleConvert(r.Upload.UploadImage))...)

	swagger.Router(router, "/api", api.Openapi)
	router.GET("api/health/check", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return router.Run(addr)
}

func apiHandleConvert(svc func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, err := svc(c)
		tail := time.Now().UnixMilli() % 100000
		if err != nil {
			if wrapError, ok := errx.As(err); ok {
				h := gin.H{
					"code":    wrapError.ErrorCode(),
					"message": fmt.Sprintf("%s[%d]", err.Error(), tail),
				}
				if wrapError.Detail() != nil {
					h["detail"] = wrapError.Detail()
				}
				c.JSON(http.StatusBadRequest, h)
				return
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"code":    "SYSTEM_ERROR",
				"message": fmt.Sprintf("%s[%d]", err.Error(), tail),
			})
			return
		}
		c.JSON(http.StatusOK, value)
	}
}
