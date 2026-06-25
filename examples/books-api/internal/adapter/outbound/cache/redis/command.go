package redisadapter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	conn *redis.Client
}

func (c *Client) Exist(ctx context.Context, key string) bool {
	return c.conn.Exists(context.Background(), key).Val() == 1
}

func (c *Client) SetEx(ctx context.Context, key string, value any, expiry time.Duration) error {
	return c.conn.Set(context.Background(), key, value, expiry).Err()
}

func (c *Client) Set(ctx context.Context, key string, value any) error {
	return c.conn.Set(context.Background(), key, value, 0).Err()
}

func (c *Client) GetString(ctx context.Context, key string) (string, error) {
	return c.conn.Get(context.Background(), key).Result()
}
