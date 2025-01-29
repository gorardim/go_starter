package admin

import (
	"context"
	"fmt"
	"time"

	"app/api/admin"
	"app/api/model"
	"app/pkg/gormx"
	"app/pkg/sqx"
	"app/services/internal/component/lang"
	"app/services/internal/repo"

	"github.com/samber/lo"
)

var _ admin.UserServer = (*User)(nil)

type User struct {
	UserRepo *repo.UserRepo
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
	// 性别
	Sex string `json:"sex"`
	// bsc地址
	BscAddress string `json:"bsc_address"`
	// bsc uid
	BscUid string `json:"bsc_uid"`
	// trc20地址
	TrcAddress string `json:"trc_address"`
	// 余额
	Balance float64 `json:"balance"`
	// 收益
	Income float64 `json:"income"`
	// 积分
	Point float64 `json:"point"`
	// 免手续费提现额度
	FreeFeeWithdrawAmount float64 `json:"free_fee_withdraw_amount"`
	// 是否有效
	IsValid string `json:"is_valid"`
	// 是否俱乐部拥有者: Y:是 N:否
	IsClubOwner string `json:"is_club_owner"`
	// KYC状态:未认证UNAUTH,已认证AUTH
	KycStatus string `json:"kyc_status"`
	// 是否开启用户等级定级
	EnableLevelGrade string `json:"enable_level_grade"`
	// 代理级别
	Level int `json:"level"`
	// LevelName
	LevelName model.LangType `json:"level_name"`
	// 投资金额
	InvestAmount float64 `json:"invest_amount"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at"`
	// 是内部号
	IsInternal string `json:"is_internal"`
}

type UserSafeQuestionListItem struct {
	// 问题id
	QuestionId int `json:"question_id"`
	// 问题
	Question model.LangType `json:"question"`
	// 答案
	Answer string `json:"answer"`
}

func (u *User) ResetPayPassword(ctx context.Context, req *admin.UserResetPayPasswordRequest) (*admin.UserResetPayPasswordResponse, error) {

	user, err := u.UserRepo.Find(ctx, "user_id = ?", req.UserId)
	if err != nil {
		return nil, err
	}

	if _, err := u.UserRepo.UpdateMap(ctx, map[string]interface{}{
		"pay_password": "",
	}, "user_id =? ", user.UserId); err != nil {
		return nil, err
	}

	return &admin.UserResetPayPasswordResponse{}, nil
}

func (u *User) Reset2fa(ctx context.Context, req *admin.UserReset2faRequest) (*admin.UserReset2faResponse, error) {

	user, err := u.UserRepo.Find(ctx, "user_id = ?", req.UserId)
	if err != nil {
		return nil, err
	}

	if _, err := u.UserRepo.UpdateMap(ctx, map[string]interface{}{
		"secret_two_fa": "",
	}, "user_id =? ", user.UserId); err != nil {
		return nil, err
	}
	return &admin.UserReset2faResponse{}, nil
}

