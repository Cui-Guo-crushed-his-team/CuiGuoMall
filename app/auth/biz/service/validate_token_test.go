package service

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/auth/biz/dal/redis"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/auth/conf"
	auth "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/auth"
	r "github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

func TestValidateToken_Run(t *testing.T) {
	content, err := os.ReadFile("../../conf/dev/conf.yaml")
	if err != nil {
		panic(err)
	}
	cf := new(conf.Config)
	err = yaml.Unmarshal(content, cf)
	redisClient := r.NewClient(&r.Options{
		Addr:     cf.Redis.Address,
		Username: "",
		Password: "",
		DB:       1,
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	ctx := context.Background()
	s := NewValidateTokenService(ctx, redis.NewRedisRepo(redisClient))
	req := &auth.ValidateTokenRequest{
		UserId:    "123456789",
		UserRole:  "user",
		UserTrait: "192.168.1.1",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
