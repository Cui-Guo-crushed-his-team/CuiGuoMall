package redis

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"time"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/auth/conf"
)

var Rdb *RedisRepo

type RedisRepo struct {
	client *redis.Client
}

func NewRedisRepo(client *redis.Client) *RedisRepo {
	return &RedisRepo{client: client}
}
func Init() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	Rdb = NewRedisRepo(redisClient)
}
func (r *RedisRepo) Exists(ctx context.Context, key string) error {
	_, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		klog.Errorf("Redis: key %s 不存在: %v", key, err)
		return errors.New("invalid redis key")
	} else if err != nil {
		klog.Errorf("Redis GET error, key %s: %v", key, err)
		return err
	}
	return nil
}

// GetTTL 获取 key 的剩余过期时间
func (r *RedisRepo) GetTTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := r.client.TTL(ctx, key).Result()
	if err != nil {
		klog.Errorf("Redis TTL error, key %s: %v", key, err)
		return ttl, err
	}
	return ttl, nil
}

// Expire 对 key 进行续约，设置新的过期时间
func (r *RedisRepo) Expire(ctx context.Context, key string, expiration time.Duration) error {
	if err := r.client.Expire(ctx, key, expiration).Err(); err != nil {
		klog.Errorf("Redis Expire error, key %s: %v", key, err)
		return err
	}
	return nil
}
