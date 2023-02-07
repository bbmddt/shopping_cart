package product

import (
	"errors"
)

var (
	ErrProductNotFound         = errors.New("商品沒有找到")
	ErrProductStockIsNotEnough = errors.New("商品庫存不足")
)
