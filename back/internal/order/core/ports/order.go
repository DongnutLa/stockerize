package order_ports

import (
	"context"

	order_domain "github.com/DongnutLa/stockio/internal/order/core/domain"
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	"github.com/gofiber/fiber/v2"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, productDto *order_domain.CreateOrderDTO, authUser *user_domain.User) (*order_domain.Order, *shared_domain.ApiError)
	UpdateOrder(ctx context.Context, productDto *order_domain.UpdateOrderDTO, authUser *user_domain.User) (*order_domain.Order, *shared_domain.ApiError)
}

type IOrderHandler interface {
	CreateOrder(c *fiber.Ctx) error
	UpdateOrder(c *fiber.Ctx) error
}
