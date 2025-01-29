package cachex

import (
	"time"

	"golang.org/x/sync/singleflight"
)

type Flight[T any] struct {
	singleflight.Group
}

func (g *Flight[T]) Do(key string, fn func() (T, error), expired time.Duration) (T, error) {
	v, has := DefaultCache.Get(key)
	if has {
		if v1, ok := v.(T); ok {
			return v1, nil
		}
	}
	v1, err, _ := g.Group.Do(key, func() (interface{}, error) {
		return fn()
	})
	if err != nil {
		DefaultCache.Set(key, v1, expired)
	}
	return v.(T), err
}
