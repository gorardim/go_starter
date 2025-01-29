package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	password := "123"
	hashed, err := GeneratePassword(password)
	assert.NoError(t, err)
	t.Logf("password: %s", hashed)
	assert.NoError(t, ValidatePassword(password, string(hashed)))
}
