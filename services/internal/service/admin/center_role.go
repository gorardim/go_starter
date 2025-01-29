package admin

import (
	"app/api/admin"
	"app/api/model"
	"app/pkg/errx"
	"app/pkg/gormx"
	"app/services/internal/repo"
	"context"
)

var _ admin.CenterRoleServer = (*CenterRole)(nil)

type CenterRole struct {
	CenterRoleRepo     *repo.CenterRoleRepo
	CenterRoleMenuRepo *repo.CenterRoleMenuRepo
	CenterRoleApiRepo  *repo.CenterRoleApiRepo
}

func (c *CenterRole) All(ctx context.Context, req *admin.CenterRoleAllRequest) (*admin.CenterRoleAllResponse, error) {
	find, err := c.CenterRoleRepo.FindAllBy(ctx, gormx.
		Where("status = ?", model.CenterRoleStatusOn))
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	items := make([]*admin.CenterRoleAllItem, 0)
	for _, v := range find {
		items = append(items, &admin.CenterRoleAllItem{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	return &admin.CenterRoleAllResponse{
		Items: items,
	}, nil
}

func (c *CenterRole) List(ctx context.Context, req *admin.CenterRoleListRequest) (*admin.CenterRoleListResponse, error) {
	paginate, err := c.CenterRoleRepo.Paginate(ctx, gormx.Order("id desc"), req.Page, 10)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	items := make([]*admin.CenterRoleInfo, 0)
	for _, v := range paginate.List {
		items = append(items, &admin.CenterRoleInfo{
			Id:        v.Id,
			Name:      v.Name,
			Status:    v.Status,
			IsSuper:   v.IsSuper,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &admin.CenterRoleListResponse{
		Total: int(paginate.Total),
		Items: items,
	}, nil
}

func (c *CenterRole) Create(ctx context.Context, req *admin.CenterRoleCreateRequest) (*admin.CenterRoleCreateResponse, error) {
	err := c.CenterRoleRepo.Create(ctx, &model.CenterRole{
		Name:    req.Name,
		IsSuper: model.No,
		Status:  model.CenterRoleStatusOn,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterRoleCreateResponse{}, nil
}

func (c *CenterRole) Update(ctx context.Context, req *admin.CenterRoleUpdateRequest) (*admin.CenterRoleUpdateResponse, error) {
	_, err := c.CenterRoleRepo.UpdateById(ctx, &model.CenterRole{
		Id:   req.Id,
		Name: req.Name,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterRoleUpdateResponse{}, nil
}

func (c *CenterRole) SwitchStatus(ctx context.Context, req *admin.CenterRoleSwitchStatusRequest) (*admin.CenterRoleSwitchStatusResponse, error) {
	_, err := c.CenterRoleRepo.UpdateById(ctx, &model.CenterRole{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterRoleSwitchStatusResponse{}, nil
}

func (c *CenterRole) BindMenuList(ctx context.Context, req *admin.CenterRoleBindMenuListRequest) (*admin.CenterRoleBindMenuListResponse, error) {
	appRoleList, err := c.CenterRoleMenuRepo.FindAllBy(ctx, gormx.Where("role_id = ?", req.RoleId))
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	items := make([]int, 0)
	for _, v := range appRoleList {
		items = append(items, v.MenuId)
	}
	return &admin.CenterRoleBindMenuListResponse{
		MenuIdList: items,
	}, nil
}

func (c *CenterRole) BindMenuUpdate(ctx context.Context, req *admin.CenterRoleBindMenuUpdateRequest) (*admin.CenterRoleBindMenuUpdateResponse, error) {
	// 先删除
	_, err := c.CenterRoleMenuRepo.Delete(ctx, "role_id = ?", req.RoleId)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	for _, menuId := range req.MenuIdList {
		// 再添加
		err = c.CenterRoleMenuRepo.Create(ctx, &model.CenterRoleMenu{
			RoleId: req.RoleId,
			MenuId: menuId,
		})
	}
	return &admin.CenterRoleBindMenuUpdateResponse{}, nil
}

func (c *CenterRole) BindApiList(ctx context.Context, req *admin.CenterRoleBindApiListRequest) (*admin.CenterRoleBindApiListResponse, error) {
	find, err := c.CenterRoleApiRepo.FindAllBy(ctx, gormx.Where("role_id = ?", req.RoleId))
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	items := make([]int, 0)
	for _, v := range find {
		items = append(items, v.ApiId)
	}
	return &admin.CenterRoleBindApiListResponse{
		ApiIdList: items,
	}, nil
}

func (c *CenterRole) BindApiUpdate(ctx context.Context, req *admin.CenterRoleBindApiUpdateRequest) (*admin.CenterRoleBindApiUpdateResponse, error) {
	// 先删除
	_, err := c.CenterRoleApiRepo.Delete(ctx, "role_id = ?", req.RoleId)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	for _, apiId := range req.ApiIdList {
		// 再添加
		err = c.CenterRoleApiRepo.Create(ctx, &model.CenterRoleApi{
			RoleId: req.RoleId,
			ApiId:  apiId,
		})
	}
	return &admin.CenterRoleBindApiUpdateResponse{}, nil
}
