package admin

import (
	"app/api/admin"
	"app/api/model"
	"app/pkg/errx"
	"app/pkg/gormx"
	"app/services/internal/repo"
	"context"
)

var _ admin.CenterApiServer = (*CenterApi)(nil)

type CenterApi struct {
	CenterApiRepo *repo.CenterApiRepo
}

func (c *CenterApi) List(ctx context.Context, req *admin.CenterApiListRequest) (*admin.CenterApiListResponse, error) {
	paginate, err := c.CenterApiRepo.Paginate(ctx, gormx.WhereIf(req.Name != "", "name like ?", "%"+req.Name+"%").
		Order("id desc"), req.Page, 10)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	items := make([]*admin.CenterApiInfo, 0)
	for _, v := range paginate.List {
		items = append(items, &admin.CenterApiInfo{
			Id:        v.Id,
			Name:      v.Name,
			Path:      v.Path,
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &admin.CenterApiListResponse{
		Items: items,
		Total: int(paginate.Total),
	}, nil
}

func (c *CenterApi) All(ctx context.Context, req *admin.CenterApiAllRequest) (*admin.CenterApiAllResponse, error) {
	list, err := c.CenterApiRepo.FindAllBy(ctx, gormx.WhereIf(req.Status != "", "status = ?", req.Status))
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	items := make([]*admin.CenterApiAllItem, 0)
	for _, v := range list {
		items = append(items, &admin.CenterApiAllItem{
			Id:   v.Id,
			Name: v.Name,
			Path: v.Path,
		})
	}
	return &admin.CenterApiAllResponse{
		Items: items,
	}, nil
}

func (c *CenterApi) Create(ctx context.Context, req *admin.CenterApiCreateRequest) (*admin.CenterApiCreateResponse, error) {
	err := c.CenterApiRepo.Create(ctx, &model.CenterApi{
		Name:   req.Name,
		Path:   req.Path,
		Status: model.CenterApiStatusOn,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterApiCreateResponse{}, nil
}

func (c *CenterApi) Update(ctx context.Context, req *admin.CenterApiUpdateRequest) (*admin.CenterApiUpdateResponse, error) {
	_, err := c.CenterApiRepo.UpdateById(ctx, &model.CenterApi{
		Id:   req.Id,
		Name: req.Name,
		Path: req.Path,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterApiUpdateResponse{}, nil
}

func (c *CenterApi) SwitchStatus(ctx context.Context, req *admin.CenterApiSwitchStatusRequest) (*admin.CenterApiSwitchStatusResponse, error) {
	_, err := c.CenterApiRepo.UpdateById(ctx, &model.CenterApi{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err)
	}
	return &admin.CenterApiSwitchStatusResponse{}, nil
}
