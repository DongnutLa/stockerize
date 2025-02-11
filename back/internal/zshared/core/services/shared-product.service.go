package shared_services

import (
	"context"

	product_repositories "github.com/DongnutLa/stockio/internal/product/repositories"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/rs/zerolog"
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

func (s *SharedProductService) HandleStock(c context.Context, payload map[string]interface{}, topic string) {
	//
}
