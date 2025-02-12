package store_ports

import (
	"context"

	store_domain "github.com/DongnutLa/stockio/internal/store/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	"github.com/gofiber/fiber/v2"
)

type IStoreService interface {
	CreateStore(ctx context.Context, storeDto *store_domain.CreateStoreDTO) (*store_domain.Store, *shared_domain.ApiError)
}

type IStoreHandler interface {
	CreateStore(c *fiber.Ctx) error
}
