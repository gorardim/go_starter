package model

import "time"

// 用户token表

type UserToken struct {
	Id        int       `json:"id" gorm:"primaryKey;column:id"`      //
	UserId    int       `json:"user_id" gorm:"column:user_id"`       // 用户id
	Token     string    `json:"token" gorm:"column:token"`           // token
	Ip        string    `json:"ip" gorm:"column:ip"`                 // ip
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` //
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"` //
}

func (*UserToken) TableName() string {
	return "user_token"
}

func (*UserToken) PK() string {
	return "id"
}
