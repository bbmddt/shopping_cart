package api_helper

import "errors"

// 回響結構體
type Response struct {
	Message string `json:"message"`
}

// 回響錯誤結構體
type ErrorResponse struct {
	Message string `json:"errorMessage"`
}

// 自定義錯誤
var (
	ErrInvalidBody = errors.New("請檢查你的請求體")
)
