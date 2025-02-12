package product_handlers

import (
	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	product_ports "github.com/DongnutLa/stockio/internal/product/core/ports"
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	"github.com/DongnutLa/stockio/internal/zshared/constants"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService product_ports.IProductService
}

func NewProductHandlers(productService product_ports.IProductService) product_ports.IProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) SearchProducts(fiberCtx *fiber.Ctx) error {
	authUser := fiberCtx.Locals(constants.AUTH_USER_KEY).(*user_domain.User)
	if authUser == nil {
		authErr := shared_domain.ErrAuthUserNotFound
		return fiberCtx.Status(authErr.HttpStatusCode).JSON(authErr)
	}

	queryParams := shared_domain.SearchQueryParams{}
	if err := fiberCtx.QueryParser(&queryParams); err != nil {
		bodyErr := shared_domain.ErrFailedToParseBody
		return fiberCtx.Status(bodyErr.HttpStatusCode).JSON(bodyErr)
	}

	response, apiErr := h.productService.SearchProducts(fiberCtx.Context(), &queryParams, authUser)
	if apiErr != nil {
		return fiberCtx.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(response)
}

func (h *ProductHandler) CreateProduct(fiberCtx *fiber.Ctx) error {
	authUser := fiberCtx.Locals(constants.AUTH_USER_KEY).(*user_domain.User)
	if authUser == nil {
		authErr := shared_domain.ErrAuthUserNotFound
		return fiberCtx.Status(authErr.HttpStatusCode).JSON(authErr)
	}

	productDto := product_domain.CreateProductDTO{}
	err := fiberCtx.BodyParser(&productDto)
	if err != nil {
		bodyErr := shared_domain.ErrFailedToParseBody
		return fiberCtx.Status(bodyErr.HttpStatusCode).JSON(bodyErr)
	}

	product, apiErr := h.productService.CreateProduct(fiberCtx.Context(), &productDto, authUser)
	if apiErr != nil {
		return fiberCtx.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return fiberCtx.Status(fiber.StatusCreated).JSON(product)
}

func (h *ProductHandler) UpdateProduct(fiberCtx *fiber.Ctx) error {
	authUser := fiberCtx.Locals(constants.AUTH_USER_KEY).(*user_domain.User)
	if authUser == nil {
		authErr := shared_domain.ErrAuthUserNotFound
		return fiberCtx.Status(authErr.HttpStatusCode).JSON(authErr)
	}

	productDto := product_domain.UpdateProductDTO{}
	err := fiberCtx.BodyParser(&productDto)
	if err != nil {
		bodyErr := shared_domain.ErrFailedToParseBody
		return fiberCtx.Status(bodyErr.HttpStatusCode).JSON(bodyErr)
	}

	product, apiErr := h.productService.UpdateProduct(fiberCtx.Context(), &productDto, authUser)
	if apiErr != nil {
		return fiberCtx.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return fiberCtx.Status(fiber.StatusCreated).JSON(product)
}

func (h *ProductHandler) UpdateProductStock(fiberCtx *fiber.Ctx) error {
	authUser := fiberCtx.Locals(constants.AUTH_USER_KEY).(*user_domain.User)
	if authUser == nil {
		authErr := shared_domain.ErrAuthUserNotFound
		return fiberCtx.Status(authErr.HttpStatusCode).JSON(authErr)
	}

	stockDto := product_domain.UpdateStockDTO{}
	err := fiberCtx.BodyParser(&stockDto)
	if err != nil {
		bodyErr := shared_domain.ErrFailedToParseBody
		return fiberCtx.Status(bodyErr.HttpStatusCode).JSON(bodyErr)
	}

	product, apiErr := h.productService.UpdateProductStock(fiberCtx.Context(), &stockDto, authUser)
	if apiErr != nil {
		return fiberCtx.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return fiberCtx.Status(fiber.StatusCreated).JSON(product)
}
