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
	shared_repositories "github.com/DongnutLa/stockio/internal/zshared/repositories"
	"github.com/DongnutLa/stockio/internal/zshared/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"google.golang.org/api/idtoken"
)

func NewAuthMiddleware(logger *zerolog.Logger, getUser bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		mongo := shared_repositories.GetMongoConnection()
		repo := user_repositories.NewUserRepository(c.Context(), constants.USERS_COLLECTION, mongo, logger)

		header := c.Get("Authorization")
		headerSplitted := strings.Split(header, " ")
		if len(header) == 0 || len(headerSplitted) != 2 {
			apiErr := shared_domain.ErrInvalidToken
			logger.Log().Msg("TOKEN ERROR: Invalid Token")
			return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
		}

		token := headerSplitted[1]

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

		var claims shared_domain.Claims
		jsonData, _ := json.Marshal(payload.Claims)
		json.Unmarshal(jsonData, &claims)

		opts := shared_ports.FindOneOpts{
			Filter: map[string]interface{}{"email": claims.Email},
		}

		var user user_domain.User
		if getUser {
			err := repo.FindOne(c.Context(), opts, &user)
			if err != nil {
				logger.Err(err).Msg("ERROR | Getting user in auth middleware")
			} else {
				c.Locals(constants.AUTH_USER_KEY, &user)
				logger.Info().Interface("user", user).Msg("INFO | Auth middleware")
			}
		}

		c.Locals(constants.AUTH_USER_TOKEN_KEY, &claims)
		return c.Next()
	}
}
