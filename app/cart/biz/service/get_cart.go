package service

import (
	"context"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/cart/biz/dal/mysql"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/cart/biz/model"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/cart"
)

// GetCartService handles the logic to get the items in the user's cart.
type GetCartService struct {
	ctx context.Context
}

// NewGetCartService creates a new GetCartService instance.
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run is the business logic to retrieve all items in the user's cart.
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Get the database connection
	dbConn := mysql.DB

	// Query to get all items in the user's cart
	var cartItems []model.Cart
	err = dbConn.Where("user_id = ?", req.UserId).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}

	// Construct the response object
	var cartItemsResp []*cart.CartItem
	for _, item := range cartItems {
		cartItemsResp = append(cartItemsResp, &cart.CartItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
		})
	}

	// Return the response containing all cart items
	cartResp := &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.UserId,
			Items:  cartItemsResp,
		},
	}

	return cartResp, nil
}
