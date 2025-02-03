package service

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/payment/model"
	payment "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
)

type GetByOutTradeNoService struct {
	ctx context.Context
} // NewGetByOutTradeNoService new GetByOutTradeNoService
func NewGetByOutTradeNoService(ctx context.Context) *GetByOutTradeNoService {
	return &GetByOutTradeNoService{ctx: ctx}
}

// Run create note info
func (s *GetByOutTradeNoService) Run(req *payment.GetByOutTradeNoReq) (resp *payment.GetByOutTradeNoResp, err error) {
	p, err := model.GetByOutTradeNO(s.ctx, req.OutTradeNo)
	if err != nil {
		return nil, err
	}
	resp = &payment.GetByOutTradeNoResp{
		Description: p.Description,
		OutTradeNo:  p.OutTradeNO,
		TradeNo:     p.TradeNO.String,
		Amount:      p.Amount,
	}
	switch p.Status {
	case model.PaymentStatusInit:
		resp.Status = payment.Status_Init
	case model.PaymentStatusSuccess:
		resp.Status = payment.Status_Success
	case model.PaymentStatusFailed:
		resp.Status = payment.Status_Failed
	case model.PaymentStatusRefund:
		resp.Status = payment.Status_Refund
	default:
		resp.Status = payment.Status_Unknown
	}
	return resp, nil
}
