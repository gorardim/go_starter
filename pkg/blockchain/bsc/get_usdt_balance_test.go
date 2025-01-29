package bsc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetUsdtBalance(t *testing.T) {
	balance, err := newClient().GetUsdtBalance(context.Background())
	assert.NoError(t, err)
	t.Log(balance)
}
