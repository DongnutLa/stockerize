package product_repositories

import (
	"context"

	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IProductHistoryRepository interface {
	shared_ports.Repository[product_domain.History, any]
}

type ProductHistoryRepository struct {
	shared_ports.Repository[product_domain.History, any]
}

func NewProductHistoryRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) IProductHistoryRepository {
	indexes := getProductHistoryIndexes()
	connection.Collection(collection).Indexes().CreateMany(ctx, indexes)

	repo := shared_repositories.BuildNewRepository[product_domain.History, any](ctx, collection, connection, logger)
	return &ProductHistoryRepository{
		repo,
	}
}

func getProductHistoryIndexes() []mongo.IndexModel {
	store_product_history_index := "store_product_history_index"

	storeProductHistoryIdx := mongo.IndexModel{
		Keys: bson.D{
			{Key: "_id", Value: 1},
			{Key: "user.store._id", Value: 1},
			{Key: "createdAt", Value: 1},
		},
		Options: &options.IndexOptions{
			Name: &store_product_history_index,
		},
	}

	return []mongo.IndexModel{storeProductHistoryIdx}
}
