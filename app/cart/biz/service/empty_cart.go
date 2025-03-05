package service

import (
	"context"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/cart/biz/dal/mysql"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/cart/biz/model"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/cart"
)

// EmptyCartService handles the logic to empty a user's cart.
type EmptyCartService struct {
	ctx context.Context
}

// NewEmptyCartService creates a new EmptyCartService instance.
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run is the business logic to clear all items in the user's cart.
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Get the database connection
	dbConn := mysql.DB

	// Delete all items in the cart for the given user ID
	if err := dbConn.Where("user_id = ?", req.UserId).Delete(&model.Cart{}).Error; err != nil {
		return nil, err
	}

	// Return an empty response indicating success
	return &cart.EmptyCartResp{}, nil
}
