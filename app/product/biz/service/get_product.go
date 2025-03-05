package service

import (
	"context"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/product/biz/model"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/product" // 导入 kitex_gen/product 包
	"gorm.io/gorm"
)

type GetProductService struct {
	ctx context.Context
	Db  *gorm.DB
}

// NewGetProductService creates a new instance of GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{
		ctx: ctx,
	}
}

// Run handles the logic for getting a single product based on the given ID
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// 创建 ProductQuery 实例，传递上下文和数据库实例
	productQuery := model.ProductQuery{
		Ctx: s.ctx,
		Db:  s.Db,
	}

	// 调用 GetById 方法获取数据库中的产品
	modelProduct, err := productQuery.GetById(int(req.GetId())) // 使用 req.GetId() 获取 ID
	if err != nil {
		return nil, err
	}

	// 将数据库中的 model.Product 转换为 kitex_gen/product.Product
	resp = &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(modelProduct.Id),                    // 转换为 uint32 类型
			Name:        modelProduct.Name,                          // 映射 Name
			Description: modelProduct.Description,                   // 映射 Description
			Picture:     modelProduct.Picture,                       // 映射 Picture
			Price:       modelProduct.Price,                         // 映射 Price
			Categories:  extractCategories(modelProduct.Categories), // 使用 Categories 字段进行转换
		},
	}
	return resp, nil
}

// 修改 extractCategories 接受 []model.Category 类型
func extractCategories(categories []model.Category) []string {
	var categoryNames []string
	for _, category := range categories {
		categoryNames = append(categoryNames, category.Name) // 假设 Category 结构有 Name 字段
	}
	return categoryNames
}
