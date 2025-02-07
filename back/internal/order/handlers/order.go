package order_handlers

import (
	order_ports "github.com/DongnutLa/stockio/internal/order/core/ports"
)

type OrderHandlers struct {
	orderService order_ports.OrderService
}

// !TEST
var _ order_ports.OrderHandlers = (*OrderHandlers)(nil)

func NewOrderHandlers(orderService order_ports.OrderService) *OrderHandlers {
	return &OrderHandlers{
		orderService: orderService,
	}
}
