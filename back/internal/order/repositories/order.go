package order_repositories

import (
	"context"

	order_domain "github.com/DongnutLa/stockio/internal/order/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IOrderRepository interface {
	shared_ports.Repository[order_domain.Order, any]
}

type OrderRepository struct {
	shared_ports.Repository[order_domain.Order, any]
}

func NewOrderRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) IOrderRepository {
	indexes := getOrderIndexes()
	connection.Collection(collection).Indexes().CreateMany(ctx, indexes)

	repo := shared_repositories.BuildNewRepository[order_domain.Order, any](ctx, collection, connection, logger)
	return &OrderRepository{
		repo,
	}
}

func getOrderIndexes() []mongo.IndexModel {
	search_order_index := "search_order_index"
	store_order_index := "store_order_index"

	searchOrderIdx := mongo.IndexModel{
		Keys: bson.D{
			{Key: "user.store._id", Value: 1},
			{Key: "type", Value: 1},
			{Key: "consecutive", Value: "text"},
			{Key: "createdAt", Value: -1},
		},
		Options: &options.IndexOptions{
			Name: &search_order_index,
		},
	}
	storeOrderIdx := mongo.IndexModel{
		Keys: bson.D{
			{Key: "_id", Value: 1},
			{Key: "user.store._id", Value: 1},
			{Key: "type", Value: 1},
			{Key: "createdAt", Value: -1},
		},
		Options: &options.IndexOptions{
			Name: &store_order_index,
		},
	}

	return []mongo.IndexModel{searchOrderIdx, storeOrderIdx}
}
