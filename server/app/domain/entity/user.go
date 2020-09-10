package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint64 `gorm:"primary_key"`
	GoogleSub string `validate:"required"`
	Email     string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Validate() error {
	validate = validator.New()
	err := validate.Struct(u)
	return err
}
