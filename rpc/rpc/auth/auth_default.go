package auth

import (
	"context"
	auth "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/auth"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest, callOptions ...callopt.Option) (resp *auth.ValidateTokenResponse, err error) {
	resp, err = defaultClient.ValidateToken(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ValidateToken call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
