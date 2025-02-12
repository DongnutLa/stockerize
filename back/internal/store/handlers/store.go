package store_handlers

import (
	store_domain "github.com/DongnutLa/stockio/internal/store/core/domain"
	store_ports "github.com/DongnutLa/stockio/internal/store/core/ports"
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	"github.com/DongnutLa/stockio/internal/zshared/constants"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	"github.com/gofiber/fiber/v2"
)

type StoreHandler struct {
	storeService store_ports.IStoreService
}

func NewStoreHandlers(storeService store_ports.IStoreService) store_ports.IStoreHandler {
	return &StoreHandler{
		storeService: storeService,
	}
}

func (h *StoreHandler) CreateStore(fiberCtx *fiber.Ctx) error {
	auth := fiberCtx.Locals(constants.AUTH_USER_KEY)
	if auth == nil {
		authErr := shared_domain.ErrAuthUserNotFound
		return fiberCtx.Status(authErr.HttpStatusCode).JSON(authErr)
	}
	authUser := auth.(*user_domain.User)
	if authUser.Role != constants.SudoRole {
		authErr := shared_domain.ErrAuthUserNotFound
		return fiberCtx.Status(authErr.HttpStatusCode).JSON(authErr)
	}

	storeDto := store_domain.CreateStoreDTO{}
	err := fiberCtx.BodyParser(&storeDto)
	if err != nil {
		bodyErr := shared_domain.ErrFailedToParseBody
		return fiberCtx.Status(bodyErr.HttpStatusCode).JSON(bodyErr)
	}

	store, apiErr := h.storeService.CreateStore(fiberCtx.Context(), &storeDto)
	if apiErr != nil {
		return fiberCtx.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(store)
}
