package shared_services

import (
	"time"

	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog"
)

type JwtService struct {
	logger *zerolog.Logger
	jwtKey []byte
}

var _ shared_ports.JwtService = (*JwtService)(nil)

func NewJwtService(jwtKey []byte, logger *zerolog.Logger) *JwtService {
	return &JwtService{
		logger: logger,
		jwtKey: jwtKey,
	}
}

func (j *JwtService) GenerateJWT(id, email, name string) (string, *shared_domain.ApiError) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &shared_domain.Claims{
		ID:    id,
		Email: email,
		Name:  name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(j.jwtKey)
	if err != nil {
		return "", shared_domain.ErrGenerateToken
	}

	return tokenString, nil
}

func (j JwtService) VerifyJWT(tokenString string) (*shared_domain.Claims, *shared_domain.ApiError) {
	claims := &shared_domain.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			j.logger.Error().Msgf("Unexpected sign method: %v", token.Header["alg"])
			return nil, shared_domain.ErrInvalidCredentials
		}
		return j.jwtKey, nil
	})

	if err != nil {
		j.logger.Error().Err(err).Msg("Something went wrong validating token")
		return nil, shared_domain.ErrInvalidCredentials
	}

	if !token.Valid {
		j.logger.Error().Msg("Invalid token error")
		return nil, shared_domain.ErrInvalidCredentials
	}

	return claims, nil
}
