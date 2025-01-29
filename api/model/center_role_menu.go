package model

import "time"

// 角色菜单表

type CenterRoleMenu struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id"`      //
	RoleId    int       `json:"role_id" gorm:"column:role_id"`       // 角色id
	MenuId    int       `json:"menu_id" gorm:"column:menu_id"`       // 菜单id
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

func (*CenterRoleMenu) TableName() string {
	return "center_role_menu"
}

func (*CenterRoleMenu) PK() string {
	return "id"
}
