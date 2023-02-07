package order

import (
	"shopping_cart/domain/cart"
	"shopping_cart/domain/product"
	"shopping_cart/utils/pagination"
	"time"
)

var day14ToHours float64 = 336

type Service struct {
	orderRepository       Repository
	orderedItemRepository OrderedItemRepository
	productRepository     product.Repository
	cartRepository        cart.Repository
	cartItemRepository    cart.ItemRepository
}

// 實例化
func NewService(
	orderRepository Repository,
	orderedItemRepository OrderedItemRepository,
	productRepository product.Repository,
	cartRepository cart.Repository,
	cartItemRepository cart.ItemRepository,
) *Service {
	orderRepository.Migration()
	orderedItemRepository.Migration()
	return &Service{
		orderRepository:       orderRepository,
		orderedItemRepository: orderedItemRepository,
		productRepository:     productRepository,
		cartRepository:        cartRepository,
		cartItemRepository:    cartItemRepository,
	}

}

// 完成訂單
func (c *Service) CompleteOrder(userId uint) error {
	currentCart, err := c.cartRepository.FindOrCreateByUserID(userId)
	if err != nil {
		return err
	}
	cartItems, err := c.cartItemRepository.GetItems(currentCart.UserID)
	if err != nil {
		return err
	}
	if len(cartItems) == 0 {
		return ErrEmptyCartFound
	}
	orderedItems := make([]OrderedItem, 0)
	for _, item := range cartItems {
		orderedItems = append(orderedItems, *NewOrderedItem(item.Count, item.ProductID))
	}
	err = c.orderRepository.Create(NewOrder(userId, orderedItems))
	return err
}

// 取消訂單
func (c *Service) CancelOrder(uid, oid uint) error {
	currentOrder, err := c.orderRepository.FindByOrderID(oid)
	if err != nil {
		return err
	}
	if currentOrder.UserID != uid {
		return ErrInvalidOrderID
	}
	if time.Until(currentOrder.CreatedAt).Hours() > day14ToHours {
		return ErrCancelDurationPassed
	}
	currentOrder.IsCanceled = true
	err = c.orderRepository.Update(*currentOrder)

	return err
}

// 獲取訂單
func (c *Service) GetAll(page *pagination.Pages, uid uint) *pagination.Pages {
	orders, count := c.orderRepository.GetAll(page.Page, page.PageSize, uid)
	page.Items = orders
	page.TotalCount = count
	return page

}
