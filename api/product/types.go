package product

import "shopping_cart/domain/product"

// 創建商品請求參數
type CreateProductRequest struct {
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	Count      int     `json:"count"`
	CategoryID uint    `json:"categoryID"`
}

// 創建商品回應參數
type CreateProductResponse struct {
	Message string `json:"message"`
}

// 刪除商品請求參數
type DeleteProductRequest struct {
	SKU string `json:"sku"`
}

// 更新商品請求參數
type UpdateProductRequest struct {
	SKU        string  `json:"sku"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	Count      int     `json:"count"`
	CategoryID uint    `json:"categoryID"`
}

// 型態轉換 UpdateProductRequest to Product
func (p *UpdateProductRequest) ToProduct() *product.Product {
	return product.NewProduct(p.Name, p.Desc, p.Count, p.Price, p.CategoryID)
}
