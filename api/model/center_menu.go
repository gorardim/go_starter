package model

import "time"

// 应用菜单表

const (
	CenterMenuStatusOn  = "ON"  // 用户状态
	CenterMenuStatusOff = "OFF" // OFF
)

type CenterMenu struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id"`      // 主键
	Pid       int       `json:"pid" gorm:"column:pid"`               // 父级id
	Path      string    `json:"path" gorm:"column:path"`             // 菜单路径
	Status    string    `json:"status" gorm:"column:status"`         //
	NameJson  LangType  `json:"name_json" gorm:"column:name_json"`   // 菜单名称
	Icon      string    `json:"icon" gorm:"column:icon"`             // 菜单图标
	SortNum   int       `json:"sort_num" gorm:"column:sort_num"`     // 排序
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

func (*CenterMenu) TableName() string {
	return "center_menu"
}

func (*CenterMenu) PK() string {
	return "id"
}
