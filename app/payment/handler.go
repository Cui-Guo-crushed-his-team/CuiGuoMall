package main

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/payment/biz/service"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Prepay implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Prepay(ctx context.Context, req *payment.PrepayReq) (resp *payment.PrepayResp, err error) {
	resp, err = service.NewPrepayService(ctx).Run(req)

	return resp, err
}

// Finish implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Finish(ctx context.Context, req *payment.FinishReq) (resp *payment.FinishResp, err error) {
	resp, err = service.NewFinishService(ctx).Run(req)

	return resp, err
}

// GetByOutTradeNo implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) GetByOutTradeNo(ctx context.Context, req *payment.GetByOutTradeNoReq) (resp *payment.GetByOutTradeNoResp, err error) {
	resp, err = service.NewGetByOutTradeNoService(ctx).Run(req)

	return resp, err
}
