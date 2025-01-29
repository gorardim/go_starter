package admin

import (
	"app/api/model"
	"context"
)

// CenterMenuServer 菜单服务 (release)
type CenterMenuServer interface {
	//x:api post /admin/center/menu/all 菜单列表
	All(ctx context.Context, req *CenterMenuAllRequest) (*CenterMenuAllResponse, error)
	//x:api post /admin/center/menu/detail 菜单详情
	Detail(ctx context.Context, req *CenterMenuDetailRequest) (*CenterMenuDetailResponse, error)
	//x:api post /admin/center/menu/create 创建菜单
	Create(ctx context.Context, req *CenterMenuCreateRequest) (*CenterMenuCreateResponse, error)
	//x:api post /admin/center/menu/update 更新菜单
	Update(ctx context.Context, req *CenterMenuUpdateRequest) (*CenterMenuUpdateResponse, error)
	//x:api post /admin/center/menu/switch/status 切换菜单状态
	SwitchStatus(ctx context.Context, req *CenterMenuSwitchStatusRequest) (*CenterMenuSwitchStatusResponse, error)
	//x:api post /admin/center/menu/bind/api/create 绑定菜单API
	BindApiCreate(ctx context.Context, req *CenterMenuBindApiCreateRequest) (*CenterMenuBindApiCreateResponse, error)
	//x:api post /admin/center/menu/bind/api/list 绑定菜单api列表
	BindApiList(ctx context.Context, req *CenterMenuBindApiListRequest) (*CenterMenuBindApiListResponse, error)
	//x:api post /admin/center/menu/bind/api/all 绑定菜单api列表
	BindApiAll(ctx context.Context, req *CenterMenuBindApiAllRequest) (*CenterMenuBindApiAllResponse, error)
	//x:api post /admin/center/menu/bind/api/delete 删除绑定菜单api
	BindApiDelete(ctx context.Context, req *CenterMenuBindApiDeleteRequest) (*CenterMenuBindApiDeleteResponse, error)
}

type CenterMenuInfo struct {
	Id        int        `json:"id"`
	Pid       int        `json:"pid"`
	Name      string     `json:"name"`
	NameJson  model.Lang `json:"name_json"`
	Path      string     `json:"path"`
	Status    string     `json:"status"`
	Icon      string     `json:"icon"`
	SortNum   int        `json:"sort_num"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}
type CenterMenuDetailRequest struct {
	Id int `json:"id" binding:"required"`
}
type CenterMenuDetailResponse struct {
	Id        int        `json:"id"`
	Pid       int        `json:"pid"`
	Name      string     `json:"name"`
	NameJson  model.Lang `json:"name_json"`
	Path      string     `json:"path"`
	Status    string     `json:"status"`
	Icon      string     `json:"icon"`
	SortNum   int        `json:"sort_num"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

type CenterMenuAllRequest struct{}
type CenterMenuAllResponse struct {
	Items []*CenterMenuInfo `json:"items"`
}

type CenterMenuCreateRequest struct {
	Pid      int        `json:"pid"`
	NameJson model.Lang `json:"name_json"`
	Path     string     `json:"path"`
	Icon     string     `json:"icon"`
	AppId    string     `json:"app_id"`
	SortNum  int        `json:"sort_num"`
}

type CenterMenuCreateResponse struct{}

type CenterMenuUpdateRequest struct {
	Id       int        `json:"id"`
	Pid      int        `json:"pid"`
	NameJson model.Lang `json:"name_json"`
	Path     string     `json:"path"`
	Icon     string     `json:"icon"`
	SortNum  int        `json:"sort_num"`
}

type CenterMenuUpdateResponse struct{}

type CenterMenuSwitchStatusRequest struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type CenterMenuSwitchStatusResponse struct{}

type CenterMenuBindApiCreateRequest struct {
	MenuId int   `json:"menu_id"`
	ApiIds []int `json:"api_ids"`
}

type CenterMenuBindApiCreateResponse struct{}

type CenterMenuBindApiListRequest struct {
	MenuId int `json:"menu_id"`
}

type CenterMenuBindApiInfo struct {
	Id       int    `json:"id"`
	MenuId   int    `json:"menu_id"`
	ApiId    int    `json:"api_id"`
	Path     string `json:"path"`
	Name     string `json:"name"`
	CreateAt string `json:"create_at"`
}

type CenterMenuBindApiListResponse struct {
	Items []*CenterMenuBindApiInfo `json:"items"`
}

type CenterMenuBindApiDeleteRequest struct {
	MenuId int `json:"menu_id"`
	ApiId  int `json:"api_id"`
}

type CenterMenuBindApiDeleteResponse struct{}

type CenterMenuBindApiAllRequest struct{}

type CenterMenuBindApiAllResponse struct {
	Items []*CenterMenuBindApiInfo `json:"items"`
}
