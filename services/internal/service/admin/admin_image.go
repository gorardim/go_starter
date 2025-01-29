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

var _ admin.AdminImageServer = (*AdminImage)(nil)

type AdminImage struct {
	AdminImagesRepo *repo.AdminImageRepo
}

func (a *AdminImage) List(ctx context.Context, req *admin.AdminImageListRequest) (*admin.AdminImageListResponse, error) {
	builder := sqx.Select("*").
		From("admin_image").
		WhereIf(req.Name != "", "name like ?", "%"+req.Name+"%").
		OrderBy("id DESC")

	pagination, err := gormx.Paginate[model.AdminImage](ctx, a.AdminImagesRepo.DB(), builder, req.Page, 10)
	if err != nil {
		return nil, err
	}
	return &admin.AdminImageListResponse{
		Total: int(pagination.Total),
		Items: lo.Map(pagination.List, func(item *model.AdminImage, index int) *admin.AdminImageListItem {
			return &admin.AdminImageListItem{
				Id:      item.Id,
				Name:    item.Name,
				HttpUrl: item.Url,
			}
		}),
	}, nil
}
