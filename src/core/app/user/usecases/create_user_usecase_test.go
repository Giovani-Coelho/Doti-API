package userUseCase_test

import (
	"context"
	"testing"

	userUseCase "github.com/Giovani-Coelho/Doti-API/src/core/app/user/usecases"
	userDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	userDTO "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user/dtos"
	mock_repository "github.com/Giovani-Coelho/Doti-API/test/mocks/repository"
)

func TestCreateUserUseCase(t *testing.T) {
	mockRepo := &mock_repository.MockUserRepository{
		CreateFn: func(ctx context.Context, user userDomain.IUserDomain) (userDomain.IUserDomain, error) {
			return user, nil
		},
		CheckUserExistsFn: func(ctx context.Context, email string) (bool, error) {
			return false, nil
		},
	}

	createUser := userUseCase.NewCreateUserUseCase(mockRepo)

	ctx := context.Background()

	user := userDTO.CreateUserDTO{
		Name:     "New User",
		Email:    "newuser@example.com",
		Password: "password123",
	}

	t.Run("Create new user successfully", func(t *testing.T) {
		err := createUser.Execute(ctx, user)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}
	})

	t.Run("User already exists", func(t *testing.T) {
		mockRepo.CheckUserExistsFn = func(
			ctx context.Context, email string,
		) (bool, error) {
			return true, nil
		}

		err := createUser.Execute(ctx, user)

		if err == nil {
			t.Fatalf("expected: the user already exists. But we got: %v", err)
		}
	})
}
