package container

import (
	usercase "github.com/Giovani-Coelho/Doti-API/internal/core/app/user"
	userhandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/user"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

func (c *container) NewUser() userhandler.UserHandler {
	userRepository := repository.NewUserRepository(c.DB)

	createUserUseCase := usercase.NewCreateUserUseCase(userRepository)

	return userhandler.NewUserHandler(
		createUserUseCase,
	)
}
