package order

// 完成訂單請求
type CompleteOrderRequest struct {
}

// 取消訂單
type CancelOrderRequest struct {
	OrderId uint `json:"orderId"`
}
