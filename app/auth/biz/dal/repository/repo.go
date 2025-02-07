package repository

import (
	"context"
	"time"
)

type CacheRepository interface {
	// Exists 检查key是否存在（此处通过GET判断,可能后续会有在v值中存储用户部分信息的需求）
	Exists(ctx context.Context, key string) error
	// GetTTL 返回key的剩余过期时间
	GetTTL(ctx context.Context, key string) (time.Duration, error)
	// Expire 设置key的过期时间
	Expire(ctx context.Context, key string, expiration time.Duration) error
}
