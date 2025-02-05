package service

import (
	"context"
	auth "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/auth"
)

type ValidateTokenService struct {
	ctx context.Context
} // NewValidateTokenService new ValidateTokenService
func NewValidateTokenService(ctx context.Context) *ValidateTokenService {
	return &ValidateTokenService{ctx: ctx}
}

// Run create note info
func (s *ValidateTokenService) Run(req *auth.ValidateTokenRequest) (resp *auth.ValidateTokenResponse, err error) {

	return
}
