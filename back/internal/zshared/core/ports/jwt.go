package shared_ports

import (
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
)

type IJwtService interface {
	GenerateJWT(authUser user_domain.User) (string, *shared_domain.ApiError)
	VerifyJWT(tokenString string) (*shared_domain.Claims, *shared_domain.ApiError)
}
