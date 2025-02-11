package order_services

import (
	"context"
	"math"
	"time"

	order_domain "github.com/DongnutLa/stockio/internal/order/core/domain"
	order_ports "github.com/DongnutLa/stockio/internal/order/core/ports"
	order_repositories "github.com/DongnutLa/stockio/internal/order/repositories"
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

type OrderService struct {
	logger    *zerolog.Logger
	orderRepo order_repositories.IOrderRepository
}

func NewOrderService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository order_repositories.IOrderRepository,
) order_ports.IOrderService {
	return &OrderService{
		logger:    logger,
		orderRepo: repository,
	}
}

func (s *OrderService) CreateOrder(
	ctx context.Context,
	orderDto *order_domain.CreateOrderDTO,
	authUser *user_domain.User,
) (*order_domain.Order, *shared_domain.ApiError) {
	s.logger.Info().Interface("orderDto", orderDto).Interface("user", authUser).Msg("Attempt to create order")

	if orderDto.Products == nil {
		return nil, shared_domain.ErrEmptyOrderProducts
	}

	products := *orderDto.Products
	if len(products) == 0 {
		return nil, shared_domain.ErrEmptyOrderProducts
	}

	totals := calculateOrderTotal(products, orderDto.Discount)
	if totals.Total != orderDto.Totals.Total {
		return nil, shared_domain.ErrInconsistentTotals
	}

	newOrder := order_domain.NewOrder(orderDto.Type, &products, authUser, &totals, orderDto.PaymentMethod)

	if err := s.orderRepo.InsertOne(ctx, *newOrder); err != nil {
		s.logger.Error().Err(err).Interface("newOrder", newOrder).Msg("Order creation failed")
		return nil, shared_domain.ErrFailedOrderCreate
	}

	s.logger.Info().Interface("newOrder", newOrder).Msg("Order successfully created")

	return newOrder, nil
}

func (s *OrderService) UpdateOrder(
	ctx context.Context,
	orderDto *order_domain.UpdateOrderDTO,
	authUser *user_domain.User,
) (*order_domain.Order, *shared_domain.ApiError) {
	s.logger.Info().Interface("orderDto", orderDto).Interface("user", authUser).Msg("Attempt to create order")

	opts := shared_ports.FindOneOpts{Filter: map[string]interface{}{"_id": orderDto.ID}}
	order := order_domain.Order{}
	if err := s.orderRepo.FindOne(ctx, opts, &order); err != nil {
		s.logger.Error().Err(err).Interface("order", orderDto).Msg("Order update failed in find")
		return nil, shared_domain.ErrFailedOrderUpdate
	}

	if orderDto.Products == nil {
		return nil, shared_domain.ErrEmptyOrderProducts
	}

	products := *orderDto.Products
	if len(products) == 0 {
		return nil, shared_domain.ErrEmptyOrderProducts
	}

	totals := calculateOrderTotal(products, orderDto.Discount)
	if totals.Total != orderDto.Totals.Total {
		return nil, shared_domain.ErrInconsistentTotals
	}

	now := time.Now()

	updOpts := shared_ports.UpdateOpts{
		Filter: map[string]interface{}{
			"_id": order.ID,
		},
		Payload: &map[string]interface{}{
			"products":      products,
			"totals":        totals,
			"paymentMethod": orderDto.PaymentMethod,
			"updatedAt":     &now,
		},
	}

	res, err := s.orderRepo.UpdateOne(ctx, updOpts)
	if err != nil {
		s.logger.Error().Err(err).Interface("order", orderDto).Msg("Order update failed")
		return nil, shared_domain.ErrFailedOrderUpdate
	}

	return res, nil
}

func calculateOrderTotal(products []order_domain.OrderProduct, discount float64) order_domain.Totals {
	reduceFunc := func(
		acc float64,
		curr order_domain.OrderProduct,
		_ int,
	) float64 {
		acc += curr.Price * float64(curr.Quantity)
		return acc
	}

	subtotal := lo.Reduce(products, reduceFunc, 0)
	total := math.Max(0, subtotal-discount)

	newTotals := order_domain.NewTotals(subtotal, discount, total)
	return *newTotals
}
