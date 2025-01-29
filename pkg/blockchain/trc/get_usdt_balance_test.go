package trc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetUsdtBalance(t *testing.T) {
	balance, err := newClient().GetUsdtBalance(context.Background(), "178efcf524d0f4bf086577aa51fa10b8")
	assert.NoError(t, err)
	t.Log(balance)
}
