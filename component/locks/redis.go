package locks

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

var ErrLockFailed = errors.New("lock failed")

type RedisLock struct {
	Redis *redis.Client
}

type RedisLockEntry struct {
	Redis *redis.Client
	Key   string
}

func (r *RedisLockEntry) Unlock() error {
	return r.Redis.Del(context.Background(), r.Key).Err()
}

func (r *RedisLock) Lock(ctx context.Context, key string, expire time.Duration) (*RedisLockEntry, error) {
	result, err := r.Redis.SetNX(ctx, key, 1, expire).Result()
	if err != nil {
		return nil, err
	}
	if !result {
		return nil, ErrLockFailed
	}
	return &RedisLockEntry{
		Redis: r.Redis,
		Key:   key,
	}, nil
}

func (r *RedisLock) Unlock(ctx context.Context, key string) error {
	return r.Redis.Del(ctx, key).Err()
}
