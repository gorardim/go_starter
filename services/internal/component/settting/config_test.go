package setting

import (
	"context"
	"testing"

	"app/pkg/gormx"
	"app/services/internal/repo"

	"github.com/stretchr/testify/assert"
)

func Test_setting_GetValue(t *testing.T) {
	db := gormx.NewTestDb(t)
	s := setting[string]{
		SettingRepo: repo.NewSettingRepo(db),
	}
	value, err := s.GetValue(context.Background(), "str")
	assert.NoError(t, err)
	assert.Equal(t, "str", value)
}
