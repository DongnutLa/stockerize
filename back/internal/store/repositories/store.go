package store_repositories

import (
	"context"

	store_domain "github.com/DongnutLa/stockio/internal/store/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStoreRepository interface {
	shared_ports.Repository[store_domain.Store, any]
}

type StoreRepository struct {
	shared_ports.Repository[store_domain.Store, any]
}

func NewStoreRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) IStoreRepository {
	repo := shared_repositories.BuildNewRepository[store_domain.Store, any](ctx, collection, connection, logger)
	return &StoreRepository{
		repo,
	}
}
