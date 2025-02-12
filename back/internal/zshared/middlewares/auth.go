package middlewares

import (
	"encoding/json"
	"fmt"
	"strings"

	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	user_repositories "github.com/DongnutLa/stockio/internal/user/repositories"
	"github.com/DongnutLa/stockio/internal/zshared/constants"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/DongnutLa/stockio/internal/zshared/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"google.golang.org/api/idtoken"
)

var jwtService shared_ports.IJwtService
var userRepo user_repositories.IUserRepository

func NewAuthMiddleware(
	jwt shared_ports.IJwtService,
	repo user_repositories.IUserRepository,
	logger *zerolog.Logger,
	googleProvider bool,
) fiber.Handler {
	jwtService = jwt
	userRepo = repo

	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		headerSplitted := strings.Split(header, " ")
		if len(header) == 0 || len(headerSplitted) != 2 {
			apiErr := shared_domain.ErrInvalidToken
			logger.Log().Msg("TOKEN ERROR: Invalid Token")
			return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
		}

		token := headerSplitted[1]

		var email string
		if googleProvider {
			validator, err := idtoken.NewValidator(c.Context())
			if err != nil {
				logger.Log().Msg(fmt.Sprintf("TOKEN ERROR %s", err.Error()))
				apiErr := shared_domain.ErrInvalidToken
				return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
			}

			googleId := utils.GetConfig("google_id")
			payload, err := validator.Validate(c.Context(), token, googleId)
			if err != nil || payload == nil {
				logger.Log().Msg(fmt.Sprintf("TOKEN ERROR %s", err.Error()))
				apiErr := shared_domain.ErrInvalidToken
				return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
			}

			var googleClaims shared_domain.GoogleClaims
			jsonData, _ := json.Marshal(payload.Claims)
			json.Unmarshal(jsonData, &googleClaims)
			email = googleClaims.Email
		} else {
			claims, apiErr := jwtService.VerifyJWT(token)
			if apiErr != nil {
				return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
			}
			email = claims.Email
		}

		opts := shared_ports.FindOneOpts{
			Filter: map[string]interface{}{"email": email},
		}

		var user user_domain.User
		err := userRepo.FindOne(c.Context(), opts, &user)
		if err != nil {
			logger.Err(err).Msg("ERROR | Getting user in auth middleware")
			apiErr := shared_domain.ErrAuthUserNotFound
			return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
		}

		c.Locals(constants.AUTH_USER_KEY, &user)
		logger.Info().Interface("user", user).Msg("INFO | Auth middleware")
		return c.Next()
	}
}
