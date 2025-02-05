package service

import (
	"context"
	auth "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/auth"
	"testing"
)

func TestValidateToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewValidateTokenService(ctx)
	// init req and assert value

	req := &auth.ValidateTokenRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
