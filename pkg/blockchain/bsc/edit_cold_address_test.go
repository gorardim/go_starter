package bsc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_EditColdAddress(t *testing.T) {
	client := newClient()
	resp, err := client.EditColdAddress(context.Background(), "0x0d84FEcfD34D52b9Fc62906c3410BadBE5b645fF", "642504")
	assert.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}
