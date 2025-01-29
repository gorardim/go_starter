package admin

import (
	"app/api/admin"
	"app/api/model"
	"app/pkg/errx"
	"app/pkg/gormx"
	"app/pkg/gormx/g"
	"app/pkg/sqx"
	"app/services/internal/component/lang"
	"app/services/internal/repo"
	"context"
	"time"

	"gorm.io/gorm"
)

var _ admin.CenterMenuServer = (*CenterMenu)(nil)

type CenterMenu struct {
	CenterAppMenuRepo    *repo.CenterMenuRepo
	CenterAppMenuApiRepo *repo.CenterMenuApiRepo
	DB                   *gorm.DB
}

func (m *CenterMenu) All(ctx context.Context, req *admin.CenterMenuAllRequest) (*admin.CenterMenuAllResponse, error) {
	all, err := m.CenterAppMenuRepo.FindAllBy(ctx, gormx.
		Order("sort_num asc"))
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	items := make([]*admin.CenterMenuInfo, 0)
	for _, v := range all {
		items = append(items, &admin.CenterMenuInfo{
			Id:        v.Id,
			Pid:       v.Pid,
			NameJson:  v.NameJson.Data,
			Name:      lang.FromLangType(ctx, v.NameJson),
			Status:    v.Status,
			SortNum:   v.SortNum,
			Path:      v.Path,
			Icon:      v.Icon,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &admin.CenterMenuAllResponse{
		Items: items,
	}, nil
}
func (m *CenterMenu) Detail(ctx context.Context, req *admin.CenterMenuDetailRequest) (*admin.CenterMenuDetailResponse, error) {
	menu, err := m.CenterAppMenuRepo.GetById(ctx, req.Id)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}

	return &admin.CenterMenuDetailResponse{
		Id:        menu.Id,
		Pid:       menu.Pid,
		NameJson:  menu.NameJson.Data,
		Name:      lang.FromLangType(ctx, menu.NameJson),
		Status:    menu.Status,
		SortNum:   menu.SortNum,
		Path:      menu.Path,
		Icon:      menu.Icon,
		CreatedAt: menu.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: menu.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (m *CenterMenu) Create(ctx context.Context, req *admin.CenterMenuCreateRequest) (*admin.CenterMenuCreateResponse, error) {
	err := m.CenterAppMenuRepo.Create(ctx, &model.CenterMenu{
		Pid:      req.Pid,
		Status:   model.CenterMenuStatusOn,
		Path:     req.Path,
		NameJson: gormx.JsonObject[model.Lang]{Data: req.NameJson},
		Icon:     req.Icon,
		SortNum:  req.SortNum,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	return &admin.CenterMenuCreateResponse{}, nil
}

func (m *CenterMenu) Update(ctx context.Context, req *admin.CenterMenuUpdateRequest) (*admin.CenterMenuUpdateResponse, error) {
	_, err := m.CenterAppMenuRepo.UpdateById(ctx, &model.CenterMenu{
		Id:       req.Id,
		Pid:      req.Pid,
		Path:     req.Path,
		NameJson: gormx.JsonObject[model.Lang]{Data: req.NameJson},
		Icon:     req.Icon,
		SortNum:  req.SortNum,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	return &admin.CenterMenuUpdateResponse{}, nil
}

func (m *CenterMenu) SwitchStatus(ctx context.Context, req *admin.CenterMenuSwitchStatusRequest) (*admin.CenterMenuSwitchStatusResponse, error) {
	_, err := m.CenterAppMenuRepo.UpdateById(ctx, &model.CenterMenu{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	return &admin.CenterMenuSwitchStatusResponse{}, nil
}

func (m *CenterMenu) BindApiCreate(ctx context.Context, req *admin.CenterMenuBindApiCreateRequest) (*admin.CenterMenuBindApiCreateResponse, error) {
	for _, apiId := range req.ApiIds {
		// 先查询
		data, err := m.CenterAppMenuApiRepo.Get(ctx, g.Where("menu_id = ? and api_id = ?", req.MenuId, apiId))
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errx.New(admin.ErrInvalidParam, err.Error())
		}
		if data != nil {
			continue
		}
		err = m.CenterAppMenuApiRepo.Create(ctx, &model.CenterMenuApi{
			MenuId: req.MenuId,
			ApiId:  apiId,
		})
		if err != nil {
			return nil, errx.New(admin.ErrInvalidParam, err.Error())
		}
	}
	return &admin.CenterMenuBindApiCreateResponse{}, nil
}

type BindApiListItem struct {
	Id       int       `json:"id"`
	MenuId   int       `json:"menu_id"`
	ApiId    int       `json:"api_id"`
	Path     string    `json:"path"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
}

func (m *CenterMenu) BindApiList(ctx context.Context, req *admin.CenterMenuBindApiListRequest) (*admin.CenterMenuBindApiListResponse, error) {
	builder := sqx.Select("ma.*,a.name,a.path").
		From("center_menu_api ma").
		From("left join center_api a on ma.api_id = a.id").
		Where("ma.menu_id = ?", req.MenuId)

	all, err := gormx.FindAll[BindApiListItem](ctx, m.DB, builder)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	items := make([]*admin.CenterMenuBindApiInfo, 0)
	for _, v := range all {
		items = append(items, &admin.CenterMenuBindApiInfo{
			Id:       v.Id,
			MenuId:   v.MenuId,
			ApiId:    v.ApiId,
			Path:     v.Path,
			Name:     v.Name,
			CreateAt: v.CreateAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &admin.CenterMenuBindApiListResponse{
		Items: items,
	}, nil
}

func (m *CenterMenu) BindApiAll(ctx context.Context, req *admin.CenterMenuBindApiAllRequest) (*admin.CenterMenuBindApiAllResponse, error) {
	builder := sqx.Select("ma.*,a.name,a.path").
		From("center_menu_api ma").
		From("left join center_api a on ma.api_id = a.id")
	all, err := gormx.FindAll[BindApiListItem](ctx, m.DB, builder)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	items := make([]*admin.CenterMenuBindApiInfo, 0)
	for _, v := range all {
		items = append(items, &admin.CenterMenuBindApiInfo{
			Id:       v.Id,
			MenuId:   v.MenuId,
			ApiId:    v.ApiId,
			Path:     v.Path,
			Name:     v.Name,
			CreateAt: v.CreateAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &admin.CenterMenuBindApiAllResponse{
		Items: items,
	}, nil
}

func (m *CenterMenu) BindApiDelete(ctx context.Context, req *admin.CenterMenuBindApiDeleteRequest) (*admin.CenterMenuBindApiDeleteResponse, error) {
	_, err := m.CenterAppMenuApiRepo.Delete(ctx, "menu_id = ? and api_id = ?", req.MenuId, req.ApiId)
	if err != nil {
		return nil, errx.New(admin.ErrInvalidParam, err.Error())
	}
	return &admin.CenterMenuBindApiDeleteResponse{}, nil
}
