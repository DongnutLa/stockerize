package order_services

import (
	"context"
	"math"
	"time"

	order_domain "github.com/DongnutLa/stockio/internal/order/core/domain"
	order_ports "github.com/DongnutLa/stockio/internal/order/core/ports"
	order_repositories "github.com/DongnutLa/stockio/internal/order/repositories"
	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

type OrderService struct {
	logger        *zerolog.Logger
	orderRepo     order_repositories.IOrderRepository
	messaging     shared_ports.IEventMessaging
	consecService shared_ports.IConsecutiveService
}

func NewOrderService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository order_repositories.IOrderRepository,
	messaging shared_ports.IEventMessaging,
	consecService shared_ports.IConsecutiveService,
) order_ports.IOrderService {
	return &OrderService{
		logger:        logger,
		orderRepo:     repository,
		messaging:     messaging,
		consecService: consecService,
	}
}

func (s *OrderService) ListOrders(
	ctx context.Context,
	queryParams *order_domain.OrdersQueryParams,
	authUser *user_domain.User,
) (*shared_domain.PagingResponse[order_domain.Order], *shared_domain.ApiError) {
	skip := (queryParams.Page - 1) * queryParams.PageSize
	result := []order_domain.Order{}

	filter := map[string]interface{}{
		"user.store._id": authUser.Store.ID,
		"type":           string(queryParams.OrderType),
	}
	if queryParams.Search != "" {
		filter["$text"] = map[string]string{"$search": queryParams.Search}
	}

	opts := shared_ports.FindManyOpts{
		Take:   queryParams.PageSize,
		Skip:   skip,
		Filter: filter,
		Sort: map[string]interface{}{
			"createdAt": -1,
		},
	}

	count, err := s.orderRepo.FindMany(ctx, opts, &result, true)
	if err != nil {
		return nil, shared_domain.ErrFailedGetOrder
	}

	response := shared_domain.PagingResponse[order_domain.Order]{
		Metadata: shared_domain.Metadata{
			Page:     queryParams.Page,
			PageSize: queryParams.PageSize,
			Count:    *count,
			HasNext:  *count > (queryParams.Page * queryParams.PageSize),
		},
		Data: result,
	}

	return &response, nil
}

func (s *OrderService) GetById(ctx context.Context, id string, authUser *user_domain.User) (*order_domain.Order, *shared_domain.ApiError) {
	order, apiErr := s.findOrderById(ctx, id, authUser)
	if apiErr != nil {
		return nil, apiErr
	}

	return order, nil
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

	totals := calculateOrderTotal(products, orderDto.Discount, orderDto.Type)
	if totals.Total != orderDto.Totals.Total {
		s.logger.Error().Float64("total", totals.Total).Float64("total_dto", orderDto.Totals.Total).Msg("Order totals inconsistents")
		return nil, shared_domain.ErrInconsistentTotals
	}

	consecutive, err := s.consecService.GetConsecutive(ctx, shared_domain.ConsecutiveType(orderDto.Type))
	if err != nil {
		s.logger.Error().Err(err).Msg("Order creation failed in getting consecutive")
		return nil, shared_domain.ErrFailedOrderCreate
	}

	newOrder := order_domain.NewOrder(orderDto.Type, &products, authUser, &totals, orderDto.PaymentMethod, consecutive)

	if err := s.orderRepo.InsertOne(ctx, *newOrder); err != nil {
		s.logger.Error().Err(err).Interface("newOrder", newOrder).Msg("Order creation failed")

		// Restore consecutive
		err := s.consecService.RestoreConsecutive(ctx, shared_domain.ConsecutiveType(orderDto.Type))
		if err != nil {
			s.logger.Error().Err(err).Msg("Failed restoring consecutive")
		}

		return nil, shared_domain.ErrFailedOrderCreate
	}

	s.logger.Info().Interface("newOrder", newOrder).Msg("Order successfully created")

	s.sendStockMessage(ctx, orderDto, products, authUser)
	return newOrder, nil
}

