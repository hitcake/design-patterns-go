package order

import "time"

type Order struct {
	OrderId   int64 `json:"orderId"`
	UserId    int64 `json:"userId"`
	ProductId int64 `json:"productId"`
	SellerId  int64 `json:"sellerId"`
	Status    int
	OrderTime time.Time `json:"orderTime"`
}

type OrderService struct {
}

func (o *OrderService) GetOrderById(orderId int64) (*Order, error) {
	return &Order{OrderId: orderId, UserId: 1, SellerId: 1, ProductId: 1, Status: 1, OrderTime: time.Now()}, nil
}
