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

	ErrFailedProductCreate = NewApiError(
		"Failed creating product",
		ProductErrors,
		fiber.StatusInternalServerError,
		10,
	)
	ErrFailedProductUpdate = NewApiError(
		"Failed updating product",
		ProductErrors,
		fiber.StatusInternalServerError,
		11,
	)

	ErrFailedToParseBody = NewApiError(
		"Failed to parse body",
		GeneralErrors,
		fiber.StatusInternalServerError,
		14,
	)
	ErrInvalidParams = NewApiError(
		"Invalid params",
		GeneralErrors,
		fiber.StatusBadRequest,
		15,
	)
	ErrGenerateToken = NewApiError(
		"Unable to generate token access",
		GeneralErrors,
		fiber.StatusInternalServerError,
		16,
	)
)
