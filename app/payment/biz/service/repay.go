package service

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/payment/conf"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/payment/model"
	payment "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
	"github.com/smartwalle/alipay/v3"
)

type RepayService struct {
	ctx          context.Context
	aliPayClient *alipay.Client
} // NewRepayService new RepayService
func NewRepayService(ctx context.Context) *RepayService {
	aliPayClient, err := alipay.New(conf.GetConf().AliPay.AppID, conf.GetConf().AliPay.PrivateKey, false)
	if err != nil {
		panic(err)
	}
	err = aliPayClient.LoadAliPayPublicKey(conf.GetConf().AliPay.PublicKey)
	if err != nil {
		panic(err)
	}
	return &RepayService{ctx: ctx, aliPayClient: aliPayClient}
}

// Run create note info
func (s *RepayService) Run(req *payment.RepayReq) (resp *payment.RepayResp, err error) {
	err = model.CreateOnce(s.ctx, req.Subject, req.OutTradeNo, req.Amount)
	if err != nil {
		return nil, err
	}

	var p = alipay.TradeWapPay{}
	p.ProductCode = "QUICK_WAP_WAY"
	p.NotifyURL = conf.GetConf().AliPay.NotifyURL
	p.ReturnURL = conf.GetConf().AliPay.ReturnURL

	p.Subject = req.Subject
	p.OutTradeNo = req.OutTradeNo
	p.TotalAmount = req.Amount

	payUrl, err := s.aliPayClient.TradeWapPay(p)
	if err != nil {
		return nil, err
	}

	return &payment.RepayResp{PayUrl: payUrl.String()}, nil

}
