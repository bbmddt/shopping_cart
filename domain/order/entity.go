package order

import (
	"shopping_cart/domain/product"
	"shopping_cart/domain/user"

	"gorm.io/gorm"
)

// 訂單結構體
type Order struct {
	gorm.Model
	UserID       uint
	User         user.User     `gorm:"foreignKey:ID;references:UserID" json:"-"`
	OrderedItems []OrderedItem `gorm:"foreignKey:OrderID"`
	TotalPrice   float32
	IsCanceled   bool
}

// 訂單項目結構體
type OrderedItem struct {
	gorm.Model
	Product    product.Product `gorm:"foreignKey:ProductID"`
	ProductID  uint
	Count      int
	OrderID    uint
	IsCanceled bool
}

// 實例化訂單
func NewOrder(uid uint, items []OrderedItem) *Order {
	var totalPrice float32 = 0.0
	for _, item := range items {
		totalPrice += item.Product.Price
	}
	return &Order{
		UserID:       uid,
		OrderedItems: items,
		TotalPrice:   totalPrice,
		IsCanceled:   false,
	}
}

// 實例化訂單項目
func NewOrderedItem(count int, pid uint) *OrderedItem {
	return &OrderedItem{
		Count:      count,
		ProductID:  pid,
		IsCanceled: false,
	}
}
