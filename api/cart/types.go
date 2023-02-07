package cart

// item請求參數
type ItemCartRequest struct {
	SKU   string `json:"sku"`
	Count int    `json:"count"`
}

// 創建分類回應
type CreateCategoryResponse struct {
	Message string `json:"message"`
}
