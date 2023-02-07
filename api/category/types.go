package category

// 創建分類請求參數型態
type CreateCategoryRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// 創建分類回應參數型態
type CreateCategoryResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
