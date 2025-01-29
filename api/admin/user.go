package admin

import (
	"app/api/model"
	"context"
)

// UserServer
type UserServer interface {
	//x:api post /admin/user/list 用户列表 user list
	List(ctx context.Context, in *UserListRequest) (*UserListResponse, error)
	//x:api post /admin/user/update_status 更新用户状态 user update status
	UpdateStatus(ctx context.Context, in *UserUpdateStatusRequest) (*UserUpdateStatusResponse, error)
	//x:api post /admin/user/detail 用户详情 user detail
	Detail(ctx context.Context, in *UserDetailRequest) (*UserDetailResponse, error)
	//x:api post /admin/user/safe_question 用户安全问题详情 user safe question detail
	SafeQuestion(ctx context.Context, in *UserSafeQuestionRequest) (*UserSafeQuestionResponse, error)
	//x:api post /admin/user/reset_pay_password 重置支付密码 reset pay password
	ResetPayPassword(ctx context.Context, in *UserResetPayPasswordRequest) (*UserResetPayPasswordResponse, error)
	//x:api post /admin/user/reset_2fa 重置2fa reset 2fa
	Reset2fa(ctx context.Context, in *UserReset2faRequest) (*UserReset2faResponse, error)
}

type UserReset2faRequest struct {
	// 用户id
	UserId int `json:"user_id" form:"user_id" binding:"required"`
}

type UserReset2faResponse struct {
}

type UserResetPayPasswordRequest struct {
	// 用户id
	UserId int `json:"user_id" form:"user_id" binding:"required"`
}

type UserResetPayPasswordResponse struct {
}

type UserSafeQuestionRequest struct {
	// 用户id
	UserId int `json:"user_id" form:"user_id" binding:"required"`
}

type UserSafeQuestionResponse struct {
	// 用户id
	UserId int `json:"user_id"`
	// 用户名
	Username string `json:"username"`
	// 是否设置支付密码 Y:是 N:否
	IsSetPayPassword string `json:"is_set_pay_password"`
	// 是否开启两步验证 Y:是 N:否
	IsEnableTwoFactorAuth string `json:"is_enable_two_factor_auth"`
	// 问题列表
	Items []*UserSafeQuestionListItem `json:"items"`
}

type UserSafeQuestionListItem struct {
	// 问题id
	QuestionId int `json:"question_id"`
	// 问题
	Question model.Lang `json:"question"`
	// 答案
	Answer string `json:"answer"`
}

type UserUpdateStatusRequest struct {
	// 用户id
	UserId int `json:"user_id" form:"user_id" binding:"required"`
	// 状态:启用ON,禁用OFF
	Status string `json:"status" form:"status" binding:"required"`
}

type UserUpdateStatusResponse struct {
}

type UserDetailRequest struct {
	// 用户id
	UserId int `json:"user_id" form:"user_id" binding:"required"`
}

type UserDetailResponse struct {
	// 用户id
	UserId int `json:"user_id"`
	// User
	User *UserListItem `json:"user"`
}

type UserListRequest struct {
	// 用户id
	UserId int `json:"user_id"`
	// 用户名
	Username string `json:"username"`
	// 昵称
	Nickname string `json:"nickname"`
	// 状态:启用ON,禁用OFF
	Status string `json:"status"`
	// 是否开启用户等级定级
	EnableLevelGrade string `json:"enable_level_grade"`
	// 父级id
	Pid int `json:"pid"`
	// 代理级别
	Level int `json:"level"`
	// 是否有效
	IsValid string `json:"is_valid"`
	// 是否俱乐部拥有者: Y:是 N:否
	IsClubOwner string `json:"is_club_owner"`
	// 邀请码
	InviteCode string `json:"invite_code"`
	// custom id
	CustomerId string `json:"customer_id"`
	// PointRange
	PointRange [2]int `json:"point_range"`
	// InvestAmountRange
	InvestAmountRange [2]int `json:"invest_amount_range"`
	// 余额
	BalanceRange [2]int `json:"balance_range"`
	// 收益
	IncomeRange [2]int `json:"income_range"`
	// 免手续费提现额度
	FreeFeeWithdrawAmountRange [2]int `json:"free_fee_withdraw_amount_range"`
	// 页码 page
	Page int `json:"page"`
	// 每页数量 page_size
	PageSize int `json:"page_size"`
}

type UserListResponse struct {
	// total 总数
	Total int `json:"total"`
	// items 列表
	Items []*UserListItem `json:"items"`
}

type UserListItem struct {
	// 用户id
	UserId int `json:"user_id"`
	// 父级id
	Pid int `json:"pid"`
	// custom id
	CustomerId string `json:"customer_id"`
	// 邀请码
	InviteCode string `json:"invite_code"`
	// 状态:启用ON,禁用OFF
	Status string `json:"status"`
	// 用户名 Email
	Username string `json:"username"`
	// 昵称
	Nickname string `json:"nickname"`
	// ParentNickname
	ParentNickname string `json:"parent_nickname"`
	// ParentUsername
	ParentUsername string `json:"parent_username"`
	// ParentInviteCode
	ParentInviteCode string `json:"parent_invite_code"`
	// ParentLevel
	ParentLevel int `json:"parent_level"`
	// 头像
	Avatar string `json:"avatar"`
	// 余额
	Balance string `json:"balance"`
	// 收益
	Income string `json:"income"`
	// 积分
	Point string `json:"point"`
	// 免手续费提现额度
	FreeFeeWithdrawAmount string `json:"free_fee_withdraw_amount"`
	// 是否有效
	IsValid string `json:"is_valid"`
	// 是否俱乐部拥有者: Y:是 N:否
	IsClubOwner string `json:"is_club_owner"`
	// 是否开启用户等级定级
	EnableLevelGrade string `json:"enable_level_grade"`
	// 代理级别
	Level int `json:"level"`
	// LevelName
	LevelName string `json:"level_name"`
	// 投资金额
	InvestAmount string `json:"invest_amount"`
	// 是内部号: Y:是 N:否
	IsInternal string `json:"is_internal"`
	// 创建时间
	CreatedAt string `json:"created_at"`
	// 更新时间
	UpdatedAt string `json:"updated_at"`
}
