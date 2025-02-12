package store_services

import (
	"context"

	store_domain "github.com/DongnutLa/stockio/internal/store/core/domain"
	store_ports "github.com/DongnutLa/stockio/internal/store/core/ports"
	store_repositories "github.com/DongnutLa/stockio/internal/store/repositories"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	"github.com/rs/zerolog"
)

type StoreService struct {
	logger    *zerolog.Logger
	storeRepo store_repositories.IStoreRepository
}

func NewStoreService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository store_repositories.IStoreRepository,
) store_ports.IStoreService {
	return &StoreService{
		logger:    logger,
		storeRepo: repository,
	}
}

func (s *StoreService) CreateStore(ctx context.Context, storeDto *store_domain.CreateStoreDTO) (*store_domain.Store, *shared_domain.ApiError) {
	newStore := store_domain.NewStore(
		storeDto.Name,
		storeDto.Contact,
		storeDto.Address,
		storeDto.City,
		storeDto.Status,
	)

	err := s.storeRepo.InsertOne(ctx, *newStore)
	if err != nil {
		return nil, shared_domain.ErrFailedStoreCreate
	}

	return newStore, nil
}
