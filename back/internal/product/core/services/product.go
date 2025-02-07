package product_services

import (
	"context"

	product_ports "github.com/DongnutLa/stockio/internal/product/core/ports"
	product_repositories "github.com/DongnutLa/stockio/internal/product/repositories"
	"github.com/rs/zerolog"
)

type ProductService struct {
	logger      *zerolog.Logger
	productRepo product_repositories.IProductRepository
}

var _ product_ports.ProductService = (*ProductService)(nil)

func NewProductService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository product_repositories.IProductRepository,
) *ProductService {
	return &ProductService{
		logger:      logger,
		productRepo: repository,
	}
}
