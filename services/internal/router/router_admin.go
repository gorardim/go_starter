package router

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	adminApi "app/api/admin"
	"app/pkg/errx"
	"app/pkg/ginx"
	"app/pkg/swagger"
	"app/services/internal/component"
	"app/services/internal/config"

	ut "github.com/go-playground/universal-translator"

	"app/services/internal/component/lang"

	"app/services/internal/service/admin"
	svcadmin "app/services/internal/service/admin"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
)

type AdminRouter struct {
	Config                   *config.Config
	Auth                     *admin.Auth
	UserAuthComponent        *component.UserAuthComponent
	AdminPermissionComponent *component.AdminPermissionComponent
	CenterMenu               *admin.CenterMenu
	CenterUser               *admin.CenterUser
	CenterRole               *admin.CenterRole
	CenterAPI                *admin.CenterApi
	Upload                   *svcadmin.Upload
	AdminImage               *admin.AdminImage
	User                     *admin.User
	AdminVideo               *admin.AdminVideo
	Setting                  *admin.Setting
}

func (r *AdminRouter) Run(addr string) error {
	router := ginx.Default()
	router.Use(lang.Trans())

	router.Use(ginx.AccessLog([]string{
		"/admin/_swagger",
		"/admin/upload/image",
	}, ginx.Desensitize{
		"/admin/auth/login": []string{"password"},
	}))

	if r.Config.Env == "local" {
		router.Use(gin.Recovery())
	} else {
		router.Use(ginx.Recovery())
	}

	middlewares := []gin.HandlerFunc{
		ginx.Context(),
		lang.Middleware,
		r.UserAuthComponent.Auth(),
		r.AdminPermissionComponent.Middleware(map[string]bool{
			"/admin/auth/login":   true,
			"/admin/auth/captcha": true,
			"/admin/auth/menus":   true,
		}),
	}

	_ = middlewares

	// register service ...
	swagger.Router(router, "/admin", adminApi.Openapi)
	adminApi.RegisterAuthServer(router, r.Auth, adminHandleConvert, middlewares...)
	adminApi.RegisterCenterMenuServer(router, r.CenterMenu, adminHandleConvert, middlewares...)
	adminApi.RegisterCenterUserServer(router, r.CenterUser, adminHandleConvert, middlewares...)
	adminApi.RegisterCenterRoleServer(router, r.CenterRole, adminHandleConvert, middlewares...)
	adminApi.RegisterCenterApiServer(router, r.CenterAPI, adminHandleConvert, middlewares...)
	adminApi.RegisterAdminImageServer(router, r.AdminImage, adminHandleConvert, middlewares...)
	adminApi.RegisterUserServer(router, r.User, adminHandleConvert, middlewares...)
	adminApi.RegisterAdminVideoServer(router, r.AdminVideo, adminHandleConvert, middlewares...)
	adminApi.RegisterSettingServer(router, r.Setting, adminHandleConvert, middlewares...)

	if r.Config.Env != "prod" {
		fmt.Println("debug mode")
	}

	router.POST("/admin/upload/image", append(middlewares, apiHandleConvert(r.Upload.UploadImage))...)
	router.POST("/admin/upload/video", append(middlewares, apiHandleConvert(r.Upload.UploadVideo))...)
	return router.Run(addr)
}
func Translate(errs validator.ValidationErrors, ut ut.Translator) []string {
	return lo.Map(errs, func(err validator.FieldError, i int) string {
		return strings.Replace(err.Translate(ut), err.Field(), lang.I18n.Get(ut.Locale(), err.Field()), 1)
	})
}

func adminHandleConvert(svc func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, err := svc(c)
		tail := time.Now().UnixMilli() % 100000
		if err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				if trans, has := lang.TransContext(c.Request.Context()); has {

					c.JSON(http.StatusBadRequest, gin.H{
						"code":    adminApi.ErrInvalidParam,
						"message": strings.Join(Translate(errs, trans), "; "),
					})
					return
				}
			}

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
