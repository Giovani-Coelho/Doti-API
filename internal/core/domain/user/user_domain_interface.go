package userdomain

import "time"

type User interface {
	EncryptPassword()

	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetIsAdmin() bool
	GetCreateAt() time.Time
	GetUpdatedAt() time.Time

	IsValid() bool
	IsValidEmail() bool

	ValidatePassword() error
}

func New(
	id string,
	name string,
	email string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	return &userDomain{
		id:        id,
		name:      name,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func NewCreateUser(
	name string, email string, password string,
) User {
	return &userDomain{
		name:     name,
		email:    email,
		password: password,
	}
}

func NewSignInUser(
	email string, password string,
) User {
	return &userDomain{
		email:    email,
		password: password,
	}
}
