package repository

import (
	"context"
	"database/sql"

	dtos "github.com/Giovani-Coelho/Doti-API/src/application/dtos/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/db/sqlc"
	"github.com/google/uuid"
)

func NewUserRepository(dtb *sql.DB) IUserRepository {
	return &UserRepository{
		DB:      dtb,
		Queries: sqlc.New(dtb),
	}
}

type UserRepository struct {
	DB      *sql.DB
	Queries *sqlc.Queries
}

type IUserRepository interface {
	Create(ctx context.Context, userDto dtos.CreateUserDto) error
}

func (ur *UserRepository) Create(ctx context.Context, userDTO dtos.CreateUserDto) error {
	userEntity := sqlc.CreateUserParams{
		ID:       uuid.New(),
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}

	err := ur.Queries.CreateUser(ctx, userEntity)

	if err != nil {
		return err
	}

	return nil
}
