package service

import (
	"context"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/cart/biz/dal/mysql"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/cart/biz/model"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/cart"
	"gorm.io/gorm"
)

// AddItemService is the service that handles adding items to the cart.
type AddItemService struct {
	ctx context.Context
}

// NewAddItemService creates a new AddItemService instance.
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run is the business logic to add an item to the cart.
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Get the database connection
	dbConn := mysql.DB

	// Check if the item already exists in the user's cart
	var cartItem model.Cart
	err = dbConn.Where("user_id = ? AND product_id = ?", req.UserId, req.Item.ProductId).First(&cartItem).Error

	if err == nil {
		// If the item already exists, update the quantity
		cartItem.Quantity += uint32(req.Item.Quantity)
		if err := dbConn.Save(&cartItem).Error; err != nil {
			return nil, err
		}
	} else if err == gorm.ErrRecordNotFound {
		// If the item does not exist, create a new cart item
		newCartItem := model.Cart{
			UserID:    req.UserId,
			ProductID: req.Item.ProductId,
			Quantity:  uint32(req.Item.Quantity),
		}
		if err := dbConn.Create(&newCartItem).Error; err != nil {
			return nil, err
		}
	} else {
		// If there is an error during the query, return it
		return nil, err
	}

	// Return an empty response on success
	return &cart.AddItemResp{}, nil
}
