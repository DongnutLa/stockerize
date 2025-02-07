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

	ErrUploadFile = NewApiError(
		"Failed to upload file",
		FileErrors,
		fiber.StatusInternalServerError,
		2,
	)
	ErrFetchFile = NewApiError(
		"Failed to fetch file",
		FileErrors,
		fiber.StatusInternalServerError,
		3,
	)

	ErrFetchUser = NewApiError(
		"Failed to fetch users",
		UserErrors,
		fiber.StatusInternalServerError,
		10,
	)

	ErrFailedToParseBody = NewApiError(
		"Failed to parse body",
		GeneralErrors,
		fiber.StatusInternalServerError,
		11,
	)
	ErrInvalidParams = NewApiError(
		"Invalid params",
		GeneralErrors,
		fiber.StatusBadRequest,
		12,
	)
	ErrGenerateToken = NewApiError(
		"Unable to generate token access",
		GeneralErrors,
		fiber.StatusInternalServerError,
		13,
	)
)
