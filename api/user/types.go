package user

// 創建用戶請求結構體
type CreateUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

// 創建用戶回應
type CreateUserResponse struct {
	Username string `json:"username"`
}

// 登入請求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登入回應
type LoginResponse struct {
	Username string `json:"username"`
	UserId   uint   `json:"userId"`
	Token    string `json:"token"`
}
