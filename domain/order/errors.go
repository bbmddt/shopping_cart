package order

import "errors"

var (
	ErrEmptyCartFound       = errors.New("購物車是空的")
	ErrInvalidOrderID       = errors.New("無效訂單")
	ErrCancelDurationPassed = errors.New("已超過取消持續時間")
	ErrNotEnoughStock       = errors.New("沒有足夠庫存")
)
