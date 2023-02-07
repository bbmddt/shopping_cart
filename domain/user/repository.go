package user

import (
	"log"

	"gorm.io/gorm"
)

// Repository 結構體
type Repository struct {
	db *gorm.DB
}

// 實例化
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 創建使用者資料表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&User{})
	if err != nil {
		log.Print(err)
	}
}

// 創建使用者
func (r *Repository) Create(u *User) error {
	result := r.db.Create(u)

	return result.Error
}

// 根據Username查詢使用者
func (r *Repository) GetByName(name string) (User, error) {
	var user User
	err := r.db.Where("UserName = ?", name).Where("IsDeleted = ?", 0).First(&user, "UserName = ?", name).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// 創建測試用資料
func (r *Repository) InsertSampleData() {
	user := NewUser("admin", "admin", "admin")
	user.IsAdmin = true
	r.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username, Password: user.Password}).FirstOrCreate(&user)
	user = NewUser("user", "user", "user")
	r.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username, Password: user.Password}).FirstOrCreate(&user)

}

// 更新使用者資料
func (r *Repository) Update(u *User) error {
	return r.db.Save(&u).Error
}
