package service

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/payment/model"
	payment "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
)

type FinishService struct {
	ctx context.Context
} // NewFinishService new FinishService
func NewFinishService(ctx context.Context) *FinishService {
	return &FinishService{ctx: ctx}
}

// Run completes the payment process
func (s *FinishService) Run(req *payment.FinishReq) (resp *payment.FinishResp, err error) {
	err = model.Finish(s.ctx, req.OutTradeNo, req.TradeNo) // Assume trade_no is provided or fetched elsewhere
	if err != nil {
		return nil, err
	}
	return &payment.FinishResp{}, nil
}
