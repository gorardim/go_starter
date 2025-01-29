package model

import "time"

// 用户角色表

type CenterUserRole struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id"`      //
	UserId    string    `json:"user_id" gorm:"column:user_id"`       // 用户id
	RoleId    int       `json:"role_id" gorm:"column:role_id"`       // 角色id
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

func (*CenterUserRole) TableName() string {
	return "center_user_role"
}

func (*CenterUserRole) PK() string {
	return "id"
}
