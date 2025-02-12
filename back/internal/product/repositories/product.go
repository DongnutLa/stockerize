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

type IProductRepository interface {
	shared_ports.Repository[product_domain.Product, any]
}

type ProductRepository struct {
	shared_ports.Repository[product_domain.Product, any]
}

func NewProductRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) IProductRepository {
	indexes := getProductIndexes()
	connection.Collection(collection).Indexes().CreateMany(ctx, indexes)

	repo := shared_repositories.BuildNewRepository[product_domain.Product, any](ctx, collection, connection, logger)
	return &ProductRepository{
		repo,
	}
}

func getProductIndexes() []mongo.IndexModel {
	search_product_index := "search_product_index"

	searchProductIdx := mongo.IndexModel{
		Keys: bson.D{
			{Key: "store._id", Value: 1},
			{Key: "name", Value: "text"},
			{Key: "sku", Value: "text"},
		},
		Options: &options.IndexOptions{
			Name: &search_product_index,
		},
	}

	return []mongo.IndexModel{searchProductIdx}
}
