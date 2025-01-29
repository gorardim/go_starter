package awssdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSession(t *testing.T) {
	session, err := NewSession()
	assert.NoError(t, err)
	t.Log(session)
}
