package shared_domain

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrInvalidCredentials = NewApiError(
		"Invalid credentials",
		AuthErrors,
		fiber.StatusUnauthorized,
		1,
	)
	ErrInvalidToken = NewApiError(
		"Invalid Token",
		AuthErrors,
		fiber.StatusUnauthorized,
		2,
	)

	ErrUploadFile = NewApiError(
		"Failed to upload file",
		FileErrors,
		fiber.StatusInternalServerError,
		3,
	)
	ErrFetchFile = NewApiError(
		"Failed to fetch file",
		FileErrors,
		fiber.StatusInternalServerError,
		4,
	)

	ErrUserNotFound = NewApiError(
		"User not found",
		UserErrors,
		fiber.StatusNotFound,
		5,
	)
	ErrAuthUserNotFound = NewApiError(
		"Auth user not found",
		UserErrors,
		fiber.StatusUnauthorized,
		6,
	)
	ErrFailedUserCreate = NewApiError(
		"Failed creating user",
		UserErrors,
		fiber.StatusInternalServerError,
		7,
	)
	ErrFetchUser = NewApiError(
		"Failed to fetch users",
		UserErrors,
		fiber.StatusInternalServerError,
		8,
	)

	ErrFailedStoreCreate = NewApiError(
		"Failed creating user",
		StoreErrors,
		fiber.StatusInternalServerError,
		9,
	)

	ErrFailedSearchProducts = NewApiError(
		"Failed searching products",
		ProductErrors,
		fiber.StatusInternalServerError,
		10,
	)
	ErrFailedProductCreate = NewApiError(
		"Failed creating product",
		ProductErrors,
		fiber.StatusInternalServerError,
		11,
	)
	ErrProductSkuExists = NewApiError(
		"Product SKU already exists",
		ProductErrors,
		fiber.StatusBadRequest,
		12,
	)
	ErrFailedProductUpdate = NewApiError(
		"Failed updating product",
		ProductErrors,
		fiber.StatusInternalServerError,
		13,
	)
	ErrFailedProductStockUpdate = NewApiError(
		"Failed updating product stock",
		ProductErrors,
		fiber.StatusInternalServerError,
		14,
	)

	ErrFailedFetchProductHistory = NewApiError(
		"Failed fetching product history",
		ProductErrors,
		fiber.StatusInternalServerError,
		15,
	)

	ErrFailedGetOrder = NewApiError(
		"Failed to get order",
		OrderErrors,
		fiber.StatusInternalServerError,
		16,
	)
	ErrEmptyOrderProducts = NewApiError(
		"Order products must not be empty",
		OrderErrors,
		fiber.StatusBadRequest,
		17,
	)
	ErrInconsistentTotals = NewApiError(
		"Order totals are not consistent",
		OrderErrors,
		fiber.StatusBadRequest,
		18,
	)
	ErrFailedOrderCreate = NewApiError(
		"Failed creating order",
		OrderErrors,
		fiber.StatusInternalServerError,
		19,
	)
	ErrFailedOrderUpdate = NewApiError(
		"Failed updating order",
		OrderErrors,
		fiber.StatusInternalServerError,
		20,
	)
	ErrFailedGetSummary = NewApiError(
		"Failed getting orders summary",
		OrderErrors,
		fiber.StatusInternalServerError,
		21,
	)

	ErrFailedToParseBody = NewApiError(
		"Failed to parse body",
		GeneralErrors,
		fiber.StatusInternalServerError,
		22,
	)
	ErrInvalidParams = NewApiError(
		"Invalid params",
		GeneralErrors,
		fiber.StatusBadRequest,
		23,
	)
	ErrGenerateToken = NewApiError(
		"Unable to generate token access",
		GeneralErrors,
		fiber.StatusInternalServerError,
		24,
	)
)
