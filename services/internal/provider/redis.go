package provider

import (
	"fmt"

	"app/services/internal/config"

	"github.com/go-redis/redis/v8"
)

func NewRedis(conf *config.Config) *redis.Client {
	value, ok := conf.Redis["default"]
	if !ok {
		panic("redis config not found")
	}

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", value.Host, value.Port),
		Password: value.Password,
		DB:       value.Db,
	})
}
