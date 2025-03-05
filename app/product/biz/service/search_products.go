package service

import (
	"context"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/product/biz/model"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/product"
	"gorm.io/gorm"
)

type SearchProductsService struct {
	ctx context.Context
	Db  *gorm.DB
}

// NewSearchProductsService creates a new instance of SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run handles the logic for searching products based on a query
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	productQuery := model.ProductQuery{Ctx: s.ctx, Db: s.Db} // db is assumed to be your gorm.DB instance

	products, err := productQuery.SearchProducts(req.Query)
	if err != nil {
		return nil, err
	}

	// Convert products to the response format
	var productList []*product.Product
	for _, p := range products {
		productList = append(productList, &product.Product{
			Id:          uint32(p.Id),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			Categories:  extractCategories(p.Categories),
		})
	}

	resp = &product.SearchProductsResp{
		Results: productList,
	}
	return resp, nil
}
