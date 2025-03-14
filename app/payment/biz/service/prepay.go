package service

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/payment/conf"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/payment/model"
	payment "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
	sf "github.com/bwmarrin/snowflake"
	alipay "github.com/smartwalle/alipay/v3"
	"time"
)

type PrepayService struct {
	ctx          context.Context
	aliPayClient *alipay.Client
	SfNode       *sf.Node
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

	return &PrepayService{ctx: ctx, aliPayClient: aliPayClient, SfNode: newSfNode()}
}

func newSfNode() *sf.Node {
	var st time.Time
	st, err := time.Parse("2006-01-02", "2025-02-01")
	if err != nil {
		panic(err)
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err := sf.NewNode(1)
	if err != nil {
		panic(err)
	}
	return node
}

// Run create note info
func (s *PrepayService) Run(req *payment.PrepayReq) (resp *payment.PrepayResp, err error) {
	outTradeNo := s.SfNode.Generate().String()
	err = model.Create(s.ctx, req.Subject, outTradeNo, req.Amount)
	if err != nil {
		return nil, err
	}

	var p = alipay.TradeWapPay{}
	p.ProductCode = "QUICK_WAP_WAY"
	p.NotifyURL = conf.GetConf().AliPay.NotifyURL
	p.ReturnURL = conf.GetConf().AliPay.ReturnURL

	p.Subject = req.Subject
	p.OutTradeNo = outTradeNo
	p.TotalAmount = req.Amount

	payUrl, err := s.aliPayClient.TradeWapPay(p)
	if err != nil {
		return nil, err
	}

	return &payment.PrepayResp{PayUrl: payUrl.String(), OutTradeNo: outTradeNo}, nil
}
