package user_services

import (
	"context"

	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	user_ports "github.com/DongnutLa/stockio/internal/user/core/ports"
	user_repositories "github.com/DongnutLa/stockio/internal/user/repositories"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/rs/zerolog"
)

type UserService struct {
	logger   *zerolog.Logger
	userRepo user_repositories.IUserRepository
}

var _ user_ports.UserService = (*UserService)(nil)

func NewUserService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository user_repositories.IUserRepository,
) *UserService {
	return &UserService{
		logger:   logger,
		userRepo: repository,
	}
}

func (u *UserService) ListUsers(ctx context.Context, topic string) (*[]user_domain.User, *shared_domain.ApiError) {
	users := []user_domain.User{}

	filter := map[string]interface{}{}
	if topic != "" {
		filter["topics"] = topic
	}

	opts := shared_ports.FindManyOpts{
		Filter: filter,
	}

	_, err := u.userRepo.FindMany(ctx, opts, &users, false)
	if err != nil {
		u.logger.Error().Err(err).Msg("Failed to fetch users")
		return nil, shared_domain.ErrFetchUser
	}

	return &users, nil
}
