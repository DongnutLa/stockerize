package store_services

import (
	"context"

	store_ports "github.com/DongnutLa/stockio/internal/store/core/ports"
	store_repositories "github.com/DongnutLa/stockio/internal/store/repositories"
	"github.com/rs/zerolog"
)

type StoreService struct {
	logger    *zerolog.Logger
	storeRepo store_repositories.IStoreRepository
}

var _ store_ports.StoreService = (*StoreService)(nil)

func NewStoreService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository store_repositories.IStoreRepository,
) *StoreService {
	return &StoreService{
		logger:    logger,
		storeRepo: repository,
	}
}
