package user

import (
	"shopping_cart/utils/hash"

	"gorm.io/gorm"
)

// 保存使用者之前callback，如果密碼沒有被加密加密密碼和salt
func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	if u.Salt == "" {
		// 為salt創建一個隨機字符串
		salt := hash.CreateSalt()
		// 創建hash加密密碼
		hashPassword, err := hash.HashPassword(u.Password + salt)
		if err != nil {
			return nil
		}
		u.Password = hashPassword
		u.Salt = salt
	}

	return
}
