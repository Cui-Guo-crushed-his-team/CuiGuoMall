package checkout

import (
	"context"
	checkout "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/checkout"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() checkoutservice.Client
	Service() string
	Checkout(ctx context.Context, Req *checkout.CheckoutReq, callOptions ...callopt.Option) (r *checkout.CheckoutResp, err error)
	GetCheckoutRecordByOrderId(ctx context.Context, Req *checkout.GetCheckoutRecordByOrderIdReq, callOptions ...callopt.Option) (r *checkout.GetCheckoutRecordByOrderIdResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := checkoutservice.NewClient(dstService, opts...)
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
	kitexClient checkoutservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() checkoutservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Checkout(ctx context.Context, Req *checkout.CheckoutReq, callOptions ...callopt.Option) (r *checkout.CheckoutResp, err error) {
	return c.kitexClient.Checkout(ctx, Req, callOptions...)
}

func (c *clientImpl) GetCheckoutRecordByOrderId(ctx context.Context, Req *checkout.GetCheckoutRecordByOrderIdReq, callOptions ...callopt.Option) (r *checkout.GetCheckoutRecordByOrderIdResp, err error) {
	return c.kitexClient.GetCheckoutRecordByOrderId(ctx, Req, callOptions...)
}
