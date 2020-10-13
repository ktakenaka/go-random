package entity

import (
	"time"

	validator "github.com/go-playground/validator/v10"
)

// Sample entity
type Sample struct {
	ID        uint64 `gorm:"primary_key"`
	Title     string `validate:"max=20,required"`
	Content   string `validate:"max=100"`
	UserID    uint64 `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Validate with validator v10
func (s *Sample) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	return err
}