func (u *User) List(ctx context.Context, req *admin.UserListRequest) (*admin.UserListResponse, error) {

	// query
	builder := sqx.Select(`u.*, ul.name as level_name, (SELECT p.nickname FROM user p WHERE u.pid = p.user_id) as parent_nickname,(SELECT p.invite_code FROM user p WHERE u.pid = p.user_id) as parent_invite_code,(SELECT p.username FROM user p WHERE u.pid = p.user_id) as parent_username,(SELECT p.level FROM user p WHERE u.pid = p.user_id) as parent_level, 
		(SELECT
			CASE
			WHEN COUNT(*) > 0 THEN 'Y'
			ELSE 'N'
			END AS is_internal FROM user_tree ut WHERE ut.parent_id = 1 and ut.user_id = u.user_id) as is_internal`).
		From(`user u`).
		LeftJoin(`user_level ul ON u.level = ul.level`).
		WhereIf(req.Status != "", `u.status = ?`, req.Status).
		WhereIf(req.Username != "", `u.username = ?`, req.Username).
		WhereIf(req.Nickname != "", `u.nickname = ?`, req.Nickname).
		WhereIf(req.IsValid != "", `u.is_valid = ?`, req.IsValid).
		WhereIf(req.IsClubOwner != "", `u.is_club_owner = ?`, req.IsClubOwner).
		WhereIf(req.Pid != 0, `u.pid = ?`, req.Pid).
		WhereIf(req.InviteCode != "", `u.invite_code = ?`, req.InviteCode).
		WhereIf(req.CustomerId != "", `u.customer_id = ?`, req.CustomerId).
		WhereIf(req.EnableLevelGrade != "", `u.enable_level_grade = ?`, req.EnableLevelGrade).
		WhereIf(req.PointRange[1] != 0, "u.point BETWEEN ? AND ?", req.PointRange[0], req.PointRange[1]).
		WhereIf(req.InvestAmountRange[1] != 0, "u.invest_amount BETWEEN ? AND ?", req.InvestAmountRange[0], req.InvestAmountRange[1]).
		WhereIf(req.BalanceRange[1] != 0, "u.balance BETWEEN ? AND ?", req.BalanceRange[0], req.BalanceRange[1]).
		WhereIf(req.IncomeRange[1] != 0, "u.income BETWEEN ? AND ?", req.IncomeRange[0], req.IncomeRange[1]).
		WhereIf(req.FreeFeeWithdrawAmountRange[1] != 0, "u.free_fee_withdraw_amount BETWEEN ? AND ?", req.FreeFeeWithdrawAmountRange[0], req.FreeFeeWithdrawAmountRange[1]).
		WhereIf(req.Level != 0, `u.level = ?`, req.Level)
	if req.UserId != 0 {
		builder.Where(`u.user_id = ?`, req.UserId)
	}

	paginate, err := gormx.Paginate[UserListItem](ctx, u.UserRepo.DB(), builder, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	items := make([]*admin.UserListItem, 0)
	for _, item := range paginate.List {
		items = append(items, &admin.UserListItem{
			UserId:                item.UserId,
			CustomerId:            item.CustomerId,
			InviteCode:            item.InviteCode,
			Status:                item.Status,
			Username:              item.Username,
			Nickname:              item.Nickname,
			Pid:                   item.Pid,
			ParentNickname:        item.ParentNickname,
			ParentUsername:        item.ParentUsername,
			ParentInviteCode:      item.ParentInviteCode,
			ParentLevel:           item.ParentLevel,
			Avatar:                item.Avatar,
			Balance:               fmt.Sprintf("%.2f", item.Balance),
			Income:                fmt.Sprintf("%.2f", item.Income),
			Point:                 fmt.Sprintf("%.2f", item.Point),
			FreeFeeWithdrawAmount: fmt.Sprintf("%.2f", item.FreeFeeWithdrawAmount),
			IsValid:               item.IsValid,
			IsClubOwner:           item.IsClubOwner,
			EnableLevelGrade:      item.EnableLevelGrade,
			Level:                 item.Level,
			LevelName:             lang.FromLangType(ctx, item.LevelName),
			InvestAmount:          fmt.Sprintf("%.2f", item.InvestAmount),
			IsInternal:            item.IsInternal,
			CreatedAt:             item.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:             item.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &admin.UserListResponse{
		Total: int(paginate.Total),
		Items: items,
	}, nil
}

func (u *User) SafeQuestion(ctx context.Context, req *admin.UserSafeQuestionRequest) (*admin.UserSafeQuestionResponse, error) {

	user, err := u.UserRepo.Find(ctx, "user_id = ?", req.UserId)
	if err != nil {
		return nil, err
	}

	// query
	builder := sqx.Select(`question_id, usq.answer, sq.question`).
		From(`user_safe_question usq`).
		LeftJoin(`safe_question sq ON usq.question_id = sq.id`).
		Where(`usq.user_id = ?`, user.UserId)

	userQuestions, err := gormx.FindAll[UserSafeQuestionListItem](ctx, u.UserRepo.DB(), builder)
	if err != nil {
		return nil, err
	}

	items := make([]*admin.UserSafeQuestionListItem, 0)
	for _, item := range userQuestions {
		items = append(items, &admin.UserSafeQuestionListItem{
			QuestionId: item.QuestionId,
			Question:   item.Question.Data,
			Answer:     item.Answer,
		})
	}

	return &admin.UserSafeQuestionResponse{
		UserId:                user.UserId,
		Username:              user.Username,
		IsSetPayPassword:      lo.Ternary[string](user.PayPassword != "", "Y", "N"),
		IsEnableTwoFactorAuth: lo.Ternary[string](user.SecretTwoFA != "", "Y", "N"),
		Items:                 items,
	}, nil
}

func (u *User) UpdateStatus(ctx context.Context, req *admin.UserUpdateStatusRequest) (*admin.UserUpdateStatusResponse, error) {

	user, err := u.UserRepo.Find(ctx, "user_id = ?", req.UserId)
	if err != nil {
		return nil, err
	}

	if user.Status == req.Status {
		return &admin.UserUpdateStatusResponse{}, nil
	}

	if _, err := u.UserRepo.UpdateMap(ctx, map[string]interface{}{
		"status":     req.Status,
		"updated_at": time.Now(),
	}, "user_id = ?", user.UserId); err != nil {
		return nil, err
	}

	return &admin.UserUpdateStatusResponse{}, nil
}

func (u *User) Detail(ctx context.Context, req *admin.UserDetailRequest) (*admin.UserDetailResponse, error) {

	builder := sqx.Select(`u.*, (SELECT p.nickname FROM user p WHERE u.pid = p.user_id) as parent_nickname,
	(SELECT
		CASE
		WHEN COUNT(*) > 0 THEN 'Y'
		ELSE 'N'
		END AS is_internal FROM user_tree ut WHERE ut.parent_id = 1 and ut.user_id = u.user_id) as is_internal`).From(`user u`).
		Where(`u.user_id = ?`, req.UserId)

	user, err := gormx.Find[UserListItem](ctx, u.UserRepo.DB(), builder)
	if err != nil {
		return nil, err
	}

	return &admin.UserDetailResponse{
		UserId: req.UserId,
		User: &admin.UserListItem{
			UserId:                user.UserId,
			Pid:                   user.Pid,
			InviteCode:            user.InviteCode,
			CustomerId:            user.CustomerId,
			Status:                user.Status,
			Username:              user.Username,
			Nickname:              user.Nickname,
			ParentNickname:        user.ParentNickname,
			Avatar:                user.Avatar,
			Balance:               fmt.Sprintf("%.2f", user.Balance),
			Income:                fmt.Sprintf("%.2f", user.Income),
			Point:                 fmt.Sprintf("%.2f", user.Point),
			FreeFeeWithdrawAmount: fmt.Sprintf("%.2f", user.FreeFeeWithdrawAmount),
			IsValid:               user.IsValid,
			IsClubOwner:           user.IsClubOwner,
			IsInternal:            user.IsInternal,
			EnableLevelGrade:      user.EnableLevelGrade,
			Level:                 user.Level,
			InvestAmount:          fmt.Sprintf("%.2f", user.InvestAmount),
			CreatedAt:             user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:             user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
