package shared_repositories

import (
	"context"

	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type IOrderSummaryRepository interface {
	shared_ports.Repository[shared_domain.OrderSummary, any]
}

type OrderSummaryRepository struct {
	shared_ports.Repository[shared_domain.OrderSummary, any]
}

func NewOrderSummaryRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) IOrderSummaryRepository {
	repo := BuildNewRepository[shared_domain.OrderSummary, any](ctx, collection, connection, logger)
	return &OrderSummaryRepository{
		repo,
	}
}
