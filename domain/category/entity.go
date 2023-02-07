package category

import (
	"gorm.io/gorm"
)

// 商品分類結構體，對應資料庫表
type Category struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Desc     string
	IsActive bool
}

// 新建商品分類
func NewCategory(name string, desc string) *Category {
	return &Category{
		Name:     name,
		Desc:     desc,
		IsActive: true,
	}
}
