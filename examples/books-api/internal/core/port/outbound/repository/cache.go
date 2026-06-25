package repository

import (
	"context"
	"time"
)

type CacheRepository interface {
	Exist(ctx context.Context, key string) bool
	SetEx(ctx context.Context, key string, data any, expiry time.Duration) error
	Set(ctx context.Context, key string, data any) error
	GetString(ctx context.Context, key string) (string, error)
}
