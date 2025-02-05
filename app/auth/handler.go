package main

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/auth/biz/service"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// ValidateToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (resp *auth.ValidateTokenResponse, err error) {
	resp, err = service.NewValidateTokenService(ctx).Run(req)

	return resp, err
}
