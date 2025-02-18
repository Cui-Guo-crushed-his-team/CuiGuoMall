package service

import (
	"context"
	"errors"
	"strconv"

	"fmt"

	model "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/checkout/biz/dal/model"
	mysql "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/checkout/biz/dal/mysql"
	mq "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/checkout/infra/mq"
	rpc "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/checkout/infra/rpc"
	cart "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/cart"
	checkout "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/checkout"
	order "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/order"
	payment "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
	product "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/product"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// 有这个TradeNo说明不是第一次支付，要进行repay
	if req.PaymentOutTradeNo != "" {
		// 获取支付记录
		paymentRecord, err := model.GetPaymentRecordByPaymentOutTradeNo(mysql.DB, s.ctx, req.PaymentOutTradeNo)
		if err != nil {
			klog.Error(err)
			err = fmt.Errorf("GetPaymentRecordByOrderID.err:%v", err)
			return nil, err
		}
		paymentResp, err := rpc.PaymentClient.GetByOutTradeNo(s.ctx, &payment.GetByOutTradeNoReq{OutTradeNo: req.PaymentOutTradeNo})
		if err != nil {
			return nil, fmt.Errorf("GetPaymentRecordByOrderID.err:%v", err)
		}
		switch paymentResp.Status {
		case payment.Status_Unknown:
			//一般不会 出现
			return nil, errors.New("payment status unKnow")
		case payment.Status_Success:
			return nil, errors.New("payment already paid")
		case payment.Status_Failed:
			return nil, errors.New("payment failed")
		case payment.Status_Refund:
			return nil, errors.New("payment refund")
		default:
			//未支付
		}

		repayResp, err := rpc.PaymentClient.Repay(s.ctx, &payment.RepayReq{
			Amount:     strconv.FormatFloat(paymentRecord.Amount, 'f', 2, 64),
			OutTradeNo: req.PaymentOutTradeNo,
			Subject:    fmt.Sprintf("Order %s", paymentRecord.OrderID),
		})
		if err != nil {
			klog.Error(err)
			return nil, err
		}
		// 不用更新状态，都是正在支付

		// todo 让支付服务去进行调用
		// _, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{UserId: req.UserId, OrderId: paymentRecord.OrderID})
		// if err != nil {
		// 	klog.Error(err)
		// 	return nil, err
		// }
		return &checkout.CheckoutResp{
			OrderId:       paymentRecord.OrderID,
			TransactionId: paymentRecord.PaymentOutTradeNo,
			PayUrl:        repayResp.PayUrl,
		}, nil

	}
	//  req.PaymentOutTradeNo == "" 就执行购物车结算
	// 获取购物车信息
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err)
		err = fmt.Errorf("GetCart.err:%v", err)
		return nil, err
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		err = errors.New("cart is empty")
		return nil, err
	}
	var (
		oi    []*order.OrderItem
		total float32
	)
	// 处理购物车内商品信息
	for _, cartItem := range cartResult.Cart.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.Error(resultErr)
			err = resultErr
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		cost := p.Price * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}
	// create order
	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Email,
	}
	// 处理地址信息
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		return nil, err
	}
	klog.Info("orderResult", orderResult)
	if orderResult == nil || orderResult.Order == nil {
		klog.Error(errors.New("order creation failed"))
		return nil, errors.New("order creation failed")
	}

	// 发送延迟取消订单消息
	err = s.sendDelayedCancelMessage(req.UserId, orderResult.Order.OrderId)
	if err != nil {
		klog.Error(err)
		return nil, fmt.Errorf("sendDelayedCancelMessage.err:%v", err)
	}

	// 清空购物车
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		return nil, err
	}
	klog.Info(emptyResult)
	payReq := &payment.PrepayReq{
		Amount:  strconv.FormatFloat(float64(total), 'f', 2, 64),
		Subject: fmt.Sprintf("Order %s", orderResult.Order.OrderId),
	}

	paymentResult, err := rpc.PaymentClient.Prepay(s.ctx, payReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		return nil, err
	}

	err = model.CreatePaymentRecord(mysql.DB, s.ctx, &model.PaymentRecord{
		OrderID:           orderResult.Order.OrderId,
		Status:            model.PaymentStatusUnpaid,
		Amount:            float64(total),
		PaymentOutTradeNo: paymentResult.OutTradeNo,
	})
	// 支付
	var orderId = orderResult.Order.OrderId

	// data, _ := proto.Marshal(&email.EmailReq{
	// 	From:        "from@example.com",
	// 	To:          req.Email,
	// 	ContentType: "text/plain",
	// 	Subject:     "You just created an order in CloudWeGo shop",
	// 	Content:     "You just created an order in CloudWeGo shop",
	// })
	// msg := &nats.Msg{Subject: "email", Data: data, Header: make(nats.Header)}

	// // otel inject
	// otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

	// _ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResult)
	// change order state
	klog.Info(orderResult)
	// 调用订单服务，标记订单已支付
	// todo 让支付服务去进行调用
	// _, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{UserId: req.UserId, OrderId: orderId})
	// if err != nil {
	// 	klog.Error(err)
	// 	return nil, err
	// }
	err = model.UpdatePaymentStatus(mysql.DB, s.ctx, orderId, model.PaymentStatusPaying)
	if err != nil {
		klog.Error(err)
		return nil, fmt.Errorf("UpdatePaymentStatus.err:%v", err)
	}

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.OutTradeNo,
		PayUrl:        paymentResult.PayUrl,
	}

	return resp, nil
}

// 发送延迟取消订单消息
func (s *CheckoutService) sendDelayedCancelMessage(userId uint32, orderId string) error {
	// 构造消息内容
	msgContent := fmt.Sprintf("%d:%s", userId, orderId) // 使用冒号分隔 userId 和 orderId

	message := &primitive.Message{
		Topic: mq.OrderCancelTopic,
		Body:  []byte(msgContent),
	}
	// 设置延迟级别为30分钟
	// RocketMQ 延迟级别: 1=1s, 2=5s, 3=10s, 4=30s, 5=1m, 6=2m, 7=3m, 8=4m, 9=5m, 10=6m, 11=7m, 12=8m, 13=9m, 14=10m, 15=20m, 16=30m
	message.WithDelayTimeLevel(16)

	_, err := mq.Producer.SendSync(context.Background(), message)
	if err != nil {
		klog.Errorf("send delay message error: %v", err)
		return err
	}
	return nil
}
