package service

import (
	"context"
	"errors"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/auth/biz/dal/repository"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/auth/conf"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/common/utils/encryption"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/common/utils/token"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/auth"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

type ValidateTokenService struct {
	ctx       context.Context
	cacheRepo repository.CacheRepository
} // NewValidateTokenService new ValidateTokenService
func NewValidateTokenService(ctx context.Context, repo repository.CacheRepository) *ValidateTokenService {
	return &ValidateTokenService{ctx: ctx, cacheRepo: repo}
}

// Run create note info
func (s *ValidateTokenService) Run(req *auth.ValidateTokenRequest) (resp *auth.ValidateTokenResponse, err error) {
	// 根据用户数据获得md5
	userMD5 := encryption.Md5(req.UserId + req.UserTrait)
	// 检查Redis中md5是否存在
	if err := s.cacheRepo.Exists(s.ctx, userMD5); err != nil {
		klog.Errorf("redis key unexisted:%v", userMD5)
		return nil, err
	}

	// 获取key的过期时间
	ttl, err := s.cacheRepo.GetTTL(s.ctx, userMD5)
	if err != nil {
		return nil, err
	}
	// TTL小于0
	if ttl < 0 {
		klog.Errorf("Redis: key %s has no expiration (ttl=%v)", userMD5, ttl)
		return nil, errors.New("token has no expiration")
	}

	// 如果剩余过期时间小于1小时则续约
	if ttl < time.Hour {
		// 签发新token
		newToken, err := token.GenerateToken(req.UserId, req.UserTrait, req.UserRole, conf.GetConf().Secret)
		if err != nil {
			klog.Errorf("Generate new token error: %v", err)
			return nil, err
		}
		// 延长过期时间到2小时
		if err := s.cacheRepo.Expire(s.ctx, userMD5, 2*time.Hour); err != nil {
			return nil, err
		}
		return &auth.ValidateTokenResponse{
			IsValid: true,
			Token:   newToken,
		}, nil
	}

	// 如果剩余过期时间大于等于1小时则正常返回
	return &auth.ValidateTokenResponse{IsValid: true}, nil
}
