package shared_domain

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}
