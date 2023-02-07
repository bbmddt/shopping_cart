package category

import (
	"mime/multipart"
	"shopping_cart/utils/csv_helper"
	"shopping_cart/utils/pagination"
)

type Service struct {
	r Repository
}

// 實例化商品分類service
func NewCategoryService(r Repository) *Service {
	// 生成資料表
	r.Migration()
	// 插入測試Sample
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// 創建分類
func (c *Service) Create(category *Category) error {
	existCity := c.r.GetByName(category.Name)
	if len(existCity) > 0 {
		return ErrCategoryExistWithName
	}

	err := c.r.Create(category)
	if err != nil {
		return err
	}

	return nil
}

// 批次創建分類
func (c *Service) BulkCreate(fileHeader *multipart.FileHeader) (int, error) {
	categories := make([]*Category, 0)
	bulkCategory, err := csv_helper.ReadCsv(fileHeader)
	if err != nil {
		return 0, err
	}
	for _, categoryVariables := range bulkCategory {
		categories = append(categories, NewCategory(categoryVariables[0], categoryVariables[1]))
	}
	count, err := c.r.BulkCreate(categories)
	if err != nil {
		return count, err
	}
	return count, nil
}

// 獲得分頁商品分類
func (c *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	categories, count := c.r.GetAll(page.Page, page.PageSize)
	page.Items = categories
	page.TotalCount = count
	return page
}
