package service

import (
	"context"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/order/biz/dal/mysql"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/order/biz/model"
)

func CancelUnpaidOrder(ctx context.Context, userId uint32, orderId string) error {
	// 获取订单状态
	order, err := model.GetOrder(mysql.DB, ctx, userId, orderId)
	if err != nil {
		return err
	}

	// 如果订单未支付，则取消订单
	if order.OrderState == model.OrderStatePlaced {
		return model.UpdateOrderState(mysql.DB, ctx, userId, orderId, model.OrderStateCanceled)
	}
	return nil
}
