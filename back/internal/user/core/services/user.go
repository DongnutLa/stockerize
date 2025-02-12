package user_services

import (
	"context"

	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	user_ports "github.com/DongnutLa/stockio/internal/user/core/ports"
	user_repositories "github.com/DongnutLa/stockio/internal/user/repositories"
	"github.com/DongnutLa/stockio/internal/zshared/constants"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
	shared_ports "github.com/DongnutLa/stockio/internal/zshared/core/ports"
	"github.com/DongnutLa/stockio/internal/zshared/utils"
	"github.com/rs/zerolog"
)

type UserService struct {
	logger     *zerolog.Logger
	userRepo   user_repositories.IUserRepository
	jwtService shared_ports.IJwtService
}

func NewUserService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository user_repositories.IUserRepository,
	jwtService shared_ports.IJwtService,
) user_ports.IUserService {
	return &UserService{
		logger:     logger,
		userRepo:   repository,
		jwtService: jwtService,
	}
}

func (u *UserService) GoogleLogin(
	ctx context.Context,
	tokenUser *shared_domain.Claims,
) (*user_domain.User, *shared_domain.ApiError) {
	opts := shared_ports.FindOneOpts{
		Filter: map[string]interface{}{
			"email": tokenUser.Email,
		},
	}

	authUser := user_domain.User{}

	err := u.userRepo.FindOne(ctx, opts, &authUser)
	if err != nil {
		return nil, shared_domain.ErrUserNotFound
	}

	return &authUser, nil
}

func (u *UserService) Login(ctx context.Context, loginDto *user_domain.LoginDTO) (*user_domain.User, *shared_domain.ApiError) {
	opts := shared_ports.FindOneOpts{
		Filter: map[string]interface{}{
			"email": loginDto.Email,
		},
	}

	authUser := user_domain.User{}

	err := u.userRepo.FindOne(ctx, opts, &authUser)
	if err != nil {
		return nil, shared_domain.ErrInvalidCredentials
	}

	if authUser.Password != loginDto.Password {
		return nil, shared_domain.ErrInvalidCredentials
	}

	authUser.Password = ""
	token, apiErr := u.jwtService.GenerateJWT(authUser)
	if apiErr != nil {
		return nil, apiErr
	}

	authUser.Token = token

	return &authUser, nil
}

func (u *UserService) CreateUser(ctx context.Context, userDto *user_domain.CreateUserDTO, authUser *user_domain.User) (*user_domain.User, *shared_domain.ApiError) {
	newUserRole := constants.ManagerRole
	if authUser.Role == constants.SudoRole {
		newUserRole = constants.AdminRole
	}

	password := utils.GeneratePassword(12, true, true, true)

	newUser := user_domain.NewUser(
		userDto.Name,
		userDto.Email,
		password,
		user_domain.UserActive,
		newUserRole,
		userDto.Store,
	)

	u.logger.Info().Interface("user", newUser).Msg("New user created with data")

	err := u.userRepo.InsertOne(ctx, *newUser)
	if err != nil {
		return nil, shared_domain.ErrFailedUserCreate
	}

	//!TODO: Send email

	return newUser, nil
}
