package service

import (
	"context"
	order "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/order"
	"testing"
)

func TestGetOrderById_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetOrderByIdService(ctx)
	// init req and assert value

	req := &order.GetOrderByIdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
