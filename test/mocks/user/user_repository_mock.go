package mocks

import (
	"context"

	dtos "github.com/Giovani-Coelho/Doti-API/src/application/dtos/user"
)

type MockUserRepository struct {
	MockCreate          func(ctx context.Context, userDto dtos.CreateUserDto) error
	MockCheckUserExists func(ctx context.Context, email string) (bool, error)
}

func (m *MockUserRepository) Create(
	ctx context.Context,
	userDto dtos.CreateUserDto,
) error {
	if m.MockCreate != nil {
		return m.MockCreate(ctx, userDto)
	}

	return nil
}

func (m *MockUserRepository) CheckUserExists(
	ctx context.Context,
	email string,
) (bool, error) {
	if m.MockCheckUserExists != nil {
		return m.MockCheckUserExists(ctx, email)
	}

	return false, nil
}