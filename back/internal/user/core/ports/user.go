package user_ports

import (
	"context"

	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	"github.com/gofiber/fiber/v2"
)

type IUserService interface {
	CreateUser(ctx context.Context, userDto *user_domain.CreateUserDTO, authUser *user_domain.User) (*user_domain.User, *shared_domain.ApiError)
	GoogleLogin(ctx context.Context, tokenUser *shared_domain.Claims) (*user_domain.User, *shared_domain.ApiError)
}

type IUserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GoogleLogin(c *fiber.Ctx) error
}
