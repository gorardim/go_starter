package model

import (
	"time"
)

var _ time.Time

// CenterRoleApi 角色api表
type CenterRoleApi struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id"`      //
	RoleId    int       `json:"role_id" gorm:"column:role_id"`       // 角色id
	ApiId     int       `json:"api_id" gorm:"column:api_id"`         // api id
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

func (*CenterRoleApi) TableName() string {
	return "center_role_api"
}

func (*CenterRoleApi) PK() string {
	return "id"
}
