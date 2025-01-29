package admin

import "context"

// CenterApiServer 后台api (release)
type CenterApiServer interface {
	//x:api post /admin/center/api/list 后台api列表
	List(ctx context.Context, req *CenterApiListRequest) (*CenterApiListResponse, error)
	//x:api post /admin/center/api/all 后台api列表
	All(ctx context.Context, req *CenterApiAllRequest) (*CenterApiAllResponse, error)
	//x:api post /admin/center/api/create 创建后台api
	Create(ctx context.Context, req *CenterApiCreateRequest) (*CenterApiCreateResponse, error)
	//x:api post /admin/center/api/update 更新后台api
	Update(ctx context.Context, req *CenterApiUpdateRequest) (*CenterApiUpdateResponse, error)
	//x:api post /admin/center/api/switch/status 切换后台api状态
	SwitchStatus(ctx context.Context, req *CenterApiSwitchStatusRequest) (*CenterApiSwitchStatusResponse, error)
}

type CenterApiInfo struct {
	Id int `json:"id"`
	// api名称
	Name string `json:"name"`
	// api路径
	Path string `json:"path"`
	// 状态
	Status string `json:"status"`
	// 创建时间
	CreatedAt string `json:"created_at"`
}

type CenterApiListRequest struct {
	Page int `json:"page"`
	// api名称
	Name string `json:"name"`
}

type CenterApiListResponse struct {
	Total int              `json:"total"`
	Items []*CenterApiInfo `json:"items"`
}

type CenterApiCreateRequest struct {
	// api名称
	Name string `json:"name"`
	// api路径
	Path string `json:"path"`
}

type CenterApiCreateResponse struct{}

type CenterApiUpdateRequest struct {
	Id int `json:"id"`
	// api名称
	Name string `json:"name"`
	// api路径
	Path string `json:"path"`
}

type CenterApiUpdateResponse struct{}

type CenterApiSwitchStatusRequest struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type CenterApiSwitchStatusResponse struct{}

type CenterApiAllRequest struct {
	// 状态
	Status string `json:"status"`
}

type CenterApiAllItem struct {
	Id int `json:"id"`
	// api名称
	Name string `json:"name"`
	// api路径
	Path string `json:"path"`
}

type CenterApiAllResponse struct {
	Items []*CenterApiAllItem `json:"items"`
}
