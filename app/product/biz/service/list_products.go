package service

import (
	"context"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/product/biz/model"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/product" // 导入 kitex_gen/product 包
	"gorm.io/gorm"
)

type ListProductsService struct {
	ctx context.Context
	Db  *gorm.DB
}

// NewListProductsService creates a new instance of ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run handles the logic for listing products with pagination and category filtering
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	productQuery := model.ProductQuery{
		Ctx: s.ctx,
		Db:  s.Db,
	}

	// 调用 ListProducts 获取产品列表
	products, err := productQuery.ListProducts(int(req.Page), req.PageSize, req.CategoryName)
	if err != nil {
		return nil, err
	}

	// 将 model.Product 转换为 kitex_gen/product.Product
	var productList []*product.Product
	for _, p := range products {
		productList = append(productList, &product.Product{
			Id:          uint32(p.Id),                    // 转换为 uint32 类型
			Name:        p.Name,                          // 映射 Name
			Description: p.Description,                   // 映射 Description
			Picture:     p.Picture,                       // 映射 Picture
			Price:       p.Price,                         // 映射 Price
			Categories:  extractCategories(p.Categories), // 使用已经定义的 extractCategories 函数
		})
	}

	// 构建响应
	resp = &product.ListProductsResp{
		Products: productList,
	}
	return resp, nil
}
