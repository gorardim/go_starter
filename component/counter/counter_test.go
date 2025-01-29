package counter

import (
	"context"
	"testing"
	"time"

	"app/pkg/redisx"
	"github.com/stretchr/testify/assert"
)

func TestCounter_Incr(t *testing.T) {
	c := &Counter{
		Client: redisx.NewTestRedis(),
	}
	incr, err := c.Incr(context.Background(), "a", time.Second*10)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), incr)

	incr, err = c.Incr(context.Background(), "a", time.Second*10)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), incr)
}
