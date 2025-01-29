package errx

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("test is", func(t *testing.T) {
		e1 := errors.New("error")
		err := New("SystemError", e1)
		assert.True(t, errors.Is(err, e1))
		assert.Equal(t, "error", err.Error())
		assert.Equal(t, "SystemError", err.ErrorCode())
	})

	t.Run("sql no row", func(t *testing.T) {
		err := New("NoRows", sql.ErrNoRows)
		assert.True(t, errors.Is(err, sql.ErrNoRows))
		assert.Equal(t, "sql: no rows in result set", err.Error())
		// as
		assert.True(t, errors.As(err, &sql.ErrNoRows))
	})

	t.Run("fmt wrap", func(t *testing.T) {
		err := New("SystemError", fmt.Errorf("wrap sql error: %w", sql.ErrNoRows))
		assert.True(t, errors.Is(err, sql.ErrNoRows))
		assert.Equal(t, "wrap sql error: sql: no rows in result set", err.Error())
		// as
		assert.True(t, errors.As(err, &sql.ErrNoRows))
	})
}
