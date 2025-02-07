package user_repositories

import (
	"context"

	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	shared_ports.Repository[user_domain.User, any]
}

type UserRepository struct {
	shared_ports.Repository[user_domain.User, any]
}

func NewUserRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) IUserRepository {
	repo := shared_repositories.BuildNewRepository[user_domain.User, any](ctx, collection, connection, logger)
	return &UserRepository{
		repo,
	}
}
