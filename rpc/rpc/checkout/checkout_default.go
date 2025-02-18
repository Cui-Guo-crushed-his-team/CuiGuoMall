package checkout

import (
	"context"
	checkout "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/checkout"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Checkout(ctx context.Context, req *checkout.CheckoutReq, callOptions ...callopt.Option) (resp *checkout.CheckoutResp, err error) {
	resp, err = defaultClient.Checkout(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Checkout call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCheckoutRecordByOrderId(ctx context.Context, req *checkout.GetCheckoutRecordByOrderIdReq, callOptions ...callopt.Option) (resp *checkout.GetCheckoutRecordByOrderIdResp, err error) {
	resp, err = defaultClient.GetCheckoutRecordByOrderId(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCheckoutRecordByOrderId call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
