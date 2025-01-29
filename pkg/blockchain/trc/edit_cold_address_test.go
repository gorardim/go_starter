package trc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_EditColdAddress(t *testing.T) {
	client := newClient()
	resp, err := client.EditColdAddress(context.Background(), "TES1BN51WnKsshWLbTsqSmFYf5ajZGPVR4", "985335")
	assert.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}
