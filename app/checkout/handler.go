package main

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/checkout/biz/service"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/checkout"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	resp, err = service.NewCheckoutService(ctx).Run(req)

	return resp, err
}

// GetCheckoutRecordByOrderId implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) GetCheckoutRecordByOrderId(ctx context.Context, req *checkout.GetCheckoutRecordByOrderIdReq) (resp *checkout.GetCheckoutRecordByOrderIdResp, err error) {
	resp, err = service.NewGetCheckoutRecordByOrderIdService(ctx).Run(req)

	return resp, err
}
