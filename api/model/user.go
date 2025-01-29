package model

import "time"

// 用户表

const (
	UserStatusOn  = "ON"  // 启用
	UserStatusOff = "OFF" // 禁用

	UserKycStatusUnauth = "UNAUTH" // 未认证
	UserKycStatusAuth   = "AUTH"   // 已认证
)

type User struct {
	UserId                int       `json:"user_id" gorm:"primaryKey;column:user_id"`                        // 用户id
	CustomerId            string    `json:"customer_id" gorm:"column:customer_id"`                           // custom id
	Pid                   int       `json:"pid" gorm:"column:pid"`                                           // 父级id
	IsClubOwner           string    `json:"is_club_owner" gorm:"column:is_club_owner"`                       // 是否俱乐部拥有者: Y:是 N:否
	InviteCode            string    `json:"invite_code" gorm:"column:invite_code"`                           // 邀请码
	Status                string    `json:"status" gorm:"column:status"`                                     // 状态:启用ON,禁用OFF
	Username              string    `json:"username" gorm:"column:username"`                                 // 用户名
	Nickname              string    `json:"nickname" gorm:"column:nickname"`                                 // 昵称
	Avatar                string    `json:"avatar" gorm:"column:avatar"`                                     // 头像
	Sex                   string    `json:"sex" gorm:"column:sex"`                                           // 性别
	Password              string    `json:"password" gorm:"column:password"`                                 // 密码
	PayPassword           string    `json:"pay_password" gorm:"column:pay_password"`                         // 支付密码
	SecretTwoFA           string    `json:"secret_two_fa" gorm:"column:secret_two_fa"`                       // 二次验证密钥
	BscAddress            string    `json:"bsc_address" gorm:"column:bsc_address"`                           // bsc地址
	BscPlainAddress       string    `json:"bsc_plain_address" gorm:"column:bsc_plain_address"`               // bsc明文地址
	BscUid                string    `json:"bsc_uid" gorm:"column:bsc_uid"`                                   // bsc uid
	TrcUid                string    `json:"trc_uid" gorm:"column:trc_uid"`                                   // trc uid
	TrcAddress            string    `json:"trc_address" gorm:"column:trc_address"`                           // trc20地址
	TrcPlainAddress       string    `json:"trc_plain_address" gorm:"column:trc_plain_address"`               // trc20明文地址
	Balance               float64   `json:"balance" gorm:"column:balance"`                                   // 余额
	Income                float64   `json:"income" gorm:"column:income"`                                     // 收益
	Point                 float64   `json:"point" gorm:"column:point"`                                       // 积分
	FreeFeeWithdrawAmount float64   `json:"free_fee_withdraw_amount" gorm:"column:free_fee_withdraw_amount"` // 免手续费提现额度
	IsValid               string    `json:"is_valid" gorm:"column:is_valid"`                                 // 是否有效
	KycStatus             string    `json:"kyc_status" gorm:"column:kyc_status"`                             // KYC状态:未认证UNAUTH,已认证AUTH
	EnableLevelGrade      string    `json:"enable_level_grade" gorm:"column:enable_level_grade"`             // 是否开启用户等级定级
	Level                 int       `json:"level" gorm:"column:level"`                                       // 代理级别
	InvestAmount          float64   `json:"invest_amount" gorm:"column:invest_amount"`                       // 投资金额
	CreatedAt             time.Time `json:"created_at" gorm:"column:created_at"`                             // 创建时间
	UpdatedAt             time.Time `json:"updated_at" gorm:"column:updated_at"`                             // 更新时间
}

func (*User) TableName() string {
	return "user"
}

func (*User) PK() string {
	return "user_id"
}
