package product_repositories

import (
	"context"

	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	shared_ports.Repository[product_domain.Product, any]
}

type ProductRepository struct {
	shared_ports.Repository[product_domain.Product, any]
}

func NewproductRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) IProductRepository {
	repo := shared_repositories.BuildNewRepository[product_domain.Product, any](ctx, collection, connection, logger)
	return &ProductRepository{
		repo,
	}
}
