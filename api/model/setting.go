package model

import "time"

type Setting struct {
	Id          int       `json:"id" gorm:"primaryKey;column:id"`        //
	Key         string    `json:"key" gorm:"column:key"`                 //
	Value       string    `json:"value" gorm:"column:value"`             //
	Description string    `json:"description" gorm:"column:description"` // 描述
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`   //
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`   //
}

func (*Setting) TableName() string {
	return "setting"
}

func (*Setting) PK() string {
	return "id"
}
