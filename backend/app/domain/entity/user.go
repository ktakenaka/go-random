package entity

import (
	"time"
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
	err := validate.Struct(u)
	return err
}
