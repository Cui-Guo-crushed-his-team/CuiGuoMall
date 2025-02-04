package service

import (
	"context"
	payment "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment"
	"testing"
)

func TestFinish_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFinishService(ctx)
	// init req and assert value

	req := &payment.FinishReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
