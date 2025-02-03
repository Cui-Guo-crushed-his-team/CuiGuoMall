package service

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/payment/conf"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/payment/model"
	payment "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
	alipay "github.com/smartwalle/alipay/v3"
)

type PrepayService struct {
	ctx          context.Context
	aliPayClient *alipay.Client
} // NewPrepayService new PrepayService
func NewPrepayService(ctx context.Context) *PrepayService {
	aliPayClient, err := alipay.New(conf.GetConf().AliPay.AppID, conf.GetConf().AliPay.PrivateKey, false)
	if err != nil {
		panic(err)
	}
	err = aliPayClient.LoadAliPayPublicKey(conf.GetConf().AliPay.PublicKey)
	if err != nil {
		panic(err)
	}
	return &PrepayService{ctx: ctx, aliPayClient: aliPayClient}
}

// Run create note info
func (s *PrepayService) Run(req *payment.PrepayReq) (resp *payment.PrepayResp, err error) {

	err = model.Create(s.ctx, req.Subject, req.OutTradeNo, req.Amount)
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

	return &payment.PrepayResp{PayUrl: payUrl.String()}, nil
}
