package order_handlers

import (
	order_domain "github.com/DongnutLa/stockio/internal/order/core/domain"
	order_ports "github.com/DongnutLa/stockio/internal/order/core/ports"
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	"github.com/DongnutLa/stockio/internal/zshared/constants"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	"github.com/gofiber/fiber/v2"
)

type OrderHandlers struct {
	orderService order_ports.IOrderService
}

func NewOrderHandlers(orderService order_ports.IOrderService) order_ports.IOrderHandler {
	return &OrderHandlers{
		orderService: orderService,
	}
}

func (h *OrderHandlers) CreateOrder(fiberCtx *fiber.Ctx) error {
	authUser := fiberCtx.Locals(constants.AUTH_USER_KEY).(*user_domain.User)
	if authUser == nil {
		authErr := shared_domain.ErrAuthUserNotFound
		return fiberCtx.Status(authErr.HttpStatusCode).JSON(authErr)
	}

	orderDto := order_domain.CreateOrderDTO{}
	err := fiberCtx.BodyParser(&orderDto)
	if err != nil {
		bodyErr := shared_domain.ErrFailedToParseBody
		return fiberCtx.Status(bodyErr.HttpStatusCode).JSON(bodyErr)
	}

	order, apiErr := h.orderService.CreateOrder(fiberCtx.Context(), &orderDto, authUser)
	if apiErr != nil {
		return fiberCtx.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(order)
}

func (h *OrderHandlers) UpdateOrder(fiberCtx *fiber.Ctx) error {
	authUser := fiberCtx.Locals(constants.AUTH_USER_KEY).(*user_domain.User)
	if authUser == nil {
		authErr := shared_domain.ErrAuthUserNotFound
		return fiberCtx.Status(authErr.HttpStatusCode).JSON(authErr)
	}

	orderDto := order_domain.UpdateOrderDTO{}
	err := fiberCtx.BodyParser(&orderDto)
	if err != nil {
		bodyErr := shared_domain.ErrFailedToParseBody
		return fiberCtx.Status(bodyErr.HttpStatusCode).JSON(bodyErr)
	}

	order, apiErr := h.orderService.UpdateOrder(fiberCtx.Context(), &orderDto, authUser)
	if apiErr != nil {
		return fiberCtx.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(order)
}
