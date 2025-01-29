package model

import "time"

// api表

const (
	CenterApiStatusOn  = "ON"  // ON
	CenterApiStatusOff = "OFF" // OFF
)

type CenterApi struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id"`      //
	Status    string    `json:"status" gorm:"column:status"`         // ON,OFF
	Name      string    `json:"name" gorm:"column:name"`             // api名称
	Path      string    `json:"path" gorm:"column:path"`             //
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` // 更新时间
}

func (*CenterApi) TableName() string {
	return "center_api"
}

func (*CenterApi) PK() string {
	return "id"
}
