package shared_repositories

import (
	"context"

	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type IConsecutiveRepository interface {
	shared_ports.Repository[shared_domain.Consecutive, any]
}

type ConsecutiveRepository struct {
	shared_ports.Repository[shared_domain.Consecutive, any]
}

func NewConsecutiveRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) IConsecutiveRepository {
	repo := BuildNewRepository[shared_domain.Consecutive, any](ctx, collection, connection, logger)
	return &ConsecutiveRepository{
		repo,
	}
}
