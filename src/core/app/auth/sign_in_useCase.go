package authcase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	authdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/auth"
	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
	authpkg "github.com/Giovani-Coelho/Doti-API/src/pkg/auth"
	"go.uber.org/zap"
)

type SignInUseCase struct {
	UserRepository repository.IUserRepository
}

type ISignInUseCase interface {
	Execute(ctx context.Context, userDTO userdomain.IUserDomain) (userdomain.IUserDomain, string, error)
}

func NewLoginUseCase(
	userRepository repository.IUserRepository,
) ISignInUseCase {
	return &SignInUseCase{
		UserRepository: userRepository,
	}
}

func (su *SignInUseCase) Execute(
	ctx context.Context,
	userEntiy userdomain.IUserDomain,
) (userdomain.IUserDomain, string, error) {
	logger.Info("Init Sign-In UseCase",
		zap.String("journey", "sign-in"),
	)

	if userEntiy.GetEmail() == "" || userEntiy.GetPassword() == "" {
		logger.Error(
			"Error: Email or Password is missing", nil,
			zap.String("journey", "sign-in"),
		)

		return nil, "", userdomain.ErrSignInValuesMissing()
	}

	userEntiy.EncryptPassword()
	user, err := su.UserRepository.FindUserByEmailAndPassword(ctx, userEntiy)

	if err != nil {
		logger.Error(
			"Error: Could not find user with those credentials", err,
			zap.String("journey", "sign-in"),
		)

		return nil, "", userdomain.ErrCouldNotFindUser()
	}

	token, err := authpkg.GenerateToken(user)

	if err != nil {
		logger.Error(
			"Error: Could not generate token", err,
			zap.String("journey", "sign-in"),
		)

		return nil, "", authdomain.ErrGeneratingToken()
	}

	logger.Info(
		"sign-in executed successfully",
		zap.String("token:", token),
		zap.String("journey", "sign-in"),
	)

	return user, token, nil
}
