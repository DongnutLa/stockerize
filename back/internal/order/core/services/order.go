package order_services

import (
	"context"

	order_ports "github.com/DongnutLa/stockio/internal/order/core/ports"
	order_repositories "github.com/DongnutLa/stockio/internal/order/repositories"
	"github.com/rs/zerolog"
)

type OrderService struct {
	logger    *zerolog.Logger
	orderRepo order_repositories.IOrderRepository
}

var _ order_ports.OrderService = (*OrderService)(nil)

func NewOrderService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository order_repositories.IOrderRepository,
) *OrderService {
	return &OrderService{
		logger:    logger,
		orderRepo: repository,
	}
}
