package product

import (
	"shopping_cart/domain/category"

	"gorm.io/gorm"
)

// 商品結構體
type Product struct {
	gorm.Model
	Name       string
	SKU        string
	Desc       string
	StockCount int
	Price      float32
	CategoryID uint              // 分類id
	Category   category.Category `json:"-"` // 分類
	IsDeleted  bool
}

// 商品結構體實例
func NewProduct(name string, desc string, stockCount int, price float32, cid uint) *Product {
	return &Product{
		Name:       name,
		Desc:       desc,
		StockCount: stockCount,
		Price:      price,
		CategoryID: cid,
		IsDeleted:  false,
	}
}
