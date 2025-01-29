package redisx

import "github.com/go-redis/redis/v8"

type Client struct {
	*redis.Client
}

func NewClient(rds *redis.Client) *Client {
	return &Client{Client: rds}
}

func (c *Client) GetAndDel(key string) (string, error) {
	pipe := c.Pipeline()
	pipe.Get(c.Context(), key)
	pipe.Del(c.Context(), key)
	exec, err := pipe.Exec(c.Context())
	if err != nil {
		return "", err
	}
	return exec[0].(*redis.StringCmd).Val(), nil
}
