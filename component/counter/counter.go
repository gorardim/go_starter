package counter

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type Counter struct {
	Client *redis.Client
}

func (c *Counter) Incr(ctx context.Context, key string, expired time.Duration) (int64, error) {
	pipeline := c.Client.Pipeline()
	pipeline.SetNX(ctx, key, 0, expired)
	pipeline.Incr(ctx, key)
	exec, err := pipeline.Exec(ctx)
	if err != nil {
		return 0, err
	}
	last := exec[len(exec)-1].(*redis.IntCmd)
	return last.Val(), last.Err()
}

func (c *Counter) IncrBy(ctx context.Context, key string, value int64, expired time.Duration) (int64, error) {
	pipeline := c.Client.Pipeline()
	pipeline.SetNX(ctx, key, 0, expired)
	pipeline.IncrBy(ctx, key, value)
	exec, err := pipeline.Exec(ctx)
	if err != nil {
		return 0, err
	}
	last := exec[len(exec)-1].(*redis.IntCmd)
	return last.Val(), last.Err()
}

func (c *Counter) Get(ctx context.Context, key string) (int64, error) {
	n, err := c.Client.Get(ctx, key).Int64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}
		return 0, err
	}
	return n, nil
}
