package product_services

import (
	"context"
	"math"

	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	product_ports "github.com/DongnutLa/stockio/internal/product/core/ports"
	product_repositories "github.com/DongnutLa/stockio/internal/product/repositories"
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

type ProductService struct {
	logger      *zerolog.Logger
	productRepo product_repositories.IProductRepository
	historyRepo product_repositories.IProductHistoryRepository
}

func NewProductService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository product_repositories.IProductRepository,
	historyRepo product_repositories.IProductHistoryRepository,
) product_ports.IProductService {
	return &ProductService{
		logger:      logger,
		productRepo: repository,
		historyRepo: historyRepo,
	}
}

func (s *ProductService) CreateProduct(
	ctx context.Context,
	productDto *product_domain.CreateProductDTO,
	authUser *user_domain.User,
) (*product_domain.Product, *shared_domain.ApiError) {
	s.logger.Info().Interface("productDto", productDto).Interface("user", authUser).Msg("Attempt to create product")

	stock := product_domain.NewStock(
		productDto.Stock.Cost,
		productDto.Stock.Price,
		productDto.Stock.Quantity,
		productDto.Stock.Quantity,
		0,
	)

	product := product_domain.NewProduct(
		productDto.Name,
		productDto.Sku,
		&authUser.Store,
		&[]product_domain.Stock{*stock},
		productDto.Unit,
	)

	if err := s.productRepo.InsertOne(ctx, *product); err != nil {
		s.logger.Error().Err(err).Interface("product", product).Msg("Product creation failed")
		return nil, shared_domain.ErrFailedProductCreate
	}

	s.logger.Info().Interface("product", product).Msg("Product successfully created")

	s.createHistory(
		ctx,
		product_domain.ProductHistoryPurchase,
		stock.Quantity,
		stock.Price,
		stock.Price-stock.Cost,
		product.Unit,
		product.ID.Hex(),
		product.Name,
		product.Sku,
		authUser,
	)

	return product, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, productDto *product_domain.UpdateProductDTO, authUser *user_domain.User) (*product_domain.Product, *shared_domain.ApiError) {
	s.logger.Info().Interface("productDto", productDto).Interface("user", authUser).Msg("Attempt to update product")

	opts := shared_ports.FindOneOpts{Filter: map[string]interface{}{"_id": productDto.ID}}
	product := product_domain.Product{}
	if err := s.productRepo.FindOne(ctx, opts, &product); err != nil {
		s.logger.Error().Err(err).Interface("product", productDto).Msg("Product update failed in find")
		return nil, shared_domain.ErrFailedProductUpdate
	}

	updOpts := shared_ports.UpdateOpts{
		Filter: map[string]interface{}{
			"_id": product.ID,
		},
		Payload: &map[string]interface{}{
			"name": productDto.Name,
		},
	}

	res, err := s.productRepo.UpdateOne(ctx, updOpts)
	if err != nil {
		s.logger.Error().Err(err).Interface("product", productDto).Msg("Product update failed")
		return nil, shared_domain.ErrFailedProductUpdate
	}

	s.createHistory(
		ctx,
		product_domain.ProductHistoryInfo,
		0,
		0,
		0,
		"",
		productDto.ID,
		productDto.Name,
		productDto.Sku,
		authUser,
	)

	return res, nil
}

func (s *ProductService) UpdateProductStock(ctx context.Context, stockDto *product_domain.UpdateStockDTO, authUser *user_domain.User) (*product_domain.Product, *shared_domain.ApiError) {
	s.logger.Info().Interface("stockDto", stockDto).Interface("user", authUser).Msg("Attempt to update product stock")

	opts := shared_ports.FindOneOpts{Filter: map[string]interface{}{"_id": stockDto.ID}}
	product := product_domain.Product{}
	if err := s.productRepo.FindOne(ctx, opts, &product); err != nil {
		s.logger.Error().Err(err).Interface("product", stockDto).Msg("Product update failed in find")
		return nil, shared_domain.ErrFailedProductUpdate
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
	case product_domain.StockAdd:
		newStock := stock[len(stock)-1]
		newStock.Available += stockDto.Quantity
		newStock.Quantity += stockDto.Quantity

		stock[len(stock)-1] = newStock
		historyType = product_domain.ProductHistoryIncrease
	case product_domain.StockRelease:
		toRelease := stockDto.Quantity
		newStock := lo.Map(stock, func(item product_domain.Stock, _ int) product_domain.Stock {
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
		})

		stock = newStock
		historyType = product_domain.ProductHistoryDecrease
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
		s.logger.Error().Err(err).Interface("stockDto", stockDto).Msg("Product stock update failed")
		return nil, shared_domain.ErrFailedProductUpdate
	}

	s.createHistory(
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

func (s *ProductService) createHistory(
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
