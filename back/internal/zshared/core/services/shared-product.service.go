package shared_services

import (
	"context"
	"math"
	"slices"

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

	priceItem, ok := lo.Find(*product.Prices, func(priceItem product_domain.Price) bool {
		return priceItem.SubUnit == stockDto.UnitPrice.SubUnit
	})
	if !ok {
		priceItem, ok = lo.Find(*product.Prices, func(priceItem product_domain.Price) bool {
			return priceItem.SubUnit == 1
		})
		if !ok {
			priceItem = (*product.Prices)[0]
		}
	}

	stock := *product.Stock
	slices.SortStableFunc(stock, utils.SortStock)

	historyType := product_domain.ProductHistoryInfo
	switch stockDto.UpdateType {
	case product_domain.StockInfo:
		newStock := stock[len(stock)-1]
		newStock.Cost = stockDto.Cost

		stock[len(stock)-1] = newStock
		historyType = product_domain.ProductHistoryInfo

	case product_domain.StockIncrease:
		qty := stockDto.Quantity * priceItem.SubUnit

		newStock := stock[len(stock)-1]
		newStock.Available += qty
		newStock.Quantity += qty

		stock[len(stock)-1] = newStock
		historyType = product_domain.ProductHistoryIncrease

	case product_domain.StockDecrease:
		qty := stockDto.Quantity * priceItem.SubUnit

		toRelease := qty
		newStock := lo.Map(stock, stockDecrease(toRelease))

		stock = newStock
		historyType = product_domain.ProductHistoryDecrease

	case product_domain.StockPurchase:
		qty := stockDto.Quantity * priceItem.SubUnit

		newStock := product_domain.NewStock(stockDto.Cost, qty, qty, 0)
		stock = append(stock, *newStock)
		historyType = product_domain.ProductHistoryPurchase

	case product_domain.StockSale:
		toRelease := stockDto.Quantity * priceItem.SubUnit

		newStock := lo.Map(stock, stockSale(toRelease))

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

	currCost := (*res.Stock)[len(*res.Stock)-1].Cost

	s.CreateHistory(
		ctx,
		historyType,
		stockDto.Quantity,
		(priceItem.Price)*stockDto.Quantity,
		(priceItem.Price-currCost)*stockDto.Quantity,
		product.Unit,
		priceItem.SubUnit,
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
	quantity, price, gain float64,
	unit product_domain.ProductUnit,
	subUnit float64,
	productId, productName, productSku string,
	user *user_domain.User,
) {
	history := product_domain.NewHistory(
		hType,
		quantity,
		price,
		gain,
		unit,
		subUnit,
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

func stockDecrease(toRelease float64) func(item product_domain.Stock, _ int) product_domain.Stock {
	return func(item product_domain.Stock, _ int) product_domain.Stock {
		if toRelease <= 0 || item.Available <= 0 {
			return item
		}

		newAvailable := math.Max(0, item.Available-toRelease)
		newQuantity := math.Max(0, item.Quantity-toRelease)

		if newAvailable > 0 {
			toRelease = 0
		} else {
			// New available == 0
			toRelease -= item.Available
		}

		item.Quantity = newQuantity
		item.Available = newAvailable

		return item
	}
}

func stockSale(toRelease float64) func(item product_domain.Stock, _ int) product_domain.Stock {
	return func(item product_domain.Stock, _ int) product_domain.Stock {
		if toRelease <= 0 || item.Available <= 0 {
			return item
		}

		addToSold := float64(0)
		newAvailable := math.Max(0, item.Available-toRelease)
		if newAvailable > 0 {
			addToSold = toRelease
			toRelease = 0
		} else {
			// New available == 0
			toRelease -= item.Available
			addToSold = item.Available
		}

		item.Available = newAvailable
		item.Sold += addToSold

		return item
	}
}
