package service

import (
	"context"
	model "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/checkout/biz/dal/model"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/checkout/biz/dal/mysql"
	checkout "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/checkout"
)

type GetCheckoutRecordByOrderIdService struct {
	ctx context.Context
} // NewGetCheckoutRecordByOrderIdService new GetCheckoutRecordByOrderIdService
func NewGetCheckoutRecordByOrderIdService(ctx context.Context) *GetCheckoutRecordByOrderIdService {
	return &GetCheckoutRecordByOrderIdService{ctx: ctx}
}

// Run create note info
func (s *GetCheckoutRecordByOrderIdService) Run(req *checkout.GetCheckoutRecordByOrderIdReq) (resp *checkout.GetCheckoutRecordByOrderIdResp, err error) {
	paymentRecord, err := model.GetPaymentRecordByOrderID(mysql.DB, s.ctx, req.OrderId)

	if err != nil {
		return nil, err
	}

	resp = &checkout.GetCheckoutRecordByOrderIdResp{
		Id:                paymentRecord.ID,
		OrderId:           paymentRecord.OrderID,
		Status:            int32(paymentRecord.Status),
		Amount:            paymentRecord.Amount,
		PaymentOutTradeNo: paymentRecord.PaymentOutTradeNo,
	}
	return resp, nil
}
