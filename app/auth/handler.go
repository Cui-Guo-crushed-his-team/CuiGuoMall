package main

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/auth/biz/dal/redis"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/auth/biz/service"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct {
	rdb *redis.RedisRepo
}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{rdb: redis.Rdb}
}

// ValidateToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (resp *auth.ValidateTokenResponse, err error) {
	resp, err = service.NewValidateTokenService(ctx, s.rdb).Run(req)
	return resp, err
}
