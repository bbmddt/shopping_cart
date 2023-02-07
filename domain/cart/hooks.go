package cart

import (
	"gorm.io/gorm"
)

// 如果計數為零，則刪除商品
func (item *Item) AfterUpdate(tx *gorm.DB) (err error) {

	if item.Count <= 0 {
		return tx.Unscoped().Delete(&item).Error
	}
	return

}
