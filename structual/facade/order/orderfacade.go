package order

import "time"

type OrderWrapper struct {
	OrderId   int64     `json:"orderId"`
	OrderTime time.Time `json:"orderTime"`
	Status    int
	Buyer     User    `json:"buyer"`
	Seller    User    `json:"seller"`
	Product   Product `json:"product"`
}

type OrderFacade struct {
	orderService   *OrderService
	userService    *UserService
	productService *ProductService
}

func (o *OrderFacade) GetOrder(orderId int64) (*OrderWrapper, error) {
	order, err := o.orderService.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}
	buyer, err := o.userService.GetUserById(order.UserId)
	if err != nil {
		return nil, err
	}
	seller, err := o.userService.GetUserById(order.SellerId)
	if err != nil {
		return nil, err
	}
	product, err := o.productService.GetProductById(order.ProductId)
	if err != nil {
		return nil, err
	}
	return &OrderWrapper{OrderId: orderId, OrderTime: order.OrderTime, Status: order.Status, Buyer: *buyer, Seller: *seller, Product: *product}, nil
}
