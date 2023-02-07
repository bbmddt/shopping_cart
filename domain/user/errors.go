package user

import (
	"errors"
)

// 自定義錯誤

var (
	ErrUserExistWithName = errors.New("Username已經存在")
	ErrUserNotFound      = errors.New("使用者未找到")

	ErrMismatchedPasswords = errors.New("密碼錯誤")
	ErrInvalidUsername     = errors.New("無效Username")
	ErrInvalidPassword     = errors.New("無效Password")
)
