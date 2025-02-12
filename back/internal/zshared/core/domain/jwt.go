package shared_domain

import (
	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	user_domain.User
	jwt.RegisteredClaims
}

type GoogleClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}
