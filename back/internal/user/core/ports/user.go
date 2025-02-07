package user_ports

import (
	"context"

	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	ListUsers(ctx context.Context, topic string) (*[]user_domain.User, *shared_domain.ApiError)
}

type UserHandlers interface {
	ListUsers(c *fiber.Ctx) error
}
