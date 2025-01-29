package admin

import (
	"app/api/admin"
	"app/api/model"
	"app/pkg/gormx"
	"app/pkg/sqx"
	"app/services/internal/repo"
	"context"

	"github.com/samber/lo"
)

var _ admin.AdminVideoServer = (*AdminVideo)(nil)

type AdminVideo struct {
	AdminVideoRepo *repo.AdminVideoRepo
}

func (a *AdminVideo) List(ctx context.Context, req *admin.AdminVideoListRequest) (*admin.AdminVideoListResponse, error) {
	builder := sqx.Select("*").
		From("admin_video").
		WhereIf(req.Name != "", "name like ?", "%"+req.Name+"%").
		OrderBy("id DESC")

	pagination, err := gormx.Paginate[model.AdminVideo](ctx, a.AdminVideoRepo.DB(), builder, req.Page, 10)
	if err != nil {
		return nil, err
	}
	return &admin.AdminVideoListResponse{
		Total: int(pagination.Total),
		Items: lo.Map(pagination.List, func(item *model.AdminVideo, index int) *admin.AdminVideoListItem {
			return &admin.AdminVideoListItem{
				Id:      item.Id,
				Name:    item.Name,
				HttpUrl: item.Url,
			}
		}),
	}, nil
}
