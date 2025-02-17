package mq

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/cloudwego/kitex/pkg/klog"
)

var (
	Producer rocketmq.Producer
	err      error
)

func CreateTopic() error {
	adminClient, err := admin.NewAdmin(
		admin.WithResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	)
	if err != nil {
		return err
	}
	defer adminClient.Close()

	return adminClient.CreateTopic(
		context.Background(),
		admin.WithTopicCreate("email-topic", 8),
	)
}

// 在 Init 函数中调用
func Init() {
	// 创建 topic
	if err := CreateTopic(); err != nil {
		klog.Errorf("create topic error: %v", err)
	}

	// 创建生产者
	Producer, err = rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithGroupName("checkout-producer"),
	)
	if err != nil {
		panic(err)
	}

	err = Producer.Start()
	if err != nil {
		panic(err)
	}
}
