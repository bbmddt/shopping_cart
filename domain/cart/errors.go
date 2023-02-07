package cart

import (
	"errors"
)

var (
	ErrItemAlreadyExistInCart = errors.New("Item 已經存在")
	ErrCountInvalid           = errors.New("數量不能是負值")
)
