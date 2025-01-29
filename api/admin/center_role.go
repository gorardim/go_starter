package admin

import "context"

// CenterRoleServer 角色服务 (release)
type CenterRoleServer interface {
	//x:api post /admin/center/role/list 角色列表
	List(ctx context.Context, req *CenterRoleListRequest) (*CenterRoleListResponse, error)
	//x:api post /admin/center/role/all 角色列表
	All(ctx context.Context, req *CenterRoleAllRequest) (*CenterRoleAllResponse, error)
	//x:api post /admin/center/role/create 创建用户角色
	Create(ctx context.Context, req *CenterRoleCreateRequest) (*CenterRoleCreateResponse, error)
	//x:api post /admin/center/role/update 更新用户角色
	Update(ctx context.Context, req *CenterRoleUpdateRequest) (*CenterRoleUpdateResponse, error)
	//x:api post /admin/center/role/switch/status 切换用户角色状态
	SwitchStatus(ctx context.Context, req *CenterRoleSwitchStatusRequest) (*CenterRoleSwitchStatusResponse, error)
	//x:api post /admin/center/role/bind/menu/list 用户角色绑定菜单列表
	BindMenuList(ctx context.Context, req *CenterRoleBindMenuListRequest) (*CenterRoleBindMenuListResponse, error)
	//x:api post /admin/center/role/bind/menu/update 更新用户角色绑定菜单
	BindMenuUpdate(ctx context.Context, req *CenterRoleBindMenuUpdateRequest) (*CenterRoleBindMenuUpdateResponse, error)
	//x:api post /admin/center/role/bind/api/list 用户角色绑定api列表
	BindApiList(ctx context.Context, req *CenterRoleBindApiListRequest) (*CenterRoleBindApiListResponse, error)
	//x:api post /admin/center/role/bind/api/update 更新用户角色绑定api
	BindApiUpdate(ctx context.Context, req *CenterRoleBindApiUpdateRequest) (*CenterRoleBindApiUpdateResponse, error)
}

type CenterRoleInfo struct {
	Id int `json:"id"`
	// 角色名称
	Name string `json:"name"`
	// 状态
	Status string `json:"status"`
	// 是否为超级管理员
	IsSuper string `json:"is_super"`
	// 创建时间
	CreatedAt string `json:"created_at"`
}

type CenterRoleCreateRequest struct {
	// 角色名称
	Name string `json:"name"`
}

type CenterRoleCreateResponse struct{}

type CenterRoleUpdateRequest struct {
	Id int `json:"id"`
	// 角色名称
	Name string `json:"name"`
}

type CenterRoleUpdateResponse struct{}

type CenterRoleSwitchStatusRequest struct {
	Id int `json:"id"`
	// 状态
	Status string `json:"status"`
}

type CenterRoleSwitchStatusResponse struct{}

type CenterRoleBindMenuListRequest struct {
	// 角色id
	RoleId int `json:"role_id"`
}

type CenterRoleBindMenuListResponse struct {
	MenuIdList []int `json:"menu_id_list"`
}

type CenterRoleBindMenuUpdateRequest struct {
	// 角色id
	RoleId int `json:"role_id"`
	// 菜单id列表
	MenuIdList []int `json:"menu_id_list"`
}

type CenterRoleBindMenuUpdateResponse struct{}

type CenterRoleBindApiListRequest struct {
	// 角色id
	RoleId int `json:"role_id"`
}

type CenterRoleBindApiListResponse struct {
	ApiIdList []int `json:"api_id_list"`
}

type CenterRoleBindApiUpdateRequest struct {
	// 角色id
	RoleId int `json:"role_id"`
	// api id列表
	ApiIdList []int `json:"api_id_list"`
}

type CenterRoleBindApiUpdateResponse struct{}

type CenterRoleListRequest struct {
	Page int `json:"page"`
}
type CenterRoleListResponse struct {
	Total int               `json:"total"`
	Items []*CenterRoleInfo `json:"items"`
}

type CenterRoleAllRequest struct{}

type CenterRoleAllItem struct {
	Id int `json:"id"`
	// 角色名称
	Name string `json:"name"`
}

type CenterRoleAllResponse struct {
	Items []*CenterRoleAllItem `json:"items"`
}
