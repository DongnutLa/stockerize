package user_handlers

import (
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	user_ports "github.com/DongnutLa/stockio/internal/user/core/ports"
	"github.com/DongnutLa/stockio/internal/zshared/constants"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService user_ports.IUserService
}

func NewUserHandler(userService user_ports.IUserService) user_ports.IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (c *UserHandler) GoogleLogin(fiberCtx *fiber.Ctx) error {
	tokenUser := fiberCtx.Locals(constants.AUTH_USER_TOKEN_KEY).(*shared_domain.Claims)
	user, err := c.userService.GoogleLogin(fiberCtx.Context(), tokenUser)
	if err != nil {
		return fiberCtx.Status(err.HttpStatusCode).JSON(err)
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(user)
}

func (c *UserHandler) Login(fiberCtx *fiber.Ctx) error {
	loginDto := user_domain.LoginDTO{}
	err := fiberCtx.BodyParser(&loginDto)
	if err != nil {
		bodyErr := shared_domain.ErrFailedToParseBody
		return fiberCtx.Status(bodyErr.HttpStatusCode).JSON(bodyErr)
	}

	user, apiErr := c.userService.Login(fiberCtx.Context(), &loginDto)
	if apiErr != nil {
		return fiberCtx.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(user)
}

func (c *UserHandler) CreateUser(fiberCtx *fiber.Ctx) error {
	auth := fiberCtx.Locals(constants.AUTH_USER_KEY)
	if auth == nil {
		authErr := shared_domain.ErrAuthUserNotFound
		return fiberCtx.Status(authErr.HttpStatusCode).JSON(authErr)
	}
	authUser := auth.(*user_domain.User)

	userDto := user_domain.CreateUserDTO{}
	err := fiberCtx.BodyParser(&userDto)
	if err != nil {
		bodyErr := shared_domain.ErrFailedToParseBody
		return fiberCtx.Status(bodyErr.HttpStatusCode).JSON(bodyErr)
	}

	user, apiErr := c.userService.CreateUser(fiberCtx.Context(), &userDto, authUser)
	if apiErr != nil {
		return fiberCtx.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(user)
}
