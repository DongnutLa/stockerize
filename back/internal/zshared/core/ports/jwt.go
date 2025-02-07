package shared_ports

import shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"

type JwtService interface {
	GenerateJWT(id, email, name string) (string, *shared_domain.ApiError)
	VerifyJWT(tokenString string) (*shared_domain.Claims, *shared_domain.ApiError)
}
