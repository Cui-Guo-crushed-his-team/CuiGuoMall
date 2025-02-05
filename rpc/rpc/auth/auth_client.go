package auth

import (
	"context"
	auth "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/auth"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() authservice.Client
	Service() string
	ValidateToken(ctx context.Context, Req *auth.ValidateTokenRequest, callOptions ...callopt.Option) (r *auth.ValidateTokenResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := authservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient authservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() authservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ValidateToken(ctx context.Context, Req *auth.ValidateTokenRequest, callOptions ...callopt.Option) (r *auth.ValidateTokenResponse, err error) {
	return c.kitexClient.ValidateToken(ctx, Req, callOptions...)
}
