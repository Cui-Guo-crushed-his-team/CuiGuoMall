package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Id          int        `json:"id" gorm:"primaryKey"` // 修改为 int 类型
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	Ctx context.Context
	Db  *gorm.DB // 修改为 db 以便和实际的 gorm.DB 结构体匹配
}

// GetById retrieves a product by its ID
func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.Db.WithContext(p.Ctx).Model(&Product{}).First(&product, productId).Error
	if err != nil {
		return product, fmt.Errorf("could not find product with ID %d: %w", productId, err)
	}
	return product, nil
}

// SearchProducts searches for products by name or description
func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.Db.WithContext(p.Ctx).Model(&Product{}).Find(&products, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	if err != nil {
		return nil, fmt.Errorf("error occurred while searching for products: %w", err)
	}
	return products, nil
}

// ListProducts retrieves a list of products with pagination and optional category filter
func (p ProductQuery) ListProducts(page int, pageSize int64, categoryName string) (products []*Product, err error) {
	query := p.Db.WithContext(p.Ctx).Model(&Product{})

	if categoryName != "" {
		// Assume a `Category` table exists and you can filter by category name
		query = query.Joins("JOIN product_category ON product_category.product_id = product.id").
			Joins("JOIN category ON category.id = product_category.category_id").
			Where("category.name = ?", categoryName)
	}

	// Apply pagination
	err = query.Offset((page - 1) * int(pageSize)).
		Limit(int(pageSize)).
		Find(&products).Error

	if err != nil {
		return nil, fmt.Errorf("error occurred while listing products: %w", err)
	}
	return products, nil
}
