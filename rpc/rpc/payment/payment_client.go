package payment

import (
	"context"
	payment "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() paymentservice.Client
	Service() string
	Prepay(ctx context.Context, Req *payment.PrepayReq, callOptions ...callopt.Option) (r *payment.PrepayResp, err error)
	Repay(ctx context.Context, Req *payment.RepayReq, callOptions ...callopt.Option) (r *payment.RepayResp, err error)
	Finish(ctx context.Context, Req *payment.FinishReq, callOptions ...callopt.Option) (r *payment.FinishResp, err error)
	GetByOutTradeNo(ctx context.Context, Req *payment.GetByOutTradeNoReq, callOptions ...callopt.Option) (r *payment.GetByOutTradeNoResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := paymentservice.NewClient(dstService, opts...)
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
	kitexClient paymentservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() paymentservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Prepay(ctx context.Context, Req *payment.PrepayReq, callOptions ...callopt.Option) (r *payment.PrepayResp, err error) {
	return c.kitexClient.Prepay(ctx, Req, callOptions...)
}

func (c *clientImpl) Repay(ctx context.Context, Req *payment.RepayReq, callOptions ...callopt.Option) (r *payment.RepayResp, err error) {
	return c.kitexClient.Repay(ctx, Req, callOptions...)
}

func (c *clientImpl) Finish(ctx context.Context, Req *payment.FinishReq, callOptions ...callopt.Option) (r *payment.FinishResp, err error) {
	return c.kitexClient.Finish(ctx, Req, callOptions...)
}

func (c *clientImpl) GetByOutTradeNo(ctx context.Context, Req *payment.GetByOutTradeNoReq, callOptions ...callopt.Option) (r *payment.GetByOutTradeNoResp, err error) {
	return c.kitexClient.GetByOutTradeNo(ctx, Req, callOptions...)
}
