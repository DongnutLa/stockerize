package user_handlers

import (
	"context"

	user_ports "github.com/DongnutLa/stockio/internal/user/core/ports"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userService user_ports.UserService
}

// !TEST
var _ user_ports.UserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService user_ports.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) ListUsers(c *fiber.Ctx) error {
	topic := c.Query("topic", "")
	users, apiErr := h.userService.ListUsers(context.TODO(), topic)
	if apiErr != nil {
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
