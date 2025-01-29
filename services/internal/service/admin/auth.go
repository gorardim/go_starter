package admin

import (
	"app/api/api"
	"app/pkg/gormx"
	"app/pkg/gormx/g"
	"app/services/internal/component/lang"
	"app/services/internal/config"
	"context"
	"strings"

	"github.com/mojocn/base64Captcha"

	"app/api/admin"
	"app/api/model"
	"app/pkg/errx"
	"app/pkg/sqx"
	"app/services/internal/component"
	"app/services/internal/repo"

	"gorm.io/gorm"
)

var _ admin.AuthServer = (*Auth)(nil)

type Auth struct {
	CenterUserRepo     *repo.CenterUserRepo
	CenterRoleRepo     *repo.CenterRoleRepo
	CenterMenuRepo     *repo.CenterMenuRepo
	CenterUserRoleRepo *repo.CenterUserRoleRepo
	UserAuthComponent  *component.UserAuthComponent
	DB                 *gorm.DB
	CaptchaCode        *base64Captcha.Captcha
	Config             *config.Config
}

func (o *Auth) AuthMenus(ctx context.Context, req *admin.AuthMenusRequest) (*admin.AuthMenusResponse, error) {
	user, ok := component.UserFormContext(ctx)
	if !ok {
		return nil, errx.New(api.ErrInvalidParam, "user not found in context")
	}

	// if user is super admin, return all menus
	builder := sqx.Select("r.id").
		From("center_user_role ur").
		LeftJoin("center_role r on ur.role_id = r.id").
		Where("ur.user_id = ?", user.UserId).
		Where("r.is_super = ?", "Y")
	has, err := gormx.Exists(ctx, o.DB, builder)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	if has {
		all, err := o.CenterMenuRepo.FindAllBy(ctx, gormx.
			Order("sort_num asc"))
		if err != nil {
			return nil, errx.New(admin.ErrInvalidParam, err.Error())
		}
		items := make([]*admin.CenterMenuInfo, 0, len(all))
		for _, v := range all {
			items = append(items, &admin.CenterMenuInfo{
				Id:        v.Id,
				Pid:       v.Pid,
				Name:      lang.FromLangType(ctx, v.NameJson),
				Status:    v.Status,
				SortNum:   v.SortNum,
				Path:      v.Path,
				Icon:      v.Icon,
				CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		return &admin.AuthMenusResponse{
			Items: items,
		}, nil
	}

	builder = sqx.Select("*").
		From("center_menu").
		Where("id in (select distinct menu_id from center_role_menu where role_id in (select role_id from center_user_role where user_id = ?))", user.UserId).
		OrderBy("sort_num")

	all, err := gormx.FindAll[model.CenterMenu](ctx, o.DB, builder)
	if err != nil {
		return nil, errx.New(api.ErrInvalidParam, err.Error())
	}
	items := make([]*admin.CenterMenuInfo, 0, len(all))
	for _, v := range all {
		items = append(items, &admin.CenterMenuInfo{
			Id:        v.Id,
			Pid:       v.Pid,
			Name:      lang.FromLangType(ctx, v.NameJson),
			Status:    v.Status,
			SortNum:   v.SortNum,
			Path:      v.Path,
			Icon:      v.Icon,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &admin.AuthMenusResponse{
		Items: items,
	}, nil
}

func (o *Auth) Captcha(ctx context.Context, _ *admin.AuthCaptchaRequest) (*admin.AuthCaptchaResponse, error) {
	id, b64s, err := o.CaptchaCode.Generate()
	if err != nil {
		return nil, errx.New(api.ErrBusiness, "failedToGenerateVerificationCode")
	}

	return &admin.AuthCaptchaResponse{
		CaptchaId:    id,
		CaptchaImage: b64s,
	}, nil
}

func (o *Auth) AuthLogin(ctx context.Context, req *admin.AuthLoginRequest) (*admin.AuthLoginResponse, error) {
	if o.Config.Env != "local" {
		if !o.CaptchaCode.Verify(req.CaptchaId, strings.ToLower(req.VerifyCode), true) {
			return nil, errx.New(admin.ErrInvalidParam, "verifyCodeError")
		}
	}
	// 验证密码
	user, err := o.CenterUserRepo.Get(ctx, g.Where("username = ?", req.Username))
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, "usernameOrPasswordError")
	}
	hashPassword := o.UserAuthComponent.HashPassword(req.Password, user.UserId)
	if hashPassword != user.Password {
		return nil, errx.New(admin.ErrInvalidParam, "usernameOrPasswordError")
	}
	// check status
	if user.Status != model.CenterUserStatusOn {
		return nil, errx.New(admin.ErrInvalidParam, "userIsForbidden")
	}

	// 获取角色
	builder := sqx.Select("count(*) as n").
		From("center_user_role ur").
		From("left join center_role r on r.id = ur.role_id").
		Where("ur.user_id = ?", user.UserId)

	sql, args := builder.Build()
	var sum int
	if err = o.DB.Raw(sql, args...).Scan(&sum).Error; err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}

	if sum == 0 {
		return nil, errx.New(admin.ErrInvalidParam, "user has no role, please contact the administrator")
	}

	// 生成jwt
	token, err := o.UserAuthComponent.GenerateToken(user.UserId)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}

	return &admin.AuthLoginResponse{
		Token:    token,
		UserId:   user.UserId,
		Username: user.Username,
		Nickname: user.Nickname,
	}, nil
}
