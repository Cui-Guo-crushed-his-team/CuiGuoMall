package service

import (
	"context"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/order/biz/dal/mysql"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/order/biz/model"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/cart"
	order "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/order"
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
