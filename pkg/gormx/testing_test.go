package gormx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type userModel struct {
	Id        int `gorm:"primary_key"`
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*userModel) TableName() string {
	return "user"
}

func (*userModel) PK() string {
	return "id"
}

func TestNewTestDb(t *testing.T) {
	db := NewTestDb(t)
	var databases []string
	scan := db.Raw("show databases").Scan(&databases)
	assert.NoError(t, scan.Error)
	assert.Contains(t, databases, "mysql")

	u1 := &userModel{}
	first := db.First(u1)
	assert.NoError(t, first.Error)
	assert.Equal(t, "user1", u1.Username)
}
