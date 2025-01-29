package model

import "time"

// 角色表

const (
	CenterRoleStatusOn  = "ON"  // 角色状态
	CenterRoleStatusOff = "OFF" // OFF
)

type CenterRole struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id"`      //
	Name      string    `json:"name" gorm:"column:name"`             // 角色名称
	Status    string    `json:"status" gorm:"column:status"`         // 角色状态ON,OFF
	IsSuper   string    `json:"is_super" gorm:"column:is_super"`     // 是否为超级管理员Y,N
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

func (*CenterRole) TableName() string {
	return "center_role"
}

func (*CenterRole) PK() string {
	return "id"
}
