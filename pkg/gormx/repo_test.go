package gormx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepo_Exists(t *testing.T) {
	db := NewTestDb(t)
	userRepo := NewRepo[userModel](db)
	exist, err := userRepo.Exists(context.Background(), "id > ?", 0)
	assert.NoError(t, err)
	assert.True(t, exist)
}

func TestRepo_Find(t *testing.T) {
	db := NewTestDb(t)
	userRepo := NewRepo[userModel](db)
	user, err := userRepo.Find(context.Background(), "id = ?", 1)
	assert.NoError(t, err)
	assert.Equal(t, "user1", user.Username)
}

func TestRepo_FindOption(t *testing.T) {
	db := NewTestDb(t)
	userRepo := NewRepo[userModel](db)
	user, err := userRepo.FindBy(context.Background(), Where("id = ?", 1).Order("id desc").Limit(1))
	assert.NoError(t, err)
	assert.Equal(t, "user1", user.Username)

	paginate, err := userRepo.Paginate(context.Background(), Where("id > ?", 0).Order("id desc"), 1, 10)
	assert.NoError(t, err)
	t.Log(paginate)
}

func TestRepo_SumN(t *testing.T) {
	db := NewTestDb(t)
	userRepo := NewRepo[userModel](db)
	sums, err := userRepo.SumN(context.Background(), []string{"id", "balance"}, "id > ?", 0)
	assert.NoError(t, err)
	t.Log(sums)
}
