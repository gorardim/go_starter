package component

import (
	"app/api/admin"
	"app/api/api"
	"app/pkg/errx"
	"app/pkg/gormx"
	"app/pkg/gormx/g"
	"app/pkg/sqx"
	"app/services/internal/repo"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminPermissionComponent struct {
	CenterApiRepo      *repo.CenterApiRepo
	CenterMenuApiRepo  *repo.CenterMenuApiRepo
	CenterUserRoleRepo *repo.CenterUserRoleRepo
	DB                 *gorm.DB
}

func (a *AdminPermissionComponent) CheckAdminPermission(ctx context.Context, req *http.Request) error {
	user, ok := UserFormContext(ctx)
	if !ok {
		return errx.New(api.ErrInvalidParam, "user not found in context")
	}

	// find super role
	// if it has super role, permit it
	builder := sqx.Select("r.id").
		From("center_user_role ur").
		LeftJoin("center_role r on ur.role_id = r.id").
		Where("ur.user_id = ?", user.UserId).
		Where("r.is_super = ?", "Y")
	has, err := gormx.Exists(ctx, a.DB, builder)
	if err != nil {
		return err
	}
	if has {
		return nil
	}

	path := req.URL.Path
	// find api
	// if api not in menu, permit it
	centerApi, err := a.CenterApiRepo.Get(ctx, g.Where("path = ?", path))
	if err != nil {
		// if not found, permit it
		if err == gorm.ErrRecordNotFound {
			return nil
		}
	}

	// find menu
	// if api not in menu, permit it
	has, err = a.CenterMenuApiRepo.Exists(ctx, "api_id = ?", centerApi.Id)
	if err != nil {
		return err
	}
	if !has {
		return nil
	}

	// find role api
	// if user has role api, permit it
	builder = sqx.Select("id").
		From("center_role_api").
		Where("role_id in (select role_id from center_user_role where user_id = ?)", user.UserId).
		Where("api_id = ?", centerApi.Id)
	has, err = gormx.Exists(ctx, a.DB, builder)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("no permission to access %s,please contact administrator", path)
	}
	return nil
}

func (a *AdminPermissionComponent) Middleware(skipPathMap map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if skipPathMap[c.Request.URL.Path] {
			c.Next()
			return
		}
		err := a.CheckAdminPermission(c.Request.Context(), c.Request)
		tail := time.Now().UnixMilli() % 100000
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    admin.ErrNoPermission,
				"message": fmt.Sprintf("%s[%d]", err.Error(), tail),
			})
			return
		}
		c.Next()
	}
}
