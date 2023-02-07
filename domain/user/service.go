package user

import "shopping_cart/utils/hash"

// 使用者service結構體
type Service struct {
	r Repository
}

// 實例化service
func NewUserService(r Repository) *Service {
	r.Migration()
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// 創建使用者
func (c *Service) Create(user *User) error {
	if user.Password != user.Password2 {
		return ErrMismatchedPasswords
	}
	// 使用者存在
	_, err := c.r.GetByName(user.Username)
	if err == nil {
		return ErrUserExistWithName
	}
	// 無效使用者名稱
	if ValidateUserName(user.Username) {
		return ErrInvalidUsername
	}
	// 無效密碼
	if ValidatePassword(user.Password) {
		return ErrInvalidPassword
	}
	// 創建使用者
	err = c.r.Create(user)
	return err
}

// 查詢使用者
func (c *Service) GetUser(username string, password string) (User, error) {
	user, err := c.r.GetByName(username)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	match := hash.CheckPasswordHash(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

// 更新使用者
func (c *Service) UpdateUser(user *User) error {
	return c.r.Update(user)
}
