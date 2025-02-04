package payment

import (
	"context"
	payment "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Prepay(ctx context.Context, req *payment.PrepayReq, callOptions ...callopt.Option) (resp *payment.PrepayResp, err error) {
	resp, err = defaultClient.Prepay(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Prepay call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Finish(ctx context.Context, req *payment.FinishReq, callOptions ...callopt.Option) (resp *payment.FinishResp, err error) {
	resp, err = defaultClient.Finish(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Finish call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetByOutTradeNo(ctx context.Context, req *payment.GetByOutTradeNoReq, callOptions ...callopt.Option) (resp *payment.GetByOutTradeNoResp, err error) {
	resp, err = defaultClient.GetByOutTradeNo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetByOutTradeNo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
