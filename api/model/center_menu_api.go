package model

import "time"

// 菜单api表

type CenterMenuApi struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id"`      //
	MenuId    int       `json:"menu_id" gorm:"column:menu_id"`       // 菜单id
	ApiId     int       `json:"api_id" gorm:"column:api_id"`         // api id
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

func (*CenterMenuApi) TableName() string {
	return "center_menu_api"
}

func (*CenterMenuApi) PK() string {
	return "id"
}