func (s *OrderService) UpdateOrder(
	ctx context.Context,
	orderDto *order_domain.UpdateOrderDTO,
	authUser *user_domain.User,
) (*order_domain.Order, *shared_domain.ApiError) {
	s.logger.Info().Interface("orderDto", orderDto).Interface("user", authUser).Msg("Attempt to create order")

	order, apiErr := s.findOrderById(ctx, orderDto.ID, authUser)
	if apiErr != nil {
		s.logger.Error().Err(apiErr).Interface("order", orderDto).Msg("Order update failed in find")
		return nil, shared_domain.ErrFailedOrderUpdate
	}

	if orderDto.Products == nil {
		return nil, shared_domain.ErrEmptyOrderProducts
	}

	// products := *orderDto.Products
	// if len(products) == 0 {
	// 	return nil, shared_domain.ErrEmptyOrderProducts
	// }

	// totals := calculateOrderTotal(products, orderDto.Discount)
	// if totals.Total != orderDto.Totals.Total {
	// 	return nil, shared_domain.ErrInconsistentTotals
	// }

	now := time.Now()

	updOpts := shared_ports.UpdateOpts{
		Filter: map[string]interface{}{
			"_id": order.ID,
		},
		Payload: &map[string]interface{}{
			// "products":      products,
			// "totals":        totals,
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

// PRIVATE METHODS
func (s *OrderService) findOrderById(ctx context.Context, id string, authUser *user_domain.User) (*order_domain.Order, *shared_domain.ApiError) {
	opts := shared_ports.FindOneOpts{
		Filter: map[string]interface{}{
			"_id":            id,
			"user.store._id": authUser.Store.ID,
		},
	}

	order := order_domain.Order{}
	if err := s.orderRepo.FindOne(ctx, opts, &order); err != nil {
		s.logger.Error().Err(err).Interface("order", id).Msg("Find order failed")
		return nil, shared_domain.ErrFailedGetOrder
	}

	return &order, nil
}

func calculateOrderTotal(products []order_domain.OrderProductDTO, discount float64, oType order_domain.OrderType) order_domain.Totals {
	reduceFunc := func(
		acc float64,
		curr order_domain.OrderProductDTO,
		_ int,
	) float64 {
		price := curr.Price
		if oType == order_domain.OrderTypePurchase {
			price = curr.Cost
		}

		acc += price * float64(curr.Quantity)
		return acc
	}

	subtotal := lo.Reduce(products, reduceFunc, 0)
	total := math.Max(0, subtotal-discount)

	newTotals := order_domain.NewTotals(subtotal, discount, total)
	return *newTotals
}

func (s *OrderService) sendStockMessage(
	ctx context.Context,
	orderDto *order_domain.CreateOrderDTO,
	products []order_domain.OrderProductDTO,
	authUser *user_domain.User,
) {
	var updateType product_domain.StockUpdateType
	if orderDto.Type == order_domain.OrderTypePurchase {
		updateType = product_domain.StockPurchase
	}
	if orderDto.Type == order_domain.OrderTypeSale {
		updateType = product_domain.StockSale
	}

	for _, product := range products {
		msg := shared_domain.MessageEvent{
			EventTopic: shared_domain.HandleStockTopic,
			Topic:      shared_domain.HandleEventsTopic,
			Data: map[string]interface{}{
				"updateStockDTO": &product_domain.UpdateStockDTO{
					ID:         product.ID,
					UpdateType: updateType,
					StockCreate: product_domain.StockCreate{
						Cost:      product.Cost,
						Quantity:  product.Quantity,
						UnitPrice: product.UnitPrice,
					},
				},
				"authUser": authUser,
			},
		}

		s.messaging.SendMessage(ctx, &msg)
		s.logger.Info().Interface("message", msg).Msgf("Stock message for product %s sent", product.Name)
	}
}
