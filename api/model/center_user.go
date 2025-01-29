package model

import "time"

// 用户表

const (
	CenterUserStatusOn  = "ON"  // 用户状态
	CenterUserStatusOff = "OFF" // OFF
)

type CenterUser struct {
	Id          int       `json:"id" gorm:"primaryKey;column:id"`            // 主键
	Username    string    `json:"username" gorm:"column:username"`           // 用户名称
	UserId      string    `json:"user_id" gorm:"column:user_id"`             // 用户id
	Nickname    string    `json:"nickname" gorm:"column:nickname"`           //
	Remark      string    `json:"remark" gorm:"column:remark"`               //
	Status      string    `json:"status" gorm:"column:status"`               // 用户状态ON,OFF
	Password    string    `json:"password" gorm:"column:password"`           // 密码
	SecretTwoFA string    `json:"secret_two_fa" gorm:"column:secret_two_fa"` // 二次验证密钥
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`       // 创建时间
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`       // 更新时间
}

func (*CenterUser) TableName() string {
	return "center_user"
}

func (*CenterUser) PK() string {
	return "id"
}
