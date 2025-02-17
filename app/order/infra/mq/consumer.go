package mq

import (
	"context"
	"strconv"
	"strings"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/order/biz/service"
)

var (
	Consumer rocketmq.PushConsumer
)

func InitConsumer() {
	var err error
	Consumer, err = rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("order-consumer"),
	)
	if err != nil {
		panic(err)
	}

	// 订阅订单取消主题
	err = Consumer.Subscribe("order-cancel-topic", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			// 解析消息内容
			parts := strings.Split(string(msg.Body), ":")
			if len(parts) != 2 {
				klog.Errorf("invalid message format: %s", msg.Body)
				return consumer.ConsumeSuccess, nil // 格式错误的消息直接跳过
			}

			userId, err := strconv.ParseUint(parts[0], 10, 32)
			if err != nil {
				klog.Errorf("parse userId error: %v", err)
				return consumer.ConsumeSuccess, nil
			}

			orderId := parts[1]
			if err := service.CancelUnpaidOrder(ctx, uint32(userId), orderId); err != nil {
				klog.Errorf("cancel order error: %v, userId: %d, orderId: %s", err, userId, orderId)
				return consumer.ConsumeRetryLater, err
			}
			klog.Infof("cancel order success: userId: %d, orderId: %s", userId, orderId)
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		panic(err)
	}

	// 启动消费者
	err = Consumer.Start()
	if err != nil {
		panic(err)
	}
}
