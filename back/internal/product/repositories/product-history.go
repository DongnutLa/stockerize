package product_repositories

import (
	"context"

	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductHistoryRepository interface {
	shared_ports.Repository[product_domain.History, any]
}

type ProductHistoryRepository struct {
	shared_ports.Repository[product_domain.History, any]
}

func NewProductHistoryRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) IProductHistoryRepository {
	repo := shared_repositories.BuildNewRepository[product_domain.History, any](ctx, collection, connection, logger)
	return &ProductHistoryRepository{
		repo,
	}
}
