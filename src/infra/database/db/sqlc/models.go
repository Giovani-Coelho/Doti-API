// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Email    string
	Name     string
	Password string
}
