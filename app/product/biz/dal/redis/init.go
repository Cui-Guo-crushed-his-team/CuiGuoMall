package redis

import (
	"context"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/product/conf"
	"github.com/go-redis/redis/v8" // 修改为使用 v8 版本的 go-redis
)

var (
	RedisClient *redis.Client
)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	// 使用 context.Background() 进行 Ping 操作
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
