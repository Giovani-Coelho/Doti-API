package userhandler

import (
	"context"
	"net/http"

	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	userdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/src/infra/http/responder"
)

func (uc *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	var userDto userdto.CreateUserDTO
	if !res.DecodeJSONBody(r, &userDto) {
		return
	}

	userDomain := userdomain.NewCreateUserDomain(
		userDto.Name,
		userDto.Email,
		userDto.Password,
	)

	ctx := context.Background()

	user, err := uc.CreateUserUseCase.Execute(ctx, userDomain)

	if err != nil {
		res.Error(err, 400)
		return
	}

	res.AddBody(userdto.NewUserCreatedResponse(user))
	res.Write(201)
}
