package shared_services

import (
	"context"
	"math"

	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	product_repositories "github.com/DongnutLa/stockio/internal/product/repositories"
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/DongnutLa/stockio/internal/zshared/utils"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

type SharedProductService struct {
	logger      *zerolog.Logger
	productRepo product_repositories.IProductRepository
	historyRepo product_repositories.IProductHistoryRepository
}

func NewSharedProductService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository product_repositories.IProductRepository,
	historyRepo product_repositories.IProductHistoryRepository,
) shared_ports.ISharedProductService {
	return &SharedProductService{
		logger:      logger,
		productRepo: repository,
		historyRepo: historyRepo,
	}
}

func (s *SharedProductService) HandleStock(ctx context.Context, payload map[string]interface{}, topic string) {
	s.logger.Info().Interface("body", payload).Msgf("Message received for topic %s", topic)

	stockDto := utils.EventDataToStruct[product_domain.UpdateStockDTO](payload["updateStockDTO"])
	authUser := utils.EventDataToStruct[user_domain.User](payload["authUser"])

	s.UpdateProductStock(ctx, stockDto, authUser)
}

func (s *SharedProductService) UpdateProductStock(ctx context.Context, stockDto *product_domain.UpdateStockDTO, authUser *user_domain.User) (*product_domain.Product, *shared_domain.ApiError) {
	opts := shared_ports.FindOneOpts{Filter: map[string]interface{}{"_id": stockDto.ID}}
	product := product_domain.Product{}
	if err := s.productRepo.FindOne(ctx, opts, &product); err != nil {
		apiErr := shared_domain.ErrFailedProductStockUpdate
		s.logger.Error().Err(err).Interface("error", apiErr).Interface("product", stockDto).Msg("Product update failed in find")
		return nil, apiErr
	}

	stock := *product.Stock
	historyType := product_domain.ProductHistoryInfo
	switch stockDto.UpdateType {
	case product_domain.StockInfo:
		newStock := stock[len(stock)-1]
		newStock.Cost = stockDto.Cost
		newStock.Price = stockDto.Price

		stock[len(stock)-1] = newStock
		historyType = product_domain.ProductHistoryInfo
	case product_domain.StockIncrease:
		newStock := stock[len(stock)-1]
		newStock.Available += stockDto.Quantity
		newStock.Quantity += stockDto.Quantity

		stock[len(stock)-1] = newStock
		historyType = product_domain.ProductHistoryIncrease
	case product_domain.StockDecrease:
		toRelease := stockDto.Quantity
		newStock := lo.Map(stock, s.stockRelease(toRelease))

		stock = newStock
		historyType = product_domain.ProductHistoryDecrease
	case product_domain.StockPurchase:
		newStock := product_domain.NewStock(stockDto.Cost, stockDto.Price, stockDto.Quantity, stockDto.Quantity, 0)
		stock = append(stock, *newStock)
		historyType = product_domain.ProductHistoryPurchase
	case product_domain.StockSale:
		toRelease := stockDto.Quantity
		newStock := lo.Map(stock, s.stockRelease(toRelease))

		stock = newStock
		historyType = product_domain.ProductHistorySale
	}

	updOpts := shared_ports.UpdateOpts{
		Filter: map[string]interface{}{
			"_id": product.ID,
		},
		Payload: &map[string]interface{}{
			"stock": &stock,
		},
	}

	res, err := s.productRepo.UpdateOne(ctx, updOpts)
	if err != nil {
		apiErr := shared_domain.ErrFailedProductStockUpdate
		s.logger.Error().Err(err).Interface("error", apiErr).Interface("stockDto", stockDto).Msg("Product stock update failed")
		return nil, apiErr
	} else {
		s.logger.Info().Interface("product", res).Msg("Product stock updated successfully")
	}

	s.CreateHistory(
		ctx,
		historyType,
		stockDto.Quantity,
		stockDto.Price,
		stockDto.Price-stockDto.Cost,
		"",
		product.ID.Hex(),
		product.Name,
		product.Sku,
		authUser,
	)

	return res, nil
}

func (s *SharedProductService) CreateHistory(
	ctx context.Context,
	hType product_domain.ProductHistoryType,
	quantity int64,
	price, gain float64,
	unit product_domain.ProductUnit,
	productId, productName, productSku string,
	user *user_domain.User,
) {
	history := product_domain.NewHistory(
		hType,
		quantity,
		price,
		gain,
		unit,
		productId,
		productName,
		productSku,
		user,
	)
	if err := s.historyRepo.InsertOne(ctx, *history); err != nil {
		s.logger.Error().Err(err).Interface("history", history).Msg("Product history creation failed")
	} else {
		s.logger.Info().Interface("history", history).Msg("Product history successfully created")
	}
}

func (s *SharedProductService) stockRelease(toRelease int64) func(item product_domain.Stock, _ int) product_domain.Stock {
	return func(item product_domain.Stock, _ int) product_domain.Stock {
		if toRelease <= 0 {
			return item
		}

		newQuantity := int64(math.Max(0, float64(item.Available)-float64(toRelease)))
		if newQuantity > 0 {
			toRelease = 0
		} else {
			toRelease -= item.Available
		}

		item.Quantity -= item.Available
		item.Available = newQuantity
		return item
	}
}
