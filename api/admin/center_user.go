package admin

import "context"

// CenterUserServer 后台用户服务 (release)
type CenterUserServer interface {
	//x:api post /admin/center/user/list 用户列表
	List(ctx context.Context, req *CenterUserListRequest) (*CenterUserListResponse, error)
	//x:api post /admin/center/user/create 创建用户
	Create(ctx context.Context, req *CenterUserCreateRequest) (*CenterUserCreateResponse, error)
	//x:api post /admin/center/user/update 更新用户
	Update(ctx context.Context, req *CenterUserUpdateRequest) (*CenterUserUpdateResponse, error)
	//x:api post /admin/center/user/switch/status 切换用户状态
	SwitchStatus(ctx context.Context, req *CenterUserSwitchStatusRequest) (*CenterUserSwitchStatusResponse, error)
	//x:api post /admin/center/user/modify/password 修改密码
	ModifyPassword(ctx context.Context, req *CenterUserModifyPasswordRequest) (*CenterUserModifyPasswordResponse, error)
	//x:api post /admin/center/user/role/list 用户角色列表
	RoleList(ctx context.Context, req *CenterUserRoleListRequest) (*CenterUserRoleListResponse, error)
	//x:api post /admin/center/user/role/bind 用户角色绑定
	RoleBind(ctx context.Context, req *CenterUserRoleBindRequest) (*CenterUserRoleBindResponse, error)
	//x:api post /admin/center/user/role/unbind 用户角色解绑
	RoleUnbind(ctx context.Context, req *CenterUserRoleUnbindRequest) (*CenterUserRoleUnbindResponse, error)
}

type UserInfo struct {
	Id        int    `json:"id"`
	UserId    string `json:"user_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`
	UserType  string `json:"user_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CenterUserListRequest struct {
	Page     int    `json:"page"`
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}

type CenterUserListResponse struct {
	Total int         `json:"total"`
	Items []*UserInfo `json:"items"`
}

type CenterUserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	UserType string `json:"user_type"`
	Remark   string `json:"remark"`
}

type CenterUserCreateResponse struct{}

type CenterUserUpdateRequest struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
	UserType string `json:"user_type"`
}

type CenterUserUpdateResponse struct{}

type CenterUserSwitchStatusRequest struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type CenterUserSwitchStatusResponse struct{}

type CenterUserModifyPasswordRequest struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type CenterUserModifyPasswordResponse struct{}

type CenterUserRoleListRequest struct {
	UserId string `json:"user_id"`
}

type CenterUserRole struct {
	// 角色ID
	Id int `json:"id"`
	// 角色名称
	Name string `json:"name"`
	// 是否未超级管理员
	IsSuper string `json:"is_super"`
	// 角色状态
	Status string `json:"status"`
	// 创建时间
	CreatedAt string `json:"created_at"`
}

type CenterUserRoleListResponse struct {
	Items []*CenterUserRole `json:"items"`
}

type CenterUserRoleBindRequest struct {
	UserId string `json:"user_id"`
	RoleId int    `json:"role_id"`
}

type CenterUserRoleBindResponse struct{}

type CenterUserRoleUnbindRequest struct {
	UserId string `json:"user_id"`
	RoleId int    `json:"role_id"`
}

type CenterUserRoleUnbindResponse struct{}
