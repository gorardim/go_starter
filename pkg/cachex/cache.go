package cachex

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
)

var DefaultCache = cache.New(cache.NoExpiration, cache.NoExpiration)

func Set[T any](key string, value T, expired time.Duration) {
	DefaultCache.Set(key, value, expired)
}

func Get[T any](key string) (T, bool) {
	v, has := DefaultCache.Get(key)
	if !has {
		var empty T
		return empty, false
	}
	if v1, ok := v.(T); ok {
		return v1, true
	}
	var empty T
	return empty, false
}

func GetOrSet[T any](key string, fn func() (T, error), expired time.Duration) (T, error) {
	v, has := Get[T](key)
	if has {
		return v, nil
	}
	v, err := fn()
	if err != nil {
		var empty T
		return empty, err
	}
	Set(key, v, expired)
	return v, nil
}

func ServiceCache[T, V any](ctx context.Context, req T, key string, fn func(ctx context.Context, req T) (V, error), expired time.Duration) (V, error) {
	v, has := Get[V](key)
	if has {
		return v, nil
	}
	v, err := fn(ctx, req)
	if err != nil {
		var empty V
		return empty, err
	}
	Set(key, v, expired)
	return v, nil
}
