package trc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_EditHotAddress(t *testing.T) {
	client := newClient()
	resp, err := client.EditHotAddress(context.Background(), "467394")
	assert.NoError(t, err)
	t.Log(resp)
}
