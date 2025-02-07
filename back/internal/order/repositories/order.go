package order_repositories

import (
	"context"

	order_domain "github.com/DongnutLa/stockio/internal/order/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type IOrderRepository interface {
	shared_ports.Repository[order_domain.Order, any]
}

type OrderRepository struct {
	shared_ports.Repository[order_domain.Order, any]
}

func NewOrderRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) IOrderRepository {
	repo := shared_repositories.BuildNewRepository[order_domain.Order, any](ctx, collection, connection, logger)
	return &OrderRepository{
		repo,
	}
}
