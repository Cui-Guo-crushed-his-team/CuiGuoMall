package service

import (
	"context"

	cart "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/gateway/hertz_gen/cart"
	common "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/gateway/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
