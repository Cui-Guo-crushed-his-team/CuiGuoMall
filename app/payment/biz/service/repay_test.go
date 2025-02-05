package service

import (
	"context"
	payment "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
	"testing"
)

func TestRepay_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRepayService(ctx)
	// init req and assert value

	req := &payment.RepayReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
