package cart

import (
	"shopping_cart/domain/product"
	"shopping_cart/domain/user"

	"gorm.io/gorm"
)

// 購物車結構體
type Cart struct {
	gorm.Model
	UserID uint
	User   user.User `gorm:"foreignKey:ID;references:UserID"`
}

// 實例化
func NewCart(uid uint) *Cart {
	return &Cart{
		UserID: uid,
	}
}

// Item
type Item struct {
	gorm.Model
	Product   product.Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Count     int
	CartID    uint
	Cart      Cart `gorm:"foreignKey:CartID" json:"-"`
}

// 創建Item
func NewCartItem(productId uint, cartId uint, count int) *Item {
	return &Item{
		ProductID: productId,
		Count:     count,
		CartID:    cartId,
	}
}
