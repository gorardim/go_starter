package admin

import (
	"app/api/admin"
	"app/api/model"
	"app/pkg/errx"
	"app/pkg/gormx"
	"app/pkg/gormx/g"
	"app/pkg/sqx"
	"app/services/internal/component"
	"app/services/internal/repo"
	"context"
	"time"

	"gorm.io/gorm"
)

var _ admin.CenterUserServer = (*CenterUser)(nil)

type CenterUser struct {
	CenterUserRepo     *repo.CenterUserRepo
	CenterUserRoleRepo *repo.CenterUserRoleRepo
	UserAuthComponent  *component.UserAuthComponent
	DB                 *gorm.DB
}

type RoleListItem struct {
	Id        int       `json:"id"`
	UserId    string    `json:"user_id"`
	RoleId    int       `json:"role_id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	IsSuper   string    `json:"is_super"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *CenterUser) RoleList(ctx context.Context, req *admin.CenterUserRoleListRequest) (*admin.CenterUserRoleListResponse, error) {
	builder := sqx.Select("ur.*,r.name,r.status,r.is_super").
		From("center_user_role ur").
		From("left join center_role r on ur.role_id = r.id").
		Where("ur.user_id = ?", req.UserId)

	all, err := gormx.FindAll[RoleListItem](ctx, u.DB, builder)
	if err != nil {
		return nil, err
	}
	items := make([]*admin.CenterUserRole, 0)
	for _, v := range all {
		items = append(items, &admin.CenterUserRole{
			Id:        v.RoleId,
			Name:      v.Name,
			IsSuper:   v.IsSuper,
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &admin.CenterUserRoleListResponse{
		Items: items,
	}, nil
}

func (u *CenterUser) RoleBind(ctx context.Context, req *admin.CenterUserRoleBindRequest) (*admin.CenterUserRoleBindResponse, error) {
	// 先判断是否已经绑定
	exists, err := u.CenterUserRoleRepo.Exists(ctx, "user_id = ? and role_id = ?", req.UserId, req.RoleId)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	if exists {
		return nil, errx.New(admin.ErrInvalidParam, "角色已经绑定")
	}
	err = u.CenterUserRoleRepo.Create(ctx, &model.CenterUserRole{
		UserId: req.UserId,
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterUserRoleBindResponse{}, nil
}

func (u *CenterUser) RoleUnbind(ctx context.Context, req *admin.CenterUserRoleUnbindRequest) (*admin.CenterUserRoleUnbindResponse, error) {
	_, err := u.CenterUserRoleRepo.Delete(ctx, "user_id = ? and role_id = ?", req.UserId, req.RoleId)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterUserRoleUnbindResponse{}, nil
}

func (u *CenterUser) List(ctx context.Context, req *admin.CenterUserListRequest) (*admin.CenterUserListResponse, error) {
	paginate, err := u.CenterUserRepo.Paginate(ctx, gormx.Order("id desc"), req.Page, 10)
	if err != nil {
		return nil, err
	}
	items := make([]*admin.UserInfo, 0)
	for _, v := range paginate.List {
		items = append(items, &admin.UserInfo{
			Id:        v.Id,
			UserId:    v.UserId,
			Username:  v.Username,
			Nickname:  v.Nickname,
			Remark:    v.Remark,
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &admin.CenterUserListResponse{
		Total: int(paginate.Total),
		Items: items,
	}, nil
}

func (u *CenterUser) Create(ctx context.Context, req *admin.CenterUserCreateRequest) (*admin.CenterUserCreateResponse, error) {
	userId := u.UserAuthComponent.GenerateUserId()
	err := u.CenterUserRepo.Create(ctx, &model.CenterUser{
		Username: req.Username,
		Nickname: req.Nickname,
		UserId:   userId,
		Password: u.UserAuthComponent.HashPassword(req.Password, userId),
		Status:   model.CenterUserStatusOn,
		Remark:   req.Remark,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterUserCreateResponse{}, nil
}

func (u *CenterUser) Update(ctx context.Context, req *admin.CenterUserUpdateRequest) (*admin.CenterUserUpdateResponse, error) {
	_, err := u.CenterUserRepo.UpdateById(ctx, &model.CenterUser{
		Id:       req.Id,
		Nickname: req.Nickname,
		Remark:   req.Remark,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterUserUpdateResponse{}, nil
}

func (u *CenterUser) SwitchStatus(ctx context.Context, req *admin.CenterUserSwitchStatusRequest) (*admin.CenterUserSwitchStatusResponse, error) {
	_, err := u.CenterUserRepo.UpdateById(ctx, &model.CenterUser{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterUserSwitchStatusResponse{}, nil
}

func (u *CenterUser) ModifyPassword(ctx context.Context, req *admin.CenterUserModifyPasswordRequest) (*admin.CenterUserModifyPasswordResponse, error) {
	// 先查询
	user, err := u.CenterUserRepo.Get(ctx, g.Where("id = ?", req.Id))
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	// 再修改
	_, err = u.CenterUserRepo.UpdateById(ctx, &model.CenterUser{
		Id:       req.Id,
		Password: u.UserAuthComponent.HashPassword(req.Password, user.UserId),
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterUserModifyPasswordResponse{}, nil
}
