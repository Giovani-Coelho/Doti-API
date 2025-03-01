package userController

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/application/user/dtos"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

func (uc *UserControllers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user userDTO.CreateUserDTO

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		httpErr := rest_err.NewBadRequestError("Unable to parse request body")

		res, err := json.Marshal(httpErr)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(res)
		return
	}

	ctx := context.Background()

	err = uc.UserServices.CreateUser(ctx, user)
	if err != nil {
		if httpErr, ok := err.(*rest_err.RestErr); ok {
			res, err := json.Marshal(httpErr)
			if err != nil {
				log.Fatal(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			w.Write(res)
			return
		}

		return
	}
}
