package setting

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"app/services/internal/repo"
)

type setting[T any] struct {
	SettingRepo *repo.SettingRepo
}

func (c *setting[T]) GetValue(ctx context.Context, key string) (T, error) {
	var v T
	m, err := c.SettingRepo.Find(ctx, "`key` = ?", key)
	if err != nil {
		if err == sql.ErrNoRows {
			return v, fmt.Errorf("setting key %s not found", key)
		}
		return v, err
	}
	if err = unmarshal(m.Value, &v); err != nil {
		return v, err
	}
	return v, nil
}

func (c *setting[T]) SetValue(ctx context.Context, key string, value T) error {
	// 读取全局配置表中内容
	itemValue, err := marshal(value)
	if err != nil {
		return err
	}
	_, err = c.SettingRepo.UpdateMap(ctx, map[string]any{
		"value": itemValue,
	}, "`key` = ?", key)
	return err
}

func unmarshal(s string, v interface{}) error {
	switch d := v.(type) {
	case *string:
		*d = s
		return nil
	default:
		return json.Unmarshal([]byte(s), v)
	}
}

func marshal(v interface{}) (string, error) {
	switch v.(type) {
	case string:
		return v.(string), nil
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
}
