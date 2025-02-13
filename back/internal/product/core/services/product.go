package product_services

import (
	"context"
	"time"

	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	product_ports "github.com/DongnutLa/stockio/internal/product/core/ports"
	product_repositories "github.com/DongnutLa/stockio/internal/product/repositories"
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	logger               *zerolog.Logger
	productRepo          product_repositories.IProductRepository
	historyRepo          product_repositories.IProductHistoryRepository
	sharedProductService shared_ports.ISharedProductService
	messaging            shared_ports.IEventMessaging
}

func NewProductService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository product_repositories.IProductRepository,
	historyRepo product_repositories.IProductHistoryRepository,
	sharedProductService shared_ports.ISharedProductService,
	messaging shared_ports.IEventMessaging,
) product_ports.IProductService {
	return &ProductService{
		logger:               logger,
		productRepo:          repository,
		historyRepo:          historyRepo,
		sharedProductService: sharedProductService,
		messaging:            messaging,
	}
}

func (s *ProductService) GetById(ctx context.Context, id string, authUser *user_domain.User) (*product_domain.Product, *shared_domain.ApiError) {
	product, apiErr := s.findProductById(ctx, id, authUser)
	if apiErr != nil {
		return nil, apiErr
	}

	product.StockSummary = totalizeStock(*product.Stock)
	product.Stock = nil
	product.Store = nil

	return product, nil
}
func (s *ProductService) GetHistory(ctx context.Context, id string, authUser *user_domain.User) (*[]product_domain.History, *shared_domain.ApiError) {
	productId, _ := primitive.ObjectIDFromHex(id)
	var history []product_domain.History
	opts := shared_ports.FindManyOpts{
		Filter: map[string]interface{}{
			"productId":      productId,
			"user.store._id": authUser.Store.ID,
		},
		Sort: map[string]interface{}{
			"createdAt": 1,
		},
	}

	_, err := s.historyRepo.FindMany(ctx, opts, &history, false)
	if err != nil {
		apiErr := shared_domain.ErrFailedFetchProductHistory
		return nil, apiErr
	}

	newHistory := lo.Map(history, func(itt product_domain.History, _ int) product_domain.History {
		item := itt
		item.User = nil
		return item
	})

	return &newHistory, nil
}

func (s *ProductService) SearchProducts(
	ctx context.Context,
	queryParams *shared_domain.SearchQueryParams,
	authUser *user_domain.User,
) (*shared_domain.PagingResponse[product_domain.Product], *shared_domain.ApiError) {
	skip := (queryParams.Page - 1) * queryParams.PageSize
	result := []product_domain.Product{}
	opts := shared_ports.FindManyOpts{
		Take: queryParams.PageSize,
		Skip: skip,
		Filter: map[string]interface{}{
			"$text":     map[string]string{"$search": queryParams.Search},
			"store._id": authUser.Store.ID,
		},
	}
	count, err := s.productRepo.FindMany(ctx, opts, &result, true)
	if err != nil {
		return nil, shared_domain.ErrFailedSearchProducts
	}

	response := shared_domain.PagingResponse[product_domain.Product]{
		Metadata: shared_domain.Metadata{
			Page:     queryParams.Page,
			PageSize: queryParams.PageSize,
			Count:    *count,
			HasNext:  *count > (queryParams.Page * queryParams.PageSize),
		},
		Data: lo.Map(result, func(item product_domain.Product, _ int) product_domain.Product {
			newItem := item
			newItem.StockSummary = totalizeStock(*newItem.Stock)
			newItem.Stock = nil
			newItem.Store = nil
			return newItem
		}),
	}

	return &response, nil
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

	s.sharedProductService.CreateHistory(
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

	product, apiErr := s.findProductById(ctx, productDto.ID, authUser)
	if apiErr != nil {
		return nil, apiErr
	}

	now := time.Now()

	updOpts := shared_ports.UpdateOpts{
		Filter: map[string]interface{}{
			"_id": product.ID,
		},
		Payload: &map[string]interface{}{
			"name":      productDto.Name,
			"sku":       productDto.Sku,
			"updatedAt": &now,
		},
	}

	res, err := s.productRepo.UpdateOne(ctx, updOpts)
	if err != nil {
		s.logger.Error().Err(err).Interface("product", productDto).Msg("Product update failed")
		return nil, shared_domain.ErrFailedProductUpdate
	}

	s.sharedProductService.CreateHistory(
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

	return s.sharedProductService.UpdateProductStock(ctx, stockDto, authUser)
}

// PRIVATE METHODS
func (s *ProductService) findProductById(ctx context.Context, id string, authUser *user_domain.User) (*product_domain.Product, *shared_domain.ApiError) {
	opts := shared_ports.FindOneOpts{
		Filter: map[string]interface{}{
			"_id":       id,
			"store._id": authUser.Store.ID,
		},
	}

	product := product_domain.Product{}
	if err := s.productRepo.FindOne(ctx, opts, &product); err != nil {
		s.logger.Error().Err(err).Interface("product", id).Msg("Find product failed")
		return nil, shared_domain.ErrFailedProductUpdate
	}

	return &product, nil
}

func totalizeStock(stocks []product_domain.Stock) *product_domain.Stock {
	reduceFunc := func(acc, curr product_domain.Stock, _ int) product_domain.Stock {
		return product_domain.Stock{
			Available: acc.Available + curr.Available,
			Quantity:  acc.Quantity + curr.Quantity,
			Sold:      acc.Sold + curr.Sold,
			Cost:      curr.Cost,
			Price:     curr.Price,
			ID:        primitive.NilObjectID,
			CreatedAt: nil,
			UpdatedAt: nil,
		}
	}

	totalizedStock := lo.Reduce(stocks, reduceFunc, product_domain.Stock{})
	return &totalizedStock
}
