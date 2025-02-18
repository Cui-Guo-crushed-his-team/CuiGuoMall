package service

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/order/biz/dal/mysql"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/order/biz/model"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/order/infra/rpc"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/cart"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/checkout"
	order "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/order"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
)

type GetOrderByIdService struct {
	ctx context.Context
} // NewGetOrderByIdService new GetOrderByIdService
func NewGetOrderByIdService(ctx context.Context) *GetOrderByIdService {
	return &GetOrderByIdService{ctx: ctx}
}

// Run create note info
func (s *GetOrderByIdService) Run(req *order.GetOrderByIdReq) (resp *order.GetOrderByIdResp, err error) {
	orderResp, err := model.GetOrder(mysql.DB, s.ctx, req.UserId, req.OrderId)
	if err != nil {
		return nil, err
	}
	// 如果是新创建，就去payment同步状态，
	if orderResp.OrderState == model.OrderStatePlaced {
		//其实我是更想 存在订单中的
		checkoutResp, err := rpc.CheckoutClient.GetCheckoutRecordByOrderId(s.ctx, &checkout.GetCheckoutRecordByOrderIdReq{
			UserId:  req.UserId,
			OrderId: req.OrderId,
		})
		if err != nil {
			return nil, err
		}
		paymentResp, err := rpc.PaymentClient.GetByOutTradeNo(s.ctx, &payment.GetByOutTradeNoReq{
			OutTradeNo: checkoutResp.PaymentOutTradeNo,
		})
		if err != nil {
			return nil, err
		}
		switch paymentResp.Status {
		case payment.Status_Success:
			orderResp.OrderState = model.OrderStatePaid
			err = model.UpdateOrderState(mysql.DB, s.ctx, req.UserId, req.OrderId, model.OrderStatePaid)
		case payment.Status_Failed:
			orderResp.OrderState = model.OrderStateCanceled
			err = model.UpdateOrderState(mysql.DB, s.ctx, req.UserId, req.OrderId, model.OrderStateCanceled)
		default:
			// success
		}
		if err != nil {
			return nil, err
		}

	}

	var orders []*order.OrderItem
	for _, v := range orderResp.OrderItems {
		orders = append(orders, &order.OrderItem{
			Cost: v.Cost,
			Item: &cart.CartItem{
				ProductId: v.ProductId,
				Quantity:  v.Quantity,
			},
		})
	}
	return &order.GetOrderByIdResp{
		Order: &order.Order{
			OrderItems:   orders,
			OrderId:      orderResp.OrderId,
			UserId:       orderResp.UserId,
			UserCurrency: orderResp.UserCurrency,
			Address:      nil, //偷懒不写
			Email:        orderResp.Consignee.Email,
			CreatedAt:    int32(orderResp.CreatedAt.UnixMilli()),
			Status:       string(orderResp.OrderState),
		},
	}, nil

}
