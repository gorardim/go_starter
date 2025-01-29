package bsc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_EditHotAddress(t *testing.T) {
	client := newClient()
	resp, err := client.EditHotAddress(context.Background(), "104371")
	assert.NoError(t, err)
	t.Log(resp)
}
