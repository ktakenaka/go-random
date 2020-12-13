package entity

import (
	"time"

	validator "github.com/go-playground/validator/v10"
)

// User entity
type User struct {
	Base
	GoogleSub string `validate:"required"`
	Email     string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Validate validation
func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	return err
}
