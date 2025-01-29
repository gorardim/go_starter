package cache

import (
	"app/services/internal/component/lang"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type Home struct {
	*redis.Client
}

const (
	HomePrefix   = "travel:home:"
	HomeBanner   = HomePrefix + "banner"
	HomeKingKong = HomePrefix + "king_kong"
	HomeMarquee  = HomePrefix + "marquee"
	HomeGif      = HomePrefix + "gif"
	HomeNotice   = HomePrefix + "notice"
)

func (h *Home) Key(ctx context.Context, key string) string {
	return key + ":" + lang.FromContext(ctx)
}

func (h *Home) HasCache(ctx context.Context, key string) (int64, error) {
	return h.Exists(ctx, h.Key(ctx, key)).Result()
}

func (h *Home) SetCache(ctx context.Context, key string, result interface{}) error {
	data, err := json.Marshal(result)
	if err != nil {
		return err
	}
	return h.Set(ctx, h.Key(ctx, key), data, time.Hour*24).Err()
}

func (h *Home) GetCache(ctx context.Context, key string, result interface{}) error {
	data, err := h.Get(ctx, h.Key(ctx, key)).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

func (h *Home) DelCache(ctx context.Context, key ...string) error {
	keys := make([]string, 0, len(key))
	for _, v := range key {
		itemKeys, err := h.Keys(ctx, v+"*").Result()
		if err != nil {
			return err
		}
		keys = append(keys, itemKeys...)
	}
	if len(keys) == 0 {
		return nil
	}
	return h.Del(ctx, keys...).Err()
}

func (h *Home) ClearAllCache(ctx context.Context) error {
	keys, err := h.Keys(ctx, HomePrefix+"*").Result()
	if err != nil {
		return err
	}
	return h.Del(ctx, keys...).Err()
}
